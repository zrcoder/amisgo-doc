---
title: "怎么使用本地 sdk"
date: 2025-01-05T20:16:18+08:00
weight: 1
---

amisgo 默认使用 cdn （ https://cdn.jsdelivr.net/npm/amis ）加载 amis 的 sdk；同时支持使用本地 sdk。

我们提供了一个编译标签 `amisgo_local_sdk`，假设你之前这样编译：

```sh
go build  -o myapp .
```

那么加上对应标签的编译命令是：

```sh
go build --tags "amisgo_local_sdk"  -o myapp .
```

使用本地 sdk 会比使用 cdn 编译后的二进制大 50M 左右。

我们建议默认行为，但如果你的网络访问 cdn 不稳定或有特别限制，那么可以使用本地 sdk 。
