# 同时开启wifi和hotspot


```bash
sudo iw phy phy0 interface add wlan0_ap type managed

sudo iw dev wlan0_ap set type __ap
sudo ip link set wlan0_ap up


sudo tee /etc/hostapd/same_channel.conf << EOF
interface=wlan0ap
driver=nl80211
ssid=MyHotspot
hw_mode=g
channel=169
wmm_enabled=1
wpa=2
wpa_passphrase=12345678
wpa_key_mgmt=WPA-PSK
rsn_pairwise=CCMP
EOF
















sudo iw dev wlan0_ap set channel 6
sudo ip addr add 192.168.42.1/24 dev wlan0_ap


cat << EOF | sudo tee /tmp/activate-hostapd.conf
interface=wlan0_ap
driver=nl80211
ssid=TestActivate
channel=9
hw_mode=g
wpa=2
wpa_passphrase=""
wpa_key_mgmt=WPA-PSK
rsn_pairwise=CCMP
ignore_broadcast_ssid=0
EOF

cat << EOF | sudo tee /tmp/activate-hostapd.conf
interface=wlan0_ap
driver=nl80211
ssid=janus-test
channel=9
hw_mode=g
rsn_pairwise=CCMP
ignore_broadcast_ssid=0
EOF

sudo hostapd -B /tmp/activate-hostapd.conf


sudo ip addr flush dev wlan0_ap
sudo ip addr add 192.168.42.1/24 dev wlan0_ap


sudo dnsmasq -i wlan0_ap     --dhcp-range=192.168.42.2,192.168.42.100,255.255.255.0,12h     --dhcp-option=option:router,192.168.42.1     --dhcp-option=option:dns-server,8.8.8.8,8.8.4.4
echo 1 | sudo tee /proc/sys/net/ipv4/ip_forward
sudo iptables -t nat -F
sudo iptables -t nat -A POSTROUTING -o wlan0 -j MASQUERADE
sudo iptables -A FORWARD -i wlan0_ap -o wlan0 -j ACCEPT
sudo iptables -A FORWARD -i wlan0 -o wlan0_ap -m state --state RELATED,ESTABLISHED -j ACCEPT
````

```bash
sudo iw dev wlan0_ap del


ip link show

nm-connection-editor

sudo ip addr del 192.168.123.45/24 dev wlan0_ap
```