package main

import (
	"encoding/json"
	"fmt"
)

type Company struct {
	Title     string
	Employees map[string]Person
}

type Person struct {
	Name    string
	Surname string
	Id      string
}

func CreateCompany(title string) *Company {
	return &Company{
		Title:     title,
		Employees: make(map[string]Person),
	}
}

func (c *Company) ApplyEmployee(person *Person) {
	c.Employees[person.Id] = *person
}

func (c *Company) FireEmployee(id string) {
	delete(c.Employees, id)
}

func CreatePerson(name, surname, id string) Person {
	return Person{
		Name:    name,
		Surname: surname,
		Id:      id,
	}
}

func main() {

	// Створити словник англійських слів з їх українським перекладом та знайти переклад певного слова.

	var Dictionary = map[string]string{
		"cat":  "кіт",
		"dog":  "собака",
		"bird": "птах",
		"fish": "риба",
	}

	fmt.Println(Dictionary["cat"])

	// Створити програму, яка буде зберігати дані про працівників певної компанії, додапє і видаляє працівників

	Apple := CreateCompany("Apple")
	john := CreatePerson("John", "Smith", "1")
	Apple.ApplyEmployee(&john)
	john.Name = "asdsadsada"
	//Apple.ApplyEmployee(CreatePerson("Jack", "Smith", "2"))
	//Apple.ApplyEmployee(CreatePerson("Jane", "Smith", "3"))

	Apple.FireEmployee("2")

	result, _ := json.MarshalIndent(Apple, "", "  ")
	fmt.Println(string(result))
}
