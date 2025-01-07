---
title: "和其他框架集成"
date: 2024-12-23T16:28:44+08:00
weight: 4
next: docs/qa/local-sdk
---

首先，amisgo 提供了 Handle 和 HandleFunc 方法来集成其他实现了 http.Handler 的实例。

假设我们的 UI 部分用 amisgo 定义，api 部分可以用标准库也可以用三方库比如 gin 来实现，可以这样写：

```go
const (
	apiPrefix = "/api/"
	date      = "date"
)

func main() {
	dateApi := apiPrefix + date
	g := gin.Default()
	g.GET(dateApi, func(c *gin.Context) {
		c.JSON(200, comp.SuccessResponse("", comp.Data{"date": time.Now()}))
	})

	app := amisgo.New().
		Handle(apiPrefix, g).
		Mount("/", comp.Page().InitApi(dateApi).Body("Now: ${date}"))

	panic(app.Run("8888"))
}
```

又因为 amisgo 的 Egine 本身是一个 http.Handler，所以可以这样写:

```go
func main() {
	dateApi := apiPrefix + date
	g := gin.Default()
	g.GET(dateApi, func(c *gin.Context) {
		c.JSON(200, comp.SuccessResponse("", comp.Data{"date": time.Now()}))
	})

	app := amisgo.New().
		Handle(apiPrefix, g).
		Mount("/", comp.Page().InitApi(dateApi).Body("Now: ${date}"))

	http.Handle("/", app)
	http.Handle(apiPrefix, g)
	panic(http.ListenAndServe(":8888", nil))
}
```

还可以将 amisgo 的 Engine 包装为 gin.HandlerFunc：

```go
func main() {
	dateApi := apiPrefix + date
	g := gin.Default()
	g.GET(dateApi, func(c *gin.Context) {
		c.JSON(200, comp.SuccessResponse("", comp.Data{"date": time.Now()}))
	})

	app := amisgo.New().
		Handle(apiPrefix, g).
		Mount("/", comp.Page().InitApi(dateApi).Body("Now: ${date}"))

	g.GET("/", func(c *gin.Context) {
		app.ServeHTTP(c.Writer, c.Request)
	})

	panic(g.Run(":8888"))
}
```
