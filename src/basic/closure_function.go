package main

import "fmt"
import "strings"

func main() {
    var getMinMax = func(n []int) (int, int) {
        var min, max int
        for i, e := range n {
            switch {
                case i == 0:
                    max, min = e, e
                case e > max:
                    max = e
                case e < min:
                    min = e
            }
        }
        return min, max
    }

    var numbers = []int{2, 3, 4, 5, 1, 2, 3}
    var lastMin, lastMax = getMinMax(numbers)
    fmt.Printf("data : %v\nlastMin : %v\nlastMax : %v\n", numbers, lastMin, lastMax)
    // %v diatas utk memanggil data yg tipenya dinamis

    var newNumbers = func(min int) []int {
        var r []int
        for _, e := range numbers {
            if e < min {
                continue
            }
            r = append(r, e)
        }
        return r
    }(3) // IIFE teknik, ada parameter tambahan ()

    fmt.Println("original number:", numbers)
    fmt.Println("filtered number:", newNumbers)
    
    fmt.Println("//////")

    var max = 3
    var howMany, getNumbers = findMax(numbers, max)
    var theNumbers = getNumbers()

    fmt.Println("numbers\t:", numbers) 
    fmt.Printf("find \t: %d\n\n", max) 

    fmt.Println("found \t:", howMany) 
    fmt.Println("value \t:", theNumbers)

    fmt.Println("//////")

    var data = []string{"andy", "budi", "ari"}
    var dataContains0 = filter(data, func(each string) bool {
        return strings.Contains(each, "i") // mencari huruf i dari each
    }) 
    var dataLength5 = filter(data, func(each string) bool {
        return len(each) == 3
    })

    fmt.Println("data asli \t\t:", data)
    fmt.Println("filter ada huruf \"i\"\t:", dataContains0)
    fmt.Println("filter jumlah huruf \"5\"\t:", dataLength5)
}

func findMax(numbers []int, max int) (int, func() []int) {
    var res []int
    for _, e := range numbers {
        if e <= max {
            res = append(res, e)
        }
    }

    return len(res), func() []int { // parameter ke-2 fungsi kosong
        return res
    }
}

type FilterCallback func(string) bool // alternative penulisan callback dgn type

// closure sebagai parameter
func filter(data []string, callback FilterCallback) []string {
    var result [] string
    for _, each := range data {
        if filtered := callback(each); filtered {
            result = append(result, each)
        }
    }
    return result
}