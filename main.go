package main

import "course/biography"

func main() {
	b := biography.NewBiography("Кирило", "Сулімовський").
		SetIsMarried(true).
		SetDateOfBirth("1991-10-21").
		SetChildren(2).
		SetMajor("Computer Science").
		AddPhoneNumber("+380931234567").
		AddPhoneNumber("+380931234568").
		AddPhoneNumber("+380931234569")

	b.PrintBiography()
}
