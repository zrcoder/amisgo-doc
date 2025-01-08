---
title: "简化交互代码"
date: 2024-12-23T16:28:01+08:00
weight: 3
---

## Api，InitApi

假设我们页面有一个表单，用户输入信息，点击提交后我们需要把信息存入数据库。

借助 form 组件的 Api 方法，可以这样做：

```go
func init() {
    http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
                input, _ := io.ReadAll(r.Body)
                defer r.Body.Close()

				m := map[string]any{}
                json.Unmarsharl(input, &m)

                name := m["name"]
                email := m["email"]
                // save the user into db ...

	})
}

...

comp.Page().Body(
	comp.Form().Api("/user").Body(
		comp.InputText().Label("姓名").Name("name"),
		comp.InputEmail().Label("邮箱").Name("email"),
	),
)
```

显然这个代码非常啰嗦，amisgo 给 form 组件增加了一个新方法 Submit：

```go
// Submit 设置表单提交后的回调逻辑，使用通用的 Data 类型处理表单数据
// 适用于需要灵活处理表单提交的场景
func (f form) Submit(callback func(model.Data) error) form
```

这样一来，用户代码可以简化成如下：

```go
comp.Page().Body(
	comp.Form().Body(
		comp.InputText().Label("姓名").Name("name"),
		comp.InputEmail().Label("邮箱").Name("email"),
	).Submit(func(m model.Data) error {
		name := m.Get("name")
		email := m.Get("email")
		// save name and email to db
		// ...
                return nil
	}),
)
```

> 同样还有个 SubmitTo 方法，使用具体类型处理表单数据
>
> ```go
> func (f form) SubmitTo(receiver any, callback func(any) error) form
> ```

类似地，我们包装了 page 组件的 InitApi，增加 InitData 方法：

```go
func (p page) InitData(getter func() (any, error)) page
```

这样，用户写页面获取当前时间的代码就简化成了：

```go
    comp.Page().
	Title("标题").
	Body("内容部分. 可以使用 \\${var} 获取变量。如: `\\$date`: ${date}").
	InitData(getDate)

func getDate() (any, error) {
	y, m, d := time.Now().Date()
	mm := time.Now().UnixNano()
	return map[string]string{"date": fmt.Sprintf("%d-%d-%d %d", y, m, d, mm)}, nil
}
```

## Action 按钮

假设页面有两个编辑器，第一个用于输入 json，第二个是只读的，当点击按钮后，以第一个编辑器的内容作为输入，转换为 yaml，渲染到第二个编辑器中，怎么做呢？

直接借助 ajax 类型的行为按钮，可以这么写：

```go
func init() {
	http.HandleFunc("/convert", func(w http.ResponseWriter, r *http.Request) {
		input, _ := io.ReadAll(r.Body)
		defer r.Body.Close()
		m := map[string]any{}
		json.Unmarshal(input, &m)
		// ...
		output := "age: 27"
		resp := model.Response{Msg: "转换成功", Data: model.Data{"output": output}} // 这里的key值必须是第二个编辑器的 name
		data, _ := json.Marshal(resp)
		w.Write(data)
	})
}

func main() {
	index := comp.Page().Body(
		comp.Form().ColumnCount(2).Body(
			comp.Editor().Language("json").Label("json").Name("input").Size("xxl"),
			comp.Editor().Label("yaml").Label("yaml").Name("output").Size("xxl").ReadOnly(true),
		).Actions(
			comp.Action().Label("Convert").Level("primary").ActionType("ajax").Api(
				model.Schema{
					"url":  "/convert",
					"data": model.Schema{"input": "${input}"},
					"responses": model.Schema{
						"200": model.Schema{
							"then": model.Schema{
								"actionType": "setValue",
								"args": model.Schema{
									"value": "${response}",
								},
							},
						},
					},
				},
			),
...
```

显然，这也比较啰嗦，我们为行为按钮新增了 Transform 方法：

```go
func (a action) Transform(src, dst, successMsg string, transfor func(input any) (any, error)) action
```

上面的代码就可以简化成：

```go
    comp.Page().Body(
	comp.Form().ColumnCount(2).Body(
	    comp.Editor().Language("json").Label("json").Name("input").Size("xxl"),
	    comp.Editor().Label("yaml").Label("yaml").Name("output").Size("xxl").ReadOnly(true),
	).Actions(
	    comp.Action().Label("Converrt").Level("primary").Transform("input", "output", "转换成功", func(input any) (any, error) {
		// transform input json to yaml
		output := "age: 27"
		return output, nil
	}),
```

实际上，我们还支持了多对多的 TransformMultiple 方法，可以实现从多个组件的输入值转换后渲染到多个组件。

可以参考 [dev-toys](https://github.com/zrcoder/amisgo-examples/tree/main/dev-toys) 中的示例，其中的 convert 组件用了 Transform， 生成多种类型的 hash 值用了 TransformMultiple。
