package average

import (
	"fmt"
	"math/rand"
)

type Process struct {
	randNumberChannel chan int
	resultsChannel    chan int
	syncChannel       chan bool
}

func StartProcess() {
	process := Process{
		randNumberChannel: make(chan int),
		resultsChannel:    make(chan int),
		syncChannel:       make(chan bool),
	}

	go process.printAverage()
	go process.average()
	go process.generateRandomNumbers(10)

	<-process.syncChannel

}

func (p Process) generateRandomNumbers(count int) {
	for i := 0; i < count; i++ {
		randNumber := rand.Intn(100)
		fmt.Println("Generated random number: ", randNumber, " from goroutine: ", "goroutine")
		p.randNumberChannel <- randNumber
	}
	fmt.Println("Try to close average channel")
	p.CloseAverage()
}

func (p Process) average() {
	sum := 0
	count := 0

	for {
		select {
		case number, ok := <-p.randNumberChannel:
			if ok {
				fmt.Println("Received number: ", number, " from goroutine: ", "average")
				sum += number
				count++
				p.resultsChannel <- sum / count
			} else {
				fmt.Println("Try to close print channel")
				p.ClosePrint()
				return
			}
		}
	}

}

func (p Process) printAverage() {
	for {
		select {
		case res, ok := <-p.resultsChannel:
			if ok {
				fmt.Println("Average: ", res, " from goroutine: ", "printAverage")
			} else {
				fmt.Println("Print channel closed")
				p.syncChannel <- true
				return
			}
		}
	}
}

func (p Process) CloseAverage() {
	close(p.randNumberChannel)
}

func (p Process) ClosePrint() {
	close(p.resultsChannel)
}
