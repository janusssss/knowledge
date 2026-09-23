好的，这是一份 OpenRC 常用命令速查表，按功能分类整理：

### 📋 OpenRC 常用命令速查表

| 功能分类 | 命令 | 说明 |
|---|---|---|
| **服务状态** | | |
| | `rc-status` | 显示当前运行级别下所有服务的状态 |
| | `rc-status -s` | 以简短格式列出所有服务状态 |
| | `rc-status --crashed` | 只显示崩溃/失败的服务 |
| | `rc-status -a` | 显示所有运行级别下的服务 |
| | `rc-service <service> status` | 查看单个服务的详细状态 |
| **服务操作** | | |
| | `rc-service <service> start` | 启动服务 |
| | `rc-service <service> stop` | 停止服务 |
| | `rc-service <service> restart` | 重启服务 |
| | `rc-service <service> reload` | 重载配置（不中断服务） |
| | `rc-service <service> zap` | 强制标记服务为已停止（清理残留状态） |
| **开机自启管理** | | |
| | `rc-update add <service>` | 添加服务到默认运行级别（开机自启） |
| | `rc-update add <service> <runlevel>` | 添加服务到指定运行级别 |
| | `rc-update del <service>` | 从运行级别移除服务（取消自启） |
| | `rc-update show` | 显示所有运行级别的服务列表 |
| | `rc-update -v show` | 显示所有可用服务及其所在运行级别 |
| **运行级别管理** | | |
| | `rc` | 切换到默认运行级别（通常用于启动时） |
| | `openrc` | 切换到指定运行级别（如 `openrc default`） |
| | `rc default` | 切换到 default 运行级别 |
| | `rc shutdown` | 切换到关机运行级别 |
| | `rc-status <runlevel>` | 查看指定运行级别的状态 |
| **服务列表** | | |
| | `ls /etc/init.d/` | 列出所有可用服务脚本 |
| | `rc-update -v show` | 列出所有服务及其运行级别归属 |
| **日志与调试** | | |
| | `rc-service <service> --verbose start` | 详细模式启动服务，便于排错 |
| | `rc-service <service> --debug start` | 调试模式启动服务 |
| | `/var/log/rc.log` | OpenRC 自身日志（需在 `/etc/rc.conf` 中设 `rc_logger="YES"`） |
| **配置文件** | | |
| | `/etc/init.d/<service>` | 服务脚本 |
| | `/etc/conf.d/<service>` | 服务专属配置变量 |
| | `/etc/rc.conf` | OpenRC 全局配置 |
| | `/etc/runlevels/` | 各运行级别的服务软链接目录 |

### 💡 几个实用小技巧

1. **不需要加 `.service` 后缀**：OpenRC 里直接写服务名即可，如 `rc-service sshd restart`，而不是 `sshd.service`。

2. **`rc-service` vs `rc-update` 的分工**：
   - `rc-service` → 管**现在**（启动、停止、重启）
   - `rc-update` → 管**开机**（是否自启）

3. **查看服务属于哪个运行级别**：
   ```bash
   rc-update -v show | grep <service>
   ```

4. **强制清理服务状态**：如果服务状态卡住（比如显示 started 但进程已死），可以用：
   ```bash
   rc-service <service> zap
   rc-service <service> start
   ```

5. **Gentoo 用户注意**：Gentoo 上 OpenRC 是默认 init 系统，Alpine Linux 也默认使用 OpenRC，两者命令基本一致。

需要我再补充某个具体场景（比如网络配置、日志排查）的用法吗？
