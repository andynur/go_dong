package main

import "fmt"

func main() {
    // map have key and value -> map[key]value
    var chicken map[string]int
    chicken = map[string]int {} // inisialisasi default harus ada 

    chicken["januari"] = 50
    chicken["februari"] = 40
    
    fmt.Println("januari",  chicken["januari"])
    fmt.Println("mei",      chicken["mei"]) // blm diinisialisasi kembali ke nilai default

    fmt.Println("//////")

    // inisialisasi vertical
    var chicken1 = map[string]int{"januari": 50, "februari": 40}

    // inisialisasi horizontal
    var chicken2 = map[string]int{
        "januari":  50,
        "februari": 40,
        "maret":    30,
        "april":    20,
        "mei":      10, // coma in end of item is required!
    }

    // inisialisasi make dan new
    var chicken3 = make(map[string]int)
    var chicken4 = *new(map[string]int) // menggunakan asterisk (*) karna data pointer

    fmt.Println(chicken1, chicken2, chicken3, chicken4)

    fmt.Println("//////")

    // iterasi / mengeluarkan item map

    for key, val := range chicken2 {
        fmt.Println(key, "    \t:",	val)
    }

    fmt.Println("//////")

    // hapus item map
    fmt.Println("jumlah item:", len(chicken2))	//	2
    fmt.Println(chicken2)

    delete(chicken2,	"januari")

    fmt.Println("jumlah item:", len(chicken2))	//	1
    fmt.Println(chicken2)

    fmt.Println("//////")

    // deteksi item dengan 2 variabel
    var value, isExist = chicken2["agustus"]

    if isExist {
        fmt.Println(value)
    } else {
        fmt.Println("item is not exist")
    }

    fmt.Println("//////")

    var chickens = []map[string]string {
        map[string]string{"name": "chicken yellow", "gender": "male"},        
        map[string]string{"name": "chicken blue",   "gender": "male"},
        map[string]string{"name": "chicken red",    "gender": "male"},
    }

    for _, chicken := range chickens {
        fmt.Println(chicken["gender"], chicken["name"])
    }

    // dengan go lang versi baru
    var chickens2 = []map[string]string{
        {"name": "chicken red", "gender": "male"},
        {"name": "chicken blue", "gender": "male"},
        {"name": "chicken yellow", "gender": "female"},
    }    

    for _, chicken := range chickens2 {
        fmt.Println(chicken["gender"], chicken["name"])
    }

    // kalau string string bisa berbeda key
    var data = []map[string]string{
        {"name": "chicken blue", "gender": "male", "color":	"brown"},
        {"address":	"mangga	street", "id": "k001"},
        {"community": "chicken lovers"},
    }

    for _, sample := range data {
        fmt.Println(sample["id"], sample["name"])
    }

}