package main

import "fmt"

func main() {
    // horizontal array
    var fruits = [4]string{"apple", "grape", "banana", "melon"}

    // vertical array
    var days = [4]string{
        "sunday",
        "monday",
        "tuesday",
        "friday", // coma in end of array is required!!!!!
    }

    fmt.Println("Jumlah element \t\t", len(fruits))
    fmt.Println("Isi semua element \t", fruits)
    fmt.Println("Isi semua element \t", days)

    fmt.Println("///////")

    // without amount an array
    var numbers = [...]int{2, 3, 4, 5, 2, 1}
    fmt.Println("Jumlah element \t\t", len(numbers))
    fmt.Println("Isi semua element \t", numbers)    

    fmt.Println("///////")

    // array multidimention
    var numbers1 = [2][3]int{[3]int{3, 2, 1}, [3]int{3, 4, 5}}
    var numbers2 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
    fmt.Println("Numbers 1 :", numbers1)
    fmt.Println("Numbers 2 :", numbers2)

    fmt.Println("///////")

    // array loop
    var new_fruits = [4]string{"apple", "grape", "banana", "melon"}

    for i := 0; i < len(new_fruits); i++ {
        fmt.Printf("element %d : %s\n", i, new_fruits[i])
    }

    fmt.Println("///////")

    // array loop with range
    for i, fruit := range new_fruits {
        fmt.Printf("element %d : %s\n", i, fruit)        
    }

    // array loop - only access value
    for _, fruit := range new_fruits {
        fmt.Printf("nama buah : %s\n", fruit)
    } 
    
    // array loop - only access index
    for i := range new_fruits {
        fmt.Printf("no buah : %d\n", i)
    }     

    fmt.Println("///////")

    // allocation array element
    var last_fruits = make([]string, 2)
    last_fruits[0] = "apple"
    last_fruits[1] = "manggo"

    fmt.Println(last_fruits)    
}