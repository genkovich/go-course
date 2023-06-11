package biography

import (
	"fmt"
	"time"
)

type Biography struct {
	Name         string
	Age          int
	Children     byte
	Major        string
	isMarried    bool
	Information  string
	PhoneNumbers []string
}

func NewBiography(Name string, Surname string) *Biography {
	name := fmt.Sprintf("%s %s", Surname, Name)

	return &Biography{
		Name:        name,
		Information: "",
	}
}

func getYear(t time.Time) int {
	year, _, _ := t.Date()
	return year
}

func GetAge(birthdate time.Time) int {
	today := time.Now()

	if today.Before(birthdate) {
		return 0
	}

	age := getYear(today) - getYear(birthdate)

	anniversary := birthdate.AddDate(age, 0, 0)
	if anniversary.After(today) {
		age--
	}

	return age
}

func (b *Biography) greetings() *Biography {
	b.Information = fmt.Sprintf("Привіт, мене звати %s, мені %d.\n", b.Name, b.Age)
	return b
}

func (b *Biography) printIsMarried() *Biography {
	if b.isMarried {
		b.Information += "Я одружений. \n"
	} else {
		b.Information += "Я не одружений. \n"
	}
	return b
}

func (b *Biography) SetChildren(Children byte) *Biography {
	b.Children = Children
	return b
}

func (b *Biography) printChildren() *Biography {
	if b.Children > 0 {
		b.Information += fmt.Sprintf("У мене %d дітей.\n", b.Children)
	} else {
		b.Information += "Дітей немає.\n"
	}
	return b
}

func (b *Biography) SetMajor(Major string) *Biography {
	b.Major = Major
	return b
}

func (b *Biography) printMajor() *Biography {
	b.Information += fmt.Sprintf("%s.\n", b.Major)
	return b
}

func (b *Biography) AddPhoneNumber(PhoneNumber string) *Biography {
	b.PhoneNumbers = append(b.PhoneNumbers, PhoneNumber)
	return b
}

func (b *Biography) printContacts() *Biography {
	b.Information += "Мої контакти:"
	for _, number := range b.PhoneNumbers {
		b.Information += fmt.Sprintf("\n%s", number)
	}
	return b
}

func (b *Biography) SetDateOfBirth(DateOfBirth string) *Biography {
	birth, _ := time.Parse("2006-01-02", DateOfBirth)
	age := GetAge(birth)
	b.Age = age
	return b
}

func (b *Biography) SetIsMarried(isMarried bool) *Biography {
	b.isMarried = isMarried
	return b
}

func (b *Biography) String() string {
	return b.Information
}

func (b *Biography) PrintBiography() {
	b.greetings().printMajor().printIsMarried().printChildren().printContacts()
	fmt.Println(b)
}
