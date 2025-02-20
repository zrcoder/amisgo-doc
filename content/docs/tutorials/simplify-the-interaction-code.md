---
title: "简化交互代码"
date: 2024-12-23T16:28:01+08:00
weight: 3
---

## 1. Api 与 InitApi 的优化

假设我们有一个表单页面，用户输入信息后点击提交，我们需要将信息存入数据库。

### 传统方式

使用 `form` 组件的 `Api` 方法，代码可能如下：

```go {hl_lines=[3,8]}
app := amisgo.New()
index := app.Page().Body(
	app.Form().Api("/user").Body(
		app.InputText().Label("姓名").Name("name"),
		app.InputEmail().Label("邮箱").Name("email"),
	),
)
app.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
	input, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	m := map[string]string{}
	json.Unmarshal(input, &m)

	name := m["name"]
	email := m["email"]
	fmt.Println(name, email)
	// 将用户信息存入数据库 ...
})
app.Mount("/", index)
```

### 优化方式

`amisgo` 为 `form` 组件新增了 `Submit` 方法，简化了代码：

```go
func (f Form) Submit(callback func(schema.Schema) error) Form
```

优化后的代码如下：

```go
app := amisgo.New()
index := app.Page().Body(
	app.Form().Api("/user").Body(
		app.InputText().Label("姓名").Name("name"),
		app.InputEmail().Label("邮箱").Name("email"),
	).Submit(
		func(s schema.Schema) error {
			name := s.Get("name").(string)
			email := s.Get("email").(string)
			fmt.Println(name, email)
			// 将用户信息存入数据库 ...
			return nil
		},
	),
)
```

> 此外，另有 `SubmitTo` 方法允许使用具体类型处理表单数据：
>
> ```go
> func (f form) SubmitTo(receiver any, callback func(any) error) form
> ```

### InitData 方法

类似地，`page` 组件的 `InitApi` 也有对应的优化方法 `InitData`：

```go
func (p page) InitData(getter func() (any, error)) page
```

例如，获取当前时间的代码可以简化为：

```go
func main() {
	app := amisgo.New()
	index := app.Page().Body("Now: ${date}").InitData(getDate)
	app.Mount("/", index)
	app.Run(":8888")
}

func getDate() (any, error) {
	y, m, d := time.Now().Date()
	mm := time.Now().UnixNano()
	return map[string]string{"date": fmt.Sprintf("%d-%d-%d %d", y, m, d, mm)}, nil
}
```

## 2. Action 按钮的优化

假设页面有两个输入框，第一个用于输入人名，第二个只读。点击按钮后，将输入做一定转换，并渲染到第二个文本框中。

### 传统方式

使用 `ajax` 类型的行为按钮，代码如下：

```go {hl_lines=[12,27]}
app := amisgo.New()
index := app.Page().Body(
	app.Form().WrapWithPanel(false).Body(
		app.InputText().Name("input"),
		app.InputText().Name("output").ReadOnly(true),
		app.Action().
			Label("Greet").
			Level("primary").
			ActionType("ajax").
			Api(
				app.Api().
					Url("/convert").
					Data(schema.Schema{"input": "${input}"}).
					Set(
						"resp",
						schema.Schema{
							"200": schema.Schema{
								"then": app.EventAction().ActionType("setValue").
									Args(app.EventActionArgs().Value("${resp}")),
							},
						},
					),
			),
	),
)
app.Mount("/", index)
app.HandleFunc("/convert", func(w http.ResponseWriter, r *http.Request) {
	input, _ := io.ReadAll(r.Body)
	defer r.Body.Close()
	m := map[string]string{}
	json.Unmarshal(input, &m)
	output := "hello " + m["input"]
	resp := schema.SuccessResponse("", schema.Schema{"output": output}) // 这里的 key 值必须是第二个编辑器的 name
	w.Write(resp.Json())
})
```

### 优化方式

`amisgo` 为行为按钮新增了 `Transform` 方法：

```go
func (a Action) Transform(transfor func(input any) (any, error), src, dst string) Action
```

优化后的代码如下：

```go
app := amisgo.New()
index := app.Page().Body(
	app.Form().WrapWithPanel(false).Body(
		app.InputText().Name("input"),
		app.InputText().Name("output").ReadOnly(true),
		app.Action().Label("Greet").
			Level("primary").
			Transform(func(input any) (any, error) {
				return "hello " + input.(string), nil
			}, "input", "output"),
	),
)
app.Mount("/", index)
```

### 多对多转换

另有 `TransformMultiple` 方法支持从多个组件的输入值转换后渲染到多个组件。可以参考示例库中的 dev-toys，其中的 convert 组件使用了 Transform，生成多种类型的 hash 值使用了 TransformMultiple。

---

通过以上优化，代码变得更加简洁易读，减少了冗余的 HTTP 请求处理逻辑，提升了开发效率。
