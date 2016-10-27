package main

import "fmt"
import "time"

func main() {
    var timer = time.NewTimer(4 * time.Second)
    fmt.Println("start")
    <-timer.C // return objek *time.Timer dgn method C
    fmt.Println("after 4 seconds")
}