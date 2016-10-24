package initLibrary

import "fmt"

var Student = struct {
    Name  string
    Grade int
}{} // anonymous struct dgn default inisialisasi kosong {}

// fungsi init() dipanggil duluan daripada main()
func init() {
    Student.Name  = "Arif Law"
    Student.Grade = 1

    fmt.Println("---> library/initLibrary.go imported")
}