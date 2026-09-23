### 概念

​	RESTful是一种软件架构风格，核心是面向资源的设计思想，将一切抽象为资源，通过标准的HTTP方法对资源进行操作。

### 设计原则

1. 资源命令规范
   - 使用名词复数形式：`/api/users`而不是`/api/user`
   - URL层次清晰：`/api/users/{userId}/orders`表示用户下的订单
   - 使用小写字符和连字符：`/api/user-profils`

2. HTTP方法语义化

   - GET	`/api/users`	查询用户列表
   - GET	`/api/users/123`	获取特定用户
   - POST	`/api/users`	创建新用户
   - PUT	`/apii/users/123`	完整更新用户
   - PATHC	`/api/users/123`	部分更新用户
   - DELETE	`/api/users/123`	删除用户

3. 状态吗规范

   1. 成功状态码

      - 200	查询成功

      - 201	创建成功

      - 204	删除成功，无返回内容

   2. 客户端错误

      - 400	请求参数错误

      - 401	未认证

      - 403	无权限

      - 404	资源不存在

   3. 服务端错误

      - 500	服务器内部错误