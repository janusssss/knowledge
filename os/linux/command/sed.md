### 概述

`sed`（**Stream Editor**）是一个强大的**流编辑器**，主要用于对文本进行**查找、替换、删除、插入**等操作。它逐行处理输入（文件或管道数据），并输出修改后的结果。

### 用法

```bash
Usage: sed [OPTION]... {script-only-if-no-other-script} [input-file]...
options:
  -i[SUFFIX], --in-place[=SUFFIX]  edit files in place (makes backup if SUFFIX supplied)
  -e script, --expression=script	add the script to the commands to be executed
	-n, --quiet, --silent	
```

### script

```bash
# [addr]X[options]

X is a single-letter sed command. [addr] is an optional line address. 
If [addr] is specified, the command X will be executed only on the matched lines. 
[addr] can be a single line number, a regular expression, or a range of lines (see sed addresses). 
Additional [options] are used for some sed commands. 
```

| addr          | 作用                       |
| ------------- | -------------------------- |
| `n`           | 第 n 行                    |
| `n,m`         | 第n~m 行                   |
| $             | 最后一行                   |
| `/pattern/`   | 匹配 pattern 的行          |
| `n,/pattern/` | 从第 n 行到匹配 "end" 的行 |

| X            | 作用                          |
| ------------ | ----------------------------- |
| `s/old/new/` | 替换，`///`分隔符可以任何字符 |
| `d`          | 删除                          |
| `p`          | 打印                          |

| options | 作用                                 |
| ------- | ------------------------------------ |
| `g`     | 全局替换（默认只替换每行第一个匹配） |
| `i`     | 忽略大小写                           |
|         |                                      |

