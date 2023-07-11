package main

import (
	"errors"
	"github.com/rs/zerolog/log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

type Calc int

func (c *Calc) Add(args *Args, result *int) error {
	log.Info().Msgf("operation: Add")
	*result = args.A + args.B
	return nil
}

func (c *Calc) Sub(args *Args, result *int) error {
	log.Info().Msgf("operation: Sub")
	*result = args.A - args.B
	return nil
}

func (c *Calc) Multiply(args *Args, result *int) error {
	log.Info().Msgf("operation: Multiply")
	*result = args.A * args.B
	return nil
}

func (c *Calc) Divide(args *Args, result *int) error {
	log.Info().Msgf("operation: Divide")
	if args.B == 0 {
		return errors.New("can't divide by zero")
	}

	*result = args.A / args.B
	return nil
}

func main() {
	log.Info().Msgf("Starting RPC Server")

	var calculator Calc

	rpcServer := rpc.NewServer()

	err := rpcServer.Register(&calculator)
	if err != nil {
		return
	}

	listener, err := net.Listen("tcp", ":8082")

	if err != nil {
		log.Fatal().Err(err).Msg("can't create tcp listener")
	}

	log.Info().Msg("Started RPC Server")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		log.Info().Msgf("accepted new connection")

		go rpcServer.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
