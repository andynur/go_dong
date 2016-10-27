package main

import "fmt"
import "reflect"

type student struct {
    Name  string
    Grade int
}

func main() {
    var number = 23
    var reflectValue = reflect.ValueOf(number)  // mengambil nilai

    fmt.Println("tipe variabel  :", reflectValue.Type()) // mengambil nilai tipe data

    if reflectValue.Kind() == reflect.Int {  // mengecek tipe datanya = int 
        fmt.Println("nilai variabel :", reflectValue.Int())
    }

    // mengambil nilai dgn interface
    // bisa mengakses segala jensi nilai
    fmt.Println("nilai variabel interface :", reflectValue.Interface().(int)) 

    fmt.Println("")

    var	s1	=	&student{Name:	"wick",	Grade:	2}
    s1.getPropertyInfo()

    fmt.Println("")

    var s2 = &student{Name: "andy nur", Grade: 2}
    fmt.Println("nama :", s2.Name)

    var newReflectValue = reflect.ValueOf(s2)
    var method = newReflectValue.MethodByName("SetName")

    method.Call([]reflect.Value{
        reflect.ValueOf("nur"),
    })    

    fmt.Println("nama :", s2.Name)
}

// akses informasi property var obyek
func (s *student) getPropertyInfo() {
    var reflectValue = reflect.ValueOf(s)

    if reflectValue.Kind() == reflect.Ptr { // Ptr = pointer
        reflectValue = reflectValue.Elem() // cari obyel reflect aslinya
    }

    var reflectType = reflectValue.Type()

    for i := 0; i < reflectValue.NumField(); i++ {
        fmt.Println("nama       :", reflectType.Field(i).Name)
        fmt.Println("tipe data  :", reflectType.Field(i).Type)
        fmt.Println("nilai      :", reflectValue.Field(i).Interface())
        fmt.Println(".")
    }
}

// akses informasi method var obyek
func (s *student) SetName(name string) {
    s.Name = name
}

