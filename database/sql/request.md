# 内联和外联的区别

SQL 的 `JOIN` 语法是关系型数据库最核心的操作之一，用于根据两个（或多个）表之间的共同字段，把它们的行组合起来。


| 连接类型 | 关键字 | 返回结果 |
| :--- | :--- | :--- |
| **内连接** | `INNER JOIN` | 只返回**两表都匹配**的行（交集）。不匹配则丢弃。 |
| **左外连接** | `LEFT JOIN` | 返回**左表全部** + 右表匹配的行。右表无匹配则用 `NULL` 填充。 |
| **右外连接** | `RIGHT JOIN` | 返回**右表全部** + 左表匹配的行。左表无匹配则用 `NULL` 填充。 |
| **全外连接** | `FULL JOIN` | 返回**两表全部**行。无论是否匹配，缺失侧用 `NULL` 填充。<br>（MySQL 不支持，需用 `LEFT UNION RIGHT` 模拟） |
| **交叉连接** | `CROSS JOIN` | 返回**笛卡尔积**（左表行数 × 右表行数）。极少直接使用，配 `WHERE` 可模拟内连接。 |

# union
| 操作 | 作用 | 方向 |
| :--- | :--- | :--- |
| **JOIN** | **水平拼接**：把两个表的**列**组合在一起（左右扩展） | 横向 |
| **UNION** | **垂直拼接**：把两个查询的**行**堆叠在一起（上下合并） | 纵向 |

**核心要求**：  
多个 `SELECT` 语句的**列数必须相同**，且**对应列的数据类型必须兼容**。

| 关键字 | 行为 | 性能 |
| :--- | :--- | :--- |
| `UNION` | **去重合并**：自动去除重复行（相当于 `SELECT DISTINCT`） | 较慢（需排序去重） |
| `UNION ALL` | **全部合并**：保留所有行，包括重复行 | 更快（直接追加） |


# 千万级别（10,000,000+）的大表增加索引

不能直接在生产库执行 `CREATE INDEX`，否则会锁表数小时，导致服务宕机。

| 策略 | 适用场景 | 锁表影响 | 风险 |
| :--- | :--- | :--- | :--- |
| **① pt-online-schema-change** | 所有场景（首选） | 几乎无锁（闪锁 < 1秒） | 需安装 Percona 工具包 |
| **② 影子表（手动重建）** | 无法安装第三方工具时 | 闪锁（切换瞬间） | 需额外存储空间（双份表） |
| **③ 直接 ALTER（业务低峰期）** | 库很小（< 500万）或能接受停服 | **全程锁表（MDL 锁）** | 高危，千万级可能锁 30分钟+ |


### 🥇 方案一：pt-online-schema-change（生产环境标准）

这是 MySQL 官方推荐的 Percona 工具，原理是**创建影子表 → 复制数据 → 建索引 → 切换表名**，过程中通过触发器同步增量数据。

**安装**（CentOS 示例）：
```bash
yum install percona-toolkit -y
```

**执行命令**（关键参数说明）：
```bash
pt-online-schema-change \
--alter "ADD INDEX idx_name (column_name)" \   # 要执行的 DDL
--execute \                                     # 真正执行（不加则只检查）
--host=localhost --user=root --password=xxx \
D=your_database,t=your_table \
--chunk-size=10000 \                            # 每次复制 1 万行（防负载过高）
--max-load="Threads_running=50" \              # 系统负载超 50 个线程则暂停
--critical-load="Threads_running=200" \        # 超 200 则终止
--set-vars="lock_wait_timeout=60"              # 防止触发器等待超时
```

> **监控进度**：执行时会显示 `Copying rows` 进度百分比，千万级表通常需要 **10~60 分钟**（取决于服务器 IO）。

---

### 🥈 方案二：影子表手动重建（无需安装工具）

适合 DBA 手动操作，逻辑与 `pt` 工具类似。

```sql
-- 1. 创建结构相同的新表（含新索引）
CREATE TABLE your_table_new LIKE your_table;
ALTER TABLE your_table_new ADD INDEX idx_name (column_name);

-- 2. 分批拷贝数据（千万级务必分批！）
INSERT INTO your_table_new SELECT * FROM your_table 
WHERE id BETWEEN 1 AND 100000;   -- 循环执行，直到全部拷贝

-- 3. 交换表名（闪锁，影响 < 1秒）
RENAME TABLE your_table TO your_table_old, 
             your_table_new TO your_table;

-- 4. 验证后删除旧表（建议等 1-2 天再删）
DROP TABLE your_table_old;
```

> ⚠️ **注意**：拷贝期间新数据会丢失，需配合**触发器**或**写暂停**。所以一般只在凌晨业务低峰期操作。

---

### 🥉 方案三：直接 ALTER（仅限紧急情况）

**不推荐**，但如果你实在没办法，必须这样做：

```sql
-- 设置超长等待时间，避免锁等待超时导致业务报错
SET SESSION lock_wait_timeout = 3600;

-- 执行（会锁表）
ALTER TABLE your_table ADD INDEX idx_name (column_name);
```

- **锁表影响**：`ALTER` 执行期间，所有对该表的 `SELECT`/`INSERT`/`UPDATE`/`DELETE` 都会被阻塞。
- **预估时间**：千万级 InnoDB 表建索引，大约需要 **（表大小 / 磁盘 IO 速度）**，比如 50GB 表在普通 SSD 约 20-40 分钟。

