# 测试网线还是转换头不满千兆
ethtool -s 网口名 speed 1000 duplex full autoneg off
转换器问题:直接断开连接
网线问题:依然能跑满百兆或超过百兆

# 网线通过无线上网
```fish
sudo sysctl -w net.ipv4.ip_forward=1

sudo iptables -t nat -A POSTROUTING -o wlan0 -j MASQUERADE

# 允许转发
sudo iptables -A FORWARD -i enp0s13f0u1u4 -o wlan0 -j ACCEPT
sudo iptables -A FORWARD -i wlan0 -o enp0s13f0u1u4 -m state --state RELATED,ESTABLISHED -j ACCEPT
```
