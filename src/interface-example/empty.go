package main

import "fmt"
import "strings"

func main() {
    var secret interface{}  // penggunaan interface kosong, menjadi tipe data

    secret = "ethan hunt"
    fmt.Println(secret)

    secret = []string{"apple", "manggo", "banana"}
    fmt.Println(secret)

    secret = 12.3
    fmt.Println(secret)

    fmt.Println()

    // casting variabel
    var newSecret interface{}

    newSecret = 2
    var number = newSecret.(int) * 10
    fmt.Println(newSecret, "multiplied by 10 is :", number)

    newSecret =[]string{"apple", "banana", "manggo"}
    var gruits = strings.Join(newSecret.([]string), ", ")
    fmt.Println(gruits, "is my favorite fruits.") 

    fmt.Println()

    // casting variabel ke obyek
    var objSecret interface{} = &person{name: "andy", age: 21}
    var name = objSecret.(*person).name 
    fmt.Println(name)

    fmt.Println()

    // kombinasi slice, map dan interface{}
    var newPerson = []map[string]interface{}{
        {"name": "budi", "age": 19},
        {"name": "arif", "age": 23},
        {"name": "naufal", "age": 20},
    }

    for _, each := range newPerson {
        fmt.Println(each["name"], "age is", each["age"])
    }

    fmt.Println()

    // slice dan interface utk array
    var newFruits = []interface{}{
        map[string]interface{}{
            "nama": "manggo", "total": 10},
            []string{"manggo", "grape", "papaya"},
            "orange",
    }

    for _, each := range newFruits {
        fmt.Println(each)
    }
}

type person struct {
    name string
    age  int
}