package main

import "fmt"

type person struct {
	age int
}

func main() {
	var T, age int

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&age)
		p := person{age: age}
		p = p.NewPerson(age)
		p.amIOld()
		for j := 0; j < 3; j++ {
			p = p.yearPasses()
		}
		p.amIOld()
		fmt.Println()
	}
}

func (p person) NewPerson(initialAge int) person {
	//Add some more code to run some checks on initialAge
	if initialAge < 0 {
		fmt.Println("Age is not valid, setting age to 0.")
		initialAge = 0
	}
	p.age = initialAge

	return p
}

func (p person) amIOld() {
	//Do some computation in here and print out the correct statement to the console
	stmt := ""
	switch {
	case p.age < 13:
		stmt = "You are young."
	case p.age >= 13 && p.age < 18:
		stmt = "You are a teenager."
	case p.age >= 18:
		stmt = "You are old."
	}
	fmt.Println(stmt)
}

func (p person) yearPasses() person {
	//Increment the age of the person in here
	p.age++
	return p
}
