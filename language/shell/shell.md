### 循环体

- `break`：立即退出循环
- `continue`：跳过当前循环体剩余部分，继续下一次循环

#### 数据行遍历

```bash
while IFS= read -r line; do
    echo "Processing: $line"
done < "$data"
```

#### `for`

```sh
for var in item1 item2 item3
do
    command1
    command2
    ...
done
```

#### C风格`for`

```sh
for (( expr1; expr2; expr3 ))
do
    command1
    command2
    ...
done
```

#### `while` 

```sh
while [ condition ]
do
    command1
    command2
    ...
done
```

#### `until` 

```sh
until [ condition ]
do
    command1
    command2
    ...
done
```

#### `select` 

```sh
select var in item1 item2 item3
do
    case $var in
        item1 )
            command1
            ;;
        item2 )
            command2
            ;;
        # etc.
    esac
done
```

### 是否

#### 简单的 `if` 语句

```sh
if [ condition ]
then
    # 如果条件为真，执行这里的命令
fi
```

#### `if-else` 语句

```sh
if [ condition ]
then
    # 如果条件为真，执行这里的命令
else
    # 如果条件为假，执行这里的命令
fi
```

#### `if-elif-else` 语句

```sh
if [ condition1 ]
then
    # 如果 condition1 为真，执行这里的命令
elif [ condition2 ]
then
    # 如果 condition1 为假而 condition2 为真，执行这里的命令
else
    # 如果所有条件都为假，执行这里的命令
fi
```

### 使用双括号的条件表达式（bash）

双括号允许更自然的数学和逻辑运算符，并且可以省略一些引号。

```sh
if (( condition ))
then
    # 执行...
fi

# 或者对于字符串比较和文件测试等，使用双方括号
if [[ condition ]]
then
    # 执行...
fi
```

### 测试命令

可以直接使用命令作为条件表达式。如果命令成功（返回状态码0），则条件为真；否则为假。

```sh
if command
then
    # 如果命令成功，执行这里的命令
fi
```

### 操作符

使用单方括号`[ ]`时，条件表达式的两边必须有空格，例如 `[ "$var" = "value" ]`。而在使用双方括号`[[ ]]`时，虽然不是严格要求，但为了保持一致性，通常也建议保留空格。

### 文件测试

- `-e FILE`：文件存在
- `-f FILE`：文件是普通文件
- `-d DIR`：目录存在
- `-r FILE`：文件可读
- `-w FILE`：文件可写
- `-x FILE`：文件可执行

### 字符串比较

- `str1 = str2`：字符串相等
- `str1 != str2`：字符串不等
- `-z str`：字符串长度为零
- `-n str` 或 `str`：字符串长度非零

### 整数比较

- `int1 -eq int2`：等于
- `int1 -ne int2`：不等于
- `int1 -lt int2`：小于
- `int1 -le int2`：小于等于
- `int1 -gt int2`：大于
- `int1 -ge int2`：大于等于

### 逻辑操作符

- `!`：非
- `-a`：与（在`test`或`[`中）
- `-o`：或（在`test`或`[`中）
- `&&`：与（在`[[`中）
- `||`：或（在`[[`中）

### 语法

### `${parameter:+word}`

```bash
${parameter:+word}
如果 parameter 已经被设置并且它的值不是空的，那么这个表达式将被替换为 word；
否则，整个表达式将被替换为空字符串
```

```bash
var=""
echo "${var:+Hello}"  # 输出为空，因为 var 是空的

var="World"
echo "${var:+Hello}"  # 输出: Hello，因为 var 不是空的
```

### 关键字

local：生命局部变量

### Shell 通配符

| 符号             | 作用             |
| ---------------- | ---------------- |
| `*`              | 匹配任意数量字符 |
| `?`              | 匹配单个字符     |
| `[abc]`或`[a-c]` | 匹配字符集合     |

