package zoo

import "fmt"

type FreeWorld struct {
	Animals map[string]Animal
}

func (fw *FreeWorld) addAnimal(animal Animal) {
	fw.Animals[animal.Title] = animal
}

func (fw *FreeWorld) removeAnimal(animal string) {
	delete(fw.Animals, animal)
}

func (fw *FreeWorld) showAnimals() {
	fmt.Println("Animals in FreeWorld:")
	for _, animal := range fw.Animals {
		fmt.Println(animal.Title)
	}
}

func (fw *FreeWorld) findAnimal(animal string) (Animal, bool) {
	a, res := fw.Animals[animal]
	return a, res
}
