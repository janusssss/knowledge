### tar

```bash
tar [option] [file]

options:
	-c, --create               create a new archive
	-f, --file=ARCHIVE         use archive file or device ARCHIVE
	-x, --extract, --get       extract files from an archive
	-v, --verbose              verbosely list files processed
```

### telnet

```bash
# 测试指定IP地址或域名的特定端口是否开放和连通
telnet ip:port
```

### cut

```bash
Usage: cut OPTION... [FILE]...

flags:
  -b, --bytes=LIST        以byte为基本单位
  -c, --characters=LIST   以字符串为基本单位
  -f, --fields=LIST       以字段为基本单位，配合-d使用
  -d, --delimiter=DELIM   use DELIM instead of TAB for field delimiter
                          that contains no delimiter character, unless
                          the -s option is specified
  -n                      (ignored)
    --complement          补集
                          or fields
  -s, --only-delimited    do not print lines not containing delimiters
    --output-delimiter=STRING  use STRING as the output delimiter
                          the default is to use the input delimiter
  -z, --zero-terminated   line delimiter is NUL, not newline
    --help        display this help and exit
    --version     output version information and exit
	
LIST:
  N     N'th byte, character or field, counted from 1
  N-    from N'th byte, character or field, to end of line
  N-M   from N'th to M'th (included) byte, character or field
  -M    from first to M'th (included) byte, character or field
  N,M

```

### grep

```bash
Usage: grep [OPTION]... PATTERNS [FILE]...

Pattern selection and interpretation:
  -E, --extended-regexp     PATTERNS are extended regular expressions
  -F, --fixed-strings       PATTERNS are strings
  -G, --basic-regexp        PATTERNS are basic regular expressions
  -P, --perl-regexp         PATTERNS are Perl regular expressions
  -e, --regexp=PATTERNS     use PATTERNS for matching
  -f, --file=FILE           take PATTERNS from FILE
  -i, --ignore-case         ignore case distinctions in patterns and data
      --no-ignore-case      do not ignore case distinctions (default)
  -w, --word-regexp         match only whole words
  -x, --line-regexp         match only whole lines
  -z, --null-data           a data line ends in 0 byte, not newline

Miscellaneous:
  -s, --no-messages         suppress error messages
  -v, --invert-match        select non-matching lines
  -V, --version             display version information and exit
      --help                display this help text and exit

Output control:
  -m, --max-count=NUM       stop after NUM selected lines
  -b, --byte-offset         print the byte offset with output lines
  -n, --line-number         print line number with output lines
      --line-buffered       flush output on every line
  -H, --with-filename       print file name with output lines
  -h, --no-filename         suppress the file name prefix on output
      --label=LABEL         use LABEL as the standard input file name prefix
  -o, --only-matching       show only nonempty parts of lines that match
  -q, --quiet, --silent     suppress all normal output
      --binary-files=TYPE   assume that binary files are TYPE;
                            TYPE is 'binary', 'text', or 'without-match'
  -a, --text                equivalent to --binary-files=text
  -I                        equivalent to --binary-files=without-match
  -d, --directories=ACTION  how to handle directories;
                            ACTION is 'read', 'recurse', or 'skip'
  -D, --devices=ACTION      how to handle devices, FIFOs and sockets;
                            ACTION is 'read' or 'skip'
  -r, --recursive           like --directories=recurse
  -R, --dereference-recursive  likewise, but follow all symlinks
      --include=GLOB        search only files that match GLOB (a file pattern)
      --exclude=GLOB        skip files that match GLOB
      --exclude-from=FILE   skip files that match any file pattern from FILE
      --exclude-dir=GLOB    skip directories that match GLOB
  -L, --files-without-match  print only names of FILEs with no selected lines
  -l, --files-with-matches  print only names of FILEs with selected lines
  -c, --count               print only a count of selected lines per FILE
  -T, --initial-tab         make tabs line up (if needed)
  -Z, --null                print 0 byte after FILE name

Context control:
  -B, --before-context=NUM  print NUM lines of leading context
  -A, --after-context=NUM   print NUM lines of trailing context
  -C, --context=NUM         print NUM lines of output context
  -NUM                      same as --context=NUM
      --group-separator=SEP  print SEP on line between matches with context
      --no-group-separator  do not print separator for matches with context
      --color[=WHEN],
      --colour[=WHEN]       use markers to highlight the matching strings;
                            WHEN is 'always', 'never', or 'auto'
  -U, --binary              do not strip CR characters at EOL (MSDOS/Windows)

When FILE is '-', read standard input.  With no FILE, read '.' if
recursive, '-' otherwise.  With fewer than two FILEs, assume -h.
Exit status is 0 if any line is selected, 1 otherwise;
if any error occurs and -q is not given, the exit status is 2.

Report bugs to: bug-grep@gnu.org
GNU grep home page: <https://www.gnu.org/software/grep/>
General help using GNU software: <https://www.gnu.org/gethelp/>
```

- options :
  - -i : 忽略大小写
  - -n : 显示行号
  - -v : 反向匹配（不显示匹配的行）
  - -w : 匹配完整单词
  - -c : 显示匹配行的数量
  - -l : 显示包含匹配模式的文件名
  - -r  : 递归搜索目录
  - -E : 使用扩展正则表达式
  - -P : 使用 Perl 正则表达式。
  - -C：显示匹配行前后行数
  - **`-q` 或 `--quiet|--silent`**：这个选项告诉 `grep` 不要输出任何匹配的行或信息到标准输出。它的主要用途是通过退出状态码来判断是否找到了匹配项。



### sed

```bash
echo "hello" | sed -e 's/h/m/g'

# / 可以替换成任意字符 g可以换成数字 -e可以省略 ''也可以省略
```


### cpu调节器
```bash
# 查看所有cpu当前调节器
grep . /sys/devices/system/cpu/cpu*/cpufreq/scaling_governor

# 查看所有cpu可用调节器
grep . /sys/devices/system/cpu/cpu*/cpufreq/scaling_available_governors
```


### cpupower
```bash
# 查看所有cpu模式
sudo cpupower -c all frequency-info  | grep "energy performance preference"
# 设置模式 performance powersave
sudo cpupower -c all frequency-set -g performance
```