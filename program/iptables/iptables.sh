#! /bin/bash

# 清除已有规则（可选，谨慎操作）
# sudo iptables -t nat -F

# 设置 POSTROUTING 链的 MASQUERADE 规则
# 将 enp12s0 出口的流量源地址伪装成 wlan0 的 IP
sudo iptables -t nat -A POSTROUTING -o wlan0 -j MASQUERADE

# 允许转发的流量通过 FORWARD 链
sudo iptables -A FORWARD -i enp12s0 -o wlan0 -j ACCEPT
sudo iptables -A FORWARD -i wlan0 -o enp12s0 -m state --state ESTABLISHED,RELATED -j ACCEPT
