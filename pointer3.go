/* RZFeeser | Alta3 Research
   Without Pointers - Why we need pointers  */    

package main

import (
 "fmt"
)

type User struct {
 Name string
 Pets []string
}

func (u User) newPet() {
 u.Pets = append(u.Pets, "Lucy")                    // Simple function should ensure "Fido" is added to User
 fmt.Println(u)                                     // This user is **NOT** the same user as the one in main()!!
}

func main() {
 u1 := User{Name: "Anna", Pets: []string{"Bailey"}}
 u1.newPet()                                         // {Anna [Bailey Fido]} -- This *should* add "Fido" to "u"
 fmt.Println(u1)                                     // We **DO NOT** see Fido -- what happened?
}

