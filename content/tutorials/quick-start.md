---
title: "快速开始"
date: 2024-12-23T16:26:55+08:00
weight: 1
---

```go
package main

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

func main() {
	index := comp.Page().Title("Amisgo").Body(
		comp.Form().
		Api("https://xxx/api/saveForm").
		Body(
			comp.InputText().Label("姓名").Name("name"),
			comp.InputEmail().Label("邮箱").Name("email"),
		),
	)

	ag := amisgo.New().Mount("/", index)

	panic(ag.Run(":8080"))
}
```

运行代码后，访问 http://localhost:8080。
