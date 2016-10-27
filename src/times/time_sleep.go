package main

import "fmt"
import "time"

func main() {
    fmt.Println("start")
    time.Sleep(time.Second * 4) // menghentikan program pada block tsb
    fmt.Println("after 4 seconds")
}