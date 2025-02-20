---
title: "常见问题"
date: 2024-12-29T20:30:56+08:00
breadcrumbs: false
weight: 3
prev: docs/tutorials/integrate-with-other-frameworks
---

## 怎么导入本地 SDK

amisgo 默认通过 CDN 加载 amis SDK。如果网络不稳定或有特殊限制，可以从本地导入。

你需要从 [amis 仓库 release](https://github.com/baidu/amis/releases) 下载 `jssdk.tar.gz` 并解压到 `jssdk` 目录。

然后在初始化时使用 `WithLocalSdk` 选项配置本地路径：

```go
amisgo.New(conf.WithLocalSdk(http.Dir("jssdk")))
```

或者可以用 Go 的 embed 特性将 jssdk 目录做成 Go 包后使用，也可直接使用我们做好的：

```go
import sdk "gitee.com/rdor/amis-sdk"

amisgo.New(conf.WithLocalSdk(http.FS(sdk.FS)))
```

> **注意**：本地 SDK 会使程序包增加约 50M。

## 怎么使用中间件

假设要做一个有用户系统的应用，访问页面时需先鉴权，失败则重定向到登录页。使用中间件是实现此功能的理想方式。

引擎的 Mount、Handle 和 HandleFunc 方法均支持中间件，示例代码如下：

```go
const loginPath = "/login"

func main() {
	app := amisgo.New()
	index := app.Page().Body("Hello, Amisgo!")
	login := app.Page().Body(
		app.Form().Body(
			app.InputEmail().Name("user"),
			app.InputPassword().Name("password"),
		),
	)
	app.Mount("/", index, checkAuthMiddleware, testMiddleware)
	app.Mount(loginPath, login)

	panic(app.Run(":8080"))
}

// 鉴权检查，失败则重定向到登录页。
func checkAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("check auth middleware")
		if r.URL.Path != loginPath && !checkAuth(r) {
			util.Redirect(w, r, loginPath, http.StatusTemporaryRedirect)
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

func checkAuth(r *http.Request) bool {
	return rand.Intn(2) == 0 // 仅示例，这里随机返回鉴权结果
}
```

更多实现可参考示例库中的 todo-app 。

## 怎么切换主题

amisgo 支持 amis 的四种内置主题：云舍、antd、ang 和 dark。您可以配置使用其中多个主题，并在页面添加 ThemeSelect 或 ThemeButtonGroupSelect 来支持用户切换主题。

```go
app := amisgo.New(
	conf.WithThemes(
		theme.Theme{Value: theme.Cxd, Label: "Light"},
		theme.Theme{Value: theme.Dark, Label: "Dark"},
	),
)
app.Mount("/", app.Page().Body(
	app.ThemeButtonGroupSelect(),
	"Hello, World!",
))
app.Run(":8888")
```

实际的例子可以参考示例库的 dev-topys 和 todp-app 。

## 怎么兼容纯 JSON

假如你想直接写 JSON 来定义页面，而不是用 comp 模块的 API 来定义（这个场景可能来自快速验证 amis 文档里的示例 JSON），仅需要向 Mount 方法传递 JSON 内容即可。例如：

```go
const amisJSON = `{
	"type": "page",
	"title": "Hello",
	"body": "World!"
}`

app := amisgo.New()
app.Mount("/", json.RawMessage(amisJSON))
app.Run(":8080")
```

> 注意把 string 转成 json.RawMessage

当然，你可以更进一步，用 JSON 文件定义页面内容，那么仅需要把文件内容传递到 Mount 的第二个参数即可； 甚至使用 embed，进一步简化代码。

```go
//go:embed pages/index.json
var index json.RawMessage

func main() {
	app := amisgo.New()
	app.Mount("/", index)
	app.Run(":8080")
}
```
