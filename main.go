package main

import (
    "github.com/gorilla/mux"
    "os"
    "fmt"
    "net/http"
    "github.com/damanhanzo/true-size/controllers"
)

func main() {

    router := mux.NewRouter()

    router.HandleFunc("/api/sneaker/new", controllers.CreateSneaker).Methods("POST")
    router.HandleFunc("/api/sneakers", controllers.GetSneakers).Methods("GET")
    router.HandleFunc("api/sneaker/true-size", controllers.GetTrueSizeByBrandAndModel).Method("GET")

    port := os.Getenv("PORT")
    if port == "" {
        port = "4000" //localhost
    }

    fmt.Println(port)

    err := http.ListenAndServe(":" + port, router) //Launch the app, visit localhost:8000/api
    if err != nil {
        fmt.Print(err)
    }
}
