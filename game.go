package goscii

import "bytes"

type Game struct {
	Width, Height uint
	Scenes        []*Scene
	CurrentScene  *Scene
	Buffer        bytes.Buffer
}

func NewGame(width, height uint) *Game {
	return &Game{
		Width:        width,
		Height:       height,
		CurrentScene: nil,
		Scenes:       []*Scene{},
		Buffer:       *bytes.NewBuffer(nil),
	}
}

func (g *Game) NewScene(s *Scene) {
	g.Scenes = append(g.Scenes, s)
}
