package main

import "fmt"
import "os"

func main() {
    defer fmt.Println("halo") // eksekusi defer tdk akan muncul duluan exit
    os.Exit(1) // exit status 1
    fmt.Println("selamat datang")
}