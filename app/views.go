package core

import (
    // "fmt"
    "net/http"
    "github.com/julienschmidt/httprouter"

    "../framework/views"
)

// func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//     fmt.Fprint(w, "Welcome!\n")
// }

// func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//     fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
// }

var Index = views.View{
    Get: func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        views.TemplateRender("test.html", w, r, ps) // repeats... not good
    },
}