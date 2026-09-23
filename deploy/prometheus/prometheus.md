- ss -tunl

  ```bash
  ss -tunl 是一个用于显示套接字统计信息的命令，它是 ss（socket statistics）工具的一个常用示例。下面是对这个命令中各个选项的解释：
  
      -t：显示TCP套接字的信息。
      -u：显示UDP套接字的信息。
      -n：显示数字形式的地址和端口号，而不是尝试解析为名称。这可以加快显示速度，并且在处理大量连接时很有用。
      -l：仅显示处于监听状态的套接字。
  ```

  

```bash
http://localhost:9100/metrics # node-exprotor
http://localhost:9090/query # prometheus
```

