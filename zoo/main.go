package zoo

import (
	"encoding/json"
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

func (z *Zookeeper) catch(animal string, world *FreeWorld) {
	if _, ok := world.Animals[animal]; !ok {
		fmt.Printf(color.RedString("There is no %s in FreeWorld\n"), animal)
		return
	}

	a := world.Animals[animal]

	luckPoints := rand.Intn(z.Luck)
	penalty := z.Cage.Size

	if (luckPoints + z.Agility - penalty) < a.Agility {
		fmt.Printf(color.RedString("You can't catch %s. Not enought agility\n"), a.Title)
		return
	}

	occupiedSpace := 0
	for _, animalInCage := range z.Cage.Animals {
		occupiedSpace = occupiedSpace + animalInCage.Size
	}

	cageEmptySpace := z.Cage.Size - occupiedSpace

	if cageEmptySpace < a.Size {
		fmt.Printf(color.RedString("You can't put %s into the cage. Not enought space\n"), a.Title)
		return
	}

	z.Cage.Animals[animal] = a
	delete(world.Animals, animal)

	fmt.Printf(color.GreenString("Great, You caught the %s \n"), a.Title)
}

func (z *Zookeeper) changeCage(c *Cage) {
	c.Animals = make(map[string]Animal)
	z.Cage = c
	fmt.Printf(color.CyanString("Your cage changed. Now your agility is %d \n"), z.Agility-z.Cage.Size)
}

func (z *Zookeeper) bringToZoo(zoo *Zoo) {
	for _, animal := range z.Cage.Animals {
		zoo.Animals = append(zoo.Animals, animal)
		delete(z.Cage.Animals, animal.Title)
		fmt.Printf(color.GreenString("Your brought %s to the Zoo.\n"), animal.Title)
	}
}

type Cage struct {
	Animals map[string]Animal
	Size    int
}

type Animal struct {
	Title   string `json:"title"`
	Size    int    `json:"size"`
	Agility int    `json:"agility"`
}

type Zoo struct {
	Animals []Animal
}

type FreeWorld struct {
	Animals map[string]Animal
}

func CatchAnimals() {
	zookeeper := Zookeeper{
		Name:  "John",
		stats: stats{Agility: 100, Luck: 50},
	}

	smallCage := Cage{
		Size: 20,
	}

	bigCage := Cage{
		Size: 100,
	}

	zookeeper.changeCage(&smallCage)

	freeWorld := FreeWorld{
		Animals: map[string]Animal{
			"lion": Animal{
				Title:   "lion",
				Size:    10,
				Agility: 50,
			},
			"elephant": Animal{
				Title:   "elephant",
				Size:    100,
				Agility: 0,
			},
			"mouse": Animal{
				Title:   "mouse",
				Size:    1,
				Agility: 80,
			},
			"tiger": Animal{
				Title:   "tiger",
				Size:    15,
				Agility: 70,
			},
			"bear": Animal{
				Title:   "bear",
				Size:    50,
				Agility: 40,
			},
		},
	}

	zoo := Zoo{}

	zookeeper.catch("lion", &freeWorld)
	zookeeper.bringToZoo(&zoo)
	zookeeper.changeCage(&bigCage)
	zookeeper.catch("lion", &freeWorld)
	zookeeper.catch("mouse", &freeWorld)
	zookeeper.catch("elephant", &freeWorld)
	zookeeper.catch("bear", &freeWorld)
	zookeeper.bringToZoo(&zoo)
	zookeeper.changeCage(&smallCage)

	zookeeper.catch("tiger", &freeWorld)
	zookeeper.catch("mouse", &freeWorld)
	zookeeper.bringToZoo(&zoo)

	zooResult, _ := json.MarshalIndent(zoo, "", "  ")
	fmt.Println(string(zooResult))
}
