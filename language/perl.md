### 打印匹配内容

```bash
 perl -nE 'while (/"[^"]+Income[^"]+"/g) { print "$&\n" }' preferences.json
 
 perl -nE 'say for /便签/g' preferences.json

 # 打印1次
 perl -nE 'say $& if /.{50}/' task.txt
```

### 原地编辑文件

```bash
perl -pi -e '正则表达式' 文件名
```

# Perl风格正则表达式
```perl
# 匹配
$str =~ /pattern/flags;

# 替换
$str =~ s/pattern/replacement/flags;

# 字符转换
$str =~ tr/searchlist/replacementlist/flags;
```
