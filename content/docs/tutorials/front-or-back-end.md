---
title: "前端还是后端"
date: 2024-12-23T16:27:19+08:00
weight: 2
---

在使用 amisgo 进行开发时，可以选择两种主要模式：纯前端模式和前后端合一模式。本文将探讨这两种模式的优缺点。

## 纯前端模式

在纯前端模式下，amisgo 主要用于编写前端代码，而后端服务则由额外的服务器提供。这种模式遵循传统的前后端分离架构。

### 示例：仿写 Go Playground

我们以仿写 [Go Playground](https://go.dev/play) 为例，主要实现 `Run` 和 `Format` 两个功能。前端代码使用 amisgo 编写，而后端功能通过调用 Go Playground 的现有 API 实现。

### 主要代码

```go {hl_lines=[2,3,4,5,6,7,8,9,10,11]}
app.Form().Body(
	app.Button().Primary(true).Label("Run").TransformMultiple(func(s schema.Schema) (schema.Schema, error) {
		res, err := compile(s.Get("body").(string))
		if err != nil {
			return schema.Schema{"result": "❌ " + err.Error()}, nil
		}
		return schema.Schema{"result": res}, nil
	}, "body"),
	app.Button().Primary(true).Label("Format").Transform(func(input any) (any, error) {
		return format(input.(string))
	}, "body", "body"),
	app.Editor().Language("c").Name("body").Size("xxl").Value("${examples}").
		AllowFullscreen(false).Options(schema.Schema{"fontSize": 15}),
	app.Code().Name("result").Language("plaintext"),
)
```

### 关键点解析

- **Transform 方法**：`Run` 和 `Format` 按钮的 `Transform` 方法用于处理用户输入并调用相应的 API。`compile` 和 `format` 函数负责与 Go+ Playground 的 API 交互，并将结果返回给前端。

- **本地代理服务**：通过本地代理服务中转请求，可以简化前端代码，避免在组件中直接编写 API 调用逻辑。

### 效果展示

![Go Playground 效果图](/goplay.png)

完整代码请参考“应用案例”章节。

## 前后端合一模式

在前后端合一模式下，前后端代码都是 Go 语言写的，所有代码都位于同一个仓库中，结合 Go 的 embed 特性，部署时也仅需一个二进制文件。

### 示例：Todo List 应用

我们以一个简单的 Todo List 应用为例，前端用 amisgo 编写，后端则使用 gin 库，数据库采用 SQLite，展示如何实现前后端合一模式。

代码请参考“应用案例”章节。

## 总结

- **纯前端模式**：适合需要与现有后端服务集成的场景，遵循传统的前后端分离架构。
- **前后端合一模式**：适合小型项目或需要简化部署流程的场景，技术栈统一，部署简便。
