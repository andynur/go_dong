package main

import "fmt"

func main() {
    // TIPE DATA
    var positiveNumber uint8 = 89
    var negativeNumber = -124342644
    
    fmt.Printf("bilangan positif: %T\n", positiveNumber)
    fmt.Printf("bilangan negatif: %T\n", negativeNumber)

    var decimalNumber = 2.62

    fmt.Printf("bilangan desimal: %f\n", decimalNumber)    
    fmt.Printf("bilangan desimal koma 3: %.3f\n", decimalNumber)

    var exist bool = true
    fmt.Printf("ada atau gk? %t \n", exist)

    // penulisan string dgn ``, bisa disusupi ""
    var message = `Nama saya "Andy Nur"
    Salam kenal,
    Mari belajar "Golang".`
    fmt.Println(message)   

    // KONSTANTA
    const uniqueName string = "badala"
    fmt.Print("halo ", uniqueName, "!\n")

    
}   