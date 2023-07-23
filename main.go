package main

import (
	"course/cw15"
	"course/cw15/order"
	"fmt"
)

type LoggerObserver struct {
}

func (l LoggerObserver) GetNotified(subject any) {
	fmt.Printf("logger_observer_notified: %s", subject)
}

func (l LoggerObserver) GetID() string {
	return "logger"
}

func main2() {
	observerRegistry := cw15.NewRegistry()
	observerRegistry.Register(LoggerObserver{})

	orderService := order.NewService(observerRegistry)

	orderService.ProcessOrder()

}

//////////////////////////

type LoginSub struct {
}

func main() {

}
