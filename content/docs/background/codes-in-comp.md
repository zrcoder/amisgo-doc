---
title: "comp 目录下的代码"
date: 2024-12-23T16:26:35+08:00
weight: 3
next: tutorials/quick-start
---

## 需要的代码

comp 目录下的代码大同小异。

1. 首先在 comp/base.go 定义了 Schema 类型， 底层为 map[strng]any：

```go
type Schema  map[string]any
```

2. 组件基本定义

```go
type xxx Schema

func XXX() xxx {
  return xxx{}.set("type": "xxx")
}

func (x xxx) set(key string, value interface{}) xxx {
  x[key] = value
  return x
}
```

3. 各个属性方法，如：

```go
func (x xxx) Title(value string) xxx {
  return x.set("title", value)
}
...
```

## 怎么生成每个组件的属性方法

1. 基于百度 amis 仓库 release 的 schema.json 生成
2. 基于百度 amis API 文档生成

可惜，#1 尝试了很多工具都失败了；#2 又非常耗费时间，要知道 amis 有一百多个组件， Api 文档又不是非常规范，很难写工具自动转换代码。

这个问题专门在 amis 仓库提了 issue：https://github.com/baidu/amis/issues/10760

最终比较费劲地完成了这个任务：发现 php 项目 https://github.com/slowlyo/owl-admin/blob/master/src/Renderers ，里边定义了大多数 amis 的组件，基于该项目，借助 chat-gpt， 花了周末两天时间做完了转换。
