package views

import (
    "fmt"
    // "log"
    "net/http"
    "github.com/julienschmidt/httprouter"

    "path/filepath"
    "os"
    "strings"
    "io/ioutil"
    "html/template"
    "github.com/fatih/structs"
    "github.com/imdario/mergo"
    "unicode"
)

const TemplateDirectories = "/templates"

type View struct {
    Get func (w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

// check then call the get method on this...
type TemplateView struct {
    View
    TemplateName string
}

func Upper(s string) string {
    a := []rune(s)
    a[0] = unicode.ToUpper(a[0])
    s = string(a)
    return s
}

func ParamsToMap(params httprouter.Params) map[string]interface{} {
    parammap := make(map[string]interface{})
    for k := range params {
        param := params[k]
        parammap[Upper(param.Key)] = param.Value
    }
    return parammap
}

// idea... have c Context struct? try that...
// http://stackoverflow.com/questions/12655464/can-functions-be-passed-as-parameters-in-go
type Context struct {
    Context map[string]interface{}
}

func (view TemplateView) Render(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    // fmt.Fprint(w, GetTemplate(view.TemplateName))
    file := GetTemplate(view.TemplateName)
    t, _ := template.New(view.TemplateName).Parse(file)
    context := make(map[string]interface{})
    requestcontext := structs.Map(r)
    urlcontext := ParamsToMap(ps)
    mergo.Merge(&context, requestcontext)
    mergo.Merge(&context, urlcontext)
    t.ExecuteTemplate(w, t.Name(), context)
}

func TemplateRender(name string, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    file := GetTemplate(name)
    t, _ := template.New(name).Parse(file)
    context := make(map[string]interface{})
    requestcontext := structs.Map(r)
    urlcontext := ParamsToMap(ps)
    mergo.Merge(&context, requestcontext)
    mergo.Merge(&context, urlcontext)
    t.ExecuteTemplate(w, t.Name(), context)
}


// TEMPLATES

func GetTemplate(name string) string {
    return GetTemplateContents(GetTemplatePath(name))
}

func GetTemplatePath(templatename string) string {
    var matchedpath string
    searchdir := "../app"
    filepath.Walk(searchdir, func(fp string, fi os.FileInfo, err error) error {
        if err != nil {
            fmt.Println(err) // can't walk here,
            return nil       // but continue walking elsewhere
        }
        if !!fi.IsDir() {
            return nil // not a file.  ignore.
        }
        fp = strings.Replace(fp, "\\", "/", -1)
        matched, err := filepath.Match("*" + TemplateDirectories + "*/" + templatename, fp)
        if err != nil {
            fmt.Println(err) // malformed pattern
            return err       // this is fatal.
        }
        if matched {
            matchedpath = fp
        }
        return nil
    })
    return matchedpath
}

func GetTemplateContents(templatepath string) string {
    contents, _ := ioutil.ReadFile(templatepath)
    return string(contents)
}