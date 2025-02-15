package main

import(
	"fmt"
	"unsafe"
)

type Person struct{
	name string // 16 bytes (all occupied)
	age int32    // 8 bytes (int32 is 4 bytes, but as my computer architecture is 8bytes, the minimum is 8 bytes. 
		         // So it occupies only the first 4 bytes and the rest is not used)
	single bool // 0 byte (bool is 1 byte, but it's sharing tha same 8bytes block (word) as age above. 
				// age gets the first 4 bytes of the block and single gets 1 and the last 3 is not used)
} // total, 24 bytes

// |---------word---------| S = occupied by string, I = occupied by int, B = occupied by bool
// |S||S||S||S||S||S||S||S|
// |I||I||I||I||B|| || || |

type Person2 struct{ // change the order of the attributes can change the size occupied by this struct
	age int32    // 8 bytes (int32 is 4 bytes and my PC architecture is 8bytes, so it allocates 8bytes but only uses the first 4. The rest is not used)
	name string // 16 bytes (all occupied)
	single bool // 8 bytes (bool is 1 byte and my PC architecture is 8bytes, so it allocates 8bytes but only uses the first 1. The rest is not used)
} // total, 32 bytes

// |---------word---------| S = occupied by string, I = occupied by int, B = occupied by bool
// |I||I||I||I|| || || || |
// |S||S||S||S||S||S||S||S|
// |B|| || || || || || || |

func main(){
	p1 := Person{name: "joao", age: 10, single: true}
	p2 := Person2{name: "joao", age: 10, single: true}
	
	fmt.Println( unsafe.Sizeof(p1) )
	fmt.Println( unsafe.Sizeof(p2) )
}