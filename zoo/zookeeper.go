package zoo

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
)

type stats struct {
	Agility int
	Luck    int
}

type Zookeeper struct {
	Name string
	Cage *Cage
	stats
}

func (z *Zookeeper) Catch(animal string, world *FreeWorld) {
	a, isExist := world.findAnimal(animal)
	if !isExist {
		fmt.Printf(color.RedString("There is no %s in FreeWorld\n"), animal)
		world.showAnimals()
		return
	}

	if !z.isAgilityBiggerThan(a.Agility) {
		fmt.Printf(color.RedString("You can't Catch %s. Not enought agility\n"), a.Title)
		return
	}

	if !z.putAnimalIntoTheCage(a, world) {
		return
	}

	fmt.Printf(color.GreenString("Great, You caught the %s \n"), a.Title)
}

func (z *Zookeeper) ChangeCage(c *Cage) {
	z.Cage = c
	fmt.Printf(color.CyanString("Your cage changed. Now your agility is %d \n"), z.Agility-z.Cage.Size)
}

func (z *Zookeeper) BringToZoo(zoo *Zoo) {
	for _, animal := range z.Cage.Animals {
		zoo.Animals = append(zoo.Animals, animal)
		delete(z.Cage.Animals, animal.Title)
		fmt.Printf(color.GreenString("Your brought %s to the Zoo.\n"), animal.Title)
	}
}

func (z *Zookeeper) isAgilityBiggerThan(agility int) bool {
	luckPoints := rand.Intn(z.Luck)
	penalty := z.Cage.Size

	return (luckPoints + z.Agility - penalty) > agility
}

func (z *Zookeeper) canPutAnimalIntoTheCage(animal Animal) bool {
	return z.Cage.FreeSpace() >= animal.Size
}

func (z *Zookeeper) putAnimalIntoTheCage(a Animal, world *FreeWorld) bool {
	return z.Cage.PutAnimalInside(a, world)
}
