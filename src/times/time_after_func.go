package main

import "fmt"
import "time"

func main() {
    var ch = make(chan bool)

    // param: durasi timer, callback (callback dieksekusi jika sdh memenuhi durasi)
    time.AfterFunc(4*time.Second, func() {
        fmt.Println("expired")
        ch <- true
    })

    fmt.Println("start")
    <-ch
    fmt.Println("after 4 seconds")
}