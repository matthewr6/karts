package core

import (
    "../framework/views"
    "fmt"
)

var _ = fmt.Println

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