package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	p := person{name: name}
	p.age = age
	return &p
}

func main() {
	fmt.Println(person{"A", 20})

	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon", 10))
	p1 := newPerson("Bon", 20)
	fmt.Println(p1.age)

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	s.age = 51
	fmt.Println(sp.age)
}
