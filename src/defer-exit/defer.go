package main

import "fmt"

func main() {
    defer fmt.Println("halo") // defer utk mengakhirkan eksekusi kode
    fmt.Println("selamat datang")
}