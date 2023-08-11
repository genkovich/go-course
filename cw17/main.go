package main

import (
	"fmt"
	"github.com/mediocregopher/radix/v3"
)

func main() {
	p, err := radix.NewPool("tcp", "127.0.0.1:6379", 3)

	if err != nil {
		panic(err.Error())
	}

	err = p.Do(radix.Cmd(nil, "SETEX", "user:num", "30", "5"))
	if err != nil {
		panic(err.Error())
	}

	var userName string

	err = p.Do(radix.Cmd(&userName, "GET", "user:test"))
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(userName, "done")

	err = p.Do(radix.Cmd(nil, "DEL", "user:test"))
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("deleted")

	var counter string

	err = p.Do(radix.Cmd(&counter, "INCR", "user:num"))
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(counter)

}
