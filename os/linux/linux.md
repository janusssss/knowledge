### chomd

```bash
# 给文件file_name设置同用户组权限为读写执行
chmod -R g=rwx file_name
# 增加file_name文件所有用户的执行权限 => chmod a+x flie_name
chmod +x file_name 
```

```bash
chmod [option] mode[,mode]... file...
```

| 符合模式           | 含义                    |
| ------------------ | ----------------------- |
| `+`                | 增加                    |
| `-`                | 移除                    |
| `=`                | 设定                    |
| `u`                | `user` 用户（所有者）   |
| `g`                | `group`文件所属的用户组 |
| `o`                | `others`其他所有人      |
| `a`                | `all`全部               |
| `r`                | `read`读                |
| `w`                | `write`写               |
| `x`                | `execute`执行           |
| `-R` `--recursive` | 递归                    |

### chown

```bash
sudo chown -R janus:janus /path/to/hadoop/
```



### grep

```bash
```

|                      |                        |
| -------------------- | ---------------------- |
| `-i, --ignore-case`  | 忽略大小写             |
| `-r, --recursive`    | 递归地处理目录中的文件 |
| ` -w, --word-regexp` | match only whole words |

### tar

| 参数              | 说明                         |
| ----------------- | ---------------------------- |
| `-C, --directory` | 执行命令前先切换到指定的目录 |
|                   |                              |
|                   |                              |

### usermod

```bash
sudo usermod -aG root janus
newgrp root
```

### 更换字体

```bash
gsettings set org.gnome.desktop.interface font-name '字体名称 [style] 字号'
# 应用程序字体
gsettings get org.gnome.desktop.interface font-name
# 文档字体
gsettings get org.gnome.desktop.interface document-font-name
# 等宽字体（通常用于终端应用）
gsettings get org.gnome.desktop.interface monospace-font-name
# 窗口标题字体
gsettings get org.gnome.desktop.wm.preferences titlebar-font



gsettings set org.gnome.desktop.interface font-name 'Fira Sans Light Italic 11'
gsettings set org.gnome.desktop.interface document-font-name 'Fira Sans Light Italic 11'
gsettings set org.gnome.desktop.interface monospace-font-name 'Fira Sans Light Italic 11'
gsettings set org.gnome.desktop.wm.preferences titlebar-font 'Fira Sans Light Italic 11'
```

### find

```bash
find [搜索路径] [匹配条件] [执行操作]

options:
	-name
	-type <d|f|l>
	-iname
	-exec <command> 对每个文件执行命令
```

### 权限位（数字）

一般为3或4位数字表示权限

- 3位：基本权限 `user group other`
- 4位：特殊权限+基本权限

| 权限 | 字母模式 | 数字模式 |
| ---- | -------- | -------- |
| 读   | r        | 4        |
| 写   | w        | 2        |
| 执行 | x        | 1        |

 

