### 增加更新时间

```sql
-- 1. 首先设置默认值
alter table sample alter column update_time set default current_timestamp(0);

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

### 查看自定义函数

```sql
SELECT 
    tgname as trigger_name,
    relname as table_name,
    pg_get_triggerdef(oid) as trigger_definition,
    tgenabled as enabled_status
FROM pg_trigger t
JOIN pg_class c ON t.tgrelid = c.oid
WHERE NOT tgisinternal; 
```

### 系统

| 表名                    | 作用                               |
| ----------------------- | ---------------------------------- |
| `pg_catalog.pg_trigger` | 存储数据库中所有触发器的元数据信息 |
|                         |                                    |
|                         |                                    |

### 命令

| 命令               | 作用                           |
| ------------------ | ------------------------------ |
| `show search_path` | 显示当前会话的模式搜索路径设置 |
|                    |                                |
|                    |                                |

### Alter

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
```

### 查看配置文件

```bash
# PostgreSQL服务器的主要配置文件，控制数据库服务器的运行参数
psql -U postgres -c "SHOW config_file;"
# 控制哪些客户端可以连接到数据库,以及使用什么认证方法
psql -U postgres -c "SHOW hba_file;"
```

