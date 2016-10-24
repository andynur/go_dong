package library

import "fmt"

// modifier Public ditandai dengan huruf Kapital diawal fungsi
func SayHello(name string) {
    fmt.Println("hello")
    introduce(name)
}    

// modifier Private ditandai dengan huruf kecil diawal fungsi
func introduce(name string) {
    fmt.Println("nama saya", name)
}