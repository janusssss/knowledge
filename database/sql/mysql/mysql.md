```bash
sudo mysql_secure_installation
```

```bash
SHOW VARIABLES LIKE 'validate_password%';
 
SET GLOBAL validate_password.length = 6;
SET GLOBAL validate_password.mixed_case_count = 0;
SET GLOBAL validate_password.number_count = 0;
SET GLOBAL validate_password.special_char_count = 0;

ALTER USER 'username'@'host' IDENTIFIED BY 'Aa1!Bb2@C';
```

```bash
root
1234
```

### 安装maridab

```
sudo pacman -S mariadb
# 初始化
sudo mysql_install_db --user=mysql --basedir=/usr --datadir=/var/lib/mysql
```

```mariadb
CREATE USER 'janus'@'%' IDENTIFIED BY 'janus';
GRANT ALL PRIVILEGES ON *.* TO 'janus'@'%';
FLUSH PRIVILEGES;

CREATE USER 'janus'@'localhost' IDENTIFIED BY 'janus';
GRANT ALL PRIVILEGES ON *.* TO 'janus'@'localhost';
FLUSH PRIVILEGES;

SELECT User, Host FROM mysql.user;
SHOW GRANTS FOR 'janus'@'%';
```

### 用户权限相关

```mysql
# 查看密码策略
show variables like 'validate_password.%';
set persist validate_password.policy=LOW;
alter user 'janus'@'localhost' identified by 'janus';
insert
```

### 连接

```bash
mysql -h 10.142.102.136 -p -u qwsn 
```

### 查看表详情

```sql
SHOW FULL COLUMNS FROM 
show 
```

### 修改字段注释

```bash
ALTER TABLE 表名 MODIFY 字段名 字段类型 COMMENT '新的注释内容';
```

### 数据类型

### DDL操作

#### create

用于创建数据库或数据库对象，如表、视图、索引、存储过程等。

```sql
-- 创建数据库
CREATE DATABASE database_name;
-- 创建表
CREATE TABLE table_name (column_definitions);
-- 复制存在的表
create table table_name as select * from exit_table_name;
```

#### ALTER

用于修改现有的数据库或其对象结构。

```sql
-- 添加列
ALTER TABLE table_name ADD column_name datatype;
-- 修改列的数据类型
ALTER TABLE table_name MODIFY column_name new_datatype;
-- 删除列
ALTER TABLE table_name DROP COLUMN column_name;
-- 重命名表
ALTER TABLE old_table_name RENAME TO new_table_name;

ALTER TABLE orders ADD PRIMARY KEY (order_id, product_id);
```

#### DROP

用于删除数据库、表或其他数据库对象。

```sql
-- 删除数据库
DROP DATABASE database_name;
-- 删除表
DROP TABLE table_name;
```

### DML

#### insert

```sql
INSERT INTO ... VALUES(...)
INSERT INTO ... SELECT ...
```

### 数据类型

#### Numeric Data Types

#####  Integer Types

| ype         | Storage (Bytes) | Minimum Value Signed | Minimum Value Unsigned | Maximum Value Signed | Maximum Value Unsigned |
| ----------- | --------------- | -------------------- | ---------------------- | -------------------- | ---------------------- |
| `TINYINT`   | 1               | `-128`               | `0`                    | `127`                | `255`                  |
| `SMALLINT`  | 2               | `-32768`             | `0`                    | `32767`              | `65535`                |
| `MEDIUMINT` | 3               | `-8388608`           | `0`                    | `8388607`            | `16777215`             |
| `INT`       | 4               | `-2147483648`        | `0`                    | `2147483647`         | `4294967295`           |
| `BIGINT`    | 8               | `-263`               | `0`                    | `263-1`              | `264-1`                |

#### String Data Types

The string data types are CHAR, VARCHAR, BINARY, VARBINARY, BLOB, TEXT, ENUM, and SET

### Storage Engines

| engine  | 概述                                                         |
| ------- | ------------------------------------------------------------ |
| InnoDB  | 默认存储引擎，事务安全（符合 ACID），具有提交、回滚和崩溃恢复功能，数据存储在聚集索引中，以减少基于主键查询的 I/O |
| MyISAM  | 占用空间小，表级锁定限制了读/写工作负载的性能，因此它通常用于 Web 和数据仓库配置中的只读或只读工作负载 |
| Memory  | 数据存储在 RAM 中，可以快速查找非关键数据的环境中快速访问    |
| CSV     | 带有逗号分隔值的文本文档                                     |
| Archive | 数据紧凑、未编制索引，用于存储和检索大量很少引用的历史、存档或安全审计信息 |
| NDB     | 也称为 NDBCLUSTER，此群集数据库引擎特别适用于需要最高正常运行时间和可用性的应用进程 |

### func

|                                                   |                                                              |
| ------------------------------------------------- | ------------------------------------------------------------ |
| JSON_SET(info, '$.address.city', 'Shanghai')      | 将所有记录的 address.city 更新为 "Shanghai"，如果该字段不存在，则自动创建。 |
| JSON_REPLACE(info, '$.address.city', 'Shanghai'); | 如果字段不存在，这条语句不会报错，也不会修改数据。           |
|                                                   |                                                              |

### mysqldump

```sql
mysqldump -h 10.143.160.145 -P 3306 -u root -p dacp_cthq_dev dx_assets_field_hx \
--no-create-info 
> backup_inserts_only.sql
```



### mysql

#### 执行sql文件

```bash
mysql < file_name
```

### mysqlimport

### 导入数据

```bash
# SET GLOBAL local_infile = OFF;
mysqlimport -p --local --ignore-lines 1 core dacp_table_info.csv
```

