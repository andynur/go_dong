package main

import "fmt"
import "flag"

func main() {
    // fungsinya sama dgn arguments utk parameterize dgn key - value
    // go run flag.go banana potato "ice cream"
    // param (key, default, description) - return pointer

     // metode 1 inisialisasi variabel
    var name = flag.String("name", "anonymous", "type your name")

    // metode 2 parameternya pointer
    var age int64    
    flag.Int64Var(&age, "age", 25, "type your age")

    flag.Parse()
    fmt.Printf("name\t: %s\n", *name) // di dereference dulu
    fmt.Printf("age\t: %d\n", age) // tdk perlu dereference
}