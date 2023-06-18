package main

import "fmt"

// Створити інтерфейс кавоварки і його імплементацію
// Навчити кавоварку варити на арабіці і на бленді (інтерфейс кави і дви стратегії)

type CoffeeMachine interface {
	CreateCoffee()
}

type CoolMachine struct{}

func (cm CoolMachine) CreateCoffee(c Coffee) {
	fmt.Printf("Brewing %s\n", c.GetTitle())
}

type Coffee interface {
	GetTitle() string
	GetBeans() string
}

type Arabica struct{}

func (a Arabica) GetTitle() string {
	return "Arabica"
}

func (a Arabica) GetBeans() string {
	return "Arabica Beans"
}

type Blend struct{}

func (b Blend) GetTitle() string {
	return "Blend"
}

func (b Blend) GetBeans() string {
	return "Blend Beans"
}

func main() {
	var coffeeMachine = CoolMachine{}
	coffeeMachine.CreateCoffee(Arabica{})
	coffeeMachine.CreateCoffee(Blend{})
}
