package core

import (
    // "github.com/julienschmidt/httprouter"

    "../framework/views"
)

var Routes = map[string]views.View {
    "/": Index,
    "/test/:param": Index,
}