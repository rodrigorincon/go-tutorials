package main

import (
    "fmt"
    "time"
)

func main() {
    // read Date and Time from computer
    fmt.Println("------------- Time methods")
    currentTime := time.Now()
    fmt.Println("The time is", currentTime)
    fmt.Println("The year is", currentTime.Year())
    fmt.Println("The month is", currentTime.Month())
    fmt.Println("The day is", currentTime.Day())
    fmt.Println("The hour is", currentTime.Hour())
    fmt.Println("The minute is", currentTime.Minute())
    fmt.Println("The second is", currentTime.Second())
    fmt.Println("The nano second is", currentTime.Nanosecond())
    fmt.Println("Number of days since the beggining of year:", currentTime.YearDay())
    fmt.Println("Week day:", currentTime.Weekday())

    // string format
    fmt.Println("------------- Time format")
    fmt.Println("time formatted:", currentTime.Format("2006-01-02 03:04:05"))

    // getting time in a timezone
    fmt.Println("------------- Timezones")
    fmt.Println("time in UTC: ", currentTime.UTC())

    timezone, _ := time.LoadLocation("Asia/Shanghai")
    fmt.Println("time in", timezone, time.Now().In(timezone) )

    // changing timezone in the system
    timezone, _ = time.LoadLocation("America/Sao_Paulo")
    fmt.Println("Timezone:", timezone)
    time.Local = timezone // set the timezone as official timezone so all methods use it now
    fmt.Println("Tempo SP:", time.Now())

    // creating date object
    fmt.Println("------------- Creating time objects")
    theTime := time.Date(2024, 10, 21, 12, 45, 50, 100, time.Local)
    fmt.Println("The time is", theTime)
    fmt.Println("Month is ", theTime.Month())

    timeString := "2021-08-15 02:30:45"
    theTime2, err := time.Parse("2006-01-02 03:04:05", timeString)
    if err != nil {
        fmt.Println("Could not parse time:", err)
    }else{
        fmt.Println("The time parsed is", theTime2)
    }

    timeString = "2021-02-32 02:50:45"
    theTime3, err2 := time.Parse("2006-01-02 03:04:05", timeString)
    if err2 != nil {
        fmt.Println("Could not parse time:", err2)
    }else{
        fmt.Println("The time parsed is", theTime3)
    }
}
