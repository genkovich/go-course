package internet_store

import (
	"context"
	"flag"
	"fmt"
	"sync"
)

var calculateChannel chan CustomerRequest

func Run() {
	count := flag.Int("customers", 1, "positive int")
	flag.Parse()

	fmt.Println(*count)
	ctx := context.WithValue(context.Background(), "count", *count)

	calculateChannel = make(chan CustomerRequest)
	var wg sync.WaitGroup

	wg.Add(2)

	go calculate(&wg)
	go generate(ctx, &wg)

	wg.Wait()
}

func generate(ctx context.Context, wg *sync.WaitGroup) {
	count := ctx.Value("count").(int)

	var customerRequest CustomerRequest
	for i := 0; i < count; i++ {
		customerRequest = generateCustomerRequest()
		fmt.Println(customerRequest)
		calculateChannel <- customerRequest
	}

	wg.Done()
	close(calculateChannel)
}

func calculate(wg *sync.WaitGroup) {
	for {
		select {
		case customerRequest, ok := <-calculateChannel:
			if !ok {
				wg.Done()
				return
			}

			sum := 0

			for _, basketItem := range customerRequest.basket {
				sum += basketItem.product.price * basketItem.quantity
			}

			fmt.Println(sum)

		}
	}
}
