package forest

import (
	"course/new_world_game/stages"
	"course/new_world_game/units"
	"fmt"
)

func buildForestStage() *StageForest {
	return &StageForest{
		Stage: stages.Stage{
			Title: "Ліс",
			Content: "Ти опиняєшся в лісі. Навколо тебе тягнуться високі дерева, вкриті густим мохом. " +
				"\nЗемля під твоїми ногами вкрита пухким шаром опалого листя. Ти чуєш крики птахів і шурхіт тварин, що ховаються в підліссі. " +
				"\nЦей ліс, попри свою віддаленість та дикість, викликає почуття миру і злагоди. Ти розумієш, що все погане позаду. Так і закінчилась твоя маленька історія",
			OptionsTitle: "І це все?",
			Options: map[string]string{
				"end": "Закінчити",
			},
		},
	}
}

type StageForest struct {
	stages.Stage
}

func Run(player *units.Player) {
	stage := buildForestStage()
	result := stage.Run()

	switch result {
	case "end":
		fmt.Printf("До зустрічі, %s", player.Name)
		return
	}
}
