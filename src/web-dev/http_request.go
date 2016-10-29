package	main

import "fmt"
import "net/http"
import "encoding/json"

var	baseURL	= "http://localhost:8080"

type student struct {
    ID		string
    Name	string
    Grade	int
}

func main() {
    var users = fetchUsers()

    for _, each := range users {
        fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
    }    
}

// menerima request dari /users
func fetchUsers() []*student {
    var err error
    var client = &http.Client{} // objek utk eksekusi request
    var data []*student

    // request_type, url_target, payload (jika ada)
    // return http.Request
    request, err := http.NewRequest("POST", baseURL+"/users", nil)
    if err != nil {
        fmt.Println(err.Error())
        return data
    }

    // eksekusi requestnya
    // return http.Response
    response, err := client.Do(request)
    if err != nil {
        fmt.Println(err.Error())
        return data
    }
    defer response.Body.Close() // close stelah tdk dipakai

    // ambil data response (string)
    // convert to JSON
    err = json.NewDecoder(response.Body).Decode(&data) 
    if err != nil {
        fmt.Println(err.Error())
        return data
    }

    return data
}