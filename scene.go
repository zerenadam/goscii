package goscii

type Scene struct {
	Name    string
	Tilemap [][]rune
	Sprites []*Sprite
}
