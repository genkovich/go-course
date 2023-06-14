package units

import (
	"course/new_world_game/units/weapons"
	"fmt"
	"math/rand"
)

type Monster struct {
	Name        string
	health      int
	handsWeapon weapons.Weapon
}

func CreateMonster(name string) Monster {
	return Monster{
		Name:        name,
		health:      45,
		handsWeapon: weapons.CreateClaws(),
	}
}

func (m *Monster) IsDied() bool {
	return m.health <= 0
}

func (m *Monster) diceOf(diceSize int) int {
	randNumber := rand.Intn(diceSize)
	fmt.Printf("Монстр кинув кості, випало %d з %d можливих\n", randNumber, diceSize)
	return randNumber
}

func (m *Monster) Attack(p *Player) {
	luck := m.diceOf(6)
	damage := 0

	if luck < 4 {
		fmt.Printf("Монстр промазав\n")
	} else if luck < 6 {
		damage = m.handsWeapon.Damage
		fmt.Printf("Вам нанесли %d шкоди\n", damage)
	} else if luck == 6 {
		additionalDamage := rand.Intn(m.handsWeapon.Damage)
		damage += additionalDamage
		fmt.Printf("Критичний удар! Вам нанесли %d шкоди\n", damage)
	}

	p.health -= damage
	fmt.Printf("у вас залишилось %d здоровья.\n", p.health)
}
