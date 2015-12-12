package views

import (
    "fmt"
    // "log"
    // "net/http"
    // "github.com/julienschmidt/httprouter"

    "path/filepath"
    "os"
    "strings"
    "io/ioutil"
    // "html/template"
    // "github.com/fatih/structs"
    // "github.com/imdario/mergo"
    // "unicode"
)

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