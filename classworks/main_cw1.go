package main

import (
	"encoding/json"
	"fmt"
)

type Tree struct {
	Name string `json:"name"`
	Nest *Nest  `json:"nest"`
}

type Nest struct {
	Birds []Bird `json:"birds"`
}

type Bird struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Song  *Song  `json:"song"`
}

type Song struct {
	Artist string `json:"artist"`
	Track  string `json:"track"`
}

func main() {

	JohnLennonSong := Song{Artist: "John Lennon", Track: "Imagine"}
	ImagineDragonsSong := Song{Artist: "Imagine Dragons", Track: "Believer"}
	BeatlesSong := Song{Artist: "The Beatles", Track: "Yesterday"}

	birds := []Bird{
		{Name: "Sparrow", Color: "Brown", Song: &JohnLennonSong},
		{Name: "Parrot", Color: "Green", Song: &ImagineDragonsSong},
		{Name: "Swan", Color: "White", Song: &BeatlesSong},
	}

	nest := Nest{Birds: birds}

	tree := Tree{Name: "Baobab", Nest: &nest}

	result, _ := json.MarshalIndent(tree, "", "  ")
	fmt.Println(string(result))

	emptyTree := Tree{}
	fmt.Println(emptyTree)

	emptyResult, _ := json.MarshalIndent(emptyTree, "", "  ")
	fmt.Println(string(emptyResult))
}
