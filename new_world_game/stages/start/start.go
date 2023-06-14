package start

import (
	"course/new_world_game/pkg"
	"course/new_world_game/stages"
	"course/new_world_game/stages/valley"
	"course/new_world_game/units"
	"fmt"
)

func buildStage() *stages.Stage {
	return &stages.Stage{
		Title:        "Привіт!",
		Content:      "Радий вітати тебе в грі!",
		OptionsTitle: "Що будемо робити?",
		Options: map[string]string{
			"start": "Почнемо!",
			"exit":  "Якось потім...",
		},
		PostStageRunFunc: nextStage,
	}
}

func nextStage(selectedOption string, nil *units.Player) {
	switch selectedOption {
	case "start":
		fmt.Println("Введи своє їм'я")

		var name string
		fmt.Scan(&name)
		player := units.CreatePlayer(name)
		valley.Run(&player)

	case "exit":
		pkg.PrintText("До зустрічі!")
	}
}

func Run() {
	stage := buildStage()
	stage.RunWithPostFunc(nil)
}
