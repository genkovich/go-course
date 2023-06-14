package fight

import (
	"course/new_world_game/pkg"
	"course/new_world_game/units"
)

func PVE(p *units.Player, m *units.Monster) bool {

	for !p.IsDied() && !m.IsDied() {
		p.Attack(m)
		m.Attack(p)

		pkg.PrintText("Продовжити...")
	}

	return m.IsDied()
}
