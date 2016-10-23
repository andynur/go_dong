package main

import "fmt"

func main() {
    // basic
    for i := 0; i < 5; i++ {
        fmt.Println("Angka", i)
    }

    fmt.Println("///////")

    // just condition
    var i = 5

    for i > 0 {
        fmt.Println("Angka", i)
        i--
    }

    fmt.Println("///////")

    // without argument
    var j = 0
    
    for {
        fmt.Println("Angka", j)
        j++

        if j == 5 {
            break
        }
    }

    fmt.Println("///////")

    // access outer loop with label
    outerLoop:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i == 2 {
                break outerLoop
            }
            fmt.Print("matriks [", i, "][", j, "]", "\n")
        }
    }

}