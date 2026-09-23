要在你的KVM配置中增加GPU直通，有两种主要方式：核显直通和NVIDIA独显直通。以下是具体的配置步骤：

## 1. 首先确认硬件信息

```bash
# 查看显卡信息
lspci | grep -E "VGA|3D"
# 查看IOMMU分组
sudo virsh nodedev-list --cap pci
```

## 2. 核显直通配置（相对简单）

在你的配置文件中添加或修改以下部分：

```xml
<domain type="kvm">
  <!-- 在<features>部分添加IOMMU支持 -->
  <features>
    <acpi/>
    <apic/>
    <ioapic driver="kvm"/>
    <hyperv>
      <relaxed state="on"/>
      <vapic state="on"/>
      <spinlocks state="on" retries="8191"/>
      <vendor_id state="on" value="1234567890ab"/>
    </hyperv>
    <kvm>
      <hidden state="on"/>
    </kvm>
    <vmport state="off"/>
    <smm state="on"/>
  </features>

  <!-- 在<devices>部分添加核显设备 -->
  <devices>
    <!-- 删除原有的<video>配置 -->
    <!-- 添加核显直通 -->
    <hostdev mode="subsystem" type="pci" managed="yes">
      <driver name="vfio"/>
      <source>
        <!-- 使用lspci查看核显的PCI地址，通常是00:02.0 -->
        <address domain="0x0000" bus="0x00" slot="0x02" function="0x0"/>
      </source>
      <address type="pci" domain="0x0000" bus="0x00" slot="0x01" function="0x0"/>
    </hostcard>
    
    <!-- 可选：添加HDMI音频设备（如果核显有） -->
    <hostdev mode="subsystem" type="pci" managed="yes">
      <driver name="vfio"/>
      <source>
        <!-- 音频设备通常是00:02.1 -->
        <address domain="0x0000" bus="0x00" slot="0x02" function="0x1"/>
      </source>
      <address type="pci" domain="0x0000" bus="0x06" slot="0x00" function="0x0"/>
    </hostdev>
  </devices>
</domain>
```

## 3. NVIDIA独显直通配置（更复杂但性能更好）

### 步骤1：修改宿主机配置

```bash
# 编辑GRUB配置
sudo nano /etc/default/grub
# 修改GRUB_CMDLINE_LINUX行，添加：
# intel_iommu=on  # Intel CPU
# 或
# amd_iommu=on    # AMD CPU

# 重新生成GRUB配置
sudo update-grub

# 加载VFIO模块
sudo nano /etc/modules-load.d/vfio.conf
# 添加：
# vfio
# vfio_iommu_type1
# vfio_pci
# vfio_virqfd

# 屏蔽NVIDIA驱动
sudo nano /etc/modprobe.d/blacklist.conf
# 添加：
# blacklist nouveau
# blacklist nvidia

# 配置VFIO
sudo nano /etc/modprobe.d/vfio.conf
# 添加：
# options vfio-pci ids=10de:你的GPU_ID,10de:你的音频ID

# 重启系统
sudo reboot
```

### 步骤2：修改虚拟机配置

```xml
<domain type="kvm">
  <!-- 在<features>部分添加（同核显配置） -->
  
  <!-- 修改CPU配置 -->
  <cpu mode="host-passthrough" check="none" migratable="on">
    <topology sockets="1" dies="1" clusters="1" cores="4" threads="2"/>
    <!-- 添加CPU功能 -->
    <feature policy="require" name="vmx"/>
    <feature policy="require" name="svm"/>
    <feature policy="require" name="invtsc"/>
  </cpu>
  
  <!-- 修改时钟配置 -->
  <clock offset="localtime">
    <timer name="rtc" tickpolicy="catchup"/>
    <timer name="pit" tickpolicy="delay"/>
    <timer name="hpet" present="no"/>
    <timer name="hypervclock" present="yes"/>
    <timer name="kvmclock" present="no"/>
  </clock>
  
  <!-- 在<devices>部分添加NVIDIA GPU -->
  <devices>
    <!-- 删除原有的<video>和<sound>配置 -->
    
    <!-- 添加NVIDIA GPU -->
    <hostdev mode="subsystem" type="pci" managed="yes">
      <driver name="vfio"/>
      <source>
        <!-- 使用lspci查看NVIDIA GPU的地址 -->
        <address domain="0x0000" bus="0x01" slot="0x00" function="0x0"/>
      </source>
      <!-- ROM文件通常位于 /sys/bus/pci/devices/0000:01:00.0/rom -->
      <rom file="/path/to/gpu/rom.bin"/>
      <address type="pci" domain="0x0000" bus="0x07" slot="0x00" function="0x0"/>
    </hostdev>
    
    <!-- 添加GPU音频设备 -->
    <hostdev mode="subsystem" type="pci" managed="yes">
      <driver name="vfio"/>
      <source>
        <!-- NVIDIA音频设备通常是 01:00.1 -->
        <address domain="0x0000" bus="0x01" slot="0x00" function="0x1"/>
      </source>
      <address type="pci" domain="0x0000" bus="0x08" slot="0x00" function="0x0"/>
    </hostdev>
    
    <!-- 添加USB控制器用于键鼠 -->
    <controller type="usb" index="0" model="qemu-xhci" ports="15">
      <address type="pci" domain="0x0000" bus="0x03" slot="0x00" function="0x0"/>
    </controller>
    
    <!-- 添加输入设备 -->
    <input type="mouse" bus="usb"/>
    <input type="keyboard" bus="usb"/>
  </devices>
</domain>
```

## 4. 可选优化配置

```xml
<!-- 添加Hyper-V enlightenments -->
<hyperv>
  <relaxed state="on"/>
  <vapic state="on"/>
  <spinlocks state="on" retries="8191"/>
  <vpindex state="on"/>
  <runtime state="on"/>
  <synic state="on"/>
  <stimer state="on"/>
  <reset state="on"/>
  <vendor_id state="on" value="1234567890ab"/>
  <frequencies state="on"/>
</hyperv>

<!-- 添加内存大页支持 -->
<memoryBacking>
  <hugepages/>
</memoryBacking>

<!-- 修改qemu:commandline -->
<qemu:commandline xmlns:qemu="http://libvirt.org/schemas/domain/qemu/1.0">
  <qemu:arg value="-cpu"/>
  <qemu:arg value="host,kvm=off,hv_vendor_id=null"/>
  <qemu:arg value="-object"/>
  <qemu:arg value="input-linux,id=mouse1,evdev=/dev/input/by-id/鼠标设备"/>
  <qemu:arg value="-object"/>
  <qemu:arg value="input-linux,id=kbd1,evdev=/dev/input/by-id/键盘设备,grab_all=on,repeat=on"/>
</qemu:commandline>
```

## 重要注意事项：

1. **NVIDIA直通问题**：
   - Windows 11可能需要禁用Windows驱动签名验证
   - 部分NVIDIA消费级显卡有驱动限制
   - 可能需要使用`kvm=off`隐藏虚拟化环境

2. **性能优化**：
   - 为虚拟机预留CPU核心
   - 使用大页内存
   - 启用CPU pinning

3. **显示输出**：
   - GPU直通后，虚拟机将直接使用显卡的物理输出
   - 需要为虚拟机连接单独的显示器

建议先从核显直通开始尝试，相对简单且风险较低。NVIDIA独显直通需要更多配置，且可能遇到驱动限制问题。