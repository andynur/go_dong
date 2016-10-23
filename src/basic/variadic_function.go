package main

import "fmt"
import "strings"

func main() {
    var avg = calculate(2, 4, 3, 5, 4, 3, 2, 4, 5, 3)
    var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
    fmt.Println(msg)

    fmt.Println("//////")   

    var hobbies = []string{"sleeping", "eating", "repeat"}
    yourHobbies("andy", hobbies...)
}

func calculate(numbers ...int) float64 {  // titik ... menerima inputan berupa slice
    var total int = 0
    for _, number := range numbers {
        total += number
    }

    var avg = float64(total) / float64(len(numbers))
    return avg
}

// parameter biasa digabung variadic parameter
// variadic harus ditempatkan diurutan akhir parameter
func yourHobbies(name string, hobbies ...string) {
    var hobbiesAsString = strings.Join(hobbies, ", ")

    fmt.Printf("Hello, my name is: %s\n", name)
    fmt.Printf("Hello, my hobbies are: %s\n", hobbiesAsString)
}