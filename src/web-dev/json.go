package main

import "fmt"
import "encoding/json"

type User struct {
    FullName string `json:"Name"`
    Age      int
}

func main() {
    var jsonString = `{"Name": "john wick", "Age": 27}`
    var jsonData = []byte(jsonString)

    var data User

    var err = json.Unmarshal(jsonData, &data) // decode json
    if err != nil {
        fmt.Println(err.Error())
    }

    fmt.Println("user :", data.FullName)
    fmt.Println("age :", data.Age)

    fmt.Println(".")

    // taret decode dengan map[string]interface{}
    var data1 map[string]interface{}
    json.Unmarshal(jsonData, &data1)

    fmt.Println("user :", data1["Name"])
    fmt.Println("age :", data1["Age"])

    fmt.Println(".")

    // hasil decode dgn interface{}
    var data2 interface{}
    json.Unmarshal(jsonData, &data2)

    var decodedData = data2.(map[string]interface{})
    fmt.Println("user :", decodedData["Name"])
    fmt.Println("age :", decodedData["Age"])

    fmt.Println(".")

    // decode array json ke array object        
    var newJsonString = `[
        {"Name": "andy nur", "Age": 20},
        {"Name": "budi aji", "Age": 21}
    ]` 

    var newData []*User

    var newErr = json.Unmarshal([]byte(newJsonString), &newData)
    if newErr != nil {
        fmt.Println(newErr.Error())
    }

    fmt.Println("user 1:", newData[0].FullName)
    fmt.Println("user 2:", newData[1].FullName)

    fmt.Println(".")

    // encode objek ke json
    var object = []User{{"naufal rasyid", 19}, {"Achmad Zainu", 18}}
    var lastJsonData, lastErr = json.Marshal(object)
    if lastErr != nil {
        fmt.Println(lastErr.Error())
    }

    var lastJsonString = string(lastJsonData)
    fmt.Println(lastJsonString)
}