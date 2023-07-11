package main

import (
	"github.com/rs/zerolog/log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 1. Створити клієнт і сервер, використовуючи json-rpc:
// Сервер виконує роль кошика інтернет-магазина. Кошик очікує id і name товара. Додає до кошику, оновлює і видаляє.
// Клієнт виконує роль клієнта і працює із кошиком: додате, видаляє і редагує товари

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

func (b *Basket) Add(args *Args, result *Basket) error {
	log.Info().Msg("operation: Add")
	b.items = append(b.items, args.Item)

	*result = *b
	return nil
}

func (b *Basket) Drop(args *Args, result *Basket) error {
	log.Info().Msg("operation: Drop")

	for i, v := range b.items {
		if v.Id == args.Item.Id {
			b.items = append(b.items[:i], b.items[i+1:]...)
		}
	}

	*result = *b
	return nil
}

func (b *Basket) Update(args *Args, result *Basket) error {
	log.Info().Msg("operation: Update")

	for i, v := range b.items {
		if v.Id == args.Item.Id {
			b.items[i] = args.Item
		}
	}

	*result = *b
	return nil
}

func main() {
	log.Info().Msgf("Starting RPC Server")
	basket := Basket{}

	rpcServer := rpc.NewServer()

	err := rpcServer.Register(&basket)

	if err != nil {
		log.Fatal().Err(err).Msg("can't register basket")
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
