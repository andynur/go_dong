package main

import "fmt"
import "time"

func main() {
    var time1 = time.Now()
    fmt.Printf("time1 %v\n", time1)

    // param: thn, bln, tgl, jam, menit, detik, nanodetik, timezone
    var time2 = time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC) 
    fmt.Printf("time2 %v\n", time2)

    fmt.Println(".")

    // method milk
    var now = time.Now()
    fmt.Println("year:", now.Year(), "month:", now.Month())

    fmt.Println(".")

    // parsing time
    var layoutFormat, value string
    var date time.Time

    layoutFormat = "2006-01-02 15:04:05"
    value = "2015-09-02 08:04:00"
    
    // format datanya, string yg ingin diparsing mengikuti format
    date, _ = time.Parse(layoutFormat, value) 
    fmt.Println(value, "\t->", date.String()) 

    layoutFormat = "02/01/2006 MST"
    value = "02/09/2015 WIB"
    date, _ = time.Parse(layoutFormat, value)
    fmt.Println(value, "\t->", date.String())
    
    fmt.Println(".")

    // predefned layout
    var preDate, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")
    fmt.Println(preDate.String())

    fmt.Println(".")

    // format string to time
    var strDate, _ = time.Parse(time.RFC822, "13 Jul 16 08:00 WIB")

    var dateS1 = strDate.Format("Monday 02, January 2006 15:04 MST")
    fmt.Println("dateS1", dateS1)

    var dateS2 = strDate.Format(time.RFC3339)
    fmt.Println("dateS2", dateS2)

    fmt.Println(".")

    // error handling
    var errDate, err = time.Parse("06 Jan 15", "02 Sep 15 08:00 WIB")

    if err != nil {
        fmt.Println("error", err.Error())
        return
    }

    fmt.Println(errDate)
    
}