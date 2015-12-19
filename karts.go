package karts

import (
    "fmt"
    "log"
    "net/http"
    "github.com/julienschmidt/httprouter"

    // "github.com/firedrake969/karts/staticfiles"
    "github.com/firedrake969/karts/views"
)

func RunKarts(routes map[string]views.View) {
    fmt.Println("Starting...")
    router := httprouter.New()
    
    // staticlist := staticfiles.GetStaticfiles()
    // for staticfile := range staticlist {
    //     router.GET(staticlist[staticfile].Servedpath, staticlist[staticfile].Serve)
    // }

    for k := range routes {
        route := routes[k]
        router.GET(k, route.HandleGet)
        router.POST(k, route.HandlePost)
    }

    log.Fatal(http.ListenAndServe(":3000", router))
}