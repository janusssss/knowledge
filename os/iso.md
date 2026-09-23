要将你当前 笔记本上已配置好的 Arch Linux + Niri（Wayland 合成器）系统 打包成一个 可离线安装的预配置 ISO，并能在其他电脑上通过虚拟机或物理机直接安装使用，这是一个典型的 “定制 Arch Linux 安装镜像” 需求。

虽然 Arch 官方不提供图形化安装器或预装 ISO，但社区有成熟方案。以下是最可行、稳定且支持离线安装的推荐路径：

✅ 最佳方案：使用 archiso 构建自定义 ISO

archiso 是 Arch 官方用于构建安装 ISO 的工具，完全开源、可靠，支持嵌入任意软件包、配置文件和脚本，天然支持离线安装。

🔧 整体流程概览

1. 在你的 Arch 系统中安装 archiso
2. 复制官方配置模板（releng）
3. 将你的系统配置（Niri、dotfiles、用户设置等）打包进 ISO
4. 将所需软件包缓存到 ISO 内（实现离线安装）
5. 构建 ISO
6. 用该 ISO 在虚拟机/物理机启动 → 自动部署你的完整环境

🛠️ 详细步骤

第一步：安装 archiso
sudo pacman -S archiso

第二步：创建工作目录
mkdir ~/archlive
cp -r /usr/share/archiso/configs/releng/* ~/archlive/
cd ~/archlive

releng 是官方安装 ISO 的配置模板，支持完整安装流程。

第三步：添加你的软件包（关键！）

编辑 packages.x86_64 文件，加入你需要的所有包：
基础
base
linux
linux-firmware
networkmanager
sudo
git

图形 & Niri
xorg-server
wayland
niri
niri-settings
firefox
alacritty
swaybg
wl-clipboard

中文支持（可选）
noto-fonts
noto-fonts-cjk
ibus
ibus-libpinyin

其他你用的工具
neovim
htop
...

✅ 所有列出的包都会被 打包进 ISO 的 /pkg 目录，安装时无需联网！

第四步：嵌入你的配置文件（自动化部署）

把你的配置放入 airootfs/ 目录（这是 Live 环境的根文件系统）：

创建用户家目录结构
mkdir -p airootfs/etc/skel/.config/niri
mkdir -p airootfs/etc/skel/.config/sway

复制你的 Niri 配置
cp ~/.config/niri/config.ron airootfs/etc/skel/.config/niri/
cp ~/.bashrc airootfs/etc/skel/
cp ~/.zshrc airootfs/etc/skel/

设置默认 shell（可选）
echo "export SHELL=/bin/zsh" >> airootfs/etc/skel/.profile

安装后新建用户会自动继承这些配置（因为 useradd -m 会复制 /etc/skel）。

第五步：自定义安装后脚本（可选但推荐）

你可以修改 airootfs/root/customize_airootfs.sh 或在 build.sh 中加入 post-install 脚本，例如：

- 自动创建用户
- 启用 NetworkManager
- 设置默认显示管理器（如 greetd + niri）

示例：在 airootfs/root/ 下创建 post-install.sh：
!/bin/bash
自动创建用户（安装时由 pacstrap 后 chroot 执行）
useradd -m -G wheel -s /bin/zsh myuser
echo "myuser:password" | chpasswd
sed -i 's/# %wheel ALL=(ALL) ALL/%wheel ALL=(ALL) ALL/' /etc/sudoers
systemctl enable NetworkManager

⚠️ 实际执行需在 build.sh 的 make_customize_airootfs() 阶段调用，或通过 pacstrap 后 hook 注入。

第六步：启用离线包缓存（核心！）

确保所有包都从本地缓存安装：

将你系统中已安装的包缓存复制到工作目录
mkdir -p work/x86_64/airootfs/var/cache/pacman/pkg
cp /var/cache/pacman/pkg/*.pkg.tar.zst work/x86_64/airootfs/var/cache/pacman/pkg/

或者更简单：在构建时让 mkarchiso 自动下载并打包（即使离线，只要之前在线下载过即可）。

💡 提示：首次构建需联网下载包；之后可断网重用缓存。

第七步：构建 ISO

sudo mkarchiso -v -w work -o out/ .

- -w work：工作目录（避免污染）
- -o out/：输出 ISO 到 out/ 目录

构建完成后，你会得到类似 out/archlinux-2026.03.23-x86_64.iso 的文件。

第八步：测试 ISO

用 QEMU 测试（无需写 U 盘）：
qemu-system-x86_64 -m 4G -enable-kvm -cdrom out/archlinux-*.iso -boot d

进入 Live 环境后，运行标准安装流程（或你定制的自动化脚本），即可部署你的 Niri 系统。

🚀 进阶：全自动无人值守安装（可选）

如果你希望 完全无需交互，可以：

1. 编写 install.sh 脚本放在 Live 系统中
2. 使用 fdisk + mkfs + mount + pacstrap + arch-chroot 全自动分区安装
3. 在 ISO 启动菜单中添加一个 “Auto Install” 选项

参考项目：
- archinstall（官方半自动安装器）
- archiso-profiles（Chris Titus 的自动化 ISO）

✅ 优势总结
特性   是否满足
离线安装   ✅ 所有包内嵌
保留 Niri 配置   ✅ 通过 /etc/skel
可在虚拟机使用   ✅ 标准 ISO
可在物理机安装   ✅ 支持 UEFI/BIOS
完全自由定制   ✅ 任意软件+脚本
符合 Arch 哲学   ✅ 基于官方工具链

📌 注意事项

- ISO 大小 ≈ 所有软件包总和（通常 2–4 GB）
- 不要包含敏感信息（如 SSH 密钥、密码）
- 如果用了 AUR 包（如某些 Niri 插件），需提前编译为 .pkg.tar.zst 并放入 work/.../pkg/
- 推荐在干净的 Arch 虚拟机中构建，避免污染

🔚 结语

使用 archiso 是目前构建 预配置、离线可用 Arch Linux ISO 的唯一官方推荐且生产级方案。虽然需要一点学习成本，但一旦掌握，你可以打造属于自己的 “Arch 发行版”，完美复刻你的 Niri 工作环境。

如果你需要，我可以提供一个完整的 archiso 配置模板（含 Niri + 中文 + 自动用户创建），只需替换你的配置即可一键构建！
