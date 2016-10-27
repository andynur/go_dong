package main

import "fmt"

type student struct {
    name        string
    height      float64
    age         int32
    isGraduated bool
    hobbies     []string
}

var data = student{
    name:        "wick",
    height:      182.5,
    age:         26,
    isGraduated: false,
    hobbies:     []string{"eating", "sleeping"},
}

func main() {
    fmt.Printf("%b\n", data.age) // biner numerik
    fmt.Printf("%c\n", 1400) // ascii char
    fmt.Printf("%d\n", data.age) // numerik menjadi string (1-0)
    fmt.Printf("%e\n", data.height) // numerik desimal
    fmt.Printf("%.2f\n", data.height) // numerik desimal dengan lebar digit 
    fmt.Printf("%g\n", 0.1212734) // numerik desimal sesuai  lebar datanya  
    fmt.Printf("%o\n", data.age) // oktal numerik
    fmt.Printf("%p\n", &data.name) // pointer referensi variabel
    fmt.Printf("%q\n", `" name \ height "`) // escape string
    fmt.Printf("%s\n", data.name) // format data string
    fmt.Printf("%t\n", data.isGraduated) // format data boolean
    fmt.Printf("%T\n", data.hobbies) // ambil tipe data variabel
    fmt.Printf("%v\n", data) // string nilai data hasilnya
    fmt.Printf("%+v\n", data) // format struct + property dan nilainya
    fmt.Printf("%#v\n", data) // format struct beserta deklarasi objek
    fmt.Printf("%x\n", data.age) // heksadesimal
    fmt.Printf("%%\n") // menulis karakter % :)
}