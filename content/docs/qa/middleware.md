---
title: "怎么使用中间件"
date: 2025-01-05T20:17:29+08:00
weight: 2
---

假设要做一个有用户系统的应用，访问页面时需先鉴权，失败则重定向到登录页。使用中间件是实现此功能的理想方式。

Egine 的 `Mount`、`Handle` 和 `HandleFunc` 方法均支持中间件：

```go
func (e *Engine) Mount(path string, component any, middlewares ...func(http.Handler) http.Handler) *Engine
func (e *Engine) Handle(path string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) *Engine
func (e *Engine) HandleFunc(path string, handler http.HandlerFunc, middlewares ...func(http.Handler) http.Handler) *Engine
```

#### 示例代码

```go
package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/model"
	"github.com/zrcoder/amisgo/util"
)

const (
	loginUrl   = "/login"
	echoApiUrl = "/api/echo"
)

func main() {
	index := comp.Page().InitApi(echoApiUrl).Body("${body}")
	login := comp.Page().Body(
		comp.Form().Body(
			comp.InputEmail().Name("user"),
			comp.InputPassword().Name("password"),
		),
	)
	app := amisgo.New().
		Mount("/", index, checkAuthMiddleware, testMiddleware).
		Mount(loginUrl, login).
		HandleFunc(echoApiUrl, echo)

	panic(app.Run(":8080"))
}

// 鉴权检查，失败则重定向到登录页。
func checkAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("check auth middleware")
		if r.URL.Path != loginUrl && !checkAuth(r) {
			util.Redirect(w, r, loginUrl, http.StatusTemporaryRedirect)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// 测试中间件，设置响应头并记录调试信息。
func testMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("test middleware")
		w.Header().Set("test", "test heander value")
		next.ServeHTTP(w, r)
		fmt.Println("response heander for [test]:", w.Header().Get("test"))
	})
}

func echo(w http.ResponseWriter, r *http.Request) {
	resp := model.SuccessResponse("", model.Schema{"body": "Hello, amisgo!"})
	w.Write(resp.Json())
}

func checkAuth(r *http.Request) bool {
	return rand.Intn(2) == 0 // 随机返回鉴权结果
}
```

#### 参考
更多实现可参考 [示例库](https://github.com/zrcoder/amisgo-examples) 中的 `todo-app`。
