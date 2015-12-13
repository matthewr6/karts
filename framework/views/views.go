package views

import (
    // "fmt"
    // "log"
    "net/http"
    "github.com/julienschmidt/httprouter"

    // "path/filepath"
    // "os"
    // "strings"
    // "io/ioutil"
    "html/template"
    "github.com/fatih/structs"
    // "github.com/imdario/mergo"
    "unicode"
)

const TemplateDirectories = "/templates"

type View struct {
    TemplateName string
    Get func (c *Context)
    GetContext func (map[string]interface{}) map[string]interface{}
}

func (view View) HandleGet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    context := Context{make(map[string]interface{}), w}
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

type Context struct {
    Data map[string]interface{}
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