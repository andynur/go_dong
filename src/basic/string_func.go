package main

import "fmt"
import "strings"

func main() {
    // cek string return bool
    var isExists = strings.Contains("andy nur", "nur")
    fmt.Println(isExists)

    fmt.Println(".")

    // cek string diawali kata tertentu
    var isPrevix1 = strings.HasPrefix("andy nur", "an")    
    var isPrevix2 = strings.HasPrefix("andy nur", "nu")
    fmt.Println(isPrevix1)
    fmt.Println(isPrevix2)   

    fmt.Println(".")

    // cek string diakhir kata tertentu
    var isSuffix1 = strings.HasSuffix("andy nur", "nu")
    var isSuffix2 = strings.HasSuffix("andy nur", "ur")
    fmt.Println(isSuffix1)
    fmt.Println(isSuffix2)    

    fmt.Println(".")

    var howMany = strings.Count("ethan hunt", "t")
    fmt.Println(howMany)

    fmt.Println(".")    

    // mencari index string
    var index1 = strings.Index("ethan hunt", "hu")
    var index2 = strings.Index("ethan hunt", "n")
    fmt.Println(index1)
    fmt.Println(index2) // jika dideteksi 2 diambil yg paling kecil

    fmt.Println(".")

    // mengganti bagian string dgn string lain
    // param akhir utk menentukan jml string yg diganti
    // apabila -1 mengganti semua string yg dicari
    var text = "banana"
    var find = "a"
    var replaceWith = "o"

    var newText1 = strings.Replace(text, find, replaceWith, 1)
    fmt.Println(newText1)
    
    var newText2 = strings.Replace(text, find, replaceWith, 2)
    fmt.Println(newText2)

    var newText3 = strings.Replace(text, find, replaceWith, -1)
    fmt.Println(newText3)

    fmt.Println(".")    

    // mengulang string sebanyak data yg ditentukan
    var str = strings.Repeat("na", 4)
    fmt.Println(str)

    fmt.Println(".")

    // memisah string berdasar param menjadi array
    var string1 = strings.Split("the dark knight", " ")
    var string2 = strings.Split("batman", "")
    fmt.Println(string1)    
    fmt.Println(string2)

    fmt.Println(".")

    // menggabung array string
    var data = []string{"banana", "papaya", "tomato"}
    var strJoin = strings.Join(data, "-")
    fmt.Println(strJoin)

    fmt.Println(".")
    
    // mengubah string menjadi lower case
    var strLower = strings.ToLower("aLay bEutZz")
    fmt.Println(strLower)

    fmt.Println(".")
    
    // mengubah string menjadi upper case
    var strUpper = strings.ToUpper("mAem Tah")
    fmt.Println(strUpper)

    fmt.Println(".")    
}