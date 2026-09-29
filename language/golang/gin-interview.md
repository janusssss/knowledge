# Go + Gin 面试常问问题

> 说明：Gin 在面试里很少被当成「框架用法题」考，绝大多数问题都会拐到 **HTTP 服务怎么在生产里活下来**：路由怎么匹配、中间件为什么能 `Next`、Context 为什么不能跨协程、参数校验怎么写、服务怎么优雅退出。下面按这个思路组织，源码级细节也标出来了。

---

## 一、基础与设计

1. **Gin 是什么，和 `net/http` 什么关系？**
   Gin 只是一个 **`http.Handler` 的实现**：`gin.Engine` 实现了 `ServeHTTP(w, r)`，所以 `r.Run(":8080")` 本质上就是 `http.ListenAndServe(":8080", engine)`。它没有替换标准库，只是包了一层：路由用自研的基数树、请求上下文用池化复用的 `*gin.Context`、外加参数绑定/校验/渲染/中间件这些标准库不给的糖。

   推论：Gin 可以和标准库共存——`http.Handle("/legacy", oldHandler)` 再把 gin 挂在别的路径上完全可行；反过来也能把 `net/http` 的中间件（比如 `http.StripPrefix`）包在 gin 的 handler 外面。

2. **为什么 Gin 比 `net/http` + 手写 mux 快？**
   - **路由**：自研压缩前缀树（radix tree），匹配是 O(路径长度)，不依赖 map 查找和反射
   - **Context 池化**：每个请求不是 new 一个 Context，而是 `sync.Pool` 取，用完 `reset()` 归还。这是它比 Echo 早期版本更省 GC 压力的主因
   - **零反射的参数绑定路径**：`ShouldBind` 只在有 tag 时才走反射，纯 `c.Query` 完全不反射
   - **`c.Writer` 是接口包装**，避免了 `http.ResponseWriter` 的 type assertion 热路径

   但要注意：**框架本身通常不是瓶颈**，真实业务里 JSON 序列化、DB、下游 RPC 才是。面试里说清这点比背 benchmark 数字加分。

3. **`gin.Default()` 和 `gin.New()` 的区别？**
   `New()` 返回一个没有任何中间件的 Engine；`Default()` = `New()` + `Logger()` + `Recovery()`。生产里更常见的是用 `New()` 再自己挂 zap/slog 的 logger 中间件，因为 `gin.Logger()` 是同步写 `io.Writer`，高 QPS 下会成为锁竞争点。

4. **`gin.H` 是什么？`gin.Context` 和标准库 `context.Context` 是一回事吗？**
   - `gin.H` 就是 `map[string]any` 的别名，纯粹是省字
   - **不是一回事**。`gin.Context` 是请求作用域的数据袋（参数、Writer、中间件链、`Set/Get`），只在单个请求生命周期内有效；`context.Context` 是跨调用边界的取消/超时/传值机制。两者通过 `c.Request.Context()` 桥接——所以下游 DB 调用应该用 `c.Request.Context()`，而不是 `c` 本身

5. **`Engine` 和 `RouterGroup` 的关系？**
   `Engine` 内嵌 `RouterGroup`，`Group()` 返回的是**值拷贝**出来的新 RouterGroup（共享同一份 `trees` 和 `pool`），只是把 `basePath` 和 `Handlers` 组合进去。所以：
   - 分组里的中间件是**追加**到副本上的，不会污染父分组
   - `engine.Group()` 之后再 `engine.GET()` 也不会带上分组的中间件
   - 这个「值语义」是很多人被坑的地方：把 RouterGroup 当指针传来传去会得到意料之外的链

---

## 二、路由

6. **Gin 的路由树是怎么工作的？**
   每个 HTTP method 一棵独立的树（`methodTrees`），节点是**压缩前缀树**：只存能区分开的最短前缀，公共前缀共享一个节点。
   ```
   /user/list   /user/login   /user/:id
   →  /user/   ├─ li ── st
               │        └─ ogin
               └─ :id   (param 节点)
   ```
   匹配优先级：**静态 > 参数 `:name` > 通配 `*filepath`**。所以 `/user/new` 和 `/user/:id` 可以共存，访问 `/user/new` 命中静态路由。

7. **`/user/:id` 和 `/user/*action` 有什么区别？能同时注册吗？**
   - `:id` 是**单个路径段**的参数，不跨 `/`
   - `*action` 是 catch-all，吃掉剩下所有内容（包含 `/`），只能出现在路径末尾
   - 同一个前缀下**不能**同时有 `:id` 和 `*action` 冲突的分支，注册时会 `panic: wildcard route conflicts`
   - 参数冲突也会 panic，比如同时注册 `/user/:id` 和 `/user/:name`（同名位置不同参数名）→ `panic: ':name' in new path conflicts with existing wildcard ':id'`

   面试延伸：**这种冲突是启动期 panic，属于好事**——比运行时才发现路由没生效强。

8. **怎么拿到路径参数？`c.Param` 有性能问题吗？**
   ```go
   r.GET("/user/:id", func(c *gin.Context) {
       id := c.Param("id")   // 线性扫描 Params 切片
   })
   ```
   `Params` 是个 `[]Param` 切片，`Param()` 是 O(n) 遍历，但 n 就是个位数，可以忽略。真正要注意的是 **`Params` 会被 Context 复用**——不要在 goroutine 里持有 `c.Param(...)` 切片本身（`c.Params` 是切片，取字符串值是安全的）。

9. **`NoRoute` / `NoMethod` 怎么用？**
   ```go
   r.NoRoute(func(c *gin.Context) { c.JSON(404, gin.H{"msg": "not found"}) })
   r.HandleMethodNotAllowed = true            // 默认 false，不匹配 method 时走 NoRoute
   r.NoMethod(func(c *gin.Context) { c.JSON(405, ...) })
   ```
   想把 404 统一成业务 JSON、想返回 405 而不是 404，就得配这两个。

10. **路由注册的最佳实践？**
    - 按业务拆 `RegisterRoutes(rg *gin.RouterGroup)` 函数，一个模块一个文件，避免 `main.go` 里几百行
    - 路径参数用 `c.ShouldBindUri` 而不是手写 `strconv.Atoi(c.Param("id"))`
    - 版本前缀用 `r.Group("/api/v1")`
    - **不要在 handler 里注册路由**（`r.GET` 不是并发安全的）

---

## 三、中间件（最核心的一节）

11. **中间件的本质是什么？**
    Handler 链就是一个 `[]HandlerFunc`（`gin.HandlersChain`），`Context.index` 指向当前执行到哪。`c.Next()` 就是：
    ```go
    func (c *Context) Next() {
        c.index++
        for c.index < int8(len(c.handlers)) {
            c.handlers[c.index](c)
            c.index++
        }
    }
    ```
    Engine 收到请求时把「全局中间件 + 分组中间件 + 路由 handler」拼成一条链，从 `index=0` 开始跑。所以**中间件、handler 全是同一个类型，没有特权**。

    注意 `c.index` 是 **`int8`**：一条链最多 63 个 handler，超了会 panic `too many handlers`。深层嵌套分组 + 每层挂中间件时有概率撞上。

12. **`c.Next()` 和 `c.Abort()` 的区别？是怎么「包住」后续逻辑的？**
    - `Next()` 会在一行代码里**把后面的整条链跑完**再返回，所以你在 `Next()` 前写的逻辑 = 请求前处理，后面写的 = 响应后处理，天然的洋葱模型
    ```go
    func Timing() gin.HandlerFunc {
        return func(c *gin.Context) {
            start := time.Now()
            c.Next()                      // 链在这里往下走到底
            log.Printf("cost=%v", time.Since(start))
        }
    }
    ```
    - `Abort()` 把 `index` 设成常量 `abortIndex = 63`，循环条件立刻不成立，**后续 handler 全部跳过**，但当前函数剩余代码还会执行（因为 `Abort` 不是 panic）
    - 鉴权中间件里标配：`c.AbortWithStatusJSON(401, ...)` = `Abort()` + 写响应

13. **不调用 `c.Next()` 会怎样？**
    中间件函数返回后，外层 `Next()` 的 for 循环会继续 `index++` 往下执行——**所以「不调 Next」并不会阻断链**，只是你放弃了在链中间插入自己后半段逻辑的位置。想阻断只能 `Abort()`。这是最经典的面试陷阱题。

14. **多个中间件的执行顺序是怎样的？**
    注册顺序 = 链顺序 = 洋葱由外到内：
    ```
    A前 → B前 → handler → B后 → A后
    ```
    全局中间件在最外层，分组中间件次之，路由级最后。所以：
    - `Recovery` 必须注册在最外层（`Default()` 的顺序是 Logger → Recovery，但两者都受 panic 影响，自研时要最先挂 Recovery）
    - 需要「响应之后」统计耗时的中间件，必须放在业务之前注册

15. **中间件里怎么在 handler 之间传值？**
    `c.Set(key, value)` / `c.Get(key)` / `c.MustGet(key)`，底层是一个 `map[string]any`（gin 1.10 起改成了小优化过的 map）。
    - key 不要用裸字符串，定义常量或自定义类型避免撞车
    - `MustGet` 取不到会 **panic**，会被 Recovery 捕获成 500，慎用
    - **不要用它传大对象**，`c.Keys` 在整个请求生命周期内存活

16. **为什么中间件里起 goroutine 会出问题？`c.Copy()` 是什么？**
    因为 `*gin.Context` 是**池化复用**的：handler 返回后 Engine 会 `reset()` 它再给下一个请求用（gin 1.11+ 还直接 Put 回 pool）。你在 goroutine 里拿着原始 Context 读写，会读到**别的请求的数据**，或者让响应写到错的连接上。

    正确写法：
    ```go
    func Async() gin.HandlerFunc {
        return func(c *gin.Context) {
            cc := c.Copy()               // 拷贝出只读副本：Request/Writer/Keys 复制，handlers 置空
            go func() {
                time.Sleep(time.Second)
                log.Println(cc.Request.URL.Path)   // 只能用副本
            }()
        }
    }
    ```
    注意 `Copy()` 出来的是**只读**的：不能再用它写响应（响应早就返回了），只适合记录日志、埋点、异步上报。异步写响应必须用 `c.Stream` 或者自己持有 `http.ResponseWriter`——但那是另一种设计（长连接/SSE）。

17. **中间件怎么写才「生产级」？**
    - **Recovery**：`gin.CustomRecovery` 把 panic 转成 500 并打日志 + 上报（`gin.Recovery()` 默认只打栈，会打爆日志盘）
    - **RequestID**：从 Header 取或生成 UUID，`c.Set` 后塞进日志字段
    - **Logger**：zap/slog + 结构化字段（method、path、status、cost、requestID、clientIP）
    - **CORS**：`github.com/gin-contrib/cors`，注意 `AllowCredentials` 不能配 `AllowAllOrigins`
    - **认证**：`gin-contrib/jwt` 或自己写，校验后 `c.Set("uid", ...)`
    - **限流**：`golang.org/x/time/rate` 令牌桶，key 用 uid 或 IP
    - **超时**：`context.WithTimeout` 注入 `c.Request`，或在 handler 里对下游调用设超时（gin 本身**不提供** http.TimeoutHandler 那种机制）
    - **gzip**：`gin-contrib/gzip`

---

## 四、参数绑定与校验

18. **`Bind` 和 `ShouldBind` 系列怎么选？**
    | 方法 | 出错行为 | 什么时候用 |
    |---|---|---|
    | `c.Bind*` / `MustBindWith` | 自动 400 + `Abort` | 懒，但响应格式不受控 |
    | `c.ShouldBind*` | 只返回 error，自己处理 | **推荐**，能返回统一错误结构 |
    | `c.ShouldBindBodyWith` | 同 ShouldBind，但会把 body 缓存 | 需要读两次 body（先校验后落库） |

    `Bind` 内部就是 `ShouldBind` + 失败时 `AbortWithError(400)`，所以业务要统一响应格式时一律用 `ShouldBind`。

19. **`ShouldBind` 怎么知道该用 JSON 还是 form？**
    它看 `Content-Type`：`application/json` → JSON 绑定器；`application/x-www-form-urlencoded` / `multipart/form-data` → Form；找不到就按 query 处理。所以 `ShouldBind` 是「自动挡」，需要精确控制时用 `ShouldBindJSON` / `ShouldBindQuery` / `ShouldBindUri`。

20. **为什么 body 只能读一次？怎么读两次？**
    `c.Request.Body` 是流，读完就到底了。`ShouldBindJSON` 之后 body 已经空了，再 `ShouldBindBodyWith` 会失败。
    ```go
    // 可以绑定多次，body 会被缓存进 gin 内部的 bodyBytesKey
    c.ShouldBindBodyWith(&a, binding.JSON)
    c.ShouldBindBodyWith(&b, binding.JSON)
    ```
    代价是多一次内存拷贝，大 body 慎用。另一种做法是在中间件里 `io.ReadAll` 后 `io.NopCloser(bytes.NewReader(b))` 换回去。

21. **`binding` tag 常见用法有哪些？**
    ```go
    type CreateUserReq struct {
        Name   string `json:"name"   binding:"required,min=2,max=20"`
        Email  string `json:"email"  binding:"required,email"`
        Age    int    `json:"age"    binding:"gte=0,lte=150"`
        Role   string `json:"role"   binding:"oneof=admin user guest"`
        Tags   []string `json:"tags" binding:"max=5,dive,min=1"`  // dive 深入切片元素
        Phone  string `json:"phone"  binding:"omitempty,len=11"`
        Inner  *Sub   `json:"inner"  binding:"required"`
    }
    ```
    - `required` 对数字类型是「非零值」，`0` 会被判定为缺失 —— 想区分「没传」和「传了 0」要用指针
    - 字段名匹配优先看 `json` tag（`binding` 与 `json` 名字不一致时解析会按 `json` 名报错）
    - 嵌套结构体默认递归校验，切片要 `dive`

22. **怎么注册自定义校验规则？**
    validator 的 `binding.Validator` 是**全局单例**，注册一次全局生效：
    ```go
    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        _ = v.RegisterValidation("alnum", func(fl validator.FieldLevel) bool {
            return regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(fl.Field().String())
        })
    }
    ```
    改成中文报错则注册 `universal-translator` + `zh_translations`，再把 `validator.ValidationErrors` 转成业务文案（推荐在统一错误处理里做）。

23. **校验失败的错误怎么统一返回？**
    ```go
    if err := c.ShouldBindJSON(&req); err != nil {
        var ve validator.ValidationErrors
        if errors.As(err, &ve) {
            // 逐字段翻译成 "name: 长度需在 2-20 之间"
        }
        c.JSON(400, ErrResp(err))
        return
    }
    ```
    千万不要把 validator 的原始英文错误直接吐给前端。

---

## 五、响应与渲染

24. **`c.JSON` 有什么坑？**
    - **HTML 转义**：Gin 默认 `SetEscapeHTML(true)`，`<`、`>`、`&` 会被转成 `\u003c`。前端拿到一堆 Unicode 转义就是这里。关掉：`r.SetEscapeHTML(false)` 或用 `c.PureJSON`
    - **写 header 的时机**：第一次 `c.JSON` 时状态码和 header 就发出去了，之后改 `c.Status()` 无效。**「中途 write 过再返回 500」是经典 bug**
    - 想统一响应结构（`{code,msg,data}`），最好封一层 `response.OK(c, data)`，避免每个 handler 手写

25. **`c.JSON` / `PureJSON` / `AsciiJSON` / `SecureJSON` / `JSONP` 的区别？**
    - `JSON`：默认转义 HTML
    - `PureJSON`：不转义，用于返回带 HTML/URL 的字段
    - `AsciiJSON`：所有非 ASCII 转成 `\uXXXX`，用于只接受 ASCII 的老客户端
    - `SecureJSON`：给顶层是数组的响应加 `while(1);` 前缀，防 JSON 劫持（老方案，现在靠 `X-Content-Type-Options: nosniff` 和 CORS）
    - `JSONP`：带 callback 的 JSON，仅供跨域兼容

26. **怎么返回流式数据？**
    ```go
    // SSE，长连接推送
    c.Stream(func(w io.Writer) bool {
        c.SSEvent("message", gin.H{"t": time.Now()})
        time.Sleep(time.Second)
        return true          // false 则结束
    })
    // 直接把 Reader 流给客户端（大文件下载）
    c.DataFromReader(200, size, "application/octet-stream", reader, nil)
    ```
    SSE 注意：客户端断开时 `c.Request.Context().Done()` 会关闭，`Stream` 内部靠它判断退出；Nginx 要关 buffering（`X-Accel-Buffering: no`）。

27. **静态文件和模板怎么配？**
    - `r.Static("/static", "./static")`、`r.StaticFile("/favicon.ico", ...)`、`r.StaticFS("/assets", http.FS(embedFS))`
    - 用 `embed.FS` 打包前端产物是目前单二进制部署的常见做法
    - 模板：`r.LoadHTMLGlob("templates/*")`，**glob 是启动时解析、生产环境不能热更新**；用 `SetHTMLTemplate(parseFS(*embed.FS, ...))` 可以彻底避免运行时读盘

---

## 六、错误处理与 panic

28. **`c.Error(err)` 是干嘛的？**
    把错误挂到 `c.Errors`（一条错误链）上，不中断流程，最后由 `r.Use(gin.ErrorLogger())` 或自定义 `ErrorHandler` 统一处理。实践中很多人干脆不用它，直接返回 error：
    ```go
    func wrap(fn func(*gin.Context) error) gin.HandlerFunc {
        return func(c *gin.Context) {
            if err := fn(c); err != nil { respondErr(c, err) }
        }
    }
    ```
    这样 handler 有真正的 error 返回值，业务错误也能用 `errors.Is/As` 判断——比 `c.Error` 更符合 Go 习惯。

29. **Recovery 中间件怎么工作？**
    `defer recover()`，捕获后：`c.AbortWithStatus(500)`（默认）或自定义响应 + 打印栈。要点：
    - 只在**被引用**的 goroutine 里有效，业务自己起的 goroutine panic 照样让整个进程挂掉，必须自己 recover
    - `Recovery` 默认不区分 broken pipe / connection reset 这类「客户端断开」的错误，会刷一堆无意义日志（`gin` 内部对 `ErrBrokenPipe` 有特殊处理，自研时要自己判 `errors.Is(err, syscall.EPIPE)`）
    - `RecoveryWithWriter` 可以指定输出；生产建议同时上报到日志系统 + 报警

30. **handler 里怎么正确返回业务错误？**
    套路是「定义业务错误码 + 统一 resp 包」：
    ```go
    type BizErr struct{ Code int; Msg string }
    func (e *BizErr) Error() string { return e.Msg }

    func Handle(fn func(*gin.Context) error) gin.HandlerFunc {
        return func(c *gin.Context) {
            err := fn(c)
            if err == nil { return }
            var be *BizErr
            if errors.As(err, &be) {
                c.JSON(200, gin.H{"code": be.Code, "msg": be.Msg})
                return
            }
            c.JSON(500, gin.H{"code": 500, "msg": "internal error"})
        }
    }
    ```
    注意 500 的详情**不能带原始 err**（可能含 SQL、内部路径），要打日志不带响应。

---

## 七、并发与性能

31. **`sync.Pool` 在 Gin 里怎么用的？**
    Engine 里有个 `pool sync.Pool`，`ServeHTTP` 时 `Get` 一个 Context、`reset()` 后复用，请求结束再 `Put`。Context 里持有 `Writer`、`Params`、`Keys` 等，复用能显著减少 GC 压力。

    对开发的约束（面试常考）：**任何跨越请求生命周期的引用都必须是副本**——`c.Copy()`、`c.Param()` 取出的 string、`c.GetString()` 都是安全的；`c.Params` 切片本身、`c` 指针不是。

32. **Gin 服务的性能瓶颈一般在哪？怎么排查？**
    - 框架层：路由匹配基本不是瓶颈，除非注册了上万个路由（树深了有影响）
    - 常见真实瓶颈：JSON 序列化（`encoding/json` 反射慢 → 换 `json-iterator` / `sonic` / `bytedance/sonic`）、DB 慢查询与连接池（`sql.DB.SetMaxOpenConns` 配小了会排队）、下游 RPC 超时、日志同步写盘、锁竞争
    - 手段：`net/http/pprof` 挂上去（`r.GET("/debug/pprof/*pprof", gin.WrapH(http.DefaultServeMux))`）、`go tool pprof` 看 CPU/heap/block/mutex，`go test -bench` 压单个 handler，`wrk`/`vegeta` 压整体
    - 别忘了 `GOMAXPROCS` 与容器 CPU 限额：Go 1.25 起才自动识别 cgroup 配额，老版本要用 `automaxprocs`

33. **`c.Request.Context()` 在下游调用里怎么用？**
    ```go
    ctx, cancel := context.WithTimeout(c.Request.Context(), 800*time.Millisecond)
    defer cancel()
    row := db.QueryRowContext(ctx, "select ...")
    ```
    好处：客户端断连时 `c.Request.Context()` 会 cancel，DB/Redis 调用会立刻中断，不用继续做无用功。**这是面试官最爱问的「上下文怎么贯穿」的答案**。

    注意：gin 不自动给整个 handler 加超时，要的是中间件里 `context.WithTimeout` 替换 `c.Request`（`c.Request = c.Request.WithContext(ctx)`）。

34. **Gin 支持 HTTP/2 吗？**
    `r.RunTLS` 会走标准库的 TLS 配置，只要证书是 HTTP/2 允许的就自动协商（需要 `NextProtos`，`RunTLS` 内部已处理）；明文 h2c 要自己用 `golang.org/x/net/http2/h2c` 包一层：`http.Server{Handler: h2c.NewHandler(r, &http2.Server{})}`。`http.Pusher` 通过 `c.Writer.Pusher()` 拿。

---

## 八、线上部署与可靠性（高频）

35. **`r.Run()` 为什么不能用在生产？**
    它内部是 `http.ListenAndServe`，创建的 `http.Server` **没有任何超时**：没有 `ReadTimeout`/`ReadHeaderTimeout`/`WriteTimeout`/`IdleTimeout`。慢客户端（slowloris）能一直占着连接不放手，几万个这种连接就能把服务器拖死。

    生产写法：
    ```go
    srv := &http.Server{
        Addr:              ":8080",
        Handler:           r,
        ReadHeaderTimeout: 5 * time.Second,
        ReadTimeout:       15 * time.Second,
        WriteTimeout:      30 * time.Second,
        IdleTimeout:       60 * time.Second,
    }
    ```

36. **怎么优雅关闭？**
    ```go
    go func() {
        if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
            log.Fatal(err)
        }
    }()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit                                     // 收到信号后

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {  // 停止接收新连接，等存量请求跑完
        _ = srv.Close()                        // 超时强杀
    }
    db.Close()                                 // 再关资源
    ```
    要点：`Shutdown` **不等待** hijacked 连接（WebSocket）和后台 goroutine，那些要自己用 `WaitGroup` 管；K8s 里还要先 `preStop` sleep 几秒，等 Endpoints 摘除生效再退，否则会丢请求。

37. **`ClientIP()` 为什么要配 `SetTrustedProxies`？**
    Gin 1.7 之后默认信任**所有**代理，意味着客户端可以随便伪造 `X-Forwarded-For` 来伪装 IP —— 限流、风控、审计全废。
    ```go
    _ = r.SetTrustedProxies([]string{"10.0.0.0/8"})   // 只信内网 LB
    // 或者 r.SetTrustedProxies(nil) 表示完全不信任，直接用 RemoteAddr
    ```
    这是很容易被面试官挑出来的**安全考点**。

38. **如何在同一个进程里跑多个端口的服务？**
    Gin 只负责 HTTP handler，多端口就是起多个 `http.Server` 或者用 `SO_REUSEPORT` 的库；普通做法是：
    - 业务端口 + 管理端口（`/metrics`、`/healthz`、`/debug/pprof`）分开，管理端口只绑 127.0.0.1，避免暴露
    - 健康检查别用 `/`，用 `/healthz` 只检查自身存活、`/readyz` 检查依赖（DB/Redis）就绪

39. **限流怎么做？**
    ```go
    limiter := rate.NewLimiter(rate.Limit(100), 200)   // 100 QPS，桶容量 200
    func RateLimit() gin.HandlerFunc {
        return func(c *gin.Context) {
            if !limiter.Allow() {
                c.AbortWithStatusJSON(429, gin.H{"msg": "too many requests"})
                return
            }
        }
    }
    ```
    单机令牌桶只能保护自己；全局限流要放 Redis（滑动窗口 / 令牌桶 Lua 脚本），或者交给网关（Kong/APISIX/Envoy）。注意区分**限流**（保护自己）和**熔断**（保护下游）。

---

## 九、源码与细节陷阱

40. **Context 池化的 `reset()` 都重置了什么？**
    请求结束后，Engine 会经 `reset()` 复用 Context：把 `Params` 截断为 0 长度、`Keys` 里的 map 清空、`handlers` 置 nil、`index` 归 -1、`Errors` 清空等。**这就是为什么复用是安全的，也是为什么跨请求持有引用会串数据。**

41. **`HandlersChain` 的 `int8` 上限为什么是 63？**
    `abortIndex = math.MaxInt8 / 2 = 63`，即用 63 表示「已终止」这一个哨兵值，所以真正可用的中间件/处理器最多 63 个。超过就 `panic("too many handlers")`。面试里问这个的人是在验证你有没有读过源码。

42. **`any` 路由和 `Handle` 的区别？**
    `r.Any("/x", h)` 会为**所有已注册 HTTP Method** 逐个注册（gin 1.9+ 扩展到所有方法），`r.Handle("CUSTOM", "/x", h)` 能注册任意方法名。CORS 预检要自己处理 `OPTIONS`：`r.OPTIONS` 或用 cors 中间件。

43. **路径尾斜杠（trailing slash）怎么处理？**
    Gin 默认开启 `RedirectTrailingSlash`：请求 `/foo/` 而只注册了 `/foo`，会 301 重定向。但 **POST 遇到 301 变 GET**（浏览器行为），前端调接口遇到尾斜杠问题常常就出在这里。可以通过 `r.RedirectTrailingSlash = false` 关掉自己控制。同时 `RedirectFixedPath` 会做大小写、`../` 修正，也是默认开。

44. **`c.Query` / `c.DefaultQuery` / `c.GetQuery` 的区别？**
    - `Query(k)`：不存在返回空字符串
    - `DefaultQuery(k, def)`：不存在返回 def
    - `GetQuery(k)`：返回 `(string, bool)`，能区分「空值」和「不存在」——**语义最正确**
    批量取用 `c.Request.URL.Query()`（每次调用会重新 parse，别在循环里反复调）。

45. **表单与文件上传？**
    ```go
    file, _ := c.FormFile("file")
    // 默认存在内存（gin 默认 32MB 阈值，超过落临时盘）
    c.SaveUploadedFile(file, dst)
    form, _ := c.MultipartForm()   // 多文件
    ```
    注意 `MaxMultipartMemory`（默认 32MB）只是内存阈值，**不是大小限制**——真限制大小要在 `http.Server` 或中间件里包 `http.MaxBytesReader`，否则大文件能把磁盘写满。

---

## 十、框架选型对比（面试官爱问「为什么选 Gin」）

46. **Gin / Echo / chi / Fiber / 标准库怎么选？**
    | | 路由 | 特点 | 适合 |
    |---|---|---|---|
    | Gin | 自研 radix tree | 生态最大、中间件多、Context 池化 | 大多数业务 API |
    | Echo | radix tree | 接口设计更整洁、内置更多（bind/validate/自定义 HTTPError） | 喜欢干净 API 的团队 |
    | chi | 基于 `http.Handler` 组合 | 100% 标准库风格，可与任何中间件组合 | 追求零魔法、库洁癖 |
    | Fiber | fasthttp | 非 `net/http`，性能高但兼容性有坑（如 HTTP/2、部分中间件） | IO 密集、愿意换生态 |
    | 标准库 | 1.22 起支持 `GET /x/{id}` | 零依赖、无中间件 | 小服务、CLI 内置 API |

    一句话总结方向：**要生态选 Gin，要克制选 chi，要性能且能接受 fasthttp 选 Fiber。**

47. **Go 1.22 的标准库路由变强了，还会选 Gin 吗？**
    1.22 的 `ServeMux` 支持方法与路径变量（`mux.HandleFunc("GET /items/{id}", h)`，`r.PathValue("id")`），路由这块的差距被大幅拉近。Gin 剩下的优势是：**中间件模型（洋葱 + Abort）、参数绑定与校验、响应渲染、Context 池化、超大生态**。如果团队只用标准库，这些就得自己写一遍。

    加分回答：两者的中间件思想不同——标准库是 `func(http.Handler) http.Handler` 的装饰器（可以跨框架复用），Gin 是链式 `HandlerFunc`（耦合于 `*gin.Context`，写起来短但不能跨框架）。

48. **如果让你从零设计一个 Web 框架，会怎么做？**
    面试官想看思路，按 Gin 的骨架答：
    1. 路由：压缩前缀树 + 按 method 分树，支持 `:param` / `*catch-all`
    2. 上下文：`sync.Pool` 复用，持有 `http.ResponseWriter` 与 `*http.Request`
    3. 中间件：`[]func(*Context)` + `index` 游标，`Next`/`Abort` 实现洋葱
    4. 绑定与校验：tag 驱动 + validator，按 Content-Type 分派
    5. 渲染：JSON/XML/HTML/Proto 的 `Render` 接口
    6. 错误与恢复：`defer recover`，错误链统一出口
    7. 部署：暴露 `http.Handler` 而不是自造 server，超时/优雅退出交给标准库

---

## 附：一分钟速记

- 本质：Engine 就是 `http.Handler`，`Run()` = `ListenAndServe`，生产必须自建 `http.Server` 配超时
- 路由：按 method 分树的压缩前缀树；静态 > `:param` > `*catchall`；冲突在启动期 panic
- 中间件：`[]HandlerFunc` + `index` 游标，`Next()` 往下跑完再返回（洋葱），`Abort()` 设 index=63 阻断；不调 `Next` 也不会阻断
- 上限：`index` 是 `int8`，一条链最多 63 个 handler
- Context：`sync.Pool` 复用 → 跨 goroutine 必须 `c.Copy()`；`gin.Context` ≠ `context.Context`，下游用 `c.Request.Context()`
- 绑定：`ShouldBind` 按 Content-Type 自动选，失败自己处理；body 只能读一次，读两次用 `ShouldBindBodyWith`
- 校验：validator 全局单例，`binding:"required,oneof,dive"`，`required` 对数字 0 判定为缺失
- JSON：默认转义 HTML（`SetEscapeHTML(false)` 或 `PureJSON`）；写过响应就改不了状态码
- 安全：`SetTrustedProxies` 必须配，否则 XFF 可伪造
- 优雅退出：`signal.Notify` → `srv.Shutdown(ctx)` → 关 DB
- 排查：`pprof` 挂管理端口，`wrk` 压测，JSON 序列化换 sonic/json-iterator
- 选型：要生态 Gin，要克制 chi，`net/http` 1.22 起路由够用但缺中间件与绑定
