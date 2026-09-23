# 概述
HTTP (明文传输) + SSL/TLS (加密) = HTTPS (安全传输)

### HTTP

​	HTTP（HyperText Transfer Protocol，超文本传输协议）是互联网上应用最为广泛的一种网络协议。 

- 核心特点 

  - **应用层协议**：用于Web浏览器和服务器之间的通信

  - **无状态**：每次请求都是独立的，服务器不保存客户端信息

  - **基于请求-响应模型**：客户端发送请求，服务器返回响应

  - **明文传输**：数据以文本形式传输（HTTP/2开始支持二进制）

### HTTPS协议概述 

​	**HTTPS**（HyperText Transfer Protocol Secure，安全超文本传输协议）是HTTP的安全版本。 

- 核心特点 

  - **HTTP + SSL/TLS**：在HTTP基础上加入SSL/TLS加密层

  - **数据加密**：传输内容被加密，防止窃听

  - **身份认证**：验证服务器身份，防止假冒

  - **数据完整性**：确保数据在传输过程中不被篡改

### 请求方法

| 请求方法     | 用途                       |
| ------------ | -------------------------- |
| 信息性状态码 |                            |
| `GET`        | 获取资源                   |
| `POST`       | 创建资源/提交数据          |
| `PUT`        | 更新或创建资源（完整替换） |
| `DELETE`     | 删除资源                   |
| `PATCH`      | 部分更新资源               |
| `HEAD`       | 获取响应头信息             |
| `OPTIONS`    | 获取服务器支持的HTTP方法   |
| `TRACE`      | 回显收到的请求（用于测试） |

### 状态码

| 状态码           | 说明                                         |
| ---------------- | -------------------------------------------- |
| 100              | Continue（继续）                             |
| 101              | Switching Protocols（切换协议）              |
| 成功状态码       |                                              |
| **200**          | **OK（请求成功）✅**                          |
| **201**          | **Created（创建成功）✅**                     |
| 202              | Accepted（已接受）                           |
| **204**          | **No Content（无内容）✅**                    |
| 客户端错误状态码 |                                              |
| **400**          | **Bad Request（请求错误）✅**                 |
| **401**          | **Unauthorized（未授权）✅**                  |
| **403**          | **Forbidden（禁止访问）✅**                   |
| **404**          | **Not Found（资源不存在）✅**                 |
| 405              | Method Not Allowed（方法不允许）             |
| 408              | Request Timeout（请求超时）                  |
| 409              | Conflict（冲突）                             |
| 413              | Payload Too Large（请求体过大）              |
| 422              | Unprocessable Entity（无法处理的实体）       |
| 429              | Too Many Requests（请求过多）                |
| 服务器错误状态码 |                                              |
| **500**          | **Internal Server Error（服务器内部错误）✅** |
| **502**          | **Bad Gateway（网关错误）✅**                 |
| **503**          | **Service Unavailable（服务不可用）✅**       |
| **504**          | **Gateway Timeout（网关超时）✅**             |


# TLS 握手协商
- 客户端发送 "Client Hello"
- 服务器回复 "Server Hello" + 证书
- 密钥交换和验证
- 生成会话密钥

# HTTP1.0和HTTP2.0的区别
特性  HTTP/1.0  HTTP/2.0
连接  短连接（默认） 长连接，多路复用
协议格式  文本  二进制
头部压缩  无 有（HPACK）
服务器推送 不支持 支持
流控制 无 有
优先级 无 有
安全性 不要求TLS  实际要求TLS