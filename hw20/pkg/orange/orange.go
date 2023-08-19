package orange

import "math/rand"

type Orange struct {
	Size int `json:"size"`
}

func NewOrange() *Orange {
	size := rand.Intn(15-1) + 1
	return &Orange{Size: size}
}
