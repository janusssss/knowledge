# 字符串（String）
```bash
# 基本操作
SET name "John"
GET name
INCR counter
DECR counter
APPEND message "Hello"
```
- 二进制安全，最大 512MB
- 可存储字符串、整数、浮点数
- 支持原子操作（INCR/DECR）

# 列表（List）
```bash
# 基本操作
LPUSH mylist "item1"
RPUSH mylist "item2" 
LPOP mylist
RPOP mylist
LRANGE mylist 0 -1
```
- 双向链表
- 支持阻塞操作（BLPop/BRPop）
- 适合队列、栈场景

# 集合（Set）
```bash
# 基本操作
SADD myset "member1"
SADD myset "member2"
SMEMBERS myset
SISMEMBER myset "member1"
SUNION set1 set2
```
- 无序、不重复
- 支持集合运算（交集、并集、差集）
- 适合标签、好友关系

# 有序集合（Sorted Set / ZSet）
```bash
# 基本操作
ZADD myzset 1 "member1"
ZADD myzset 2 "member2"
ZRANGE myzset 0 -1 WITHSCORES
ZRANK myzset "member1"
ZREVRANGE myzset 0 10 WITHSCORES
```
- 有序、不重复
- 每个元素关联分数
- 适合排行榜、时间序列


# 哈希（Hash）
```bash
# 基本操作
HSET user:1000 name "John" age 30
HGET user:1000 name
HGETALL user:1000
HMGET user:1000 name age
HINCRBY user:1000 age 1
```
- 键值对集合
- 适合存储对象
- 内存效率高


