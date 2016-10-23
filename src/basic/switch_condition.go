package main

import "fmt"

func main() {
    var point = 8

    switch point {
        case 8:
            fmt.Println("perfect")
        case 7:
            fmt.Println("awesome")
        default:
            fmt.Println("not bad")    
    }

    var new_point = 6

    switch new_point {
        case 8:
            fmt.Println("==perfect")
        case 7, 6, 5, 4:
            fmt.Println("==awesome")
        default:             
            fmt.Println("==not bad")
    }

    // switch gaya if - else
    var last_point = 6

    switch {
        case last_point == 8:
            fmt.Println("===perfect")
        case  (last_point < 8) && (last_point > 3):
            fmt.Println("===awesome")
            fallthrough  // melanjutkan pengecekan selanjutnya (defaultnya gak)
        default:
            { // bisa dibuat bersarang kalau ada banyak statement
                fmt.Println("===not bad")
                fmt.Println("===you need too learn more")
            }
    }
    
}