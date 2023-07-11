package main

import (
	"github.com/rs/zerolog/log"
	"net/rpc/jsonrpc"
)

type Args struct {
	Item Item
}

type Item struct {
	Id    string
	Title string
}

type Basket struct {
	items []Item
}

func main() {
	client, err := jsonrpc.Dial("tcp", "localhost:8082")
	if err != nil {
		log.Fatal().Msgf("cant create")
	}

	item := Item{
		Id:    "1",
		Title: "Good",
	}

	args := Args{item}
	var result Basket

	err = client.Call("Basket.Add", args, &result)
	if err != nil {
		log.Fatal().Err(err).Msgf("cant call add method")
	}

	log.Info().Msgf("%v", result)

}
