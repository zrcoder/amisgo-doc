---
title: "简化交互代码"
date: 2024-12-23T16:28:01+08:00
weight: 3
---

## 1. Api 与 InitApi 的优化

假设我们有一个表单页面，用户输入信息后点击提交，我们需要将信息存入数据库。

### 传统方式

使用 `form` 组件的 `Api` 方法，代码可能如下：

```go
func init() {
    http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
        input, _ := io.ReadAll(r.Body)
        defer r.Body.Close()

        m := map[string]any{}
        json.Unmarshal(input, &m)

        name := m["name"]
        email := m["email"]
        // 将用户信息存入数据库 ...
    })
}

app.Page().Body(
    app.Form().Api("/user").Body(
        app.InputText().Label("姓名").Name("name"),
        app.InputEmail().Label("邮箱").Name("email"),
    ),
)
```

### 优化方式

`amisgo` 为 `form` 组件新增了 `Submit` 方法，简化了代码：

```go
func (f form) Submit(callback func(model.Schema) error) form
```

优化后的代码如下：

```go
app.Page().Body(
    app.Form().Body(
        app.InputText().Label("姓名").Name("name"),
        app.InputEmail().Label("邮箱").Name("email"),
    ).Submit(func(m model.Schema) error {
        name := m.Get("name")
        email := m.Get("email")
        // 将 name 和 email 存入数据库
        // ...
        return nil
    }),
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
app.Page().
    Title("标题").
    Body("内容部分. 可以使用 \\${var} 获取变量。如: `\\$date`: ${date}").
    InitData(getDate)

func getDate() (any, error) {
    y, m, d := time.Now().Date()
    mm := time.Now().UnixNano()
    return map[string]string{"date": fmt.Sprintf("%d-%d-%d %d", y, m, d, mm)}, nil
}
```

## 2. Action 按钮的优化

假设页面有两个编辑器，第一个用于输入 JSON，第二个是只读的。点击按钮后，将第一个编辑器的内容转换为 YAML，并渲染到第二个编辑器中。

### 传统方式

使用 `ajax` 类型的行为按钮，代码如下：

```go
func init() {
    http.HandleFunc("/convert", func(w http.ResponseWriter, r *http.Request) {
        input, _ := io.ReadAll(r.Body)
        defer r.Body.Close()
        m := map[string]any{}
        json.Unmarshal(input, &m)
        // ...
        output := "age: 27"
        resp := model.Response{Msg: "转换成功", Data: model.Schema{"output": output}} // 这里的 key 值必须是第二个编辑器的 name
        data, _ := json.Marshal(resp)
        w.Write(data)
    })
}

func main() {
    index := app.Page().Body(
        app.Form().ColumnCount(2).Body(
            app.Editor().Language("json").Label("json").Name("input").Size("xxl"),
            app.Editor().Label("yaml").Label("yaml").Name("output").Size("xxl").ReadOnly(true),
        ).Actions(
            app.Action().Label("Convert").Level("primary").ActionType("ajax").Api(
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
        ),
    )
}
```

### 优化方式

`amisgo` 为行为按钮新增了 `Transform` 方法：

```go
func (a action) Transform(transform func(input any) (any, error), src, dst string) action
```

优化后的代码如下：

```go
app.Page().Body(
    app.Form().ColumnCount(2).Body(
        app.Editor().Language("json").Label("json").Name("input").Size("xxl"),
        app.Editor().Label("yaml").Label("yaml").Name("output").Size("xxl").ReadOnly(true),
    ).Actions(
        app.Action().Label("Convert").Level("primary").Transform(func(input any) (any, error) {
            // 将输入的 JSON 转换为 YAML
            output := "age: 27"
            return output, nil
        }, "input", "output"),
    ),
)
```

### 多对多转换

另有 `TransformMultiple` 方法支持从多个组件的输入值转换后渲染到多个组件。可以参考示例库中的 dev-toys，其中的 convert 组件使用了 Transform，生成多种类型的 hash 值使用了 TransformMultiple。

---

通过以上优化，代码变得更加简洁易读，减少了冗余的 HTTP 请求处理逻辑，提升了开发效率。
