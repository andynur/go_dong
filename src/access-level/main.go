package main

import "access-level/library"
import "access-level/structLibrary"
import "access-level/initLibrary"
import f "fmt" // bisa menggunakan alias ketika import

func main() {
    library.SayHello("andy nur")

    f.Println()

    var s1 = structLibrary.Student{"budi", 21}
    f.Println("name ", s1.Name)
    f.Println("grade", s1.Grade)

    f.Println()

    /* 
        partialHello("naufal acid") 
        
        fungsi daiatas dari partial.go
        harus dipanggil bersamaan dengan file main.go
        ex: go run main.go partial.go 
    */

    // panggil fungsi dari package initlibrary
    f.Printf("Name  : %s\n", initLibrary.Student.Name) 
    f.Printf("Grade : %d\n", initLibrary.Student.Grade)
    
}