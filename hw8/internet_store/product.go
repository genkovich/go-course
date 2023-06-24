package internet_store

import "math/rand"

type Product struct {
	title string
	price int
}

func generateRandomProduct() Product {
	randomTitles := []string{"Apple", "Banana", "Orange", "Pineapple", "Mango", "Kiwi", "Peach", "Pear", "Plum", "Grape"}

	return Product{
		title: randomTitles[rand.Intn(len(randomTitles)-1)],
		price: rand.Intn(100),
	}
}
