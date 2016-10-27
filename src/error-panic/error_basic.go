package main

import "fmt"
import "strconv"

func main() {
    var input string
    fmt.Print("Type some number: ")
    fmt.Scanln(&input) // membaca inputan user sebelum dienter

    var number int
    var err error
    number, err = strconv.Atoi(input) // konversi string ke int

    if err == nil {
        fmt.Println(number, "is number")
    } else {
        fmt.Println(input, "is not number")
        fmt.Println(err.Error()) // tampilkan error
    }
}