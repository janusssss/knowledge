### 结果说明

| 列名              | 说明           | 详细解释                                   |
| ----------------- | -------------- | ------------------------------------------ |
| **id**            | 查询序列号     | 数字越大优先级越高，相同数字表示同一组操作 |
| **select_type**   | 查询类型       | SIMPLE、PRIMARY、SUBQUERY等                |
| **table**         | 表名           | 操作的表名                                 |
| **partitions**    | 匹配的分区     | 分区表使用                                 |
| **type**          | 连接类型       | 性能关键指标                               |
| **possible_keys** | 可能使用的索引 | 理论上可以使用的索引                       |
| **key**           | 实际使用的索引 | 真正使用的索引                             |
| **key_len**       | 索引长度       | 使用的索引字节数                           |
| **ref**           | 索引比较的列   | 与索引比较的值来源                         |
| **rows**          | 扫描行数       | 估算需要扫描的行数                         |
| **filtered**      | 过滤百分比     | 通过条件过滤的行数百分比                   |
| **Extra**         | 额外信息       | 重要的优化信息                             |

### type

| value           | 说明                                    |
| --------------- | --------------------------------------- |
| system/const    | 最优，相当于常量                        |
| eq_ref          | `equal reference`主键或唯一索引等值查询 |
| ref             | 非唯一索引等值查询                      |
| ref_or_null     | 索引查找包含null                        |
| index_merge     | 索引合并                                |
| unique_subquery | in子查询使用唯一索引                    |
| index_subquery  | in子查询使用普通索引                    |
| range           | 范围查询                                |
| index           | 索引全扫描                              |
| all             | 全表扫描                                |

### extra

| value                      | 说明                             |
| -------------------------- | -------------------------------- |
| `Using index`              | 覆盖索引，值扫描索引             |
| `Using index condition`    | 索引条件下推                     |
| `Using where`              | 使用where过滤                    |
|                            |                                  |
| `Using temporary`          | 使用临时表`group by` `order by ` |
| `Using filesort`           | 文件排序，`order by`没有使用索引 |
| `Using where; Using index` | 索引过滤后还需要where过滤        |

