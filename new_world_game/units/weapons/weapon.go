package weapons

type Item struct {
	Name string
}

type Weapon struct {
	Item
	Damage   int
	ManaCost int
}

func CreateSword() Weapon {
	return Weapon{
		Item: Item{
			Name: "sword",
		},
		Damage: 10,
	}
}

func CreateAxe() Weapon {
	return Weapon{
		Item: Item{
			Name: "axe",
		},
		Damage: 15,
	}
}

func CreateBookOfMagic() Weapon {
	return Weapon{
		Item: Item{
			Name: "book",
		},
		Damage:   20,
		ManaCost: 10,
	}
}

func CreateClaws() Weapon {
	return Weapon{
		Item: Item{
			Name: "Claws",
		},
		Damage: 30,
	}
}
