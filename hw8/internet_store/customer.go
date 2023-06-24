package internet_store

import "math/rand"

type CustomerRequest struct {
	name   string
	basket []BasketItem
}

type BasketItem struct {
	product  Product
	quantity int
}

func generateCustomerRequest() CustomerRequest {
	randomNames := []string{"John", "Jack", "James", "Jill", "Jane", "Joe", "Jenny", "Jade", "Jasmine", "Jasper"}
	min := 1
	max := 10
	basketItemsQuantity := rand.Intn(max-min) + min

	var basket []BasketItem
	for i := 0; i < basketItemsQuantity; i++ {
		basket = append(basket, BasketItem{
			product:  generateRandomProduct(),
			quantity: rand.Intn(max-min) + min,
		})
	}

	return CustomerRequest{
		name:   randomNames[rand.Intn(len(randomNames)-1)],
		basket: basket,
	}
}
