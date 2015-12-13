package core

import (
    "../framework/views"
)

var Index = views.View{
    Get: func(c *views.Context) {
        views.TemplateRender("test.html", c)
    },
    // TemplateName: "test.html",
}