### 介绍

Kerberos 是一种网络认证协议，它使用票据（tickets）来允许节点在网络上相互通信时证明其身份。Kerberos 协议的设计目的是提供强认证，同时尽量减少对密码的直接传输，从而提高安全性。该协议最初由麻省理工学院（MIT）开发，并已成为许多操作系统和应用程序的标准安全功能之一。

- Kerberos 的主要特点包括：
  1. **基于票据的认证**：用户在登录时从认证服务器（通常是 KDC, Key Distribution Center）获取一个票据授予票据（TGT, Ticket-Granting Ticket）。之后，用户可以使用 TGT 来请求访问其他服务的票据（Service Tickets），而无需再次输入密码。
  2. **单点登录 (SSO)**：一旦用户通过 Kerberos 认证，他们就可以访问所有支持 Kerberos 的服务，而无需重复进行身份验证。
  3. **时间戳和加密**：为了防止重放攻击，Kerberos 使用时间戳和加密技术。每个票据都有一个有效期，并且所有的通信都是加密的。
  4. **密钥分发中心 (KDC)**：KDC 是 Kerberos 系统的核心组件，负责管理和分发票据。KDC 通常分为两个部分：认证服务器（AS, Authentication Server）和票据授予服务器（TGS, Ticket-Granting Server）。
  5. **跨域认证**：Kerberos 支持跨域认证，允许用户在一个域中认证后访问另一个域中的资源。

### 简单配置使用

#### KDC

1. 安装KDC

   ```bash
   # 系统版本：24.04.2 LTS (Noble Numbat)
   sudo apt install krb5-kdc krb5-admin-server
   ```

2. 配置`/etc/krb5kdc/kdc.conf`

   ```bash
   [kdcdefaults]
       kdc_ports = 750,88
   
   [realms]
       LOCALDOMAIN = {
           database_name = /var/lib/krb5kdc/principal
           admin_keytab = FILE:/etc/krb5kdc/kadm5.keytab
           acl_file = /etc/krb5kdc/kadm5.acl
           key_stash_file = /etc/krb5kdc/stash
           kdc_ports = 750,88
           max_life = 10h 0m 0s
           max_renewable_life = 7d 0h 0m 0s
           #master_key_type = aes256-cts
           #supported_enctypes = aes256-cts:normal aes128-cts:normal
           default_principal_flags = +preauth
       }
   ```

3. 配置`/etc/krb5.conf`

   ```bash
   [libdefaults]
           default_realm = LOCALDOMAIN
   
   # The following krb5.conf variables are only for MIT Kerberos.
           kdc_timesync = 1
           ccache_type = 4
           forwardable = true
           proxiable = true
           rdns = false
   
   
   # The following libdefaults parameters are only for Heimdal Kerberos.
           fcc-mit-ticketflags = true
   
   [realms]
           LOCALDOMAIN = {
                   kdc = janus
                   admin_server = janus
           }
        
   [domain_realm]
           .localdomain = LOCALDOMAIN
           localdomain = LOCALDOMAIN
   ```

4.  创建/初始化Kerberos database

   ```bash
   kdb5_util create -s -r LOCALDOMAIN
   
   kadmin.local -q "addprinc admin/admin" # 新增用户
   kadmin.local # 密码
   ```

5. 设置ACL权限，`/etc/krb5kdc/kadm5.acl`

   ```bash
   */admin@LOCALDOMAIN    * 
   ```

6. 启动

   ```bash
   systemctl start krb5-kdc
   systemctl start krb5-admin-server
   ```

#### Client

1. 安装

   ```bash
   # 系统版本：
   apt install krb5-user libkrb5support0 libpam-krb5
   ```

2. 配置`/etc/krb5.conf`与KDC一致

3. 认证

   ```bash
   kinit admin/admin@LOCALDOMAIN
   
   klist # 看结果
   ```

   


### hive使用

```bash
# 使用 beeline 连接
beeline -u "jdbc:hive2://172.20.70.253:10000/default;principal=admin/admin@LOCALDOMAIN"
```



