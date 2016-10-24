package main

import "fmt"
import "strings"


type student struct {
    name  string 
    grade int
}

func (s student) sayHello() {
    fmt.Println("halo", s.name)
}

func (s student) getNameAt(i int) string {
    return strings.Split(s.name, " ")[i-1]
}

// method pointer
func (s *student) sayPointHello() {
    fmt.Println("halo", s.name)
}

func main() {
    var s1 = student{"andy nur", 21}
    s1.sayHello()

    var name = s1.getNameAt(1)
    fmt.Println("nama panggilan :", name)

    var s2 = student{"budi irawan", 12}
    s2.sayPointHello()
}

