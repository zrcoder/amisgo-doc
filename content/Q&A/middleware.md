---
title: "怎么使用中间件"
date: 2025-01-05T20:17:29+08:00
weight: 2
---

假设我们做一个需要登录的 app，在浏览器访问任意页面，需要先鉴权，鉴权失败则重定向到登录页面。显然这里用中间件较好。

Egine 结构的 Mount、Handle 和 HandleFunc 方法均支持中间件，函数签名如下：

```go
func (e *Engine) Mount(path string, component any, middlewares ...func(http.Handler) http.Handler) *Engine
func (e *Engine) Handle(path string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) *Engine
func (e *Engine) HandleFunc(path string, handler http.HandlerFunc, middlewares ...func(http.Handler) http.Handler) *Engine
```

一个简单的 demo 代码如下：

```go
package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
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
	ag := amisgo.New().
		Mount("/", index, checkAuthMiddleware, testMiddleware).
		Mount(loginUrl, login).
		HandleFunc(echoApiUrl, echo)

	panic(ag.Run())
}

func checkAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Pre-processing actions, such as logging access, authentication, etc.
		fmt.Println("check auth middleware")
		if r.URL.Path != loginUrl && !checkAuth(r) {
			util.Redirect(w, r, loginUrl, http.StatusTemporaryRedirect)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func testMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("test middleware")
		w.Header().Set("test", "test heander value")
		next.ServeHTTP(w, r)
		// Post-processing actions, such as logging debug information
		fmt.Println("response heander for [test]:", w.Header().Get("test"))
	})
}

func echo(w http.ResponseWriter, r *http.Request) {
	resp := comp.SuccessResponse("", comp.Data{"body": "Hello, amisgo!"})
	w.Write(resp.Json())
}

func checkAuth(r *http.Request) bool {
	// Parse the token from the request and process authentication.
	// This is just a demonstration; it randomly returns the authentication result.
	return rand.Intn(2) == 0
}
```
