### 配置文件格式
```bash
[Desktop Entry]#标识这是一个桌面条目
Name=Firefox Web Browser#显示的应用程序名称
Type=Application#指定条目类型为应用程序,空格敏感，不能有多余的空格
Exec=/usr/bin/firefox %u#启动应用程序的命令。`%u`表示处理URL参数
Icon=/usr/share/icons/hicolor/scalable/apps/firefox.svg #指定应用程序图标路径或主题名称
Categories=WebBrowser;Application; #分类信息，便于在菜单中组织和显示
Encoding=UTF-8 #文件编码格式，通常设为UTF-8
Comment=Start Firefox #简要描述

# 可选：设置图标主题路径，如果直接使用主题图标
#Icon=system-run #/usr/share/icons
```

### 目录
位于~/.local/share/applications/或者/usr/local/share/applications/目录下。


```bash
# 微信
[Desktop Entry] 
Name=wechat
Type=Application
Exec=/opt/WeChatLinux_x86_64.AppImage
Icon=/home/janus/Pictures/icons/wechat.png
```