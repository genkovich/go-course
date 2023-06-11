package cave

import (
	"course/new_world_game/fight"
	"course/new_world_game/pkg"
	"course/new_world_game/stages"
	"course/new_world_game/stages/forest"
	"course/new_world_game/units"
	"fmt"
)

func buildCaveStage() *StageValley {
	return &StageValley{
		Stage: stages.Stage{
			Title: "Долина",
			Content: "Ти збираєшся з силами і направляєшся до печери. Ти відчуваєш прохолоду, що виринає звідти. " +
				"\nТемрява здається безмежною. Напевно, тобі потрібне джерело світла." +
				"\nТи переглядаєш свою книгу заклинань і знаходиш просте заклинання освітлення. Чари освітлюють печеру перед тобою." +
				"\nПопереду в печері ти бачиш грубу, химерну тінь. Схоже, це великий звір, що регоче глибоко в горлі. Його очі сяють у темряві.",
			OptionsTitle: "Що будемо робити?",
			Options: map[string]string{
				"sword": "Використати свій меч, щоб атакувати звіра",
				"magic": "Використати заклинання, щоб відлякати звіра",
				"avoid": "Спробувати обійти звіра та продовжити свій шлях у печері.",
			},
		},
	}
}

type StageValley struct {
	stages.Stage
}

func Run(player *units.Player) {
	stage := buildCaveStage()
	result := stage.Run()

	fmt.Println(result)

	monster := units.CreateMonster("Grizzly")

	switch result {
	case "sword":
		player.WearWeapon("sword")
		stage.startPVE(player, &monster)
	case "magic":
		player.WearWeapon("book")
		stage.startPVE(player, &monster)
	case "avoid":
		if player.TryToAvoidFight() {
			forest.Run(player)
		} else {
			player.WearWeapon("axe")
			stage.startPVE(player, &monster)
		}
	}
}

func (s StageValley) startPVE(player *units.Player, monster *units.Monster) {
	playerWin := fight.PVE(player, monster)
	if playerWin {
		pkg.PrintText("Вітаю, тобі вдалось перемогти монстра, ти рухаєшся далі!")

		forest.Run(player)
	} else {
		pkg.PrintText("Нажаль ти помер, спробуй почати спочатку")
	}
}
