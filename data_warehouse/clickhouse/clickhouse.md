### ubuntu安装

```bash
# 下载密钥
sudo mkdir -p /usr/local/share/keyrings/
curl -fsSL https://repo.clickhouse.com/CLICKHOUSE-KEY.GPG | sudo gpg --dearmor -o /usr/local/share/keyrings/clickhouse-keyring.gpg
sudo nano /etc/apt/sources.list.d/clickhouse.list
deb [signed-by=/usr/local/share/keyrings/clickhouse-keyring.gpg] https://repo.clickhouse.com/deb/stable main/

sudo apt update
sudo apt install clickhouse-server clickhouse-client
```

```bash
/etc/clickhouse-server/users.xml
修改defalut用户权限
```

```xml
<default>
    <!-- 其他配置 -->
    <access_management>1</access_management>
</defau
```



### 小案例

```sql
CREATE TABLE animal (
    date Date DEFAULT toDate(now()),  -- 使用当前日期作为默认值
    name String,
    age Int32
) ENGINE = MergeTree(date, (name), 8192);

insert into animal(name, age) values('cat',10);
select * from animal;
```

### 创建用户

```sql
CREATE USER janus IDENTIFIED WITH sha256_password BY 'janus' HOST ANY;
grant all on *.* to 'janus';


CREATE USER janus 
IDENTIFIED WITH sha256_password BY 'janus' 
HOST ANY;
```

```b
clickhouse-client --port 9001 -d janus --user default --password
 
clickhouse-client -h 10.142.149.168 -d dim_db --user default --password

alter table dwm_eda_prd_offer_inst_month drop partition ('skafhsdkfhkjsd');
```

### 查看数据大小

```sql
select database,table,formatReadableSize(sum(bytes_on_disk)) from system.parts group by database,table order by sum(bytes_on_disk) desc limit 8;
```

