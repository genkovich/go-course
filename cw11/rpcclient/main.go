package main

import (
	"github.com/rs/zerolog/log"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

func main() {
	client, err := jsonrpc.Dial("tcp", "localhost:8082")
	if err != nil {
		log.Fatal().Msgf("cant create")
	}

	args := Args{1, 2}
	var result int

	err = client.Call("Calc.Add", args, &result)
	if err != nil {
		log.Fatal().Msgf("cant call add method")
	}

	log.Info().Msgf("%v", result)

	args = Args{5, 1}

	err = client.Call("Calc.Divide", args, &result)
	if err != nil {
		log.Fatal().Msgf("cant call add method")
	}

	log.Info().Msgf("%v", result)
}
