package main

import "fmt"
import "net/http"


func main() {
    // route target, callback
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "halo")
    })

    http.HandleFunc("/index", index) // panggil fungsi index

    fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(":8080", nil) // jalankan server di port 8080
}

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "apa kabar!")
}