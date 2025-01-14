---
title: "前端还是后端"
date: 2024-12-23T16:27:19+08:00
weight: 2
---

在使用 amisgo 进行开发时，可以选择两种主要模式：纯前端模式和前后端合一模式。本文将探讨这两种模式的优缺点。

## 纯前端模式

在纯前端模式下，amisgo 主要用于编写前端代码，而后端服务则由额外的服务器提供。这种模式遵循传统的前后端分离架构。

### 示例：仿写 Go+ Playground

我们以仿写 [Go+ Playground](https://play.goplus.org) 为例，主要实现 `Run` 和 `Format` 两个功能。前端代码使用 amisgo 编写，而后端功能通过调用 Go+ Playground 的现有 API 实现。

### 主要代码

```go
index := comp.Page().Body(
	comp.Form().WrapWithPanel(false).Body(
		comp.Flex().Justify("space-between").Items(
			comp.Group().Mode("inline").Body(
				comp.Image().Alt("Go+").Src("/static/gop.svg").Height("20px").InnerClassName("border-none"),
				comp.InputGroup().Body(
					comp.Button().Primary(true).Label("Run").Transform(func(input any) (any, error) {
						return compile(input.(string))
					}, "body", "result"),
					comp.Button().Primary(true).Label("Format").Transform(func(input any) (any, error) {
						return format(input.(string))
					}, "body", "body"),
				),
				comp.Select().Name("examples").Value(defaultExample).Options(
					examples...,
				),
			),
			comp.Button().Label("Github").ActionType("url").Icon("fa fa-github").Url("https://github.com/goplus/gop"),
		),
		comp.Editor().Language("c").Name("body").Size("xxl").Value("${examples}").
			AllowFullscreen(false).Options(model.Schema{"fontSize": 15}),
		comp.Code().Name("result").Language("plaintext"),
	),
)
```

### 关键点解析

- **Transform 方法**：`Run` 和 `Format` 按钮的 `Transform` 方法用于处理用户输入并调用相应的 API。`compile` 和 `format` 函数负责与 Go+ Playground 的 API 交互，并将结果返回给前端。

```go
comp.Button().Primary(true).Label("Run").Transform(func(input any) (any, error) {
	return compile(input.(string))
}, "body", "result"),
comp.Button().Primary(true).Label("Format").Transform(func(input any) (any, error) {
	return format(input.(string))
}, "body", "body"),
```

- **本地代理服务**：通过本地代理服务中转请求，可以简化前端代码，避免在组件中直接编写 API 调用逻辑。

### 效果展示

![Go+ Playground 效果图](/gop-play.png)

完整代码请参考示例库。

## 前后端合一模式

在前后端合一模式下，前后端代码都是 Go 语言写的，所有代码都位于同一个仓库中，结合 Go 的 embed 特性，部署时也仅需一个二进制文件。

### 示例：Todo List 应用

我们以一个简单的 Todo List 应用为例，前端用 amisgo 编写，后端则使用 gin 库，数据库采用 SQLite，展示如何实现前后端合一模式。

代码请参考示例库。


## 总结

- **纯前端模式**：适合需要与现有后端服务集成的场景，遵循传统的前后端分离架构。
- **前后端合一模式**：适合小型项目或需要简化部署流程的场景，技术栈统一，部署简便。
