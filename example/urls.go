package main

import (
    "github.com/firedrake969/karts"
    "github.com/firedrake969/karts/views"
)

var Index = views.View{
    Get: func(c *views.Context) {
        views.TemplateRender("test.html", c)
    },
    Form: views.Form{
        Fields: map[string]views.Field{
            "inputfield": {
                Required: true,
                Type: "text",
            },
            "selection": {
                Required: true,
                Type: "multiple",
            },
        },
        SuccessUrl: "/test/someurl",
    },
}

var Routes = map[string]views.View {
    "/": Index,
    "/test/:param": Index,
}

func main() {
    karts.RunKarts(Routes)
}