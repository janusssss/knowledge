
1. 连接WIFI

```bash
iwctl
device list
station <device-name> scan
station <device-name> get-networks
station <device-name> connect <wifit-name>
```

2. 设置时间

```bash
timedatectl
timedatectl set-ntp true
timedatectl set-timezone Asia/Shanghai
```

3. 磁盘分区

```bash
cfdisk /dev/nvme0n1 # 分盘


fdisk -l
mkfs.fat -F 32 /dev/<efi_system_partition>（EFI 系统分区）
mkfs.ext4 /dev/<root_partition> #格式化分区
```

4. 挂载磁盘

```bash
mount /dev/<root_partition> /mnt
mount --mkdir /dev/<efi_system_partition> /mnt/boot
```

5. 安装系统

```bash
genfstab -U /mnt >> /mnt/etc/fstab
pacstrap -K /mnt base linux base-devel linux-firmware  # 旧镜像需要 pacman -S archlinux-keyring
arch-chroot /mnt #进入系统
efibootmgr # 管理grub启动
```

6. 仓库源

```bash
# /etc/pacman.d/mirrorlist

Server = https://mirrors.ustc.edu.cn/archlinux/$repo/os/$arch # 中国科学技术大学开源镜像站
Server = https://mirrors.tuna.tsinghua.edu.cn/archlinux/$repo/os/$arch # 清华大学开源软件镜像站
Server = https://repo.huaweicloud.com/archlinux/$repo/os/$arch # 华为开源镜像站
Server = http://mirror.lzu.edu.cn/archlinux/$repo/os/$arch # 兰州大学开源镜像站
Server = http://mirrors.aliyun.com/archlinux/$repo/os/$arch # 阿里云镜像站
```

7. 本地化

```bash
# 取消注释
sed -i 's/^#en_US.UTF-8/en_US.UTF-8/' /etc/locale.gen
sed -i 's/^#zh_CN.UTF-8/zh_CN.UTF-8/' /etc/locale.gen
# 生成 Local
locale-gen

# 这里不建议将 en_US.UTF-8 改为zh_CN.UTF-8 ，这样会导致终端乱码！ 
echo "LANG=en_US.UTF-8" >> /etc/locale.conf
echo "LANG=en_US.UTF-8" >> /etc/profile
```

8. 主机名与hosts配置

```bash
# 这里可以换成自己想要的名字，将主机名写入/etc/hostname
echo "ArchLinux" >> /etc/hostname
# 将 下面hostname换成自己的主机名，与 /etc/hostname 里面的名字一样
echo -e "127.0.0.1  localhost\n::1  localhost\n127.0.1.1 ArchLinux.localdomain  ArchLinux" >> /etc/hosts
```

9. 增加用户

```bash
useradd -m -g users -G wheel -s /bin/bash <username>
passwd 
<password>

打开 /etc/sudoers 文件，找到 root ALL=(ALL) ALL 并依葫芦画瓢添加 janus ALL=(ALL) ALL 即可。
#%wheel ALL=(ALL) ALL
```

10. 安装grub

```bash
cat /proc/cpuinfo	# 查看cpu型号
# pacman -S intel-ucode	# intel 电脑安装

pacman -S grub efibootmgr efivar os-prober
grub-install --target=x86_64-efi --efi-directory=/boot --bootloader-id=Arch --recheck
# echo "GRUB_DISABLE_OS_PROBER=false" >> /etc/default/grub windows双系统
grub-mkconfig -o /boot/grub/grub.cfg
```

11. 安装gnone

```bash
pacman -S gnome gnome-extra gdm  #  gnome-tweak-tool
systemctl enable gdm

# 网络
pacman -S networkmanager
systemctl enable networkmanager

# blue
pacman -S bluez
systemctl enable bluetooth.service
```

12. 重启系统，进入GUI界面

---

# GUI配置
1. AUR配置

```bash
# 编辑pacman配置文件
vim /etc/pacman.conf

## 中国科学技术大学 (安徽合肥) (ipv4, ipv6, http, https)
[archlinuxcn]
Server = https://mirrors.ustc.edu.cn/archlinuxcn/$arch
```

2. 中文字体
```bash
sudo pacman -S yay
yay -S wqy-microhei wqy-microhei-lite wqy-bitmapfont wqy-zenhei ttf-arphic-ukai adobe-source-han-sans-cn-fonts adobe-source-han-serif-cn-fonts ttf-fira-code
```

3. 安装中文输入法
```bash
yay -S ibus-rime
ibus-setup # 设置
```


### 安装中文输入法

```bash
sudo pacman -S ibus-rime

# ibus-setup 设置

```

