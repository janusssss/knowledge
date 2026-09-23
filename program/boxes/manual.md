
# 查看GPU
```bash
sudo lspci -nn | grep -E "VGA|3D|Display"
```
# 核显直通
1. 启动iommu和vfio
```bash
# 1. 检查IOMMU是否启用
sudo dmesg | grep -i iommu
# 应该看到类似：IOMMU enabled

# 2. 检查核显是否被正确识别
sudo lspci -nn -s 00:02.0
# 记下设备ID，例如：[8086:9a49]

# 3. 检查当前驱动绑定
sudo lspci -k -s 00:02.0
# 查看使用的驱动，可能是i915

# 4. 检查VFIO模块加载
lsmod | grep vfio
```

# vfio
```bash
# 1. 查看核显的具体设备ID
sudo lspci -nn -s 00:02.0
# 输出示例：8086:9a49

# 2. 创建VFIO配置
sudo nano /etc/modprobe.d/vfio.conf
# 添加：
options vfio-pci ids=8086:9a49

# 3. 创建模块加载配置
sudo nano /etc/modules-load.d/vfio.conf
# 添加：
vfio
vfio_iommu_type1
vfio_pci
vfio_virqfd

# 4. 在早期启动时加载VFIO（重要！）
sudo nano /etc/mkinitcpio.conf
# 在MODULES行添加：
MODULES=(vfio_pci vfio vfio_iommu_type1)
# 确保i915不在MODULES中

# 5. 重新生成initramfs
sudo mkinitcpio -P

# 6. 重启系统
sudo reboot
```

```bash
# 1. 编辑GRUB配置
sudo nano /etc/default/grub

# 在GRUB_CMDLINE_LINUX中添加：
GRUB_CMDLINE_LINUX="... intel_iommu=on iommu=pt vfio-pci.ids=8086:9a49"

# 2. 重新生成GRUB配置
sudo grub-mkconfig -o /boot/grub/grub.cfg

# 3. 确保模块在启动后加载
echo "vfio" | sudo tee /etc/modules-load.d/vfio.conf
echo "vfio_iommu_type1" | sudo tee -a /etc/modules-load.d/vfio.conf
echo "vfio_pci" | sudo tee -a /etc/modules-load.d/vfio.conf

# 4. 重启
sudo reboot
```