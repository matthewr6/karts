package views

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "html/template"
    "github.com/fatih/structs"
    "unicode"
    "fmt"
    "net/url"
)

var _ = fmt.Println

const TemplateDirectories = "/templates"

type View struct {
    TemplateName string
    Form Form
    Get func (c *Context)
    Post func (c *Context, e []string)
    GetContext func (map[string]interface{}) map[string]interface{}
}

func (view View) HandleGet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    context := Context{
        Data: make(map[string]interface{}),
        Writer: w,
    }
    context.Data["URL"] = UrlParamsToMap(ps)
    context.Data["Request"] = structs.Map(r)
    if view.GetContext != nil {
        context.Data = view.GetContext(context.Data)
    }
    if view.Get != nil {
        view.Get(&context)
    } else if view.TemplateName != "" {
        TemplateRender(view.TemplateName, &context)
    }
}

// flow:
// check Post func
// then handleget if no post
func (view View) HandlePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    r.ParseForm()
    context := Context{
        Data: make(map[string]interface{}),
        Form: r.Form,
        Writer: w,
    }
    context.Data["URL"] = UrlParamsToMap(ps)
    context.Data["Request"] = structs.Map(r)
    if view.GetContext != nil {
        context.Data = view.GetContext(context.Data)
    }
    var errors []string
    if &view.Form != nil {
        errors = view.Form.HandleValidate(r.Form)
    }
    if view.Post != nil {
        view.Post(&context, errors)
    } else {
        view.HandleGet(w, r, ps) // default - NOT RECOMMENDED
    }
}

type Context struct {
    Data map[string]interface{}
    Form url.Values
    Writer http.ResponseWriter
}

func Upper(s string) string {
    a := []rune(s)
    a[0] = unicode.ToUpper(a[0])
    s = string(a)
    return s
}

func UrlParamsToMap(params httprouter.Params) map[string]interface{} {
    parammap := make(map[string]interface{})
    for k := range params {
        param := params[k]
        parammap[Upper(param.Key)] = param.Value
    }
    return parammap
}

func TemplateRender(name string, c *Context) {
    file := GetTemplate(name)
    t, _ := template.New(name).Parse(file)
    t.ExecuteTemplate(c.Writer, t.Name(), c.Data)
}