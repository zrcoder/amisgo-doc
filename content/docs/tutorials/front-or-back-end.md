---
title: "前端还是后端"
date: 2024-12-23T16:27:19+08:00
weight: 2
---

可以用 amisgo 写纯前端，需要的后端是额外的服务器，这样会是传统的前后端分离的模式。但我们更推荐用 amisgo 写前后端合一的项目，这样技术栈更统一， 实际部署也仅仅是一个进程，而不是前后端各一个进程。

## 纯前端

比如我们来仿写下 [Go+ 的 playground](https://play.goplus.org) 。

主要实现 `Run` 和 `Format` 两个功能。我们用 amisgo 主要写前端代码，实际的编译调 https://play.goplus.org 已有的 api。

主要代码如下：

```go
func main() {
	examples, defaultExample, err := example.Get()
	check(err)

	index := comp.Page().Body(
		comp.Form().WrapWithPanel(false).Body(
			comp.Group().Mode("inline").Body(
				comp.Image().Alt("Go+").Src("/static/gop.svg").Height("20px").InnerClassName("border-none"),
				comp.InputGroup().Body(
					comp.Button().Primary(true).Label("Run").Transform("body", "result", "Done", func(input any) (any, error) {
						return compile(input.(string))
					}),
					comp.Button().Primary(true).Label("Format").Transform("body", "body", "Done", func(input any) (any, error) {
						return format(input.(string))
					}),
				),
				comp.Select().Name("examples").Value(defaultExample).Options(
					examples...,
				),
			),
			comp.Editor().Language("c").Name("body").Size("xl").Value("${examples}").
				AllowFullscreen(false).Options(model.Schema{"fontSize": 15}),
			comp.Code().Name("result").Language("plaintext"),
		),
	)

	app := amisgo.New().Mount("/", index).StaticFS("/static/", http.FS(static.FS))

	err = app.Run()
	check(err)
}

```

注意其中两个主要按钮的 transform 代码：

```go
comp.Button().Primary(true).Label("Run").Transform("body", "result", "Done", func(input any) (any, error) {
    return compile(input.(string))
}),
comp.Button().Primary(true).Label("Format").Transform("body", "body", "Done", func(input any) (any, error) {
    return format(input.(string))
}),
```

其中 compile 和 format 的实现也非常简单，就是根据 form 请求中 body 的值，调用对应的 goplus 的 api，然后返回结果， amisgo 封装的 Transform 方法会自动把返回值渲染到 result 组件中。

> 这里有个小技巧：写本地代理服务来中转请求，这样可以简化前端代码。不然要在组件里写 amis 的 api 对象，比较繁琐。而且本地代理服务也被 amisigo 封装了，比如这里的 transform 方法，我们只需要实现数据转换的逻辑 compile 和 format，而不用管前端是怎么调到这两个函数的。

我们实现的 playground 效果如下：
![gop-playground](/gop-play.png)

完整代码见：[amisgo-examples/gop-playground](https://github.com/zrcoder/amisgo-examples/tree/main/gop-playground)

## 前后端合一

比如我们写一个 todo list 应用，前端用 amisgo，后端用 gin， 数据库用 sqlite。所有代码都在一个仓库中，部署也仅仅是一个二进制，非常方便。

这个应用的代码见:[amisgo-examples/todo-app](https://github.com/zrcoder/amisgo-examples/tree/main/todo-app)。
