package main

import "fmt"

type stringFunc func(string)
type Name string
type Number int

func (n Number) String() string {
	return fmt.Sprintf("result is %d", n)
}

func (n Number) Square() Number {
	return n * n
}

func (n Number) Add(number Number) Number {
	return n + number
}

func (n Number) Minus(number Number) Number {
	return n - number
}

func (n Number) Apply(fn func(number Number)) {
	fn(n)
}

func (n Name) String() string {
	return "My name is " + string(n)
}

func (n Name) SayName() {
	fmt.Println(n)
}

func (n Name) SayNameWithPrefix(prefix string) {
	fmt.Println(prefix, string(n))
}

func main() {

	var name Name = "Kirill"
	printI(13)

	const a = "a"
	name.SayName()
	name.SayNameWithPrefix("Mr. ")
	Name.SayName("John")

	var sayName stringFunc = name.SayNameWithPrefix
	sayName("MR.")

	var number Number = 3
	var secondNumber Number = 4
	var thirdNumber Number = 5

	result := number.Add(secondNumber).Minus(thirdNumber)
	printFn := func(n Number) {
		fmt.Println("ok ", n)
	}
	number.Apply(printFn)
	fmt.Println(result)
}

func printI(i int) {
	fmt.Println("this is i =", i)
}
