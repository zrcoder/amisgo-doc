---
title: "怎么从本地导入 sdk"
date: 2025-01-05T20:16:18+08:00
weight: 1
---

amisgo 默认通过 [cdn](https://cdn.jsdelivr.net/npm/amis) 加载依赖的 amis sdk ，这是推荐的方式；但如果你的网络访问 cdn 不稳定或有特别限制，那么可以从本地导入 sdk 。

amisgo 支持如下配置选项：

```go
func WithSdkFS(fs http.FileSystem) Option
```

你需要下载百度 [amis 仓库 release](https://github.com/baidu/amis/releases) 里的 jssdk.tar.gz，解压得到 jssdk 目录，然后在初始化时配置：

```go
amisgo.New(conf.WithSdkFS(http.Dir("jssdk")))
```

或者可以用 Go 的 embed 特性，将 jssdk 目录做成一个 Go 包： 在 jssdk 目录下添加一个 fs.go:

```go
package jssdk

import (
	"embed"
)

//go:embed *
var FS embed.FS
```

然后使用该 FS 配置 amisgo：

```go
amisgo.New(conf.WithSdkFS(http.FS(jssdk.FS)))
```

> 使用本地 sdk 会比使用 cdn 编译后的二进制大 50M 左右。

```

```
