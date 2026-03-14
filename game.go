package goscii

import "bytes"

type Game struct {
	Width, Height uint
	Scenes        []*Scene
	CurrentScene  *Scene
	Buffer        bytes.Buffer
}
