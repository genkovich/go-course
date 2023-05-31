package biography

import (
	"fmt"
	"time"
)

const (
	Name        string = "Кирило"
	Surname     string = "Сулімовський"
	DateOfBirth string = "1991-10-21"
	Major       string = "Працюю на позиції програміста"
	isMarried   bool   = true
	Children    byte   = 0
)

var PhoneNumbers = []string{"+380671234567", "+380631234567"}

type Biography struct {
	Name        string
	Age         int
	Information string
}

func NewBiography() Biography {
	name := fmt.Sprintf("%s %s", Surname, Name)
	birth, _ := time.Parse("2006-01-02", DateOfBirth)
	age := GetAge(birth)

	return Biography{
		Name:        name,
		Age:         age,
		Information: "",
	}
}

func GetAge(birthdate time.Time) int {
	today := time.Now()
	ty, tm, td := today.Date()
	today = time.Date(ty, tm, td, 0, 0, 0, 0, time.UTC)
	by, bm, bd := birthdate.Date()
	birthdate = time.Date(by, bm, bd, 0, 0, 0, 0, time.UTC)
	if today.Before(birthdate) {
		return 0
	}
	age := ty - by
	anniversary := birthdate.AddDate(age, 0, 0)
	if anniversary.After(today) {
		age--
	}
	return age
}

func (b Biography) Greetings(n Biography) Biography {
	n.Information = fmt.Sprintf("Привіт, мене звати %s, мені %d.\n", b.Name, b.Age)
	return n
}

func (b Biography) isMarried(n Biography) Biography {
	if isMarried {
		n.Information += "Я одружений. \n"
	} else {
		n.Information += "Я не одружений. \n"
	}
	return n
}

func (b Biography) Children(n Biography) Biography {
	if Children > 0 {
		n.Information += fmt.Sprintf("У мене %d дітей.\n", Children)
	} else {
		n.Information += "Дітей немає.\n"
	}
	return n
}

func (b Biography) Major(n Biography) Biography {
	n.Information += fmt.Sprintf("%s.\n", Major)
	return n
}

func (b Biography) Contacts(n Biography) Biography {
	n.Information += "Мої контакти:"
	for _, number := range PhoneNumbers {
		n.Information += fmt.Sprintf("\n%s", number)
	}
	return n
}

func (b Biography) String() string {
	return b.Information
}

func PrintBiography() {
	biography := NewBiography()
	biography = biography.Greetings(biography)
	biography = biography.isMarried(biography)
	biography = biography.Children(biography)
	biography = biography.Major(biography)
	biography = biography.Contacts(biography)
	fmt.Println(biography)
}
