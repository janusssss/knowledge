# 概述
xclip 是一个命令行剪贴板工具，用于在终端中访问和操作 X Window 系统的剪贴板。

# 操作
```bash
# 读取系统剪贴板（Ctrl+C/Ctrl+V使用的）
xclip -selection clipboard -o

# 读取主选择（鼠标选中即复制的内容）
xclip -selection primary -o

# 将文件内容复制到剪贴板
xclip -selection clipboard -i < file.txt

# 将命令输出复制到剪贴板
echo "hello" | xclip -selection clipboard

# 将文本直接复制到剪贴板
xclip -selection clipboard <<< "hello world"
```

- 常用选项
	- selection：指定剪贴板类型
		- clipboard：系统剪贴板
		- primary：主选择（鼠标选中）
		- secondary：次选择
    - o：输出（读取剪贴板）
    - i：输入（写入剪贴板）