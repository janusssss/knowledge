### 安装

```bash
# 图形化界面有时权限不够，需要使用命令安装
sudo woeusb --target-filesystem NTFS --device /home/janus/Downloads/Windows11_x64.iso /dev/sda


sudo ./Ventoy2Disk.sh -i /dev/sda
sudo mkfs.vfat -F 32 /dev/sda
```

```bash
genisoimage -iso-level 4 -l -U -J -V "Surface_Recovery" \
-o surface_recovery.iso -b bootmgr.efi surface_recovery/
```

```bash
# 使用 QEMU 模拟 UEFI 启动（需安装 qemu 和 edk2-ovmf）
qemu-system-x86_64 -bios /usr/share/edk2-ovmf/x64/OVMF_CODE.fd -cdrom surface_recovery.iso
```

```bash
# 格式化 U 盘为 FAT32（若文件 ≤4GB）或 NTFS
sudo mkfs.fat -F32 /dev/sdX1  # 或 sudo mkfs.ntfs /dev/sdX1

# 挂载并复制文件
sudo mount /dev/sdX1 /mnt
cp -r surface_recovery/* /mnt/

# 设置 UEFI 引导（关键！）
sudo grub-install --target=x86_64-efi --removable --boot-directory=/boot /dev/sda1
grub-mkconfig -o /boot/grub/grub.cfg
sudo umount /mnt
```

