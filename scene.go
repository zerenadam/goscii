package goscii

import "strings"

type Scene struct {
	Name    string
	Tilemap [][]rune
	Sprites []*Sprite
}

func TilemapFromString(mapstr string) [][]rune {
	lines := strings.Split(strings.TrimSpace(mapstr), "\n")
	tilemap := make([][]rune, len(lines))
	for y, line := range lines {
		tilemap[y] = []rune(line)
	}
	return tilemap
}
