---
title: "设计原理"
date: 2024-12-23T16:26:06+08:00
weight: 2
---

amisgo 是基于百度 [amis](https://aisuda.bce.baidu.com/amis) 的一个 Go 语言实现。amis 允许通过 JSON 配置来定义页面，而 amisgo 则进一步优化了这一做法，通过 Go 的类型系统定义各种组件，并将其转换为 JSON，最终通过 Go template 渲染出 amis SDK 支持的 HTML 页面。针对原生 amis 较复杂的交互部分，也增加了基于回调的简化方法。

## 核心模块 comp
这个模块用于定义各种组件，代码非常简单，遵循统一模式。

### 1. 组件的基本定义
每个组件的基本定义和构造方法如下：

```go
type form model.Schema

func Form() form {
  return form{"type": "form"}
}
```

> 其中 model.Shema 底层是 map[string]any: 
> ```go
> type Schema map[string]any
> ```

### 2. 组件的属性方法

每个组件都有一系列属性设置方法，如：

```go
func (f form) Title(value string) form {
  return f.set("title", value)
}

func (f form) Body(value ...any) form {
	return f.set("body", value)
}
```

> 其中，set 是一个辅助方法，用于设置属性值并返回当前组件实例：
> 
> ```go
> func (f form) set(key string, value any) form {
>   f[key] = value
>   return f
> }
> ```

## 示例代码

以下是一个简单示例，展示了 amisgo 渲染页面的写法：

```go
comp.Page().Title("amisgo").Body(
	comp.Form().
	Api("https://xxx/api/saveForm").
	Body(
		comp.InputText().Label("姓名").Name("name"),
		comp.InputEmail().Label("邮箱").Name("email"),
	),
)
```

相较于原生 amis，amisgo 具有以下优势：

1. **强类型系统**：通过 Go 的 map 别名定义组件，方法定义组件属性，减少了 JSON 拼写错误的风险。
2. **简化表达**：amis 中交互事件描述较为复杂，amisgo 借助回调函数等简化了这一部分，使其变得明白晓畅。详见后续章节。
