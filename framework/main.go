package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/julienschmidt/httprouter"

    core "../app"
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
        router.GET(k, route.HandleGet)
    }

    log.Fatal(http.ListenAndServe(":3000", router))
}