package main

import "fmt"
import "html/template"
import "net/http"
import "path/filepath"

func main() {
    realPath, _ := filepath.Abs("src/web-dev")

    // route target, callback
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        var data = map[string]string{
            "Name":     "andy nur",
            "Message":  "have a nice day",
        }

        var t, err = template.ParseFiles(realPath + "/template.html")
        if err != nil {
            fmt.Println(err.Error())
        }

        t.Execute(w, data)
    })

    http.HandleFunc("/index", index) // panggil fungsi index

    fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(":8080", nil) // jalankan server di port 8080
}

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "apa kabar!")
}