### 问题过高排查
```bash
[janus@Arch ~]$ ps -eo pid,ppid,cmd,%mem,%cpu --sort=-%cpu | head -20
    PID    PPID CMD                         %MEM %CPU
   7722    1258 /usr/lib/localsearch-3       0.1 56.1
    503       1 /usr/bin/mount.ntfs-3g /dev  0.0 42.8
   1348    1258 /usr/bin/gnome-shell         0.8  4.8
   3836    3768 /usr/bin/clickhouse-server   3.1  4.2
   2752    1348 /usr/lib/firefox/firefox     1.4  1.5
   3125    2834 /usr/lib/firefox/firefox -c  0.9  1.1
   1751    1258 /usr/bin/gnome-software --g  0.3  1.1
   2620    1565 /opt/mihomo-party/mihomo-pa  0.7  0.7
   5099    1258 /usr/bin/kgx --gapplication  0.3  0.5
   2568    1565 /opt/mihomo-party/resources  0.2  0.4
   2249    1810 /opt/mihomo-party/mihomo-pa  0.3  0.4
   1565    1341 /opt/mihomo-party/mihomo-pa  0.6  0.3
   2251    1552 /usr/lib/webkit2gtk-4.1/Web  1.0  0.3
   2466    1552 /usr/lib/webkit2gtk-4.1/Web  0.9  0.2
     94       2 [kworker/u33:0-i915_flip]    0.0  0.1
   1681    1540 /usr/lib/ibus/ibus-extensio  0.0  0.1
      1       0 /sbin/init                   0.0  0.1
    467       2 [kworker/u32:5-i915]         0.0  0.1
   1283    1282 dbus-broker --log 10 --cont  0.0  0.1
```