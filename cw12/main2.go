package main

import "fmt"

//Створити тип студента з можливістю декорування як відмінник, спортсмен і подібне. Студент має оцінку та імʼя, які змінюються в залежності від декоратора

func main() {
	var student StudentStatusPrinter = Student{
		displayName: "Student",
		name:        "John",
	}

	var answer string

	fmt.Println("add excellent?")
	fmt.Scanln(&answer)

	if answer == "y" {
		student = ExcellentStudent{
			parent: student,
		}
	}

	fmt.Println("add sportsman?")
	fmt.Scanln(&answer)

	if answer == "y" {
		student = Sportsman{
			parent: student,
		}
	}

	student.printName()
}

type StudentStatusPrinter interface {
	printName()
}

type Student struct {
	displayName string
	name        string
}

func (s Student) printName() {
	fmt.Println(s.displayName + " " + s.name)
}

type ExcellentStudent struct {
	parent StudentStatusPrinter
}

func (e ExcellentStudent) printName() {
	fmt.Println("ExcellentStudent")
	e.parent.printName()
}

type Sportsman struct {
	parent StudentStatusPrinter
}

func (s Sportsman) printName() {
	fmt.Println("Sportsman")
	s.parent.printName()
}
