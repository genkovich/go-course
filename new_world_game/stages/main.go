package stages

import (
	"course/new_world_game/pkg"
	"fmt"
)

type Stage struct {
	Title        string
	Content      string
	OptionsTitle string
	Options      map[string]string
}

func (s *Stage) Run() string {
	fmt.Println(s.Title)
	pkg.PrintText(s.Content)
	_, result := pkg.PrintOptions(s.OptionsTitle, s.Options)

	return result
}

func main() {

}
