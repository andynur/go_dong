package main

import "fmt"

func main() {
    // mirip seperti array tetapi [] tidak diisi
    var fruits = []string{"apple", "grape", "banana", "melon"}

    var aFruits = fruits[0:3] // cara pengaksesan slice, mulai index 0 - urut ke 3
    var bFruits = fruits[1:3] 

    var aaFruits = aFruits[1:2]
    var bbFruits = bFruits[0:1]

    fmt.Println(fruits)
    fmt.Println(aFruits)
    fmt.Println(bFruits)
    fmt.Println(aaFruits)
    fmt.Println(bbFruits)

    // buah "grape" diubah menjadi "pineapple" -> index 0
    bbFruits[0] = "pineapple"
    fmt.Println("\nUbah 'grape' menjadi 'pineapple'\n")

    fmt.Println(fruits)
    fmt.Println(aFruits)
    fmt.Println(bFruits)
    fmt.Println(aaFruits)
    fmt.Println(bbFruits)

    fmt.Println("\nBuilt-in Function:\n")
    // go-lang built in function

    var lFruits = fruits[0:2]

    fmt.Println("lebar slicenya:", cap(lFruits))
    fmt.Println("kapasitas slicenya:", len(lFruits))

    fmt.Println(fruits)
    fmt.Println(lFruits)

    var cFruits = append(lFruits, "papaya") // tambah slice baru

    fmt.Println("\nTambah papaya:\n")

    fmt.Println(fruits)
    fmt.Println(lFruits)
    fmt.Println(cFruits)
   
    fmt.Println("\nMenggabungkan dua slice:\n")

    var lastFruits = []string{"manggo", "apple", "pineapple"}
    var newFruits = []string{"watermelon", "orange"}

    var copiedElement = copy(lastFruits, newFruits) // menimpa berdasar index

    fmt.Println(lastFruits)
    fmt.Println(newFruits)
    fmt.Println(copiedElement) // nilai terkecil dari target & tujuan

    fmt.Println("\nAkses Slice 3 Indeks:\n")

    var twoFruits = lastFruits[0:2]
    var threeFruits = lastFruits[0:2:2] // akses slice sesuai dgn angka kapasistas nya

    fmt.Println(lastFruits)
    fmt.Println(len(lastFruits))
    fmt.Println(cap(lastFruits))

    fmt.Println(twoFruits)
    fmt.Println(len(twoFruits))
    fmt.Println(cap(twoFruits))

    fmt.Println(threeFruits)
    fmt.Println(len(threeFruits))
    fmt.Println(cap(threeFruits))

}   