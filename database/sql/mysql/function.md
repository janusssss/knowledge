# JSON
## 创建JSON数据
```sql
-- JSON_OBJECT() - 创建 JSON 对象
SELECT JSON_OBJECT('name', 'John', 'age', 30, 'city', 'New York');
-- 结果: {"age": 30, "city": "New York", "name": "John"}

-- JSON_ARRAY() - 创建 JSON 数组
SELECT JSON_ARRAY('apple', 'banana', 'orange');
-- 结果: ["apple", "banana", "orange"]

-- JSON_QUOTE() - 将字符串转为 JSON 字符串
SELECT JSON_QUOTE('Hello "World"');
-- 结果: "Hello \"World\""
```
## 查询和提取数据
- JSON_EXTRACT() - 提取 JSON 数据
```sql
-- JSON_EXTRACT() - 提取 JSON 数据
SELECT JSON_EXTRACT('{"name": "John", "age": 30}', '$.name');
-- 结果: "John"
```
- 简写运算符
```sql
-- -> 返回 JSON 类型
SELECT data->'$.name' FROM users;

-- ->> 返回字符串类型（去除引号）
SELECT data->>'$.name' FROM users;

-- 实际示例
SELECT 
    JSON_EXTRACT('{"name": "John"}', '$.name') as with_quotes,
    '{"name": "John"}'->>'$.name' as without_quotes;
-- with_quotes: "John", without_quotes: John
```
- JSON_SEARCH() - 搜索 JSON 数据
```sql
-- 查找值的路径
SELECT JSON_SEARCH('{"user": {"name": "John"}}', 'one', 'John');
-- 结果: "$.user.name"

-- 在数组中查找
SELECT JSON_SEARCH('["apple", "banana", "orange"]', 'one', 'banana');
-- 结果: "$[1]"
```
## 修改JSON数据
- JSON_SET() - 设置值（存在则更新，不存在则添加）
```sql
SELECT JSON_SET('{"a": 1, "b": 2}', '$.c', 3, '$.a', 10);
-- 结果: {"a": 10, "b": 2, "c": 3}
```
- JSON_INSERT() - 插入值（仅当路径不存在时）
```sql
SELECT JSON_INSERT('{"a": 1}', '$.b', 2, '$.a', 99);
-- 结果: {"a": 1, "b": 2}  -- $.a 已存在，不更新
```
- JSON_REPLACE() - 替换值（仅当路径存在时）
```sql
SELECT JSON_REPLACE('{"a": 1, "b": 2}', '$.a', 10, '$.c', 3);
-- 结果: {"a": 10, "b": 2}  -- $.c 不存在，不插入
```
- JSON_REMOVE() - 删除值
```sql
SELECT JSON_REMOVE('{"a": 1, "b": 2, "c": 3}', '$.a', '$.c');
-- 结果: {"b": 2}

-- 删除数组元素
SELECT JSON_REMOVE('["a", "b", "c"]', '$[1]');
-- 结果: ["a", "c"]
```
## JSON信息函数
- JSON_TYPE() - 返回 JSON 值的类型
```sql
SELECT 
    JSON_TYPE('{"name": "John"}') as obj,        -- OBJECT
    JSON_TYPE('[1,2,3]') as arr,                -- ARRAY
    JSON_TYPE('"hello"') as str,                -- STRING
    JSON_TYPE('123') as num,                    -- INTEGER
    JSON_TYPE('true') as bool;                  -- BOOLEAN
```
- JSON_VALID() - 验证 JSON 是否有效
```sql
SELECT 
    JSON_VALID('{"name": "John"}') as valid,    -- 1
    JSON_VALID('{"name": "John"') as invalid;   -- 0
```
- JSON_LENGTH() - 返回 JSON 文档长度
```sql
SELECT 
    JSON_LENGTH('{"a": 1, "b": 2}') as obj_len,     -- 2
    JSON_LENGTH('["a", "b", "c"]') as arr_len,      -- 3
    JSON_LENGTH('"hello"') as str_len;              -- 1
```
- JSON_KEYS() - 返回对象的所有键
```sql
SELECT JSON_KEYS('{"a": 1, "b": 2, "c": 3}');
-- 结果: ["a", "b", "c"]

-- 获取嵌套对象的键
SELECT JSON_KEYS('{"user": {"name": "John", "age": 30}}', '$.user');
-- 结果: ["name", "age"]
```
- JSON_ARRAYAGG() - 将多行聚合为 JSON 数组
```sql
SELECT department, JSON_ARRAYAGG(employee_name) as employees
FROM employees
GROUP BY department;
```
- JSON_OBJECTAGG() - 将键值对聚合为 JSON 对象










|                       |          |
| --------------------- | -------- |
| `length(column_name)` | 字符长度 |
| GROUP_CONCAT(科目 ORDER BY 科目 SEPARATOR ', ')  | 行换字符串          |
| CASE WHEN condition THEN value1 ELSE value2 END  | 相当于if...else...语句         |

### date
#### 函数
```sql
-- 加法
SELECT DATE_ADD('2023-12-01', INTERVAL 7 DAY);
SELECT DATE_ADD(NOW(), INTERVAL 1 HOUR);

-- 减法
SELECT DATE_SUB('2023-12-01', INTERVAL 1 MONTH);
SELECT DATE_SUB(NOW(), INTERVAL 30 MINUTE);

-- 获取当前时间函数
SELECT NOW();           -- 2023-12-01 10:30:45
SELECT CURDATE();       -- 2023-12-01
SELECT CURTIME();       -- 10:30:45
SELECT UTC_DATE();      -- UTC日期
SELECT UTC_TIME();      -- UTC时间
SELECT UTC_TIMESTAMP(); -- UTC时间戳

-- 格式化日期
SELECT DATE_FORMAT(NOW(), '%Y-%m-%d %H:%i:%s');  -- 2023-12-01 10:30:45
SELECT DATE_FORMAT(NOW(), '%W, %M %e, %Y');     -- Friday, December 1, 2023

-- 解析日期字符串
SELECT STR_TO_DATE('2023-12-01', '%Y-%m-%d');

-- 日期差（天数）
SELECT DATEDIFF('2023-12-10', '2023-12-01'); -- 9

-- 时间差
SELECT TIMEDIFF('10:30:00', '09:15:00');     -- 01:15:00

-- 时间戳差值（秒数）
SELECT TIMESTAMPDIFF(DAY, '2023-12-01', '2023-12-10'); -- 9
SELECT TIMESTAMPDIFF(HOUR, '2023-12-01 10:00:00', '2023-12-01 12:00:00'); -- 2
SELECT TIMESTAMPDIFF(MONTH, '2023-01-01', '2023-12-01'); -- 11

-- 日期提取
SELECT YEAR(NOW());     -- 2023
SELECT MONTH(NOW());    -- 12
SELECT DAY(NOW());      -- 1
SELECT HOUR(NOW());     -- 10
SELECT MINUTE(NOW());   -- 30
SELECT SECOND(NOW());   -- 45
SELECT DAYOFWEEK(NOW()); -- 6 (星期几，1=周日)
SELECT DAYOFMONTH(NOW()); -- 1 (月中的第几天)
SELECT DAYOFYEAR(NOW()); -- 335 (年中的第几天)
SELECT WEEK(NOW());     -- 48 (周数)
SELECT QUARTER(NOW());  -- 4 (季度)
```

#### INTERVAL表达式
```sql
-- 语法结构
INTERVAL expr unit

-- 日期相关
INTERVAL 1 DAY          -- 天
INTERVAL 7 WEEK         -- 周
INTERVAL 1 MONTH        -- 月
INTERVAL 3 MONTH        -- 3个月
INTERVAL 1 YEAR         -- 年
INTERVAL 2 YEAR         -- 2年

-- 时间相关
INTERVAL 1 SECOND       -- 秒
INTERVAL 30 MINUTE      -- 分钟
INTERVAL 6 HOUR         -- 小时
INTERVAL 24 HOUR        -- 24小时

-- 复合时间
INTERVAL '1 12' DAY_HOUR        -- 1天12小时
INTERVAL '1:30:00' HOUR_SECOND  -- 1小时30分钟0秒
INTERVAL '1:30' MINUTE_SECOND   -- 1分30秒
INTERVAL '1 12:30:00' DAY_SECOND -- 1天12小时30分钟0秒
```

