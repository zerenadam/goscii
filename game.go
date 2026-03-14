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
	if g.CurrentScene == nil {
		fmt.Println("No scene to render")
		return
	}

	g.Buffer.Reset()
	scene := g.CurrentScene

	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			tile := scene.Tilemap[y][x]
			g.Buffer.WriteRune(tile)
		}
		g.Buffer.WriteRune('\n')
	}

	for _, sprite := range scene.Sprites {
		pos := sprite.Pos
		if int(pos.Y) < g.Height && int(pos.X) < g.Width {
			index := int(pos.Y)*g.Width + int(pos.X)
			bufBytes := g.Buffer.Bytes()
			bufBytes[index] = byte(sprite.Rep)
			g.Buffer = *bytes.NewBuffer(bufBytes)
		}
	}

	fmt.Print(g.Buffer.String())
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
