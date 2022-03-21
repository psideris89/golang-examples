package personService

import "fmt"

type person struct {
	name string
	age  int
}

func Create(name string, age int) person {
	return person{name, age}
}

func Update(p *person, name string, age int) {
	p.name = name
	p.age = age
}

func GetUpdated(p person, name string, age int) person {
	p.name = name
	p.age = age
	return p
}

func Print(p person) {
	fmt.Println(p)
}
