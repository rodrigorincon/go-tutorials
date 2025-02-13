package main

import (
    "fmt"
)

type Person struct {
    name string
    age int32
}

func accessingNonExistKey(people map[int32]string){
    fmt.Println("----------- accessing non-existent keys")
    fmt.Println("accessing a non-existent key in STRING:", people[9] == "") // if we access an non-existent key it will return the zero value for the value type
    // in this case, zero value for string is ""
    fmt.Println("accessing deleted key:", people[3] == "")
    fmt.Println("size of map:", len(people))

    var graphic = map[int32]int32 {
        1: 2,
        2: 4,
        3: 8,
        4: 16, // always finish with comma
    }
    fmt.Println("accessing a non-existent key in INT:", graphic[9]) // in this case, the zero value is 0

    var peopleStruct = map[int32]Person {
        1: Person{"joao", 13},
        2: Person{"maria", 10},
        3: Person{"ana", 7},
    }
    fmt.Println("person ID 2:", peopleStruct[2].name)
    fmt.Println("accessing a non-existent key in STRUCT:", peopleStruct[9]) // in this case, the zero value is empty struct (with name "" and age 0)

    var pointerStruct = map[int32]*Person {
        1: &Person{"joao", 13},
        2: &Person{"maria", 10},
        3: &Person{"ana", 7},
    }
    fmt.Println("person ID 2:", (*pointerStruct[2]).name)
    fmt.Println("accessing a non-existent key in POINTER:", pointerStruct[9]) // in this case, the zero value is nil
}

func receivingReferenceMap(param map[int32]string){
    param[2] = "changed in the function"
}


func main() {
    // map is a dictionary/hash/json, a structure with a list of keys and values
    // variable = map[keyType]valueType = {
    //      key: value,
    //      key: value
    // }

    var people = map[int32]string {
        1: "joao",
        2: "maria",
        3: "jose", // always finish with comma
    }
    people[4] = "ana"
    people[1] = "felipe" // update the value for this key
    delete(people, 3) // remove a key-value from map

    fmt.Println("dictionary:", people)
    fmt.Println("person ID 1:", people[1])
    fmt.Println("person ID 2:", people[2])

    accessingNonExistKey(people)

    fmt.Println("----------- checking if a key exists")
    value, exist := people[9]
    if exist {
      fmt.Println("The key exists:", value)
    } else {
      fmt.Println("The key doesnt exist!")
    }
    
    fmt.Println("----------- references")

    // maps are always a pointer to a hash table, so if we pass it for other var or in a function it can be changed
    var original = map[int32]string {
        1: "joao",
        2: "maria",
        3: "jose",
    }
    otherMap := original
    otherMap[1] = "changed"
    fmt.Println("original map:", original[1])
    receivingReferenceMap(original)
    fmt.Println("original map:", original[2])
    fmt.Println("copy:", otherMap)

    fmt.Println("----------- array (slice in the truth) of maps")
    var mapArray [2]map[string]int
    mapArray[0] = map[string]int{
        "joao": 1,
        "maria": 2,
    }
    mapArray[1] = map[string]int{
        "jose": 3,
        "ana": 4,
    }
    fmt.Println(mapArray)
}