package main

import "fmt"
import "net/url"

func main() {
    var urlString = "http://depeloper.com:80/hello?name=john wick&age=27"
    var u, e = url.Parse(urlString) // parsing url
 
    if e != nil {
        fmt.Println(e.Error())
    }

    fmt.Printf("url \t\t : %s\n", urlString)

    fmt.Printf("protocol \t : %s\n", u.Scheme)
    fmt.Printf("host \t\t : %s\n", u.Host)
    fmt.Printf("path \t\t : %s\n", u.Path)

    var name = u.Query()["name"][0] // parsing query
    var age  = u.Query()["age"][0]

    fmt.Printf("name \t\t : %s\n", name)
    fmt.Printf("age \t\t : %s\n", age)

}