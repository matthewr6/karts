package core

import (
    // "fmt"
    // "net/http"
    // "github.com/julienschmidt/httprouter"

    "../framework/views"
)

var Index = views.View{
    Get: func(c *views.Context) {
        views.TemplateRender("test.html", c)
    },
}