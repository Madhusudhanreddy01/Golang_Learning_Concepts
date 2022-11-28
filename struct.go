// // type struct_name struct {       |----------->Syntax
// 	member1 datatype;
// 	member2 datatype;
// 	member3 datatype;
// 	...
//   }

// package main

// import (
// 	"fmt"
// )

// type person struct {
// 	firstName string
// 	lastName  string
// 	age       int
// }

// func main() {
// 	x := person{age: 30, firstName: "John", lastName: "Anderson"}
// 	fmt.Println(x)
// 	fmt.Println(x.firstName)
// }

//------------------------------------------------------------------

package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var pers1 Person
	var pers2 Person

	// Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	// Pers2 specification
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500

	// Print Pers1 info by calling a function
	printPerson(pers1)

	// Print Pers2 info by calling a function
	printPerson(pers2)
}

func printPerson(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
}
