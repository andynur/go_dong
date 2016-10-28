package main

import "fmt"
import "regexp"

func main() {
    var text = "banana burger soup"
    var regex, err = regexp.Compile(`[a-z]+`) // alphabet a-z

    if err != nil {
        fmt.Println(err.Error())
    }

    // cari yg sesuai regex & batasnya
    var res1 = regex.FindAllString(text, 2) 
    var res2 = regex.FindAllString(text, -1)
    fmt.Printf("%#v \n", res1)
    fmt.Printf("%#v \n", res2)

    fmt.Println(".")

    // cek string memenuhi pola regexp
    var isMatch = regex.MatchString(text)
    fmt.Println(isMatch)

    fmt.Println(".")

    // cari yg memenuhi kriteria regexp
    // apabila banyak, hanya return yg pertama saja
    var findStr = regex.FindString(text)
    fmt.Println(findStr)

    fmt.Println(".")

    // cari index string
    var idx = regex.FindStringIndex(text)
    fmt.Println(idx)

    var strTest = text[0:6] // tes eksekusi index
    fmt.Println(strTest)

    fmt.Println(".")

    // ganti semua string
    var strReplace = regex.ReplaceAllString(text, "potato")
    fmt.Println(strReplace)

    fmt.Println(".")

    // ganti string dgn kondisi tertentu
    var strReplaceFunc = regex.ReplaceAllStringFunc(text, func(each string) string {
        if each == "burger" { // jika ada kata burger
            return "potato"
        }
        return each
    })
    fmt.Println(strReplaceFunc)

    fmt.Println(".")

    // memisang string
    var strSplit = regex.Split(text, -1)
    fmt.Println("%#v \n", strSplit)

    fmt.Println(".")


}