/* RZFeeser | Alta3 Research
   Address-of(&) operators  */    


package main
import (
   "fmt"
)

func noPointer() string {
   return "string"
}

func pointerTest() *string {
   return nil                    // You cannot return nil from a function that returns a string
                                 // but you can return nil from a function that returns a pointer to a string!
                                 // The zero value of a pointer is nil
}

                      // this is the return value
func pointerTestTwo() *string {
   s := "istring"                 // &"string" doesn't work
   return &s
}

func main() {
   fmt.Println(noPointer())      // prints string
   fmt.Println(pointerTest())    // prints <nil>
   fmt.Println(pointerTestTwo()) // prints something like 0x40c158 (an address in memory)
   
   s := pointerTestTwo()         // now s holds an address in memory for a variable
   fmt.Println(s)                // something like 0x40c138
   sp := *s                      // now sp holds the value found at the address s holds
   m:=sp
   fmt.Println(&sp)               // string
   fmt.Println(m)               // string

}

