package main

import "fmt"

// Define a struct named 'Person'
type Person struct {
	Name string
	Age  int
}

// Dependency Non-Injection
func infoPerson() {
	// Create an instance of Person
	person1 := Person{Name: "Non-In", Age: 3}

	// Create another instance of Person
	person2 := Person{Name: "Inj Non", Age: 2}

	// Access and modify instance variables (fields)
	fmt.Printf("Person 1: Name=%s, Age=%d\n", person1.Name, person1.Age)
	fmt.Printf("Person 2: Name=%s, Age=%d\n", person2.Name, person2.Age)

	person1.Age = 4 // Modifying person1's Age does not affect person2
	fmt.Printf("Person 1 (updated): Name=%s, Age=%d\n", person1.Name, person1.Age)
}

// Dependency Injection
func newPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

func main() {
	// Dependency Non-Injection
	infoPerson()

	// Dependency Injection
	person3 := newPerson("Inj", 30)

	// Create another instance of Person
	person4 := newPerson("Dep Inj", 25)

	// Access and modify instance variables (fields)
	fmt.Printf("Person 1: Name=%s, Age=%d\n", person3.Name, person3.Age)
	fmt.Printf("Person 2: Name=%s, Age=%d\n", person4.Name, person4.Age)

	person3.Age = 31 // Modifying person1's Age does not affect person2
	fmt.Printf("Person 1 (updated): Name=%s, Age=%d\n", person3.Name, person3.Age)
}
