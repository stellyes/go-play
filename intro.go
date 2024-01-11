package main

import (
	"log"
	//"sort"
	//"time"
)

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Chicken struct {
	Name  string
	Color string
	Chill bool
}

func main() {
	dog := Dog{
		Name:  "Fido",
		Breed: "Insert Generic Retriever",
	}

	printInfo(&dog)

	chicken := Chicken{
		Name:  "Cookie",
		Color: "White",
		Chill: true,
	}

	printInfo(&chicken)
}

func printInfo(a Animal) {
	log.Println("This animal has", a.NumberOfLegs(), "legs and", a.Says(), "is what it says")
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (c *Chicken) Says() string {
	return "Cluck"
}

func (d *Chicken) NumberOfLegs() int {
	return 2
}

// func main() {
// 	for i := 0; i <= 10; i++ {
// 		log.Println("hi", i)
// 	}

// 	animals := []string{"cat", "dog", "mouse", "duck", "chicken"}

// 	// Underscore is blank identifier, compiler ignores data in _
// 	for _, animal := range animals {
// 		log.Println(animal)
// 	}

// 	var eap = "Once upon a midnight dreary"

// 	for i, l := range eap {
// 		log.Println(i, ":", l)
// 	}
// }

// Lesson 6 - decision statements

// func main() {
// 	// If statement
// 	isTrue := true
// 	if isTrue {
// 		log.Println("isTrue is", isTrue)
// 	} else {
// 		log.Println("isTrue is false")
// 	}

// 	// switch statement
// 	myVar := "cat"
// 	switch(myVar) {
// 	case "cat":
// 		log.Println("myVar is cat")

// 	case "dog":
// 		log.Println("myVar is dog")

// 	case "bird":
// 		log.Println("bird")
// 	}
// }

// Lesson 5 - slices

// func main() {
// 	numbers := []int{1,2,3,4,5,6,7,8,9,10}  // shorthand slice syntax
// 	log.Println(numbers[6:9]) //slice from six to nine indexes
// }

// func main() {
// 	var mySlice []string

// 	mySlice = append(mySlice, "Ryan")
// 	mySlice = append(mySlice, "Kevin")

// 	log.Println(mySlice)

// 	var myInt []int

// 	myInt = append(myInt, 1)
// 	myInt = append(myInt, 4)
// 	myInt = append(myInt, 2)
// 	myInt = append(myInt, 3)

// 	log.Println("myInt before sort:", myInt)

// 	sort.Ints(myInt)

// 	log.Println("myInt after sort:", myInt)
// }

// Maps dont need to be accessed globally with pointers
// reference is always unqiue

// type User struct {
// 	FirstName string
// 	LastName  string
// }

// func main() {
// 	myMap := make(map[string]string)

// 	myMap["friend"] = "Kevin"

// 	log.Println(myMap["friend"])

// 	me := User{
// 		FirstName: "Ryan",
// 		LastName: "England",
// 	}

// 	myMap["me"] = me.FirstName

// 	log.Println(myMap["me"])

// }

// Lesson 4 - struct functions and var declarations
//
// type myStruct struct {
// 	FirstName string
// }

// // m *myStruct prefix assigns function to struct, meaning any
// // instance of myStruct has access to printFirstname()
// func (m *myStruct) printFirstName() string {
// 	return m.FirstName
// }

// func main() {
// 	var myVar myStruct

// 	myVar.FirstName = "Ryan"

// 	myVar2 := myStruct{
// 		FirstName: "Ryan",
// 	}

// 	log.Println("myVar is set to:", myVar.FirstName)
// 	log.Println("myVar2 is set to:", myVar2.FirstName)
// }

// Lesson 3 - Structs
//
// var s = "seven"

// // In Go, Variable names with capital letters
// // can be used outside the main file
// type User struct {
// 	FirstName string
// 	LastName string
// 	PhoneNumber string
// 	Age int
// 	Birthdate time.Time
// }

// func main() {
// 	user := User{
// 		FirstName: "Ryan",
// 		LastName: "England",
// 	}

// 	log.Println(user.LastName)
// }

// // func w/ multiple outputs
// func saySomething(s string) (string, string) {
// 	return s, "world"
// }

// Lesson 2 - Pointers
//
// import "log"

// // CORE TAKE AWAY: AMPERSAND GETS YOU THE REFERENCE, POINTERS POINT TO REFERENCES

// func main() {
// 	var exampleString string
// 	exampleString = "!! No way at all !!"

// 	log.Println("Pointers work something like:", exampleString)
// 	// Sending pointer function the address of exampleString in memory
// 	globalChangeWithPointer(&exampleString)
// 	log.Println("Pointers work something like:", exampleString)
// }

// func globalChangeWithPointer(s *string) {
// 	newVal := "THIS!"
// 	*s = newVal
// }

// Lesson 1 - Hello World
//
// import "fmt"

// func main() {

// 	fmt.Println("This is my first Go program!")

// 	var x int = 5
// 	var y int = 7
// 	var z int = x * y
// 	fmt.Println(z)

// 	var foo string = "Guess what?"
// 	var bar string = "Chicken butt!"
// 	fmt.Println(foo)
// 	fmt.Println(bar)

// 	bumpInTheNight := didYouHearThat()
// 	fmt.Println("The text from this point forward is:", bumpInTheNight)
// }

// func didYouHearThat() string {
// 	return "from another function!"
// }
