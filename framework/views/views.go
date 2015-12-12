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
    Get func (w http.ResponseWriter, r *http.Request, ps httprouter.Params)
    GetContext func (map[string]interface{}) map[string]interface{}
}

// idea... have c Context struct? try that...
// http://stackoverflow.com/questions/12655464/can-functions-be-passed-as-parameters-in-go
func (view View) HandleGet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    context := make(map[string]interface{})
    context["URL"] = UrlParamsToMap(ps)
    context["Request"] = structs.Map(r)
    if view.GetContext != nil {
        context = view.GetContext(context)
    }
    if view.Get != nil {
        view.Get(w, r, ps)
    }
    // return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    //     //things
    // }
}

type Context struct {
    Context map[string]interface{}
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

func TemplateRender(name string, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    file := GetTemplate(name)
    t, _ := template.New(name).Parse(file)
    context := make(map[string]interface{})
    requestcontext := structs.Map(r)
    urlcontext := UrlParamsToMap(ps)
    context["Request"] = requestcontext
    context["URL"] = urlcontext
    t.ExecuteTemplate(w, t.Name(), context)
}