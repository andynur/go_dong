package main

import "fmt"
import "time"

func main() {
    fmt.Println("start")
    <-time.After(time.Second * 4) // menghentikan program dgn return channel
    fmt.Println("after 4 seconds")
}