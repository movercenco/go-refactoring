package main

import "fmt"

type Employee struct {
	name, position string
	annualIncome   int
}

func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

type EmployeeFactory struct {
	position     string
	annualIncome int
}

func (e *EmployeeFactory) Create(name string) *Employee {
	return &Employee{
		name:         name,
		position:     e.position,
		annualIncome: e.annualIncome,
	}
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

func main() {
	developerFactory := NewEmployeeFactory("developer", 20000)
	managerFactory := NewEmployeeFactory("manager", 50000)

	developer := developerFactory("Marin")
	manager := managerFactory("Cris")

	fmt.Println(developer)
	fmt.Println(manager)

	ceoFactory := NewEmployeeFactory2("ceo", 350000)
	ceo := ceoFactory.Create("Adam")
	fmt.Println(ceo)
}
