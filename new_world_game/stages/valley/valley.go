package valley

import (
	"course/new_world_game/stages"
	"course/new_world_game/stages/cave"
	"course/new_world_game/stages/forest"
	"course/new_world_game/units"
)

func buildStage() *stages.Stage {
	return &stages.Stage{
		Title: "Долина",
		Content: "Ти прокидаєшся в центрі таємничої долини, серед містичних гір, " +
			"\nкотрі розташовані навколо тебе. Твоє тіло вкрите тонким шаром роси. " +
			"\nТи відчуваєш знайомий ваговий тягар на своєму спині - це твій рюкзак. " +
			"\n\rРаптом ти усвідомлюєш, що не пам'ятаєш хто ти і як сюди потрапив. " +
			"\nТи перевіряєш свій рюкзак. Всередині ти знаходиш книгу заклинань, " +
			"\nстарий меч та невеликий топор.",
		OptionsTitle: "Що будемо робити?",
		Options: map[string]string{
			"cave":   "Відправитись до печери",
			"forest": "Обрати стежку через ліс",
		},
		PostStageRunFunc: nextStage,
	}
}

func nextStage(selectedOption string, player *units.Player) {
	switch selectedOption {
	case "cave":
		cave.Run(player)
	case "forest":
		forest.Run(player)
	}
}

func Run(player *units.Player) {
	stage := buildStage()
	stage.RunWithPostFunc(player)
}
