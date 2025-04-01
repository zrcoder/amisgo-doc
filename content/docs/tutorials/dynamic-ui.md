---
title: "动态 UI"
date: 2025-03-28T17:47:23+08:00
weight: 5
---

我们可以用 service 的 GetSchema方法，实现完全动态的 UI，让我们看看如下例子：

```go
var app *amisgo.App

func main() {
	app = amisgo.New()
	index := app.Page().Body(
		app.Service().Name("ui").GetSchema(getDynamicUI),
	)
	app.Mount("/", index)
	app.Run(":8080")
}

func getDynamicUI() any {
	return app.Tpl().Tpl("Hello, world!")
}
```

Service 设置 name，动态 UI 甚至可以据此做刷新该 Service 的交互。

这有什么用？基于这个机制，可以简化代码，做类似游戏的逻辑，即根据应用状态，动态更新 UI。

实际的例子见应用示例章节的 amisgo-games。