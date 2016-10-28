package main

import "encoding/base64"
import "fmt"

func main() {
    var data = "john wick"
    
    // encode dan decode data ke string
    var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println("encoded:", encodedString)

    var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
    var decodedString = string(decodedByte)
    fmt.Println("decoded:", decodedString)

    fmt.Println(".")

    // encode dan decode data
    var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
    base64.StdEncoding.Encode(encoded, []byte(data))
    var encodedStringNew = string(encoded)
    fmt.Println(encodedStringNew)

    var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encodedStringNew)))
    var _, err = base64.StdEncoding.Decode(decoded, encoded)

    if err != nil {
        fmt.Println(err.Error())
    }

    var decodedStringNew = string(decoded)
    fmt.Println(decodedStringNew)

    fmt.Println(".")

    // encode dan decode urk
    var dataUrl = "http://qodr.or.id/"

    var encodedUrl = base64.URLEncoding.EncodeToString([]byte(dataUrl))
    fmt.Println(encodedUrl)

    var decodedUrl, _ = base64.URLEncoding.DecodeString(encodedUrl)
    var decodedStringUrl = string(decodedUrl)
    fmt.Println(decodedStringUrl)
}