package main

import "fmt"
import "strconv"

func main() {
    var str = "124"
    var num, err = strconv.Atoi(str) // string to int

    if err == nil {
        fmt.Println(num)
    }

    var num1 = 5000
    var str1 = strconv.Itoa(num1) // int to string

    fmt.Println(str1)

    var num3, err3 = strconv.ParseInt(str, 10, 64) // str, basis num, type

    if err3 == nil {
        fmt.Println(num3)
    }

    var num4 = int64(24)
    var str4 = strconv.FormatInt(num4, 8) // convert to string

    fmt.Println(str4)

    var str5 = "24.12"
    var num5, err5 = strconv.ParseFloat(str5, 32) // convert to float 32 digit

    if err5 == nil {
        fmt.Println(num5)
    }

    var str6 = "true"
    var bul, err6 = strconv.ParseBool(str6) // convert to bool

    if err6 == nil {
        fmt.Println(bul)
    }

    // convert using casting
    var a float64 = float64(24)
    fmt.Println(a)
    
    var b int32 = int32(21.00)
    fmt.Println(b)

    // casting string x byte
    var text1 = "halo"
    var byte1 = []byte(text1)
    var byte2 = []byte{104, 97, 108, 111}
    var str7  = string(byte2)

    fmt.Printf("%d %d %d %d \n", byte1[0], byte1[1], byte1[2], byte1[3])
    fmt.Printf("%s \n", str7)

    // convert interface{} with type assertions technic
    var data = map[string]interface{}{
        "name":     "andy nur",
        "grade":    2,
        "height":   156.5,
        "isMale":   true,
        "hobbies":  []string{"eating", "sleeping"},
    }

    fmt.Println(".")
    fmt.Println(data["name"].(string))
    fmt.Println(data["grade"].(int))
    fmt.Println(data["height"].(float64))
    fmt.Println(data["isMale"].(bool))
    fmt.Println(data["hobbies"].([]string))

    fmt.Println(".")

    // mencari tahu tipe data interface{}
    for _, val := range data {
        switch val.(type) {  // casting ke tipe 'type' 
            case string:
                fmt.Println(val.(string))
            case int:
                fmt.Println(val.(int))
            case float64:
                fmt.Println(val.(float64))
            case bool:
                fmt.Println(val.(bool))
            case []string:
                fmt.Println(val.([]string))
            default:
                fmt.Println(val.(int))
        }
    }
}