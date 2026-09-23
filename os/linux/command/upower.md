### 概述

​	`UPower`是一个系统守护进程和命令行工具，用于统一管理电源设备和提供电源信息。

### 查看电池设备地址

```bash
upower -e

# BAT(Battery)后缀的电源
# ADP(Adapter)后缀为电源适配器
# DisplayDevice为管理电源的虚拟设备

upower -i /org/freedesktop/UPower/devices/battery_BAT1
```

