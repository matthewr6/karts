package core

import (
    "../framework/views"
)

var Index = views.View{
    Get: func(c *views.Context) {
        views.TemplateRender("test.html", c)
    },
    Form: views.Form{
        Fields: map[string]bool{
            "inputfield": true,
        },
    },
}