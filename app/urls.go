package core

import (
    // "github.com/julienschmidt/httprouter"

    // "../framework/views"
)

var Routes = map[string]interface{} {
    "/": Index,
    "/test/:param": Index,
}