# ADB

## 命令

- adb devices : 检查设备是否连接
- adb shell pm list packages : 查找应用的包名
- adb shell pm uninstall <包名> :  卸载应用
- adb shell dumpsys activity : 查看当前运行的应用
- adb shell dumpsys package <包名> : 查看应用的详细信息
- adb shell am monitor : 监视应用程序的活动，打印出应用程序的启动、切换和相关活动信息，包括包名
- adb shell pm disable-user --user 0 <name> : 禁用软件
- adb shell pm list packages -d : 禁用软件列表

```bash
adb shell pm disable-user --user 0  com.oneplus.bbs
```

- 禁用软件重新启用

  ```bash
  adb shell pm enable com.oneplus.member
  ```

  
