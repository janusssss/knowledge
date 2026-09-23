### 安装教程

1. 下载二进制可执行文件

2. 将下载的二进制可执行文件重名名为 `mihomo` 并移动到 `/usr/local/bin/`

   ```bash
   cp mihomo /usr/local/bin
   cp config.yaml /etc/mihomo
   ```

3. 以守护进程的方式，运行 mihomo。

   ```bash
   /etc/systemd/system/mihomo.server
   systemctl daemon-reload
   systemctl enable mihomo
   systemctl start mihomo
   systemctl reload mihomo
   systemctl status mihomo
   journalctl -u mihomo -o cat -fs
   ```

```bash
# 源码包
paru/yay -S mihomo-party
# 二进制包
paru/yay -S mihomo-party-bin
# Git 版本源码包
paru/yay -S mihomo-party-git
# 使用系统electron的源码包
paru/yay -S mihomo-party-electron
# 使用系统electron的二进制包
paru/yay -S mihomo-party-electron-bin
