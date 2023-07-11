package minmax

import (
	"fmt"
	"math/rand"
)

type processStruct struct {
	generatedChannel chan []int
	printChannel     chan minMax
	syncChannel      chan bool
}

type minMax struct {
	min int
	max int
}

func StartProcess() {
	process := processStruct{
		generatedChannel: make(chan []int),
		printChannel:     make(chan minMax),
		syncChannel:      make(chan bool),
	}

	go process.getMinMax()
	go process.generateAndPrint(10, 100, 5)

	<-process.syncChannel
}

func (p processStruct) generateAndPrint(min int, max int, count int) {
	var randomNumbers []int
	for i := 0; i < count; i++ {
		randomNumbers = append(randomNumbers, rand.Intn(max-min)+min)
	}

	fmt.Println(randomNumbers)
	p.generatedChannel <- randomNumbers

	minMax := <-p.printChannel
	fmt.Println(minMax)

	p.syncChannel <- true
}

func (p processStruct) getMinMax() {
	test := <-p.generatedChannel
	max := test[0]
	min := test[0]
	for _, value := range test {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	result := minMax{
		min: min,
		max: max,
	}

	p.printChannel <- result
}
