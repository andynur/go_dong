package main

import "fmt"
import "runtime"

func main() {
    runtime.GOMAXPROCS(2)

    // channel utk menghubungkan goroutine
    var messages = make(chan string)

    var sayHelloTo = func(who string) {
        var data = fmt.Sprintf("hello %s", who)
        messages <- data
    }

    go sayHelloTo("andy nur")
    go sayHelloTo("budi aji")
    go sayHelloTo("saipul hidayat")

    var message1 = <-messages
    fmt.Println(message1)

    var message2 = <-messages
    fmt.Println(message2)
    
    var message3 = <-messages
    fmt.Println(message3)
}