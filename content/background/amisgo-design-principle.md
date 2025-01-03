---
title: "amisgo 设计原理"
date: 2024-12-23T16:26:06+08:00
weight: 2
---

amisgo 基于百度的 [Amis 框架](https://aisuda.bce.baidu.com/amis)，这个框架允许用 JSON 配置来定义页面。

amisgo 的想法很简单：用 Go 的 struct 定义各种组件（最终实现时更简单，底层都是 map[string]any， 只是给各种组件加上了链式调用的方法），然后把组件转成 json，用 Go template 的方式渲染一个 Amis SDK 支持的 html。

```html
...
    <div id="root" class="app-wrapper"></div>
    <script type="text/javascript">
      (function () {
...
        let amisScoped = amis.embed(
        '#root',
        {{ .AmisJson }},
...
```

比 Amis 本身的优势:

1. 强类型系统，减少写 json 属性的笔误；
2. 简化 Amis 的表达，Amis 中各种交互事件的描述还是比较复杂，amisgo 希望借助回调函数等简化这一部分，变得明白晓畅。
