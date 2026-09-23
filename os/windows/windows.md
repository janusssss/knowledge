### 命令

- netstat -ano | findstr ":0.0.0.0:<port>" : 这个端口是否使用情况

### PowerShell 

- 配置文件路径 

  ```bash
  $PROFILE
  # C:\Users\Administrator\Documents\WindowsPowerShell\Microsoft.PowerShell_profile.ps1
  ```

  ```bash
  Set-Alias etcd "C:\Janus\soft\etcd-v3.5.17-windows-amd64\etcd.exe" # 命令别名
  
  # 背景颜色
  $Host.UI.RawUI.BackgroundColor = 'Black'
  $Host.UI.RawUI.ForegroundColor = 'White'
  Clear-Host  # 刷新屏幕以应用颜色
  ```

- Start-Process powershell ：多开窗口

### bootmgfw.efi文件被误删恢复

1. 创建Windows安装介质并启动

2. 进入修复环境

   ```bash
   在Windows安装界面中，选择语言和其他首选项后点击“下一步”，然后点击屏幕底部的“修复计算机”。
   
   使用启动修复工具：
   选择“疑难解答” > “高级选项” > “启动修复”。这个过程会自动扫描并尝试修复任何可能导致系统无法启动的问题，包括恢复丢失的bootmgfw.efi文件。
   ```

# shell
```bash
# 查看序列号
Get-CimInstance -ClassName Win32_bios | select SerialNumber
```

# windows11密钥
卡号：【90】密钥：N2889-QMM3R-P98JV-V8787-JQKTY

方法①：点击此电脑-属性-更改密钥--输入密钥   点激活 

方法②：【快捷键】    win键+R键一起按，运行窗口输入 slui 3 【3前面有空格】，更改密钥--输入密钥 激活

如果无法激活：按住 win键+R键 输入  slui 4（四前面有个空格） 然后回车 国家任意选择 下一步，然后把1-9组数字,清晰的照片或截图发给我！


# windows10密钥
V2MHF-7BN3K-J4MKG-TYR3D-QRR9M