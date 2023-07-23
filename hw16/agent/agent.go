package agent

import (
	"fmt"
	"github.com/radovskyb/watcher"
	"log"
	"time"
)

type Sub interface {
	Id() string
	Consume(event any)
}

type Agent struct {
	watcher       watcher.Watcher
	subscribers   chan Sub
	eventsChannel chan any
	stop          chan struct{}
}

func NewAgent() *Agent {
	a := &Agent{
		watcher:       *watcher.New(),
		subscribers:   make(chan Sub),
		eventsChannel: make(chan any),
		stop:          make(chan struct{}),
	}

	a.setup()

	return a
}

func (a *Agent) AddSub(sub Sub) {
	a.subscribers <- sub
}

func (a *Agent) setup() {
	subscribers := make(map[string]Sub)
	fmt.Println("Start watching folder...")

	a.watcher.FilterOps(watcher.Rename, watcher.Move, watcher.Create, watcher.Remove, watcher.Chmod)

	go func() {
		for {
			select {
			case event := <-a.watcher.Event:
				for _, sub := range subscribers {
					sub.Consume(event)
				}
			case sub := <-a.subscribers:
				subscribers[sub.Id()] = sub
			case err := <-a.watcher.Error:
				log.Fatalln(err)
			case <-a.stop:
				a.stop <- struct{}{}
				return
			}
		}
	}()

	if err := a.watcher.AddRecursive("./hw16/example"); err != nil {
		log.Fatalln(err)
	}

	for path, f := range a.watcher.WatchedFiles() {
		fmt.Printf("%s: %s\n", path, f.Name())
	}

}

func (a *Agent) Watch() {
	if err := a.watcher.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}

func (a *Agent) Stop() {
	a.stop <- struct{}{}
	<-a.stop
}
