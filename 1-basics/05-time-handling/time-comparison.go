package main

import (
    "fmt"
    "time"
)

type Person struct{
    name string
    createdAt time.Time
}

func isLegalAge(person Person) bool {
    currentTime := time.Now()
    beginOfDay := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0,0,0,0, time.UTC)

    return person.createdAt.Before( beginOfDay.Add(-18*365*24*time.Hour) )
}

func filterPeopleByDate(){
    person1 := Person{"joao", time.Date(1985, 8, 15, 14, 30, 45, 100, time.UTC)}
    person2 := Person{"maria", time.Date(1978, 8, 15, 14, 30, 45, 100, time.UTC)}
    person3 := Person{"jose", time.Date(2001, 8, 15, 14, 30, 45, 100, time.UTC)}
    person4 := Person{"ana", time.Date(1987, 8, 15, 14, 30, 45, 100, time.UTC)}
    person5 := Person{"enzo", time.Date(2014, 8, 15, 14, 30, 45, 100, time.UTC)}
    person6 := Person{"valentia", time.Date(2015, 8, 15, 14, 30, 45, 100, time.UTC)}
    people := [...]Person{person1,person2,person3,person4,person5,person6}

    var legalAgeArr []Person
    var notLegalAgeArr []Person
    for _, person := range people{
        if(isLegalAge(person)){
            legalAgeArr = append(legalAgeArr, person)
        }
        if(!isLegalAge(person)){
            notLegalAgeArr = append(notLegalAgeArr, person)
        }
    }
    fmt.Println("only legal age people:", legalAgeArr)
    fmt.Println("not legal age people:", notLegalAgeArr)
}

func main() {
    // check if a time is in the future
    fmt.Println("------------- Checking what time is greater")
    currentTime := time.Now()
    timePassed := time.Date(2021, 8, 15, 14, 30, 45, 100, time.UTC)
    futureTime := time.Date(2099, 12, 25, 16, 40, 55, 200, time.UTC)

    fmt.Println("timePassed is before now?", timePassed.Before(currentTime))
    fmt.Println("futureTime is after now?", futureTime.After(currentTime))
    // compare dates like currentTime >= startTime DOESNT WORK!!

    // sub times
    fmt.Println("------------- calculating time between 2 dates")
    duration := futureTime.Sub(timePassed) // Sub returns a time.Duration object (int64 with some functions), that divide the time in hours, minutes and seconds
    fmt.Println("Duration between 2 times is", duration)

    fmt.Println("Duration in hours", duration.Hours())
    fmt.Println("Duration in minutes", duration.Minutes())
    fmt.Println("Duration in seconds", duration.Seconds())
    fmt.Println("Duration in mili seconds", duration.Milliseconds())
    fmt.Println("Duration in micro seconds", duration.Microseconds())
    fmt.Println("Duration in nano seconds", duration.Nanoseconds())

    if duration > (24 * time.Hour){ // time.Hour, Minute and Second helps to see if difference is in a time range
        fmt.Println("The days are more than 1 day of difference")
    }
        
    // check if a time in a range
    fmt.Println("------------- checking if a time is in a time range")
    startTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.UTC)
    endTime :=   startTime.Add( 2 * time.Hour )

    fmt.Println("now is between ", startTime.Format("2006-01-02 03:04:05"), "and", endTime.Format("2006-01-02 03:04:05"), "?", 
                currentTime.After(startTime) && currentTime.Before(endTime))

    // get objects created in a time range
    fmt.Println("------------- getting people created in a time range")
    filterPeopleByDate()
}
