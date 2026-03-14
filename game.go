package goscii

import (
	"bytes"
	"fmt"
)

type Game struct {
	Width, Height int
	Scenes        []*Scene
	CurrentScene  *Scene
	Buffer        bytes.Buffer
}

func (g *Game) Render() {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			g.CurrentScene.Tilemap[y][x] = ' '
		}
	}
	fmt.Println(g.CurrentScene.Tilemap)
	for _, sprite := range g.CurrentScene.Sprites {
		x := int(sprite.Pos.X)
		y := int(sprite.Pos.Y)
		if x >= 0 && x < g.Width && y >= 0 && y < g.Height {
			g.CurrentScene.Tilemap[y][x] = sprite.Rep
		}
	}
}

func NewGame(width, height int) *Game {
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
