---
title: "怎么支持多主题切换"
date: 2025-01-14T18:08:15+08:00
---

amisgo 支持 amis 的四种内置主题：云舍、antd、ang 和 dark。您可以配置使用其中多个主题，并在页面添加 ThemeSelect 或 ThemeButtonGroupSelect 来支持用户切换主题。

```go
	amisgo.New(
		conf.WithThemes(conf.ThemeCxd, conf.ThemeDark),
	).
		Mount("/", comp.Page().Body(
			comp.ThemeButtonGroupSelect(),
			"Hello, World!",
		)).
		Run(":8888")
```

实际的例子可以参考示例库的 dev-topys 和 todp-app 。
