package start

import (
	"course/new_world_game/pkg"
	"course/new_world_game/stages"
	"course/new_world_game/stages/valley"
	"course/new_world_game/units"
	"fmt"
)

func buildStartStage() *StageStart {
	return &StageStart{
		Stage: stages.Stage{
			Title:        "Привіт!",
			Content:      "Радий вітати тебе в грі!",
			OptionsTitle: "Що будемо робити?",
			Options: map[string]string{
				"start": "Почнемо!",
				"exit":  "Якось потім...",
			},
		},
	}
}

func (s *StageStart) Run() {
	selectedOption := s.Stage.Run()

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

type StageStart struct {
	stages.Stage
}

func Run() {
	stage := buildStartStage()
	stage.Run()
}
