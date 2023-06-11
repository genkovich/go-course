package zoo

import (
	"encoding/json"
	"fmt"
)

func CatchAnimals() {
	zookeeper := Zookeeper{
		Name:  "John",
		stats: stats{Agility: 100, Luck: 50},
	}

	smallCage := Cage{
		Animals: map[string]Animal{},
		Size:    20,
	}

	bigCage := Cage{
		Animals: map[string]Animal{},
		Size:    100,
	}

	zookeeper.ChangeCage(&smallCage)

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

	zookeeper.Catch("lion", &freeWorld)
	zookeeper.BringToZoo(&zoo)
	zookeeper.ChangeCage(&bigCage)
	zookeeper.Catch("lion", &freeWorld)
	zookeeper.Catch("mouse", &freeWorld)
	zookeeper.Catch("elephant", &freeWorld)
	zookeeper.Catch("bear", &freeWorld)
	zookeeper.BringToZoo(&zoo)
	zookeeper.ChangeCage(&smallCage)

	zookeeper.Catch("tiger", &freeWorld)
	zookeeper.Catch("mouse", &freeWorld)
	zookeeper.BringToZoo(&zoo)

	zooResult, _ := json.MarshalIndent(zoo, "", "  ")
	fmt.Println(string(zooResult))

	freeWorld.showAnimals()
}
