### 分类
```sql
-- 排名函数（Ranking Functions） 
ROW_NUMBER(), RANK(), DENSE_RANK(), NTILE()

-- 前后行函数（Offset Functions） 
LAG(), LEAD(), FIRST_VALUE(), LAST_VALUE(), NTH_VALUE()

-- 聚合函数（Aggregate Functions） 
SUM(), AVG(), COUNT(), MAX(), MIN()

-- 数学函数（Mathematical Functions）
ROUND(x, n)          -- 四舍五入
CEIL(x) / CEILING(x) -- 向上取整
FLOOR(x)             -- 向下取整
ABS(x)               -- 绝对值
MOD(x, y)            -- 取模
POWER(x, y) / POW(x, y) -- 幂运算
SQRT(x)              -- 平方根
TRUNCATE(x, n)       -- 截断
```

### LAG()
```sql
-- 语法结构
LAG(column, offset, default_value) OVER (
    PARTITION BY partition_expression 
    ORDER BY sort_expression
)

-- column: 要获取的列名
-- offset: 偏移量（可选，默认为1，表示前一行）
-- default_value: 默认值（可选，没有前一行时返回的值）
-- PARTITION BY: 分区字段（可选）
-- ORDER BY: 排序字段（必须）    
```

# 数学函数 Mathematical Functions
## ROUND(x,n) 
```sql
SELECT ROUND(123.456, 2);  -- 123.46
SELECT ROUND(123.456);     -- 123
```