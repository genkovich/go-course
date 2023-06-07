package zoo

import (
	"fmt"
	"github.com/fatih/color"
)

type Cage struct {
	Animals map[string]Animal
	Size    int
}

func (c *Cage) FreeSpace() int {
	occupiedSpace := 0
	for _, animalInCage := range c.Animals {
		occupiedSpace = occupiedSpace + animalInCage.Size
	}

	return c.Size - occupiedSpace
}

func (c *Cage) PutAnimalInside(animal Animal, world *FreeWorld) bool {
	if c.FreeSpace() < animal.Size {
		fmt.Printf(color.RedString("You can't put %s into the cage. Not enought space\n"), animal.Title)
		return false
	}

	c.Animals[animal.Title] = animal
	world.removeAnimal(animal.Title)

	return true
}
