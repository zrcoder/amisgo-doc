---
title: "快速开始"
date: 2024-12-23T16:26:55+08:00
weight: 1
prev: docs/background/design-principle
---

## 1. 前置条件

在开始之前，请确保已安装 **Go 1.18+**，并且已经初始化了一个 Go 模块。然后，使用以下命令安装 `amisgo`：

```bash
go get github.com/zrcoder/amisgo
```

## 2. 编写代码

以下是一个简单的示例代码，展示了如何使用 `amisgo` 创建一个包含表单的页面，并启动一个本地服务器。

```go {base_url=".",filename="quick-start.go"} 
package main

import (
	"github.com/zrcoder/amisgo"
)

func main() {
	app := amisgo.New()
	index := app.Page().Title("amisgo").Body(
		app.Form().
			Api("https://xxx/api/saveForm").
			Body(
				app.InputText().Label("姓名").Name("name"),
				app.InputEmail().Label("邮箱").Name("email"),
			),
	)
	app.Mount("/", index)

	panic(app.Run(":8080"))
}
```

## 3. 运行代码

在终端中运行以下命令来启动应用：

```bash
go run main.go
```

## 4. 访问应用

运行代码后，打开浏览器并访问 [http://localhost:8080](http://localhost:8080)，您将看到一个包含姓名和邮箱输入框的表单页面。

## 5. 下一步

现在您已经成功创建并运行了一个简单的 `amisgo` 应用。接下来，您可以尝试修改表单的内容，添加更多的组件，或者将表单提交到实际的 API 地址。
