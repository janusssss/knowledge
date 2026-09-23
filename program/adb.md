### 概述
ADB（Android Debug Bridge）是Android 的开发者、测试人员甚至高级用户都必须掌握的核心工具。
ADB 是一个功能强大的命令行工具，它扮演着电脑（开发机）与 Android 设备（或模拟器）之间的桥梁。

- Client（客户端）：就是你在一台电脑上运行的 adb 命令。
- Server（服务端）：一个在电脑后台运行的进程（adb server），它负责管理客户端与所有连接的 Android 设备之间的通信。
- Daemon（守护进程）：一个在 Android 设备或模拟器上后台运行的进程（adbd），它负责执行从 Server 端收到的命令。

### 操作
```bash
# 安装应用
adb install app.apk
# 覆盖安装（更新）
adb install -r app.apk
# 卸载应用（保留数据和缓存）
adb uninstall com.example.myapp
# 彻底卸载应用（清除数据和缓存）
adb uninstall -k com.example.myapp


# 安装应用
adb install app.apk
# 覆盖安装（更新）
adb install -r app.apk
# 卸载应用（保留数据和缓存）
adb uninstall com.example.myapp
# 彻底卸载应用（清除数据和缓存）
adb uninstall -k com.example.myapp

# 进入设备的 Linux Shell 环境
adb shell


# 查看已连接的设备
adb devices
# 查看运行中的进程
adb shell ps
# 查看详细的 CPU、内存等性能信息
adb shell top
# 查看应用日志（非常重要的调试功能）
adb logcat
# 过滤特定应用的日志
adb logcat | grep "MyApp"


# 启动一个 Activity
adb shell am start -n com.example.myapp/.MainActivity
# 发送一个广播
adb shell am broadcast -a "my.custom.action"
# 模拟点击事件
adb shell input tap 500 500
# 模拟滑动
adb shell input swipe 300 1000 300 500
```

## 文件
```bash
adb push <本地路径> <设备路径>
adb pull <设备路径> <本地路径>

# 软件目录 跳过权限
adb exec-out run-as com.janus.bookkeeper cat files/preferences.json > ./files/fyne/preferences.json
```

# 切换用户
```bash
adb shell 
run-as your.package.name
```