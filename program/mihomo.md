好的，我们来把 Mihomo 从服务端到客户端，再到规则过滤，完整捋一遍。

目的很明确：**云主机做服务端，只负责把流量转出去；台式机做客户端，决定哪些流量走云主机（比如 Steam），哪些直连。**

---

### ⚙️ 第一步：服务端（云主机）配置

服务端的角色是“纯转发出口”，不需要任何复杂的规则。配置如下：

```yaml
# /etc/mihomo/config.yaml （云主机上）
port: 7890
socks-port: 7891
allow-lan: true           # 允许来自局域网/SSH 隧道的连接
bind-address: '*'         # 监听所有地址
mode: global              # 服务端不做分流，所有请求都转发
log-level: info

dns:
  enable: true
  nameserver:
    - 8.8.8.8
    - 1.1.1.1
```

启动命令一样，直接用 service 跑就行：

```bash
systemctl start mihomo
```

---

### 💻 第二步：客户端（台式机）配置

客户端是真正做“分流”决策的地方。台式机上也需要安装 Mihomo，然后配置文件这样写：

```yaml
# 台式机上的 config.yaml
# 1. 定义监听端口（供本机程序和浏览器使用）
port: 7890          # HTTP 代理端口
socks-port: 7891    # SOCKS5 代理端口
allow-lan: false    # 只允许台式机自己用
bind-address: 127.0.0.1

# 2. 默认模式和代理组
mode: rule          # 按规则匹配，匹配不到就走 DIRECT（直连）

# 3. 定义上游服务器：用 SSH 隧道连到云主机的端口
proxies:
  - name: "vps-proxy"
    type: socks5
    server: 127.0.0.1
    port: 7891       # 通过 SSH 隧道映射到云主机的 7891
    # udp: true      # 需要 UDP（比如游戏）的话取消这一行的注释

# 4. 代理组
proxy-groups:
  - name: "分流"
    type: select
    proxies:
      - "vps-proxy"
      - "DIRECT"

# 5. 核心：规则配置（决定哪些走代理）
rules:
  # ===== Steam 相关全走代理 =====
  - DOMAIN-SUFFIX,steampowered.com,分流
  - DOMAIN-SUFFIX,steamcommunity.com,分流
  - DOMAIN-SUFFIX,steamcontent.com,分流
  - DOMAIN-SUFFIX,valvesoftware.com,分流
  - DOMAIN-SUFFIX,steamstatic.com,分流
  
  # ===== 其他常用海外服务 =====
  - DOMAIN-SUFFIX,github.com,分流
  - DOMAIN-SUFFIX,google.com,分流
  - DOMAIN-SUFFIX,openai.com,分流
  
  # ===== 国内网站和服务直连（保证速度） =====
  - DOMAIN-SUFFIX,baidu.com,DIRECT
  - DOMAIN-SUFFIX,bilibili.com,DIRECT
  - DOMAIN-SUFFIX,taobao.com,DIRECT
  - DOMAIN-SUFFIX,qq.com,DIRECT
  
  # ===== 局域网和私有 IP 直连 =====
  - IP-CIDR,10.0.0.0/8,DIRECT
  - IP-CIDR,172.16.0.0/12,DIRECT
  - IP-CIDR,192.168.0.0/16,DIRECT
  
  # ===== 兜底规则：没匹配到的全走代理 =====
  - MATCH,分流
```

---

### 🔗 第三步：建立 SSH 隧道进行端口映射

在台式机上执行这条命令，把你的本地端口和云主机的代理端口打通：

```bash
ssh -o ServerAliveInterval=60 -fNT -L 7891:localhost:7891 root@你的云主机公网IP
```

- 含义：把台式机的本地 `7891` 端口，映射到云主机的 `7891` 端口（Mihomo SOCKS5）
- 这样客户端配置里的 `server: 127.0.0.1:7891` 就能直连云主机了

**然后启动台式机的 Mihomo**。启动后，浏览器和系统就可以设置代理为 `127.0.0.1:7890`（HTTP）或 `127.0.0.1:7891`（SOCKS5）。

---

### 🎯 规则是怎么工作的？（你的核心问题）

上面配置的逻辑是这样的：

1.  **Steam 域名** 命中规则 → 流量交给 `分流` 代理组 → 走 `vps-proxy` → 云主机出国
2.  **百度、B 站等国内域名** 命中规则 → 直接走 `DIRECT`，不经过云主机
3.  **没匹配到的其他流量** 走 `MATCH,分流` → 默认也走云主机（你可以改成 `MATCH,DIRECT` 只让特定规则走代理）

规则的匹配顺序是**从上到下**的，一旦匹配就停止。所以域名匹配精度要求最高（`DOMAIN` > `DOMAIN-SUFFIX` > `DOMAIN-KEYWORD`），IP 规则放最后。

---

### 🧪 验证和调试

在台式机上，打开 Mihomo 的 Web 管理界面（默认 `http://127.0.0.1:9090/ui`，如果配置了 `external-controller`），或者直接打开浏览器访问 `https://steamcommunity.com`，看能不能打开。同时看管理界面的 Connections 部分，能看到 `steamcommunity.com` 的请求是否匹配到了“分流”规则，是否走了 `vps-proxy`。

---

### 📌 几个容易忽略的点

- **防火墙**：云主机上必须放行 `7890` 和 `7891` 端口。不用的时候最好关掉，减少被扫描的风险
- **Steam 客户端 vs 网页**：浏览器设置代理很简单，但 Steam 客户端本身不走系统代理。要让 Steam 也走，需要在台式机上用 Mihomo 的 **TUN 模式**（网络层代理），这操作稍微复杂些
- **DNS 泄露**：上面的配置只在代理层做了规则。如果你追求极致隐私，可以加 `dns` 配置块做更精细的控制

---

现在你已经有了一个基本可用的框架。先按这个配置跑起来，试试看 Steam 网页能不能正常访问。如果遇到连接失败，可以在云主机上跑 `curl -x socks5://127.0.0.1:7891 https://www.google.com` 看看代理本身是否工作正常。

跑通之后，如果你想折腾 TUN 模式实现游戏代理，我们再继续聊。
