package units

import (
	"course/new_world_game/units/weapons"
	"fmt"
	"math/rand"
)

type Player struct {
	Name        string
	health      int
	mana        int
	inventory   Inventory
	handsWeapon weapons.Weapon
}

type Inventory struct {
	Items map[string]weapons.Weapon
}

func CreatePlayer(name string) Player {
	sword := weapons.CreateSword()
	axe := weapons.CreateAxe()
	bookOfMagic := weapons.CreateBookOfMagic()

	defaultInventory := Inventory{
		Items: map[string]weapons.Weapon{
			sword.Name:       sword,
			axe.Name:         axe,
			bookOfMagic.Name: bookOfMagic,
		},
	}

	return Player{
		Name:        name,
		health:      100,
		mana:        100,
		inventory:   defaultInventory,
		handsWeapon: weapons.Weapon{},
	}
}

func (p *Player) WearWeapon(title string) {
	p.handsWeapon = p.inventory.Items[title]
}

func (p *Player) diceOf(diceSize int) int {
	randNumber := rand.Intn(diceSize)
	fmt.Printf("Ви кинули кості, випало %d з %d можливих\n", randNumber, diceSize)
	return randNumber
}

func (p *Player) IsDied() bool {
	return p.health <= 0
}

func (p *Player) Attack(m *Monster) {
	luck := p.diceOf(6)
	damage := 0

	if p.handsWeapon.ManaCost != 0 {
		if p.mana < p.handsWeapon.ManaCost {
			fmt.Printf("Недостатньо мани для використання заклинання\n")
			return
		}

		p.mana -= p.handsWeapon.ManaCost
	}

	if luck < 2 {
		fmt.Printf("Ви промахнулись\n")
	} else if luck < 6 {
		damage = p.handsWeapon.Damage
		fmt.Printf("Ви нанесли %d шкоди\n", damage)
	} else if luck == 6 {
		additionalDamage := rand.Intn(p.handsWeapon.Damage)
		damage += additionalDamage
		fmt.Printf("Критичний удар! Ви нанесли %d шкоди\n", damage)
	}

	m.health -= damage
	fmt.Printf("у монстра залишилось %d здоровья.\n", m.health)
}

func (p *Player) TryToAvoidFight() bool {
	isAvoid := false
	message := "Вам не вдалося втекти\n"

	if p.diceOf(6) > 4 {
		isAvoid = true
		message = "Ви втекли з поля бою\n"
	}

	fmt.Println(message)

	return isAvoid
}
