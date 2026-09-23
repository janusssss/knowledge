### 添加索引

```sql
ALTER TABLE table_name ADD INDEX index_name (column_name);
CREATE INDEX index_name ON table_name (column_name);
```

### 查看进程

```sql
SHOW PROCESSLIST;
kill id;
```

### 用户相关

```sql
-- 查看用户
select user,host from mysql.user;
-- 删除用户
drop user 'janus'@'%';

create user 'janus'@"%" identified by 'janus';
grant all privileges on *.* to 'janus'@'%';
flush privileges;


SELECT User, Host FROM mysql.user;
SHOW GRANTS FOR 'janus'@'%';
```

### alter

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

-- 改表名注释
alter table t_sr_category comment ='分类表';
```

