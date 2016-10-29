package main

import "fmt"
import "os"
import "io"

var path = "/home/lulz/Documents/temp/test.txt"

func main() {
    createFile() 
    writeFile()
    readFile()
    deleteFile()
}

func checkError(err error) {
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(0)
    }
}

func createFile() {
    // deteksi apakah file sudah ada
    var _, err = os.Stat(path) // path_info, error

    // buat file baru jika belum ada
    if os.IsNotExist(err) {
        var file, err = os.Create(path)
        checkError(err)
        fmt.Println("File berhasil dibuat:", path)
        defer file.Close() // di close karna otomatis open
    } else {
        fmt.Println("File sudah dibuat:", path)
    }
}   

func writeFile() {
    // buka file dengan level akses READ & writeFile
    var file, err = os.OpenFile(path, os.O_RDWR, 0644)
    checkError(err)
    defer file.Close()

    // tulis data ke file
    _, err = file.WriteString("halo\n")
    checkError(err)
    _, err = file.WriteString("mari pelajari golang\n")
    checkError(err)

    // simpan perubahan
    err = file.Sync()
    checkError(err)
    fmt.Println("File berhasil ditulis")    
}

func readFile() {
    // buka file
    var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
    checkError(err)
    defer file.Close()

    // baca file
    fmt.Println("Baca file:", path)
    var text = make([]byte, 1024)
    for {
        n, err := file.Read(text)
        if err != io.EOF {
            checkError(err)
        }
        if n == 0 {
            break
        }
    }
    fmt.Println(string(text))
    checkError(err)
}

func deleteFile() {
    var err = os.Remove(path)
    checkError(err)
    fmt.Println("File", path, "berhasil dihapus!")
}