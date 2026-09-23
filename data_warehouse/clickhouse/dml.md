### insert

```sql
CREATE TABLE bill (
    name String,
  	prov_id UInt16,
  	day_id FixedString(8)
) ENGINE = MergeTree()
PARTITION BY (day_id,prov_id)
ORDER BY (day_id,prov_id);
```

### 查看分区

```sql
SELECT partition FROM system.parts WHERE table = 'bill' ORDER BY partition;
select partition,table,database from system.parts where table='dwa_prd_pd_inst_month' and database='dwa_db';
```



```sql
FROM (
    SELECT *,
           row_number() OVER (
               PARTITION BY unique_column1, unique_column2 
               ORDER BY timestamp_column DESC
           ) as rn
    FROM original_table
) 
WHERE rn = 1;



-- 1. 创建临时表（去重后的数据）
CREATE TABLE temp_deduplicated AS original_table
ENGINE = MergeTree()
ORDER BY (primary_key, other_unique_keys)
SETTINGS index_granularity = 8192;

-- 2. 插入去重数据
INSERT INTO temp_deduplicated
SELECT *
FROM (
    SELECT *,
           row_number() OVER (PARTITION BY unique_key ORDER BY timestamp DESC) as rn
    FROM original_table
) 
WHERE rn = 1;

-- 3. 清空原表并插入去重数据
TRUNCATE TABLE original_table;
INSERT INTO original_table 
SELECT * FROM temp_deduplicated;

-- 4. 清理临时表
DROP TABLE temp_deduplicated;



-- 最快的方案：使用ReplacingMergeTree
-- 1. 创建新表
CREATE TABLE new_table AS original_table
ENGINE = ReplacingMergeTree()
ORDER BY (unique_key, timestamp);  -- 按去重字段排序

-- 2. 插入所有数据（自动去重）
INSERT INTO new_table SELECT * FROM original_table;

-- 3. 重命名（瞬间完成）
RENAME TABLE 
    original_table TO original_table_backup,
    new_table TO original_table;

-- 这个方案比窗口函数快很多




-- 优化版本，预计耗时：15-25分钟
-- 1. 创建表时指定设置
CREATE TABLE new_table AS original_table
ENGINE = ReplacingMergeTree()
ORDER BY (unique_key, timestamp)
SETTINGS 
    index_granularity = 8192,
    merge_with_ttl_timeout = 3600;

-- 2. 插入数据
INSERT INTO new_table SELECT * FROM original_table;

-- 3. 强制合并（确保去重完成）
OPTIMIZE TABLE new_table FINAL;  -- 这个操作可能需要5-10分钟

-- 4. 重命名
RENAME TABLE 
    original_table TO original_table_backup,
    new_table TO original_table;
```