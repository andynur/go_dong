package main

import (
    "errors"
    "fmt"
    "strings"
)

func main() {
    var name string

    fmt.Print("Type your name: ")
    fmt.Scanln(&name)

    if valid, err := validate(name); valid {
        fmt.Println("halo", name)
    } else {
        panic(err.Error()) // dgn panic error lebih detail
    }
}

// custom error untuk validasi data kosong
// inputan string dan 2 kembalian boolean dan error
func validate(input string) (bool, error) {
    if (strings.TrimSpace(input) == "") {
        return false, errors.New("cannot be empty")
    }

    return true, nil
}