package stages

import (
	"course/new_world_game/pkg"
	"course/new_world_game/units"
	"fmt"
)

type Stage struct {
	Title            string
	Content          string
	OptionsTitle     string
	Options          map[string]string
	PostStageRunFunc PostStageRunFunc
}

type PostStageRunFunc func(userInput string, player *units.Player)

func (s *Stage) run() string {
	fmt.Println(s.Title)
	pkg.PrintText(s.Content)
	_, result := pkg.PrintOptions(s.OptionsTitle, s.Options)

	return result
}

func (s *Stage) RunWithPostFunc(player *units.Player) {
	selectedOption := s.run()
	s.PostStageRunFunc(selectedOption, player)
}
