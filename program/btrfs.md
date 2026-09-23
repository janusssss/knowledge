# snapper
opensuse开发的快照软件，超级好用。

## 安装
```bash
sudo pacman -S snapper btrfs-assistant grub-btrfs inotify-tools boost-libs
sudo systemctl enable --now grub-btrfsd
```

- snapper 是主程序；
- btrfs-assistant 是GUI（图形化交互界面），同时提供了几个简单的命令，进一步简化快照回档需要的操作。我们还没有安装桌面环境，但是肯定会用到，先装上。
- grub-btrfs inotify-tools 在创建快照的时候自动在grub菜单里添加快照启动项

## 创建快照配置
```bash
sudo snapper -c root create-config /
sudo snapper -c home create-config /home
```

## 设置合理的快照策略
```bash
sudo vim /etc/snapper/configs/root

ALLOW_GROUPS="wheel"允许wheel组的用户无须sudo就可以操作快照
NUMBER_LIMIT="10" 设置最多保存10个快照，超出后会按照时间顺序删除旧快照。

TIME_LIMIT_HOURLY="3"每隔一小时创建的快照保存3个。

TIME_LIMIT_DAILY="1"每日快照保存1个。

其他的TIME_LIMIT数量都设置为0。这样就仅保存三个小时前的状态和昨天的状态。
```

## 开启按时间自动创建快照和自动清理
```bash
sudo systemctl enable --now snapper-timeline.timer
sudo systemctl enable --now snapper-cleanup.timer
sudo grub-mkconfig -o /boot/grub/grub.cfg
```


## 回滚
```bash
btrfs-assistant -l
btrfs-assistant -r 1

sudo snapper -c root list 
sudo snapper -c root undochange 1..0
```
```
```

# ext4转btrfs
```bash
btrfs-convert /dev/sdb1
genfstab -U /mnt >> /mnt/etc/fstab
sudo grub-mkconfig -o /boot/grub/grub.cfg
```





