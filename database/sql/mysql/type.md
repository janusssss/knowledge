# 数值类型（Numeric Types） 
```sql
TINYINT      -- 1字节，-128到127（有符号）或0到255（无符号）
SMALLINT     -- 2字节，-32,768到32,767
MEDIUMINT    -- 3字节，-8,388,608到8,388,607
INT/INTEGER  -- 4字节，-2,147,483,648到2,147,483,647
BIGINT       -- 8字节，很大范围的整数
```

# 浮点数类型 
```sql
FLOAT        -- 4字节单精度浮点数
DOUBLE       -- 8字节双精度浮点数
DECIMAL(M,D) -- 精确小数，M总位数，D小数位数
```

# 日期时间类型（Date and Time Types） 
```sql
DATE         -- 日期，'YYYY-MM-DD'格式，如'2023-12-01'
TIME         -- 时间，'HH:MM:SS'格式，如'14:30:00'
DATETIME     -- 日期时间，'YYYY-MM-DD HH:MM:SS'格式
TIMESTAMP    -- 时间戳，范围较小，自动时区转换
YEAR         -- 年份，1901-2155或0000
```

# 字符串类型（String Types） 
```sql
CHAR(n)      -- 定长字符串，最大255字符，不足用空格填充
VARCHAR(n)   -- 变长字符串，最大65,535字节
TINYTEXT     -- 最大255字符
TEXT         -- 最大65,535字符
MEDIUMTEXT   -- 最大16,777,215字符
LONGTEXT     -- 最大4,294,967,295字符
```

# 枚举和集合类型（Enum and Set Types） 
```sql
ENUM('value1', 'value2', 'value3')  -- 枚举，只能选一个值
SET('value1', 'value2', 'value3')   -- 集合，可选多个值
```

# JSON类型（JSON Type） 
```sql
JSON         -- JSON数据类型，MySQL 5.7.8+
```