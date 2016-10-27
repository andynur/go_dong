package main

import "fmt"
import "runtime"

func main() {
    runtime.GOMAXPROCS(2)

    var messages = make(chan string)
    var names = []string{"wick", "hunt", "bourne"}

    for _, each := range names {
        go func(who string) {
            var data = fmt.Sprintf("hello %s", who)
            messages <- data
        }(each)
    }

    for i := 0; i < 3; i++ {
        printMessage(messages)
    }
}

func printMessage(what chan string) {
    fmt.Println(<-what)
}