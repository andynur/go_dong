package main

import "fmt"

func main() {
    var numberA int = 4
    var numberB *int = &numberA // deklarasi var pointer menggunakan *type_data

    fmt.Println("numberA (value)    :", numberA)
    fmt.Println("numberA (address)  :", &numberA) // & mengambil nilai pointer var biasa
    fmt.Println("numberB (value)    :", *numberB) // * mengambil nilai asli var pointer
    fmt.Println("numberB (address)  :", numberB)

    fmt.Println("")

    numberA = 5

    fmt.Println("numberA (value)    :", numberA)
    fmt.Println("numberA (address)  :", &numberA)
    fmt.Println("numberB (value)    :", *numberB)
    fmt.Println("numberB (address)  :", numberB)

    fmt.Println("")

    var number = 4
    fmt.Println("before : ", number)    

    change(&number, 10)
    fmt.Println("after : ", number)
}

// pointer bisa menjadi parameter
func change(original *int, value int) { 
    *original = value
}