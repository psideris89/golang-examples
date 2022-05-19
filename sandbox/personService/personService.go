package personService

import "fmt"

type Person interface {
	Update(firstName string, lastName string, age int)
	UpdateAge(age int)
	UpdateFirstName(firstName string)
	UpdateLastName(lastName string)
	Print()
}

type person struct {
	firstName string
	lastName  string
	age       int
}

func NewPerson(firstName string, lastName string, age int) Person {
	return &person{firstName, lastName, age}
}

func (p *person) UpdateAge(age int) {
	p.age = age
}

func (p *person) UpdateFirstName(firstName string) {
	p.firstName = firstName
}

func (p *person) UpdateLastName(lastName string) {
	p.lastName = lastName
}

func (p *person) Update(firstName string, lastName string, age int) {
	p.firstName = firstName
	p.lastName = lastName
	p.age = age
}

func (p *person) Print() {
	fmt.Printf("\n%s %s, %d years old", p.firstName, p.lastName, p.age)
}
