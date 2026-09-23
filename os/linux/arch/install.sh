echo "Server = https://mirrors.ustc.edu.cn/archlinux/$repo/os/$arch # 中国科学技术大学开源镜像站" >>/etc/pacman.d/mirrorlist
echo "Server = https://mirrors.tuna.tsinghua.edu.cn/archlinux/$repo/os/$arch # 清华大学开源软件镜像站" >>/etc/pacman.d/mirrorlist

sed -i 's/^#en_US.UTF-8/en_US.UTF-8/' /etc/locale.gen
sed -i 's/^#zh_CN.UTF-8/zh_CN.UTF-8/' /etc/locale.gen
locale-gen
echo "LANG=en_US.UTF-8" >>/etc/locale.conf
echo "LANG=en_US.UTF-8" >>/etc/profile

echo "ArchLinux" >>/etc/hostname
echo -e "127.0.0.1  localhost\n::1  localhost\n127.0.1.1 ArchLinux.localdomain  ArchLinux" >>/etc/hosts

useradd -m -g users -G wheel -s /bin/bash janus
passwd janus

sed -i 's/^#%wheel ALL=(ALL) ALL/%wheel ALL=(ALL) ALL/' /etc/sudoers

pacman -S grub efibootmgr efivar
grub-install --target=x86_64-efi --efi-directory=/boot --bootloader-id=Arch --recheck
grub-mkconfig -o /boot/grub/grub.cfg

pacman -S gnome gnome-extra gdm gnome-tweak-tool
systemctl enable gdm

pacman -S networkmanager
systemctl enable networkmanager

pacman -S bluez
systemctl enable bluetooth.service

echo "[archlinuxcn]" >>/etc/pacman.conf
echo "Server = https://mirrors.ustc.edu.cn/archlinuxcn/$arch" >>/etc/pacman.conf

sudo pacman -S yay base-devel linux-firmware
yay -S wqy-microhei wqy-microhei-lite wqy-bitmapfont wqy-zenhei ttf-arphic-ukai adobe-source-han-sans-cn-fonts adobe-source-han-serif-cn-fonts ttf-fira-code
yay -S ibus-rime
sudo pacman -S ttf-jetbrains-mono-nerd
sudo pacman -S noto-fonts noto-fonts-cjk noto-fonts-emoji ttf-liberation ttf-dejavu
sudo pacman -S pipewire wireplumber pipewire-audio pipewire-alsa pipewire-pulse --noconfirm
sudo pacman -S pipewire wireplumber pipewire-audio pipewire-alsa pipewire-pulse
su
