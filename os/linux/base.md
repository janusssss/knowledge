### 结构

- /
  - var : 存放可变数据文件(动态)
    - lib: 存放与特定应用程序相关的状态信息和数据文件
    - log : 存放系统和应用程序的日志文件
  - etc : 存放系统的配置文件(静态)
    - os-release : 存储操作系统的识别信息(key-value)
    - shadow : 存储用户密码和其他与用户帐户相关的安全信息的重要文件
    - fstab : `file system table`包含系统启动时或执行 mount -a 命令时需要挂载的文件系统的静态信息 
    - hosts: 主机名映射到IP地址的本地网络配置文件
  - opt : 存放可选的第三方软件包
  - tmp : 存储临时文件和数据(temporary)
  - dev : 表示设备文件
    - /null : 任何写入 `/dev/null` 的数据都会被丢弃

### 命令

#### 目录&文件操作

### journalctl

```bash
journalctl [OPTIONS...] [MATCHES...]
# 访问和查看由 systemd 管理的系统日志，包含系统启动信息、服务日志、内核消息、应用程序日志等
```

- options : 

  - -S --since=DATE : 显示date之后的日志

  - -U --until=DATE : 显示date之前的日志

  - -u --unit=UNIT : 查看特定服务的日志

  - -f : 实时查看日志

  - xe: 查看当前登录会话的日志

  - -p : 按日志的优先级（日志级别）进行过滤

    ```bash
    0 - emergency（紧急）
    1 - alert（警报）
    2 - critical（严重）
    3 - error（错误）
    4 - warning（警告）
    5 - notice（通知）
    6 - info（信息
    7 - debug（调试）
    ```

    

### awk

```bash

```

```bash
awk [options...][program][file ...]
 
options:
  -F <fs>						fs可以是单个字符、字符串或正则表达式，默认为空格
  -v <var=value>		awk程序定义变量并赋给指定的值，可以在程序开始之前设置多个变量
  -f <path>					指定执行脚本

program: ‘
//{command1; command2} END{}’
  BEGIN			初始化代码块，在对每一行进行处理之前，初始化代码，主要是引用全局变量，设置FS分隔符
  //				匹配代码块，可以是字符串或正则表达式
  {}				命令代码块，包含一条或多条命令
  ;					多条命令使用分号分隔
  END				结尾代码块，在对每一行进行处理之后再执行的代码块，主要是进行最终计算或输出结尾摘要信息

特殊符
  -F '[:#/]'   			定义三个分隔符
  $0           			表示整个当前行
  $1/$n        			每行第一个字段/第n个字段
  NF					 			字段数量变量
  NR         	 			每行的记录号，多文件记录递增
  FNR        	 			与NR类似，不过多文件记录不递增，每个文件都从1开始
  \t           			制表符
  \n           			换行符
  FS           			BEGIN时定义分隔符
  RS       		 			输入的记录分隔符， 默认为换行符(即文本是按一行一行输入)
  ~            			匹配，与==相比不是精确比较
  !~           			不匹配，不精确比较
  ==         	 			等于，必须全部相等，精确比较
  !=           			不等于，精确比较
  &&							  逻辑与
  ||     			      逻辑或
  + 		            匹配时表示1个或1个以上
  /[0-9][0-9]+/   	两个或两个以上数字
  /[0-9][0-9]*/    	一个或一个以上数字
  FILENAME 					文件名
  OFS      					输出字段分隔符， 默认也是空格
  ORS       	 			输出的记录分隔符，默认为换行符

commands:
	func: 
		print 			打印
		system()		执行命令行命令

rule：
	1.当两个字符串或字符串与变量之间没有运算符时，默认会进行字符串拼接操作
```

- 选项: 可选参数，如 -F 指定字段分隔符
- 模式: 指定何时执行动作的条件
- 动作: 当模式匹配时要执行的操作
  - print: 用于打印整行或指定字段
  - $[n]: 字段变量
- 文件名: 需要处理的文件名

### find

```bash
find [path...] [expression]
-print	将当前找到的文件路径打印到标准输出。
-quit	立即终止 find 命令的执行。
```

- expression

  - -name : 指定名字文件，可以使用正则

  - -type : 指定类型

    ```bash
    d : 目录
    f :	文件
    l : 符号连接
    ```

  - size : 按大小寻找

    ```bash
    +100M : 大于100M
    -100M : 小于100M
    ```

  - -exec : 对找到的文件执行某个命令

### awk

```bash
Usage: mawk [Options] [Program] [file ...]

Program:
    The -f option value is the name of a file containing program text.
    If no -f option is given, a "--" ends option processing; the following
    parameters are the program text.

Options:
    -f program-file  Program  text is read from file instead of from the
                     command-line.  Multiple -f options are accepted.
    -F value         sets the field separator, FS, to value.
    -v var=value     assigns value to program variable var.
    --               unambiguous end of options.

    Implementation-specific options are prefixed with "-W".  They can be
    abbreviated:

    -W version       show version information and exit.
    -W compat        pre-POSIX 2001.
    -W dump          show assembler-like listing of program and exit.
    -W help          show this message and exit.
    -W interactive   set unbuffered output, line-buffered input.
    -W exec file     use file as program as well as last option.
    -W posix         stricter POSIX checking.
    -W random=number set initial random seed.
    -W sprintf=number adjust size of sprintf buffer.
    -W usage         show this message and exit.
```



### 基本正则

- .（点）: 匹配任何单个字符（除了换行符），a.c 可以匹配 "abc"、"axc" 
- `*`（星号）: 匹配前面的字符或子表达式零次或多次,，a*b 匹配 "b"、"ab"、"aab"、"aaab" 等
- ^（脱字符）: 匹配行的开头，^abc 匹配以 "abc" 开头的行
- $（美元符号）: 匹配行的结尾，abc$ 匹配以 "abc" 结尾的行
- []（方括号）: 匹配方括号内的任何一个字符，可以使用范围， [a-z] 匹配任何小写字母，[abc] 匹配 "a", "b" 或 "c"
- [^]（否定方括号）: 匹配不在方括号内的任何字符，`[^0-9] `匹配任何非数字字符
- \（反斜杠）: 用于转义特殊字符，使其失去特殊含义，`\.` 匹配一个实际的句点，而不是“任意字符”。

### cp

```
cp [OPTIONS...] source destination
```

- options :
  - -r : 递归
  - -f : 强制
  - -v : 输出过程
  - -u : 更新

### wc

```bash
wc [options] [flie...]
# word count 用于统计文件的字数、行数、字符数和字节数
```

- options
  - `-l` :  统计行数（Line count）
  - `-w : 统计单词数（Word count）
  - `-c : 统计字符数（Character count）
  - `-m` : 统计字符数（按字符计）
  - `-L` : 统计文件中最长一行的长度

### dirname

```bash
dirname [option] name...
-z, --zero		输出不换行
--help				帮助
--version			版本
```

```bash
Usage: dirname [OPTION] NAME...
Output each NAME with its last non-slash component and trailing slashes
removed; if NAME contains no /'s, output '.' (meaning the current directory).

  -z, --zero     end each output line with NUL, not newline
      --help        display this help and exit
      --version     output version information and exit

Examples:
  dirname /usr/bin/          -> "/usr"
  dirname dir1/str dir2/str  -> "dir1" followed by "dir2"
  dirname stdio.h            -> "."
  
GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
Full documentation <https://www.gnu.org/software/coreutils/dirname>
or available locally via: info '(coreutils) dirname invocation'
```

### uniq

```bash
[root@gz-txjs-control-55e243e31e31 janus]# uniq --help
Usage: uniq [OPTION]... [INPUT [OUTPUT]]
Filter adjacent matching lines from INPUT (or standard input),
writing to OUTPUT (or standard output).

With no options, matching lines are merged to the first occurrence.

Mandatory arguments to long options are mandatory for short options too.
  -c, --count           prefix lines by the number of occurrences
  -d, --repeated        only print duplicate lines, one for each group
  -D, --all-repeated[=METHOD]  print all duplicate lines
                          groups can be delimited with an empty line
                          METHOD={none(default),prepend,separate}
  -f, --skip-fields=N   avoid comparing the first N fields
      --group[=METHOD]  show all items, separating groups with an empty line
                          METHOD={separate(default),prepend,append,both}
  -i, --ignore-case     ignore differences in case when comparing
  -s, --skip-chars=N    avoid comparing the first N characters
  -u, --unique          only print unique lines
  -z, --zero-terminated  end lines with 0 byte, not newline
  -w, --check-chars=N   compare no more than N characters in lines
      --help     display this help and exit
      --version  output version information and exit

A field is a run of blanks (usually spaces and/or TABs), then non-blank
characters.  Fields are skipped before chars.

Note: 'uniq' does not detect repeated lines unless they are adjacent.
You may want to sort the input first, or use 'sort -u' without 'uniq'.
Also, comparisons honor the rules specified by 'LC_COLLATE'.

GNU coreutils online help: <http://www.gnu.org/software/coreutils/>
For complete documentation, run: info coreutils 'uniq invocation'

```

### sort

```bash
```

### ln

```bash
ls [option...] <target> <link>
-f, --force：如果链接文件已存在，则将其覆盖。
-i, --interactive：在覆盖前询问。
-n, --no-dereference：如果目标是一个符号链接，则不对该符号链接进行解引用。
```

### nload
nload是一个实时监控网络流量和带宽使用的命令行工具。它提供了直观的界面，用于显示通过每个网络接口的数据进出情况，非常适合用来监测网络连接的使用状况。
```bash
nload version 0.7.4
Copyright (C) 2001 - 2012 by Roland Riegel <feedback@roland-riegel.de>
nload comes with ABSOLUTELY NO WARRANTY. This is free software, and you are
welcome to redistribute it under certain conditions. For more details see the
GNU General Public License Version 2 (http://www.gnu.org/copyleft/gpl.html).

Command line syntax:
nload [options] [devices]
nload --help|-h

Options:
-a period       Sets the length in seconds of the time window for average
                calculation.
                Default is 300.
-i max_scaling  Specifies the 100% mark in kBit/s of the graph indicating the
                incoming bandwidth usage. Ignored if max_scaling is 0 or the
                switch -m is given.
                Default is 10240.
-m              Show multiple devices at a time; no traffic graphs.
-o max_scaling  Same as -i but for the graph indicating the outgoing bandwidth
                usage.
                Default is 10240.
-t interval     Determines the refresh interval of the display in milliseconds.
                Default is 500.
-u h|b|k|m|g    Sets the type of unit used for the display of traffic numbers.
   H|B|K|M|G    h: auto, b: Bit/s, k: kBit/s, m: MBit/s etc.
                H: auto, B: Byte/s, K: kByte/s, M: MByte/s etc.
                Default is h.
-U h|b|k|m|g    Same as -u, but for a total amount of data (without "/s").
   H|B|K|M|G    Default is H.
devices         Network devices to use.
                Default is to use all auto-detected devices.
--help
-h              Print this help.

example: nload -t 200 -i 1024 -o 128 -U M

The options above can also be changed at run time by pressing the 'F2' key.

```


## 系统操作

### systemctl

```bash
systemctl [OPTIONS...] COMMAND ...
# 管理和控制 systemd 系统和服务管理器的命令行工具
```

- command

  - status <server-name> : 查看特定服务的当前状态
  - start <server-name> : 启动服务
  - restart <server-name> : 重启服务
  - reload <server-name> : 重新加载服务配置
  - stop <server-name> : 停止服务
  - enable <server-name> : 系统启动时自动启动
  - disable <server-name> : 禁用服务
  - daemon-reload: 重新加载 systemd 系统和服务管理器的配置文件

- options : 

  - --type : 指定要列出的单位类型（unit type）。常见的单位类型包括 service, socket, target, mount

  - --state : 过滤和显示具有特定状态的服务单元

    ```bash
    active: 单元正在运行或已完成
    inactive: 单元未运行且未激活
    reloading: 单元正在重新加载其配置
    failed: 单元失败了，可能是因为启动过程中出现错误
    activating: 单元正在启动中
    deactivating: 单元正在停止中
    ```

  - --force : 强制执行


```bash
systemctl list-units --type=service --all # 查看全部服务
systemctl list-units --type=service --state=running # 查看运行服务，包括avtive
```



### lsblk

```bash
lsblk
# list block devices，用于列出系统中所有可用的块设备信
```

- 

## install

```bash
install [选项] 源文件 目标文件
install [选项] 源文件... 目标目录
install -d 目标目录...
# 安装或复制文件及目录，并设置适当的权限和所有权
```

- 选项

  - -d 或 --directory: 创建目录。如果目录已存在，则不会报错

  - -m 或 --mode: 设置文件或目录的权限模式。模式可以是八进制数字或符号表示法

## curl

```bash
 curl -fsSL https://ollama.com/install.sh | sh
```

```bash
Usage: curl [options...] <url>
     --abstract-unix-socket <path> Connect via abstract Unix domain socket
     --alt-svc <file name> Enable alt-svc with this cache file
     --anyauth     Pick any authentication method
 -a, --append      Append to target file when uploading
     --aws-sigv4 <provider1[:provider2[:region[:service]]]> Use AWS V4 signature authentication
     --basic       Use HTTP Basic Authentication
     --ca-native   Use CA certificates from the native OS
     --cacert <file> CA certificate to verify peer against
     --capath <dir> CA directory to verify peer against
 -E, --cert <certificate[:password]> Client certificate file and password
     --cert-status Verify the status of the server cert via OCSP-staple
     --cert-type <type> Certificate type (DER/PEM/ENG/P12)
     --ciphers <list of ciphers> SSL ciphers to use
     --compressed  Request compressed response
     --compressed-ssh Enable SSH compression
 -K, --config <file> Read config from a file
     --connect-timeout <fractional seconds> Maximum time allowed for connection
     --connect-to <HOST1:PORT1:HOST2:PORT2> Connect to host
 -C, --continue-at <offset> Resumed transfer offset
 -b, --cookie <data|filename> Send cookies from string/file
 -c, --cookie-jar <filename> Write cookies to <filename> after operation
     --create-dirs Create necessary local directory hierarchy
     --create-file-mode <mode> File mode for created files
     --crlf        Convert LF to CRLF in upload
     --crlfile <file> Use this CRL list
     --curves <algorithm list> (EC) TLS key exchange algorithm(s) to request
 -d, --data <data> HTTP POST data
     --data-ascii <data> HTTP POST ASCII data
     --data-binary <data> HTTP POST binary data
     --data-raw <data> HTTP POST data, '@' allowed
     --data-urlencode <data> HTTP POST data URL encoded
     --delegation <LEVEL> GSS-API delegation permission
     --digest      Use HTTP Digest Authentication
 -q, --disable     Disable .curlrc
     --disable-eprt Inhibit using EPRT or LPRT
     --disable-epsv Inhibit using EPSV
     --disallow-username-in-url Disallow username in URL
     --dns-interface <interface> Interface to use for DNS requests
     --dns-ipv4-addr <address> IPv4 address to use for DNS requests
     --dns-ipv6-addr <address> IPv6 address to use for DNS requests
     --dns-servers <addresses> DNS server addrs to use
     --doh-cert-status Verify the status of the DoH server cert via OCSP-staple
     --doh-insecure Allow insecure DoH server connections
     --doh-url <URL> Resolve host names over DoH
 -D, --dump-header <filename> Write the received headers to <filename>
     --egd-file <file> EGD socket path for random data
     --engine <name> Crypto engine to use
     --etag-compare <file> Pass an ETag from a file as a custom header
     --etag-save <file> Parse ETag from a request and save it to a file
     --expect100-timeout <seconds> How long to wait for 100-continue
 -f, --fail        Fail fast with no output on HTTP errors
     --fail-early  Fail on first transfer error, do not continue
     --fail-with-body Fail on HTTP errors but save the body
     --false-start Enable TLS False Start
 -F, --form <name=content> Specify multipart MIME data
     --form-escape Escape multipart form field/file names using backslash
     --form-string <name=string> Specify multipart MIME data
     --ftp-account <data> Account data string
     --ftp-alternative-to-user <command> String to replace USER [name]
     --ftp-create-dirs Create the remote dirs if not present
     --ftp-method <method> Control CWD usage
     --ftp-pasv    Use PASV/EPSV instead of PORT
 -P, --ftp-port <address> Use PORT instead of PASV
     --ftp-pret    Send PRET before PASV
     --ftp-skip-pasv-ip Skip the IP address for PASV
     --ftp-ssl-ccc Send CCC after authenticating
     --ftp-ssl-ccc-mode <active/passive> Set CCC mode
     --ftp-ssl-control Require SSL/TLS for FTP login, clear for transfer
 -G, --get         Put the post data in the URL and use GET
 -g, --globoff     Disable URL sequences and ranges using {} and []
     --happy-eyeballs-timeout-ms <milliseconds> Time for IPv6 before trying IPv4
     --haproxy-clientip Sets client IP in HAProxy PROXY protocol v1 header
     --haproxy-protocol Send HAProxy PROXY protocol v1 header
 -I, --head        Show document info only
 -H, --header <header/@file> Pass custom header(s) to server
 -h, --help <category> Get help for commands
     --hostpubmd5 <md5> Acceptable MD5 hash of the host public key
     --hostpubsha256 <sha256> Acceptable SHA256 hash of the host public key
     --hsts <file name> Enable HSTS with this cache file
     --http0.9     Allow HTTP 0.9 responses
 -0, --http1.0     Use HTTP 1.0
     --http1.1     Use HTTP 1.1
     --http2       Use HTTP/2
     --http2-prior-knowledge Use HTTP 2 without HTTP/1.1 Upgrade
     --http3       Use HTTP v3
     --http3-only  Use HTTP v3 only
     --ignore-content-length Ignore the size of the remote resource
 -i, --include     Include protocol response headers in the output
 -k, --insecure    Allow insecure server connections
     --interface <name> Use network INTERFACE (or address)
     --ipfs-gateway <URL> Gateway for IPFS
 -4, --ipv4        Resolve names to IPv4 addresses
 -6, --ipv6        Resolve names to IPv6 addresses
     --json <data> HTTP POST JSON
 -j, --junk-session-cookies Ignore session cookies read from file
     --keepalive-time <seconds> Interval time for keepalive probes
     --key <key>   Private key file name
     --key-type <type> Private key file type (DER/PEM/ENG)
     --krb <level> Enable Kerberos with security <level>
     --libcurl <file> Dump libcurl equivalent code of this command line
     --limit-rate <speed> Limit transfer speed to RATE
 -l, --list-only   List only mode
     --local-port <num/range> Force use of RANGE for local port numbers
 -L, --location    Follow redirects
     --location-trusted Like --location, and send auth to other hosts
     --login-options <options> Server login options
     --mail-auth <address> Originator address of the original email
     --mail-from <address> Mail from this address
     --mail-rcpt <address> Mail to this address
     --mail-rcpt-allowfails Allow RCPT TO command to fail for some recipients
 -M, --manual      Display the full manual
     --max-filesize <bytes> Maximum file size to download
     --max-redirs <num> Maximum number of redirects allowed
 -m, --max-time <fractional seconds> Maximum time allowed for transfer
     --metalink    Process given URLs as metalink XML file
     --negotiate   Use HTTP Negotiate (SPNEGO) authentication
 -n, --netrc       Must read .netrc for user name and password
     --netrc-file <filename> Specify FILE for netrc
     --netrc-optional Use either .netrc or URL
 -:, --next        Make next URL use its separate set of options
     --no-alpn     Disable the ALPN TLS extension
 -N, --no-buffer   Disable buffering of the output stream
     --no-clobber  Do not overwrite files that already exist
     --no-keepalive Disable TCP keepalive on the connection
     --no-npn      Disable the NPN TLS extension
     --no-progress-meter Do not show the progress meter
     --no-sessionid Disable SSL session-ID reusing
     --noproxy <no-proxy-list> List of hosts which do not use proxy
     --ntlm        Use HTTP NTLM authentication
     --ntlm-wb     Use HTTP NTLM authentication with winbind
     --oauth2-bearer <token> OAuth 2 Bearer Token
 -o, --output <file> Write to file instead of stdout
     --output-dir <dir> Directory to save files in
 -Z, --parallel    Perform transfers in parallel
     --parallel-immediate Do not wait for multiplexing (with --parallel)
     --parallel-max <num> Maximum concurrency for parallel transfers
     --pass <phrase> Pass phrase for the private key
     --path-as-is  Do not squash .. sequences in URL path
     --pinnedpubkey <hashes> FILE/HASHES Public key to verify peer against
     --post301     Do not switch to GET after following a 301
     --post302     Do not switch to GET after following a 302
     --post303     Do not switch to GET after following a 303
     --preproxy [protocol://]host[:port] Use this proxy first
 -#, --progress-bar Display transfer progress as a bar
     --proto <protocols> Enable/disable PROTOCOLS
     --proto-default <protocol> Use PROTOCOL for any URL missing a scheme
     --proto-redir <protocols> Enable/disable PROTOCOLS on redirect
 -x, --proxy [protocol://]host[:port] Use this proxy
     --proxy-anyauth Pick any proxy authentication method
     --proxy-basic Use Basic authentication on the proxy
     --proxy-ca-native Use CA certificates from the native OS for proxy
     --proxy-cacert <file> CA certificate to verify peer against for proxy
     --proxy-capath <dir> CA directory to verify peer against for proxy
     --proxy-cert <cert[:passwd]> Set client certificate for proxy
     --proxy-cert-type <type> Client certificate type for HTTPS proxy
     --proxy-ciphers <list> SSL ciphers to use for proxy
     --proxy-crlfile <file> Set a CRL list for proxy
     --proxy-digest Use Digest authentication on the proxy
     --proxy-header <header/@file> Pass custom header(s) to proxy
     --proxy-http2 Use HTTP/2 with HTTPS proxy
     --proxy-insecure Do HTTPS proxy connections without verifying the proxy
     --proxy-key <key> Private key for HTTPS proxy
     --proxy-key-type <type> Private key file type for proxy
     --proxy-negotiate Use HTTP Negotiate (SPNEGO) authentication on the proxy
     --proxy-ntlm  Use NTLM authentication on the proxy
     --proxy-pass <phrase> Pass phrase for the private key for HTTPS proxy
     --proxy-pinnedpubkey <hashes> FILE/HASHES public key to verify proxy with
     --proxy-service-name <name> SPNEGO proxy service name
     --proxy-ssl-allow-beast Allow security flaw for interop for HTTPS proxy
     --proxy-ssl-auto-client-cert Use auto client certificate for proxy (Schannel)
     --proxy-tls13-ciphers <ciphersuite list> TLS 1.3 proxy cipher suites
     --proxy-tlsauthtype <type> TLS authentication type for HTTPS proxy
     --proxy-tlspassword <string> TLS password for HTTPS proxy
     --proxy-tlsuser <name> TLS username for HTTPS proxy
     --proxy-tlsv1 Use TLSv1 for HTTPS proxy
 -U, --proxy-user <user:password> Proxy user and password
     --proxy1.0 <host[:port]> Use HTTP/1.0 proxy on given port
 -p, --proxytunnel Operate through an HTTP proxy tunnel (using CONNECT)
     --pubkey <key> SSH Public key file name
 -Q, --quote <command> Send command(s) to server before transfer
     --random-file <file> File for reading random data from
 -r, --range <range> Retrieve only the bytes within RANGE
     --rate <max request rate> Request rate for serial transfers
     --raw         Do HTTP "raw"; no transfer decoding
 -e, --referer <URL> Referrer URL
 -J, --remote-header-name Use the header-provided filename
 -O, --remote-name Write output to a file named as the remote file
     --remote-name-all Use the remote file name for all URLs
 -R, --remote-time Set the remote file's time on the local output
     --remove-on-error Remove output file on errors
 -X, --request <method> Specify request method to use
     --request-target <path> Specify the target for this request
     --resolve <[+]host:port:addr[,addr]...> Resolve the host+port to this address
     --retry <num> Retry request if transient problems occur
     --retry-all-errors Retry all errors (use with --retry)
     --retry-connrefused Retry on connection refused (use with --retry)
     --retry-delay <seconds> Wait time between retries
     --retry-max-time <seconds> Retry only within this period
     --sasl-authzid <identity> Identity for SASL PLAIN authentication
     --sasl-ir     Enable initial response in SASL authentication
     --service-name <name> SPNEGO service name
 -S, --show-error  Show error even when -s is used
 -s, --silent      Silent mode
     --socks4 <host[:port]> SOCKS4 proxy on given host + port
     --socks4a <host[:port]> SOCKS4a proxy on given host + port
     --socks5 <host[:port]> SOCKS5 proxy on given host + port
     --socks5-basic Enable username/password auth for SOCKS5 proxies
     --socks5-gssapi Enable GSS-API auth for SOCKS5 proxies
     --socks5-gssapi-nec Compatibility with NEC SOCKS5 server
     --socks5-gssapi-service <name> SOCKS5 proxy service name for GSS-API
     --socks5-hostname <host[:port]> SOCKS5 proxy, pass host name to proxy
 -Y, --speed-limit <speed> Stop transfers slower than this
 -y, --speed-time <seconds> Trigger 'speed-limit' abort after this time
     --ssl         Try SSL/TLS
     --ssl-allow-beast Allow security flaw to improve interop
     --ssl-auto-client-cert Use auto client certificate (Schannel)
     --ssl-no-revoke Disable cert revocation checks (Schannel)
     --ssl-reqd    Require SSL/TLS
     --ssl-revoke-best-effort Ignore missing/offline cert CRL dist points (Schannel)
 -2, --sslv2       Use SSLv2
 -3, --sslv3       Use SSLv3
     --stderr <file> Where to redirect stderr
     --styled-output Enable styled output for HTTP headers
     --suppress-connect-headers Suppress proxy CONNECT response headers
     --tcp-fastopen Use TCP Fast Open
     --tcp-nodelay Use the TCP_NODELAY option
 -t, --telnet-option <opt=val> Set telnet option
     --tftp-blksize <value> Set TFTP BLKSIZE option
     --tftp-no-options Do not send any TFTP options
 -z, --time-cond <time> Transfer based on a time condition
     --tls-max <VERSION> Set maximum allowed TLS version
     --tls13-ciphers <ciphersuite list> TLS 1.3 cipher suites to use
     --tlsauthtype <type> TLS authentication type
     --tlspassword <string> TLS password
     --tlsuser <name> TLS user name
 -1, --tlsv1       Use TLSv1.0 or greater
     --tlsv1.0     Use TLSv1.0 or greater
     --tlsv1.1     Use TLSv1.1 or greater
     --tlsv1.2     Use TLSv1.2 or greater
     --tlsv1.3     Use TLSv1.3 or greater
     --tr-encoding Request compressed transfer encoding
     --trace <file> Write a debug trace to FILE
     --trace-ascii <file> Like --trace, but without hex output
     --trace-config <string> Details to log in trace/verbose output
     --trace-ids   Add transfer and connection identifiers to trace/verbose output
     --trace-time  Add time stamps to trace/verbose output
     --unix-socket <path> Connect through this Unix domain socket
 -T, --upload-file <file> Transfer local FILE to destination
     --url <url>   URL to work with
     --url-query <data> Add a URL query part
 -B, --use-ascii   Use ASCII/text transfer
 -u, --user <user:password> Server user and password
 -A, --user-agent <name> Send User-Agent <name> to server
     --variable <[%]name=text/@file> Set variable
 -v, --verbose     Make the operation more talkative
 -V, --version     Show version number and quit
 -w, --write-out <format> Use output FORMAT after completion
     --xattr       Store metadata in extended file attributes
```

- 选项
  - -f: 禁止返回 HTTP 错误页面，如果请求失败，`curl` 会返回一个错误而不是显示错误页面
  - -s: 静默模式，不显示进度条或其他输出
  - -S: 即使使用 `-s` 选项，仍然在发生错误时显示错误信息
  - -L: 跟随重定向。如果服务器返回一个重定向响应，`curl` 会自动跟随重定向
  - -o: 指定下载的文件保存的路径和文件名

## chomd

```bash
chmod [选项] 权限模式 文件或目录
# 更改文件或目录权限
```

- 选项
  - -R 或 --recursive: 递归地更改目录及其内容的权限
  - -v 或 --verbose: 显示详细的输出信息
- 权限模式
  - 符号表示法
  - 八进制表示法

## tee

```bash
tee [选项] 文件...
# 将命令的输出同时显示在终端上并保存到文件中
```

- 选项
  - -a: 追加到文件末尾而不是覆盖文件

## sudo

```bash
# sudo - execute a command as another user
sudo [optins]

Options:
  -i, --login                   run login shell as the target user; a command may also be specified
```

## source

​	在当前shell环境中执行一个指定的脚本文件（和 . 同效）。

## journalctl

​	journalctl 是 systemd 系统和服务管理器的一部分，用于查询和显示系统日志。systemd 日志系统（也称为 journald）收集来自内核、系统服务和用户进程的日志信息，并将其存储在一个二进制日志文件中。journalctl 命令提供了多种选项来过滤和格式化这些日志信息，使其更易于阅读和分析。

```bash
journalctl [OPTIONS...] [MATCHES...]

Query the journal.

Flags:
     --system              Show the system journal
     --user                Show the user journal for the current user
  -M --machine=CONTAINER   Operate on local container
  -S --since=DATE          Show entries not older than the specified date
  -U --until=DATE          Show entries not newer than the specified date
  -c --cursor=CURSOR       Show entries starting at the specified cursor
     --after-cursor=CURSOR Show entries after the specified cursor
     --show-cursor         Print the cursor after all the entries
  -b --boot[=ID]           Show current boot or the specified boot
     --list-boots          Show terse information about recorded boots
  -k --dmesg               Show kernel message log from the current boot
  -u --unit=UNIT           Show logs from the specified unit
  -t --identifier=STRING   Show entries with the specified syslog identifier
  -p --priority=RANGE      Show entries with the specified priority
  -e --pager-end           Immediately jump to the end in the pager
  -f --follow              Follow the journal
  -n --lines[=INTEGER]     Number of journal entries to show
     --no-tail             Show all lines, even in follow mode
  -r --reverse             Show the newest entries first
  -o --output=STRING       Change journal output mode (short, short-iso,
                                   short-precise, short-monotonic, verbose,
                                   export, json, json-pretty, json-sse, cat)
     --utc                 Express time in Coordinated Universal Time (UTC)
  -x --catalog             Add message explanations where available
     --no-full             Ellipsize fields
  -a --all                 Show all fields, including long and unprintable
  -q --quiet               Do not show privilege warning
     --no-pager            Do not pipe output into a pager
  -m --merge               Show entries from all available journals
  -D --directory=PATH      Show journal files from directory
     --file=PATH           Show journal file
     --root=ROOT           Operate on catalog files underneath the root ROOT
     --interval=TIME       Time interval for changing the FSS sealing key
     --verify-key=KEY      Specify FSS verification key
     --force               Override of the FSS key pair with --setup-keys

Commands:
  -h --help                Show this help text
     --version             Show package version
  -F --field=FIELD         List all values that a specified field takes
     --new-id128           Generate a new 128-bit ID
     --disk-usage          Show total disk usage of all journal files
     --vacuum-size=BYTES   Reduce disk usage below specified size
     --vacuum-time=TIME    Remove journal files older than specified date
     --flush               Flush all journal data from /run into /var
     --header              Show journal header information

```

## tree

```bash
tree --help
usage: tree [-acdfghilnpqrstuvxACDFJQNSUX] [-L level [-R]] [-H  baseHREF]
        [-T title] [-o filename] [-P pattern] [-I pattern] [--gitignore]
        [--gitfile[=]file] [--matchdirs] [--metafirst] [--ignore-case]
        [--nolinks] [--hintro[=]file] [--houtro[=]file] [--inodes] [--device]
        [--sort[=]<name>] [--dirsfirst] [--filesfirst] [--filelimit #] [--si]
        [--du] [--prune] [--charset[=]X] [--timefmt[=]format] [--fromfile]
        [--fromtabfile] [--fflinks] [--info] [--infofile[=]file] [--noreport]
        [--version] [--help] [--] [directory ...]
  ------- Listing options -------
  -a            All files are listed.
  -d            List directories only.
  -l            Follow symbolic links like directories.
  -f            Print the full path prefix for each file.
  -x            Stay on current filesystem only.
  -L level      Descend only level directories deep.
  -R            Rerun tree when max dir level reached.
  -P pattern    List only those files that match the pattern given.
  -I pattern    Do not list files that match the given pattern.
  --gitignore   Filter by using .gitignore files.
  --gitfile X   Explicitly read gitignore file.
  --ignore-case Ignore case when pattern matching.
  --matchdirs   Include directory names in -P pattern matching.
  --metafirst   Print meta-data at the beginning of each line.
  --prune       Prune empty directories from the output.
  --info        Print information about files found in .info files.
  --infofile X  Explicitly read info file.
  --noreport    Turn off file/directory count at end of tree listing.
  --charset X   Use charset X for terminal/HTML and indentation line output.
  --filelimit # Do not descend dirs with more than # files in them.
  -o filename   Output to file instead of stdout.
  ------- File options -------
  -q            Print non-printable characters as '?'.
  -N            Print non-printable characters as is.
  -Q            Quote filenames with double quotes.
  -p            Print the protections for each file.
  -u            Displays file owner or UID number.
  -g            Displays file group owner or GID number.
  -s            Print the size in bytes of each file.
  -h            Print the size in a more human readable way.
  --si          Like -h, but use in SI units (powers of 1000).
  --du          Compute size of directories by their contents.
  -D            Print the date of last modification or (-c) status change.
  --timefmt <f> Print and format time according to the format <f>.
  -F            Appends '/', '=', '*', '@', '|' or '>' as per ls -F.
  --inodes      Print inode number of each file.
  --device      Print device ID number to which each file belongs.
  ------- Sorting options -------
  -v            Sort files alphanumerically by version.
  -t            Sort files by last modification time.
  -c            Sort files by last status change time.
  -U            Leave files unsorted.
  -r            Reverse the order of the sort.
  --dirsfirst   List directories before files (-U disables).
  --filesfirst  List files before directories (-U disables).
  --sort X      Select sort: name,version,size,mtime,ctime.
  ------- Graphics options -------
  -i            Don't print indentation lines.
  -A            Print ANSI lines graphic indentation lines.
  -S            Print with CP437 (console) graphics indentation lines.
  -n            Turn colorization off always (-C overrides).
  -C            Turn colorization on always.
  ------- XML/HTML/JSON options -------
  -X            Prints out an XML representation of the tree.
  -J            Prints out an JSON representation of the tree.
  -H baseHREF   Prints out HTML format with baseHREF as top directory.
  -T string     Replace the default HTML title and H1 header with string.
  --nolinks     Turn off hyperlinks in HTML output.
  --hintro X    Use file X as the HTML intro.
  --houtro X    Use file X as the HTML outro.
  ------- Input options -------
  --fromfile    Reads paths from files (.=stdin)
  --fromtabfile Reads trees from tab indented files (.=stdin)
  --fflinks     Process link information when using --fromfile.
  ------- Miscellaneous options -------
  --version     Print version and exit.
  --help        Print usage and this help message and exit.
  --            Options processing terminator.
```

```bash
用法: tree [-acdfghilnpqrstuvxACDFJQNSUX] [-L 层级 [-R]] [-H  baseHREF]
        [-T 标题] [-o 文件名] [-P 模式] [-I 模式] [--gitignore]
        [--gitfile[=]文件] [--matchdirs] [--metafirst] [--ignore-case]
        [--nolinks] [--hintro[=]文件] [--houtro[=]文件] [--inodes] [--device]
        [--sort[=]<名称>] [--dirsfirst] [--filesfirst] [--filelimit #] [--si]
        [--du] [--prune] [--charset[=]X] [--timefmt[=]格式] [--fromfile]
        [--fromtabfile] [--fflinks] [--info] [--infofile[=]文件] [--noreport]
        [--version] [--help] [--] [目录 ...]
  ------- 列表选项 -------
  -a            列出所有文件。
  -d            仅列出目录。
  -l            将符号链接当作目录来跟随。
  -f            打印每个文件的完整路径前缀。
  -x            仅停留在当前文件系统上。
  -L 层级       仅递归到指定层级的目录深度。
  -R            当达到最大目录层级时重新运行 tree。
  -P 模式       仅列出符合给定模式的文件。
  -I 模式       不列出符合给定模式的文件。
  --gitignore   使用 .gitignore 文件进行过滤。
  --gitfile X   显式读取 gitignore 文件。
  --ignore-case 忽略模式匹配时的大小写。
  --matchdirs   在 -P 模式匹配中包含目录名称。
  --metafirst   在每行的开头打印元数据。
  --prune       从输出中剪除空目录。
  --info        打印在 .info 文件中找到的信息。
  --infofile X  显式读取 info 文件。
  --noreport    关闭树状列表末尾的文件/目录计数。
  --charset X   使用字符集 X 进行终端/HTML 和缩进线输出。
  --filelimit # 如果目录中的文件超过 #，则不深入其中。
  -o 文件名     输出到文件而不是标准输出。
  ------- 文件选项 -------
  -q            将不可打印的字符打印为 '?'。
  -N            将不可打印的字符按原样打印。
  -Q            用双引号引用文件名。
  -p            打印每个文件的权限。
  -u            显示文件所有者或用户ID编号。
  -g            显示文件组所有者或组ID编号。
  -s            打印每个文件的大小（以字节为单位）。
  -h            以更易读的方式打印文件大小。
  --si          类似于 -h，但使用 SI 单位（1000 的幂）。
  --du          通过其内容计算目录的大小。
  -D            打印最后修改时间或 (-c) 状态更改时间。
  --timefmt <f> 按照格式 <f> 打印和格式化时间。
  -F            根据 ls -F 添加 '/', '=', '*', '@', '|' 或 '>'。
  --inodes      打印每个文件的 inode 编号。
  --device      打印每个文件所属的设备 ID 编号。
  ------- 排序选项 -------
  -v            按版本对文件进行字母数字排序。
  -t            按最后修改时间对文件进行排序。
  -c            按最后状态更改时间对文件进行排序。
  -U            不对文件进行排序。
  -r            反转排序顺序。
  --dirsfirst   在文件之前列出目录 (-U 禁用此功能)。
  --filesfirst  在目录之前列出文件 (-U 禁用此功能)。
  --sort X      选择排序方式：名称、版本、大小、mtime、ctime。
  ------- 图形选项 -------
  -i            不打印缩进线。
  -A            使用 ANSI 线条图形进行缩进线打印。
  -S            使用 CP437（控制台）图形进行缩进线打印。
  -n            总是关闭颜色化（-C 覆盖此选项）。
  -C            总是开启颜色化。
  ------- XML/HTML/JSON 选项 -------
  -X            以 XML 形式打印树状结构。
  -J            以 JSON 形式打印树状结构。
  -H baseHREF   以 HTML 格式打印，并将 baseHREF 作为顶级目录。
  -T 字符串     用字符串替换默认的 HTML 标题和 H1 头部。
  --nolinks     在 HTML 输出中关闭超链接。
  --hintro X    使用文件 X 作为 HTML 引入部分。
  --houtro X    使用文件 X 作为 HTML 结束部分。
  ------- 输入选项 -------
  --fromfile    从文件中读取路径 (.=标准输入)
  --fromtabfile 从制表符缩进的文件中读取树 (.=标准输入)
  --fflinks     在使用 --fromfile 时处理链接信息。
  ------- 杂项选项 -------
  --version     打印版本并退出。
  --help        打印用法和本帮助消息并退出。
  --            选项处理终止符。
```

## fallocate

```bash
root@janus-a:~# fallocate --help

Usage:
 fallocate [options] <filename>

Preallocate space to, or deallocate space from a file.

Options:
 -c, --collapse-range remove a range from the file
 -d, --dig-holes      detect zeroes and replace with holes
 -l, --length <num>   length for range operations, in bytes
 -n, --keep-size      maintain the apparent size of the file
 -o, --offset <num>   offset for range operations, in bytes
 -p, --punch-hole     replace a range with a hole (implies -n)
 -z, --zero-range     zero and ensure allocation of a range
 -v, --verbose        verbose mode

 -h, --help     display this help and exit
 -V, --version  output version information and exit

For more details see fallocate(1).
root@janus-a:~# 
```

### sed

```bash
Usage: sed [OPTION]... {script-only-if-no-other-script} [input-file]...

  -n, --quiet, --silent
                 suppress automatic printing of pattern space
  -e script, --expression=script
                 add the script to the commands to be executed
  -f script-file, --file=script-file
                 add the contents of script-file to the commands to be executed
  --follow-symlinks
                 follow symlinks when processing in place
  -i[SUFFIX], --in-place[=SUFFIX]
                 edit files in place (makes backup if SUFFIX supplied)
  -c, --copy
                 use copy instead of rename when shuffling files in -i mode
  -b, --binary
                 does nothing; for compatibility with WIN32/CYGWIN/MSDOS/EMX (
                 open files in binary mode (CR+LFs are not treated specially))
  -l N, --line-length=N
                 specify the desired line-wrap length for the `l' command
  --posix
                 disable all GNU extensions.
  -r, --regexp-extended
                 use extended regular expressions in the script.
  -s, --separate
                 consider files as separate rather than as a single continuous
                 long stream.
  -u, --unbuffered
                 load minimal amounts of data from the input files and flush
                 the output buffers more often
  -z, --null-data
                 separate lines by NUL characters
  --help
                 display this help and exit
  --version
                 output version information and exit

If no -e, --expression, -f, or --file option is given, then the first
non-option argument is taken as the sed script to interpret.  All
remaining arguments are names of input files; if no input files are
specified, then the standard input is read.

GNU sed home page: <http://www.gnu.org/software/sed/>.
General help using GNU software: <http://www.gnu.org/gethelp/>.
E-mail bug reports to: <bug-sed@gnu.org>.
Be sure to include the word ``sed'' somewhere in the ``Subject:'' field.

```

## netstate

```bash
usage: netstat [-vWeenNcCF] [<Af>] -r         netstat {-V|--version|-h|--help}
       netstat [-vWnNcaeol] [<Socket> ...]
       netstat { [-vWeenNac] -I[<Iface>] | [-veenNac] -i | [-cnNe] -M | -s [-6tuw] } [delay]

        -r, --route              display routing table
        -I, --interfaces=<Iface> display interface table for <Iface>
        -i, --interfaces         display interface table
        -g, --groups             display multicast group memberships
        -s, --statistics         display networking statistics (like SNMP)
        -M, --masquerade         display masqueraded connections

        -v, --verbose            be verbose
        -W, --wide               don't truncate IP addresses
        -n, --numeric            don't resolve names
        --numeric-hosts          don't resolve host names
        --numeric-ports          don't resolve port names
        --numeric-users          don't resolve user names
        -N, --symbolic           resolve hardware names
        -e, --extend             display other/more information
        -p, --programs           display PID/Program name for sockets
        -o, --timers             display timers
        -c, --continuous         continuous listing

        -l, --listening          display listening server sockets
        -a, --all                display all sockets (default: connected)
        -F, --fib                display Forwarding Information Base (default)
        -C, --cache              display routing cache instead of FIB
        -Z, --context            display SELinux security context for sockets

  <Socket>={-t|--tcp} {-u|--udp} {-U|--udplite} {-S|--sctp} {-w|--raw}
           {-x|--unix} --ax25 --ipx --netrom
  <AF>=Use '-6|-4' or '-A <af>' or '--<af>'; default: inet
  List of possible address families (which support routing):
    inet (DARPA Internet) inet6 (IPv6) ax25 (AMPR AX.25) 
    netrom (AMPR NET/ROM) ipx (Novell IPX) ddp (Appletalk DDP) 
    x25 (CCITT X.25) 
```

## tar

```bash
tar -tvf archive.tar 
```

## rz

```
rz version 0.12.20
Usage: rz [options] [filename.if.xmodem]
Receive files with ZMODEM/YMODEM/XMODEM protocol
    (X) = option applies to XMODEM only
    (Y) = option applies to YMODEM only
    (Z) = option applies to ZMODEM only
  -+, --append                append to existing files
  -a, --ascii                 ASCII transfer (change CR/LF to LF)
  -b, --binary                binary transfer
  -B, --bufsize N             buffer N bytes (N==auto: buffer whole file)
  -c, --with-crc              Use 16 bit CRC (X)
  -C, --allow-remote-commands allow execution of remote commands (Z)
  -D, --null                  write all received data to /dev/null
      --delay-startup N       sleep N seconds before doing anything
  -e, --escape                Escape control characters (Z)
  -E, --rename                rename any files already existing
      --errors N              generate CRC error every N bytes (debugging)
  -h, --help                  Help, print this usage message
  -m, --min-bps N             stop transmission if BPS below N
  -M, --min-bps-time N          for at least N seconds (default: 120)
  -O, --disable-timeouts      disable timeout code, wait forever for data
      --o-sync                open output file(s) in synchronous write mode
  -p, --protect               protect existing files
  -q, --quiet                 quiet, no progress reports
  -r, --resume                try to resume interrupted file transfer (Z)
  -R, --restricted            restricted, more secure mode
  -s, --stop-at {HH:MM|+N}    stop transmission at HH:MM or in N seconds
  -S, --timesync              request remote time (twice: set local time)
      --syslog[=off]          turn syslog on or off, if possible
  -t, --timeout N             set timeout to N tenths of a second
  -u, --keep-uppercase        keep upper case filenames
  -U, --unrestrict            disable restricted mode (if allowed to)
  -v, --verbose               be verbose, provide debugging information
  -w, --windowsize N          Window is N bytes (Z)
  -X  --xmodem                use XMODEM protocol
  -y, --overwrite             Yes, clobber existing file if any
      --ymodem                use YMODEM protocol
  -Z, --zmodem                use ZMODEM protocol
```





### 权限模式

##  四位八进制表示法

​	八进制表示法使用三位八进制数字来表示文件或目录的权限。每一位数字代表一组用户的权限，分别是文件所有者（user）、文件所属组（group）和其他用户（others）。

- 位数作用

  - 第一位: 特殊权限位

  - 第二位: 文件所有者的权限

  - 第三位: 文件所属组的权限

  - 第四位: 其他用户的权限

- 权限位的含义

  - 4: 读权限（read, r）

  - *: 写权限（write, w）

  - 1: 执行权限（execute, x）

- 权限组合

  - 7: 读、写、执行（rwx）

  - 6: 读、写（rw-）

  - 5: 读、执行（r-x）

  - 4: 读（r--）

  - 3: 写、执行（-wx）

  - 2: 写（-w-）

  - 1: 执行（--x）

  - 0: 无权限（---）

##  符号表示法

符号表示法使用字符来表示用户类别和权限。

- 用户类别: 
  - u: 文件所有者（user）
  - g: 文件所属组（group）
  - o: 其他用户（others）
  - a: 所有用户（all，等同于 `ugo`）
- 操作符: 
  - +: 添加权限
  - -: 移除权限
  - =: 设置权限
- 权限: 
  - r: 读权限（read）
  - w: 写权限（write）
  - x: 执行权限（execute）

# shell 特殊符号

| 符号 | 作用                                         |
| ---- | -------------------------------------------- |
| \|   | 将前一个命令的输出作为下一个命令的输入       |
| >    | 将命令的输出重定向到文件，覆盖文件内容       |
| >>   | 将命令的输出重定向到文件，追加到文件末尾     |
| &    | 将命令放入后台执行                           |
| &&   | 只有前一个命令执行成功时，才会执行下一个命令 |
| *    | 匹配任意数量的字符（包括零个字符）           |
| ？   | 匹配一个字符                                 |

# 默认安装命令

## curl

```bash
curl [options...] <url>
-X, --request <method> Specify request method to use
```

## tr

```bash
Usage: tr [OPTION]... STRING1 [STRING2]
Translate, squeeze, and/or delete characters from standard input,
writing to standard output.  STRING1 and STRING2 specify arrays of
characters ARRAY1 and ARRAY2 that control the action.

  -c, -C, --complement    use the complement of ARRAY1
  -d, --delete            delete characters in ARRAY1, do not translate
  -s, --squeeze-repeats   replace each sequence of a repeated character
                            that is listed in the last specified ARRAY,
                            with a single occurrence of that character
  -t, --truncate-set1     first truncate ARRAY1 to length of ARRAY2
      --help        display this help and exit
      --version     output version information and exit

ARRAYs are specified as strings of characters.  Most represent themselves.
Interpreted sequences are:

  \NNN            character with octal value NNN (1 to 3 octal digits)
  \\              backslash
  \a              audible BEL
  \b              backspace
  \f              form feed
  \n              new line
  \r              return
  \t              horizontal tab
  \v              vertical tab
  CHAR1-CHAR2     all characters from CHAR1 to CHAR2 in ascending order
  [CHAR*]         in ARRAY2, copies of CHAR until length of ARRAY1
  [CHAR*REPEAT]   REPEAT copies of CHAR, REPEAT octal if starting with 0
  [:alnum:]       all letters and digits
  [:alpha:]       all letters
  [:blank:]       all horizontal whitespace
  [:cntrl:]       all control characters
  [:digit:]       all digits
  [:graph:]       all printable characters, not including space
  [:lower:]       all lower case letters
  [:print:]       all printable characters, including space
  [:punct:]       all punctuation characters
  [:space:]       all horizontal or vertical whitespace
  [:upper:]       all upper case letters
  [:xdigit:]      all hexadecimal digits
  [=CHAR=]        all characters which are equivalent to CHAR

Translation occurs if -d is not given and both STRING1 and STRING2 appear.
-t is only significant when translating.  ARRAY2 is extended to length of
ARRAY1 by repeating its last character as necessary.  Excess characters
of ARRAY2 are ignored.  Character classes expand in unspecified order;
while translating, [:lower:] and [:upper:] may be used in pairs to
specify case conversion.  Squeezing occurs after translation or deletion.

GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
Full documentation <https://www.gnu.org/software/coreutils/tr>
or available locally via: info '(coreutils) tr invocation'
```

```bash
`tr` 命令的使用格式为 `tr [OPTION]... STRING1 [STRING2]`，它可以从标准输入读取字符，并根据指定的操作对这些字符进行翻译（替换）、压缩（合并重复字符）和/或删除，最终将结果输出到标准输出。`STRING1` 和 `STRING2` 分别指定了两个字符数组 `ARRAY1` 和 `ARRAY2`，它们用来控制 `tr` 的具体行为。

以下是 `tr` 命令支持的一些选项：

- `-c`, `-C`, `--complement`：使用 `ARRAY1` 的补集，即处理所有不属于 `ARRAY1` 中列出的字符。
- `-d`, `--delete`：删除 `ARRAY1` 中列出的所有字符，而不执行任何转换。
- `-s`, `--squeeze-repeats`：将最后一个指定的 `ARRAY` 中连续出现的字符序列替换为该字符的一个实例。
- `-t`, `--truncate-set1`：首先截断 `ARRAY1` 使其长度与 `ARRAY2` 相等。
- `--help`：显示帮助信息并退出。
- `--version`：显示版本信息并退出。

字符数组可以通过字符串形式指定，大多数情况下字符直接表示自己。但也有几种特殊的解释序列：

- `\NNN`：八进制值为 NNN 的字符（1 到 3 位八进制数字）。
- `\\`：反斜杠。
- `\a`：可听的 BEL 铃声。
- `\b`：退格键。
- `\f`：换页符。
- `\n`：新行。
- `\r`：回车。
- `\t`：水平制表符。
- `\v`：垂直制表符。
- `CHAR1-CHAR2`：从 `CHAR1` 到 `CHAR2` 的所有字符，按 ASCII 码升序排列。
- `[CHAR*]`：在 `ARRAY2` 中，复制字符直到与 `ARRAY1` 的长度相同。
- `[CHAR*REPEAT]`：复制字符 REPEAT 次，如果 REPEAT 以 0 开头，则认为是八进制数。
- `[:alnum:]`：所有字母和数字。
- `[:alpha:]`：所有字母。
- `[:blank:]`：所有水平空白字符。
- `[:cntrl:]`：所有控制字符。
- `[:digit:]`：所有数字。
- `[:graph:]`：所有可打印字符，不包括空格。
- `[:lower:]`：所有小写字母。
- `[:print:]`：所有可打印字符，包括空格。
- `[:punct:]`：所有标点符号。
- `[:space:]`：所有水平或垂直空白字符。
- `[:upper:]`：所有大写字母。
- `[:xdigit:]`：所有十六进制数字。
- `[=CHAR=]`：所有与 `CHAR` 等价的字符。

当未给出 `-d` 选项且同时提供了 `STRING1` 和 `STRING2` 时会发生翻译操作。
`-t` 选项仅在翻译时有意义。如果需要，`ARRAY2` 会通过重复其最后一个字符来扩展至与 `ARRAY1` 长度相同。
`ARRAY2` 中超出 `ARRAY1` 长度的多余字符将被忽略。
只有 `[:lower:]` 和 `[:upper:]` 在扩展时保证按升序展开；在翻译过程中，它们可以成对使用来指定大小写的转换。
压缩操作发生在翻译或删除之后。

GNU coreutils 的在线帮助可以在 <https://www.gnu.org/software/coreutils/> 获取，
完整的文档可以在 <https://www.gnu.org/software/coreutils/tr> 或者

通过本地命令 `info '(coreutils) tr invocation'` 查看。
```

## cut

```bash
Usage: cut OPTION... [FILE]...
Print selected parts of lines from each FILE to standard output.

With no FILE, or when FILE is -, read standard input.

Mandatory arguments to long options are mandatory for short options too.
  -b, --bytes=LIST        select only these bytes
  -c, --characters=LIST   select only these characters
  -d, --delimiter=DELIM   use DELIM instead of TAB for field delimiter
  -f, --fields=LIST       select only these fields;  also print any line
                            that contains no delimiter character, unless
                            the -s option is specified
  -n                      (ignored)
      --complement        complement the set of selected bytes, characters
                            or fields
  -s, --only-delimited    do not print lines not containing delimiters
      --output-delimiter=STRING  use STRING as the output delimiter
                            the default is to use the input delimiter
  -z, --zero-terminated   line delimiter is NUL, not newline
      --help        display this help and exit
      --version     output version information and exit

Use one, and only one of -b, -c or -f.  Each LIST is made up of one
range, or many ranges separated by commas.  Selected input is written
in the same order that it is read, and is written exactly once.
Each range is one of:

  N     N'th byte, character or field, counted from 1
  N-    from N'th byte, character or field, to end of line
  N-M   from N'th to M'th (included) byte, character or field
  -M    from first to M'th (included) byte, character or field

GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
Full documentation <https://www.gnu.org/software/coreutils/cut>
or available locally via: info '(coreutils) cut invocation'
```

## head

```bash
head -n -1 finename # 不现实最后一行
```

```bash
Usage: head [OPTION]... [FILE]...
Print the first 10 lines of each FILE to standard output.
With more than one FILE, precede each with a header giving the file name.

With no FILE, or when FILE is -, read standard input.

Mandatory arguments to long options are mandatory for short options too.
  -c, --bytes=[-]NUM       print the first NUM bytes of each file;
                             with the leading '-', print all but the last
                             NUM bytes of each file
  -n, --lines=[-]NUM       print the first NUM lines instead of the first 10;
                             with the leading '-', print all but the last
                             NUM lines of each file
  -q, --quiet, --silent    never print headers giving file names
  -v, --verbose            always print headers giving file names
  -z, --zero-terminated    line delimiter is NUL, not newline
      --help        display this help and exit
      --version     output version information and exit

NUM may have a multiplier suffix:
b 512, kB 1000, K 1024, MB 1000*1000, M 1024*1024,
GB 1000*1000*1000, G 1024*1024*1024, and so on for T, P, E, Z, Y, R, Q.
Binary prefixes can be used, too: KiB=K, MiB=M, and so on.

GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
Full documentation <https://www.gnu.org/software/coreutils/head>
or available locally via: info '(coreutils) head invocation'
```

## tail

```bash
tail -n +2 # 从第2行开始到最后一行
```

```bash
Usage: tail [OPTION]... [FILE]...
Print the last 10 lines of each FILE to standard output.
With more than one FILE, precede each with a header giving the file name.

With no FILE, or when FILE is -, read standard input.

Mandatory arguments to long options are mandatory for short options too.
  -c, --bytes=[+]NUM       output the last NUM bytes; or use -c +NUM to
                             output starting with byte NUM of each file
  -f, --follow[={name|descriptor}]
                           output appended data as the file grows;
                             an absent option argument means 'descriptor'
  -F                       same as --follow=name --retry
  -n, --lines=[+]NUM       output the last NUM lines, instead of the last 10;
                             or use -n +NUM to skip NUM-1 lines at the start
      --max-unchanged-stats=N
                           with --follow=name, reopen a FILE which has not
                             changed size after N (default 5) iterations
                             to see if it has been unlinked or renamed
                             (this is the usual case of rotated log files);
                             with inotify, this option is rarely useful
      --pid=PID            with -f, terminate after process ID, PID dies
  -q, --quiet, --silent    never output headers giving file names
      --retry              keep trying to open a file if it is inaccessible
  -s, --sleep-interval=N   with -f, sleep for approximately N seconds
                             (default 1.0) between iterations;
                             with inotify and --pid=P, check process P at
                             least once every N seconds
  -v, --verbose            always output headers giving file names
  -z, --zero-terminated    line delimiter is NUL, not newline
      --help        display this help and exit
      --version     output version information and exit

NUM may have a multiplier suffix:
b 512, kB 1000, K 1024, MB 1000*1000, M 1024*1024,
GB 1000*1000*1000, G 1024*1024*1024, and so on for T, P, E, Z, Y, R, Q.
Binary prefixes can be used, too: KiB=K, MiB=M, and so on.

With --follow (-f), tail defaults to following the file descriptor, which
means that even if a tail'ed file is renamed, tail will continue to track
its end.  This default behaviour is not desirable when you really want to
track the actual name of the file, not the file descriptor (e.g., log
rotation).  Use --follow=name in that case.  That causes tail to track the
named file in a way that accommodates renaming, removal and creation.

GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
Full documentation <https://www.gnu.org/software/coreutils/tail>
or available locally via: info '(coreutils) tail invocation'
```

### which

```bash
which <command>
# 查看命令文件的位置
```

### chmod

```bash
 sudo chmod -R a+rw /mnt/c/Users/Janus_
```

```bash
# 讲指定usr添加到指定group
sudo usermod -aG docker janus
# 刷新
newgrp docker
```



### 驱动安装

```bash
sudo mkinitcpio -P 是一个用于更新 mkinitcpio 脚本的命令，主要用于初始化块设备配置。
```

```bash
# 查看gpu的动态
watch -n 1 nvidia-smi
```

### chmod

