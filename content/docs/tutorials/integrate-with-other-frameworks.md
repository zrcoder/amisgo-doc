---
title: "和其他框架集成"
date: 2024-12-23T16:28:44+08:00
weight: 4
---

amisgo 提供了 Handle 和 HandleFunc 方法来集成其他实现了 http.Handler 的实例; 同时，其引擎也是一个标准的 http.Handler, 可方便集成进其他框架。

我们来看看三种示例：

### 1. 使用 `Handle` 方法集成 gin

假设 UI 部分由 amisgo 定义，而 API 部分使用 gin 实现。可以通过 `Handle` 方法将 gin 的路由挂载到 amisgo 中。

```go
const (
	apiPrefix = "/api/"
	datePath  = "date"
	dateApi   = apiPrefix + datePath
)

func main() {
	// 初始化 gin
	g := gin.Default()
	g.GET(dateApi, func(c *gin.Context) {
		c.JSON(200, gin.H{"date": time.Now()})
	})
	// 初始化 amisgo
	app := amisgo.New()
	app.Handle(apiPrefix, g) // 将 g 挂载到 /api/ 路径
	app.Mount("/", app.Page().InitApi(dateApi).Body("Now: ${date}"))
	// 启动服务
	panic(app.Run(":8888"))
}
```

---

### 2. 将 amisgo 作为 `http.Handler` 使用

amisgo 引擎本身实现了 `http.Handler` 接口，因此可以直接与标准库的 `http` 包集成。

```go
func main() {
	// 初始化 gin
	g := gin.Default()
	g.GET(dateApi, func(c *gin.Context) {
		c.JSON(200, gin.H{"date": time.Now()})
	})
	// 初始化 amisgo
	app := amisgo.New()
	app.Mount("/", app.Page().InitApi(dateApi).Body("Now: ${date}"))
	// 使用标准库的 http 包
	http.Handle("/", app)
	http.Handle(apiPrefix, g)
	panic(http.ListenAndServe(":8888", nil))
}
```

---

### 3. 将 amisgo 包装为 gin 的 `HandlerFunc`

如果希望以 gin 为主框架，可以将 amisgo 实例包装为 gin 的 `HandlerFunc`，并在 gin 中处理请求。

```go
func main() {
	// 初始化 gin
	g := gin.Default()
	g.GET(dateApi, func(c *gin.Context) {
		c.JSON(200, gin.H{"date": time.Now()})
	})
	// 初始化 amisgo
	app := amisgo.New()
	app.Mount("/", app.Page().InitApi(dateApi).Body("Now: ${date}"))
	// 将 amisgo 包装为 gin 的 HandlerFunc
	g.GET("/", func(c *gin.Context) {
		app.ServeHTTP(c.Writer, c.Request)
	})
	// 启动 gin
	panic(g.Run(":8888"))
}
```

---

### 总结

通过以上方式，amisgo 可以灵活地与 gin 或其他框架集成，满足不同场景的需求。无论是将 amisgo 作为主框架，还是将其嵌入到其他框架中，都能轻松实现。
