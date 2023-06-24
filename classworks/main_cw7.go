package main

import (
	"fmt"
	"time"
)

// Створити дві горутини і обʼєднати їх каналами (string). Деяка строка мусить пройти дві горутини,
// де кожна її допише, і ми отримаємо її на виход
func main() {
	var firstChannel chan []string
	firstChannel = make(chan []string)

	var secondChannel chan []string
	secondChannel = make(chan []string)

	var thirdChannel chan []string
	thirdChannel = make(chan []string)

	message := []string{"Hello"}

	go func() {
		m := <-firstChannel
		strings := append(m, "John")
		secondChannel <- strings
	}()

	go func() {
		m := <-secondChannel
		strings := append(m, "Bob")
		thirdChannel <- strings
	}()

	firstChannel <- message

	fmt.Println(<-thirdChannel)

}

func example() {
	var c chan string
	c = make(chan string)

	go func() {
		fmt.Println("Hello goroutine")

		time.Sleep(2 * time.Second)

		fmt.Println("After sleep ")
		c <- "test"
		c <- "end"
	}()
	fmt.Println("Hello from Main")

	time.Sleep(time.Second)

	fmt.Println("After sleep in main")

	message := <-c
	sec := <-c

	fmt.Println(message, sec)

}
