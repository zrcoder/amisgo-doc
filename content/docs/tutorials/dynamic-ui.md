---
title: "动态 UI"
date: 2025-03-28T17:47:23+08:00
weight: 5
---

我们可以用 service 和 amis 两个组件，实现完全动态的 UI，让我们看看如下例子：

```go
var app *amisgo.App

func main() {
	app = amisgo.New()
	index := app.Page().Body(
		app.Service().Name("ui").GetData(func() (any, error) {
			return map[string]any{
				"ui": getDynamicUI(),
			}, nil
		}).Body(
			app.Amis().Name("ui"),
		),
	)
	app.Mount("/", index)
	app.Run(":8080")
}

func getDynamicUI() any {
	return app.Tpl().Tpl("Hello, world!")
}
```

这有什么用？

基于这个机制，可以简化代码，做类似游戏的逻辑：根据应用状态，动态更新 UI。实际的例子见示例库的 ball-sort。