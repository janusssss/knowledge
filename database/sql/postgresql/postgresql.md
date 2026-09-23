### 安装

```bash
# ubuntu
sudo apt install postgresql postgresql-contribs
```

### 创建新用户

```sql
sudo -u postgres psql
CREATE ROLE janus WITH LOGIN SUPERUSER PASSWORD 'janus';
psql -U janus -d postgres
```

```sql
CREATE DATABASE asia；
DROP TABLE asia;
```

### 增删改查

```sql
create table asia(name varchar(8),age integer);
insert into asia(name,age) values ('janus',7);
select * from asia;
update asia set age = 28 where age = 21;
delete from asia where age = 7;
```

### 查看相关数据库情况

```bash
\du # 查看用户
\dt  \d # 查看表
\c mydatabase # 连接数据库
\q # 退出
\l # 查看数据库列表
\d # 查看表格结构
\d+ # 查看详细表格结构
```

### 列增加注释

### 数据类型

- 数值类型

  | **名字**  | **存储空间** | **描述**           | **范围**                                    |
  | --------- | ------------ | ------------------ | ------------------------------------------- |
  | smallint  | 2 字节       | 小范围整数         | -32768 到 +32767                            |
  | integer   | 4 字节       | 常用的整数         | -2147483648 到 +2147483647                  |
  | bigint    | 8 字节       | 大范围的整数       | -9223372036854775808 到 9223372036854775807 |
  | decimal   | 变长         | 用户声明精度，精确 | 无限制                                      |
  | numeric   | 变长         | 用户声明精度，精确 | 无限制                                      |
  | real      | 4 字节       | 变精度，不精确     | 6 位十进制数字精度                          |
  | double    | 8 字节       | 变精度，不精确     | 15 位十进制数字精度                         |
  | serial    | 4 字节       | 自增整数           | 1 到 +2147483647                            |
  | bigserial | 8 字节       | 大范围的自增整数   | 1 到 9223372036854775807                    |

- 字符类型

  | **名字**   | **描述**                           |
  | ---------- | ---------------------------------- |
  | varchar(n) | 变长，有长度限制, 没有n是varchar() |
  | char(n)    | 定长,不足补空白，没有n是char(1)    |
  | text       | 变长，无长度限制                   |

- 日期/时间类型

  | **名字**                 | **存储空间** | **描述**           | **最低值**   | **最高值**  | **分辨率** |
  | ------------------------ | ------------ | ------------------ | ------------ | ----------- | ---------- |
  | timestamp[无时区] 默认   | 8字节        | 包括日期和时间     | 4713 BC      | 5874897AD   | 1毫秒/14位 |
  | TIMESTAMP WITH TIME ZONE | 8字节        | 日期和时间，带时区 | 4713 BC      | 5874897AD   | 1毫秒/14位 |
  | interval                 | 12字节       | 时间间隔           | -178000000年 | 178000000年 | 1毫秒/14位 |
  | date                     | 4字节        | 只用于日期         | 4713 BC      | 32767AD     | 1天        |
  | time[无时区]             | 8字节        | 只用于一日内时间   | 00:00:00     | 24:00:00    | 1毫秒/14位 |

  ​	

```postgresql
 create table simplified(id smallint, creat_time timestamp, update_time timestamp,cangjie varchar(5), cantonese varchar(8))
```

### 备份与恢复

```bash
# 备份
pg_dump -U 用户名 -F c -b -v -f "备份文件路径" 数据库名
# 恢复
pg_restore -U 用户名 -h 主机地址 -d 数据库名 -v "备份文件路径"
```

### 执行sql文件

```bash
 psql -U janus -d postgres -c file.sql
```

### 查看表数据大小

```postgresql
select pg_size_pretty(pg_total_relation_size('res_daas_scene_task_statistics')) AS "总大小";
```

### 获取表结构重建

```bash
pg_dump -U your_username -h your_host -d your_database -t your_table --schema-only > output.sql
```

### 更改表名

```postgresql
alter table old_table_name rename to new_table_name;
```

### pg_dump

```bash
# 仅导出指定表的结构，不包含数据
pg_dump --dbname database_name --table table_name --schema-only --file output_file --verbose  --no-owner
```

```bash
pg_dump dumps a database as a text file or to other formats.

Usage:
  pg_dump [OPTION]... [DBNAME]

General options:
  -f, --file=FILENAME          output file or directory name
  -F, --format=c|d|t|p         output file format (custom, directory, tar,
                               plain text (default))
  -j, --jobs=NUM               use this many parallel jobs to dump
  -v, --verbose                verbose mode
  -V, --version                output version information, then exit
  -Z, --compress=METHOD[:DETAIL]
                               compress as specified
  --lock-wait-timeout=TIMEOUT  fail after waiting TIMEOUT for a table lock
  --no-sync                    do not wait for changes to be written safely to disk
  -?, --help                   show this help, then exit

Options controlling the output content:
  -a, --data-only              dump only the data, not the schema
  -b, --large-objects          include large objects in dump
  --blobs                      (same as --large-objects, deprecated)
  -B, --no-large-objects       exclude large objects in dump
  --no-blobs                   (same as --no-large-objects, deprecated)
  -c, --clean                  clean (drop) database objects before recreating
  -C, --create                 include commands to create database in dump
  -e, --extension=PATTERN      dump the specified extension(s) only
  -E, --encoding=ENCODING      dump the data in encoding ENCODING
  -n, --schema=PATTERN         dump the specified schema(s) only
  -N, --exclude-schema=PATTERN do NOT dump the specified schema(s)
  -O, --no-owner               skip restoration of object ownership in
                               plain-text format
  -s, --schema-only            dump only the schema, no data
  -S, --superuser=NAME         superuser user name to use in plain-text format
  -t, --table=PATTERN          dump only the specified table(s)
  -T, --exclude-table=PATTERN  do NOT dump the specified table(s)
  -x, --no-privileges          do not dump privileges (grant/revoke)
  --binary-upgrade             for use by upgrade utilities only
  --column-inserts             dump data as INSERT commands with column names
  --disable-dollar-quoting     disable dollar quoting, use SQL standard quoting
  --disable-triggers           disable triggers during data-only restore
  --enable-row-security        enable row security (dump only content user has
                               access to)
  --exclude-table-and-children=PATTERN
                               do NOT dump the specified table(s), including
                               child and partition tables
  --exclude-table-data=PATTERN do NOT dump data for the specified table(s)
  --exclude-table-data-and-children=PATTERN
                               do NOT dump data for the specified table(s),
                               including child and partition tables
  --extra-float-digits=NUM     override default setting for extra_float_digits
  --if-exists                  use IF EXISTS when dropping objects
  --include-foreign-data=PATTERN
                               include data of foreign tables on foreign
                               servers matching PATTERN
  --inserts                    dump data as INSERT commands, rather than COPY
  --load-via-partition-root    load partitions via the root table
  --no-comments                do not dump comments
  --no-publications            do not dump publications
  --no-security-labels         do not dump security label assignments
  --no-subscriptions           do not dump subscriptions
  --no-table-access-method     do not dump table access methods
  --no-tablespaces             do not dump tablespace assignments
  --no-toast-compression       do not dump TOAST compression methods
  --no-unlogged-table-data     do not dump unlogged table data
  --on-conflict-do-nothing     add ON CONFLICT DO NOTHING to INSERT commands
  --quote-all-identifiers      quote all identifiers, even if not key words
  --rows-per-insert=NROWS      number of rows per INSERT; implies --inserts
  --section=SECTION            dump named section (pre-data, data, or post-data)
  --serializable-deferrable    wait until the dump can run without anomalies
  --snapshot=SNAPSHOT          use given snapshot for the dump
  --strict-names               require table and/or schema include patterns to
                               match at least one entity each
  --table-and-children=PATTERN dump only the specified table(s), including
                               child and partition tables
  --use-set-session-authorization
                               use SET SESSION AUTHORIZATION commands instead of
                               ALTER OWNER commands to set ownership

Connection options:
  -d, --dbname=DBNAME      database to dump
  -h, --host=HOSTNAME      database server host or socket directory
  -p, --port=PORT          database server port number
  -U, --username=NAME      connect as specified database user
  -w, --no-password        never prompt for password
  -W, --password           force password prompt (should happen automatically)
  --role=ROLENAME          do SET ROLE before dump

If no database name is supplied, then the PGDATABASE environment
variable value is used.

Report bugs to <pgsql-bugs@lists.postgresql.org>.
PostgreSQL home page: <https://www.postgresql.org/>
```

### psql

```bash
# 执行sql exit文件
psql --dbname database_name --flie file_name
```

```bash
psql is the PostgreSQL interactive terminal.

Usage:
  psql [OPTION]... [DBNAME [USERNAME]]

General options:
  -c, --command=COMMAND    run only single command (SQL or internal) and exit
  -d, --dbname=DBNAME      database name to connect to (default: "janus")
  -f, --file=FILENAME      execute commands from file, then exit
  -l, --list               list available databases, then exit
  -v, --set=, --variable=NAME=VALUE
                           set psql variable NAME to VALUE
                           (e.g., -v ON_ERROR_STOP=1)
  -V, --version            output version information, then exit
  -X, --no-psqlrc          do not read startup file (~/.psqlrc)
  -1 ("one"), --single-transaction
                           execute as a single transaction (if non-interactive)
  -?, --help[=options]     show this help, then exit
      --help=commands      list backslash commands, then exit
      --help=variables     list special variables, then exit

Input and output options:
  -a, --echo-all           echo all input from script
  -b, --echo-errors        echo failed commands
  -e, --echo-queries       echo commands sent to server
  -E, --echo-hidden        display queries that internal commands generate
  -L, --log-file=FILENAME  send session log to file
  -n, --no-readline        disable enhanced command line editing (readline)
  -o, --output=FILENAME    send query results to file (or |pipe)
  -q, --quiet              run quietly (no messages, only query output)
  -s, --single-step        single-step mode (confirm each query)
  -S, --single-line        single-line mode (end of line terminates SQL command)

Output format options:
  -A, --no-align           unaligned table output mode
      --csv                CSV (Comma-Separated Values) table output mode
  -F, --field-separator=STRING
                           field separator for unaligned output (default: "|")
  -H, --html               HTML table output mode
  -P, --pset=VAR[=ARG]     set printing option VAR to ARG (see \pset command)
  -R, --record-separator=STRING
                           record separator for unaligned output (default: newline)
  -t, --tuples-only        print rows only
  -T, --table-attr=TEXT    set HTML table tag attributes (e.g., width, border)
  -x, --expanded           turn on expanded table output
  -z, --field-separator-zero
                           set field separator for unaligned output to zero byte
  -0, --record-separator-zero
                           set record separator for unaligned output to zero byte

Connection options:
  -h, --host=HOSTNAME      database server host or socket directory (default: "/var/run/postgresql")
  -p, --port=PORT          database server port (default: "5432")
  -U, --username=USERNAME  database user name (default: "janus")
  -w, --no-password        never prompt for password
  -W, --password           force password prompt (should happen automatically)

For more information, type "\?" (for internal commands) or "\help" (for SQL
commands) from within psql, or consult the psql section in the PostgreSQL
documentation.

Report bugs to <pgsql-bugs@lists.postgresql.org>.
PostgreSQL home page: <https://www.postgresql.org/>
```

### 判断当前数据库是否是主数据库

```postgresql
select pg_is_in_recovery();
```

如果返回false，则该节点是主数据库（因为它不处于恢复状态）。如果返回true，则表示这是一个从数据库（它正在从主数据库同步数据，处于恢复模式）

### 运算符

| 符号 | 作用                                     |
| ---- | ---------------------------------------- |
| `||` | 字符串连接运算符                         |
| `::` | 类型强转，等效于`CAST('100' AS INTEGER)` |
|      |                                          |

### DDL

### Alter

### 设置/删除默认值

```postgresql
-- 设置默认值
ALTER TABLE table_name ALTER COLUMN column_name SET DEFAULT default_value;

-- 删除默认值
ALTER TABLE table_name ALTER COLUMN column_name DROP DEFAULT;
-- 更新类型
alter table sample alter column update_time type timestamp(0);

-- 增加列
alter table sample add column update_time timestamp;

 updated_at TIMESTAMP GENERATED ALWAYS AS (CURRENT_TIMESTAMP) STORED
 
alter table sample alter column update_time set default generated always as (current_timestamp) stored;



CREATE SEQUENCE your_table_id_seq;
```

### 元命令

| 命令            | 作用                                                        |
| --------------- | ----------------------------------------------------------- |
| `\d [pattern]`  | 列出当前数据库中的所有表、视图、序列和索引，pattern空为所有 |
| `\dt [pattern]` | 列出表                                                      |
| `\d+ [pattern]` | 列出表的详细信息                                            |
|                 | 注意：pattern为shell通配符模式                              |

### 函数

| 函数                   | 作用 |
| ---------------------- | ---- |
| pg_size_pretty()       |      |
| row_number()           |      |
| to_char(value,formmat) |      |

### 增加更新时间默认值

```sql
-- 1. 首先设置默认值
alter table sample alter column update_time set default current_timestamp;

-- 2. 创建触发器函数
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.update_time = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- 3. 为表创建触发器
CREATE TRIGGER trigger_sample_update_time
    BEFORE UPDATE ON sample
    FOR EACH ROW
    EXECUTE FUNCTION update_timestamp();
```

