---
title: "怎么导入本地 SDK"
date: 2025-01-05T20:16:18+08:00
weight: 1
prev: docs/tutorials/integrate-with-other-frameworks
---

amisgo 默认通过 [CDN](https://cdn.jsdelivr.net/npm/amis) 加载 amis SDK。如果网络不稳定或有特殊限制，可以从本地导入。

### 操作步骤

1. **下载 SDK**

   从 [amis 仓库 release](https://github.com/baidu/amis/releases) 下载 `jssdk.tar.gz` 并解压到 `jssdk` 目录。

2. **配置 amisgo**

   使用 `WithLocalSdk` 选项配置本地路径：

   ```go
   amisgo.New(conf.WithLocalSdk(http.Dir("jssdk")))
   ```

3. **使用 embed 特性（可选）**

   可以用 Go 的 embed 特性将 jssdk 目录做成 Go 包后使用，也可直接使用我们做好的：

   ```go
   import "gitee.com/rdor/amis-sdk/sdk"

   amisgo.New(conf.WithLocalSdk(http.FS(sdk.FS)))
   ```

> **注意**：本地 SDK 会使程序包增加约 50M。
