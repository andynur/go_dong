package main

import "fmt"
import "runtime"

func print(till int, message string) {
    for i := 0; i < till; i++ {
        fmt.Println((i + 1), message)
    }
}

func main() {
    runtime.GOMAXPROCS(2) // menentukan jml processor yg aktif

    // goroutine merupakan mini thread
    // diawali dengan syntax `go`
    go print(5, "halo")
    print(5, "apa kabar")

    var input string
    fmt.Scanln(&input) // process blocking hingga di-enter
}