package main

import "fmt"

type student struct {
    name string
    grade int
} 

type person struct {
    name string
    age  int
}

type newStudent struct {
    grade int
    person // embed properti dari struct person
}

func main() {
    var s1 = student{}

    // macam-macam inisialisasi obyek
    s1.name = "wick"
    s1.grade = 2

    var s2 = student{"ethan", 2}

    var s3 = student{name: "jason"} // dgn cara ini tdk mesti berurutan

    fmt.Println("student 1 :", s1.name)
    fmt.Println("student 2 :", s2.name)
    fmt.Println("student 3 :", s3.name)

    fmt.Println()

    // variabel obyek pointer
    var varS1 *student = &s1

    fmt.Println("student name :", s1.name)
    fmt.Println("student var name :", varS1.name)

    varS1.name = "andy" // ubah value varS1.name

    fmt.Println("student name :", s1.name)

    fmt.Println("student var name :", varS1.name)

    fmt.Println()

    // embedded struct: penurunan properti ke struct lain
    var embedS1 = newStudent{}
    embedS1.name = "wick"
    embedS1.age = 21
    embedS1.grade = 2

    fmt.Println("embed name  :", embedS1.name)
    fmt.Println("embed age   :", embedS1.age)
    fmt.Println("embed age   :", embedS1.person.age) // akses dari parent struct-nya
    fmt.Println("embed grade :", embedS1.grade)

    fmt.Println()

    // anonymouse struct : utk dipakai sekali saja
    var anonS1 = struct {
        person
        grade int
    }{} // tanpa ada inisialisasi diawal, ttp ditulis {}

    var anonS2 = struct {
        person
        grade int
    }{ // dgn inisialisasi diawal
        person: person{"tes", 20},
        grade: 2,
    } 

    anonS1.person = person{"budi", 21}
    anonS1.grade  = 2

    fmt.Println("anonS1 name  :", anonS1.person.name)
    fmt.Println("anonS1 age   :", anonS1.person.age)
    fmt.Println("anonS1 grade :", anonS1.grade)
    fmt.Println("anonS2 name  :", anonS2.person.name)
    fmt.Println("anonS2 age   :", anonS2.person.age)
    fmt.Println("anonS2 grade :", anonS2.grade)

    fmt.Println()
    
    // kombinasi slice & struct
    var allStudents = []person{
        {name: "sultan", age: 18},
        {name: "rohmat", age: 20},
        {name: "achmad", age: 19},
    }

    for _, student := range allStudents {
        fmt.Println(student.name, "age is", student.age)
    }

    fmt.Println()

    // inisialisasi langsung anonymous struct
    var anonAllStudents = []struct {
        person
        grade int
    }{
        {person: person{"budi", 21}, grade: 2},
        {person: person{"naufal", 19}, grade: 3},
        {person: person{"akul", 21}, grade: 1},
    }

    for _, student := range anonAllStudents {
        fmt.Println(student)
    }

    fmt.Println()

    // deklarasi anonymous struct dgn var
    var anonVarStudent struct {
        person
        grade int
    }

    anonVarStudent.person = person{"gufron", 19}
    anonVarStudent.grade = 3

    // tag properti, utk keperluan json encode/decode
    type newPerson struct {
        name string `tag1`
        age  int    `tag2`
    }
}