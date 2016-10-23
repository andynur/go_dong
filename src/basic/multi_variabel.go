package main

import "fmt"

func main() {
    // berbagao macam teknik penulisan multi variabel
    var first, second, three string = "satu", "dua", "tiga"

    four, five, six := "empat", "lima", "enam"

    one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

    fmt.Printf("panggil %s %s %s!\n", first, second, three)
    fmt.Printf("panggil %s %s %s!\n", four, five, six)
    fmt.Printf("panggil juga %s %s %s!\n", one, isFriday, twoPointTwo, say)
}   