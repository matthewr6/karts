package main

import (
    "github.com/firedrake969/karts"
    "github.com/firedrake969/karts/views"
)

var Routes = map[string]views.View {
    "/": Index,
    "/test/:param": Index,
}

func main() {
    karts.RunKarts(Routes)
}