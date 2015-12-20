package views

import (
    "fmt"

    "path/filepath"
    "os"
    "strings"
    "io/ioutil"
)

func GetTemplate(name string) string {
    return GetTemplateContents(GetTemplatePath(name))
}

func GetTemplatePath(templatename string) string {
    searchdir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
    var matchedpath string
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