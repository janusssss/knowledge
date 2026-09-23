# 硬件解码
```bash
# 确定GPU型号
lspci
# 检查内核是否启动
lomod
# 安装驱动
yay -S intel-media-driver libva-utils --noconfirm
# 验证
vainfo
```

# 安装声卡驱动
```bash
# 安装相关驱动
sudo pacman -S pipewire wireplumber pipewire-audio pipewire-alsa pipewire-pulse alsa-utils
# 测试
speaker-test -c 3
```
