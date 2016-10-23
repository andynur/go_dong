package main

import (
    "fmt"
    "math/rand"
    "strings"
    "time"
)   

func main() {
    var names = []string{"Andy", "Nur"}
    printMessage("halo", names)

    fmt.Println("\n///////\n")

    rand.Seed(time.Now().Unix()) // generate random number
    var randomValue int

    randomValue = randomWithRange(2, 10)
    fmt.Println("random number 1:", randomValue)
    randomValue = randomWithRange(2, 10)
    fmt.Println("random number 2:", randomValue)
    randomValue = randomWithRange(2, 10)
    fmt.Println("\n///////\n")
    fmt.Println("random number 3:", randomValue)

    fmt.Println("\n///////\n")

    divideNumber(10, 2)    
    divideNumber(4, 0)    
    divideNumber(8, -4)    
}

func printMessage(message string, arr []string) {
    var nameString = strings.Join(arr, " ")
    fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int {
    var value = rand.Int() % (max - min + 1) + min
    return value
}

func divideNumber(m, n int) {
    if n == 0 {
        fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
        return // return disini berfungsi jg utk menghentikan proses
    }

    var res = m / n
    fmt.Printf("%d / %d = %d\n", m, n, res)
}