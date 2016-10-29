package main

import "fmt"
import "os"

func main() {
    // saat menjalankan go run di input juga arugments nya
    // go run arguments.go banana potato "ice cream"
    var argsRaw = os.Args
    fmt.Printf("-> %#v\n", argsRaw)

    var args = argsRaw[1:] // ambil slice arguments nya
    fmt.Printf("-> %#v\n", args)
}