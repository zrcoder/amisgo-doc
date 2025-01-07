---
title: "怎么导入本地 sdk"
date: 2025-01-05T20:16:18+08:00
weight: 1
prev: docs/tutorials/integrate-with-other-frameworks
---

amisgo 默认通过 [cdn](https://cdn.jsdelivr.net/npm/amis) 加载依赖的 sdk ，这是推荐的方式；但如果你的网络访问 cdn 不稳定或有特别限制，那么可以从本地导入 sdk 。

amisgo 支持如下配置选项：

```go
func WithLocalSdk(fs http.FileSystem) Option
```

你需要下载百度 [amis 仓库 release](https://github.com/baidu/amis/releases) 里的 jssdk.tar.gz，解压得到 jssdk 目录，然后在初始化时配置：

```go
amisgo.New(conf.WithLocalSdk(http.Dir("jssdk")))
```

或者可以用 Go 的 embed 特性，将 jssdk 目录做成一个 Go 包，用该包配置 amisgo。我们已经做好一个如下包，可直接引用，当然你也可以自己做这个包。

```go
import "gitee.com/rdor/amis-sdk/sdk"

...
amisgo.New(conf.WithLocalSdk(http.FS(sdk.FS)))
```

> 使用本地 sdk 会比使用 cdn 最终的程序包大 50M 左右。
