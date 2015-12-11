package main

import (
    "fmt"
    "log"
    // "strings"
    "net/http"
    "github.com/julienschmidt/httprouter"

    core "../app"
    "./views"
    "./staticfiles"
)

func main() {
    fmt.Println("Starting...")
    router := httprouter.New()
    
    staticlist := staticfiles.GetStaticfiles()
    for staticfile := range staticlist {
        router.GET(staticlist[staticfile].Servedpath, staticlist[staticfile].Serve)
    }

    for k := range core.Routes {
        route := core.Routes[k]
        route_type := fmt.Sprintf("%T", route)
        switch route_type {

            case "views.TemplateView":
                if core.Routes[k].(views.TemplateView).Get != nil {
                    router.GET(k, route.(views.TemplateView).Get)
                } else {
                    router.GET(k, route.(views.TemplateView).Render)
                }

            default:
                fmt.Printf("Unrecognized view type \"%s\"", route_type)
        }
    }

    log.Fatal(http.ListenAndServe(":3000", router))
}