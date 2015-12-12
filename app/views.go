package core

import (
    // "fmt"
    "net/http"
    "github.com/julienschmidt/httprouter"

    "../framework/views"
)

var Index = views.View{
    Get: func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        views.TemplateRender("test.html", w, r, ps) // repeats... not good
    },
}