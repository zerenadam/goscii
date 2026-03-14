package main

import (
	"fmt"

	"github.com/zerenadam/goscii"
)

func main() {
	game := goscii.NewGame(20, 10)

	tilemap := `
####################
#                  #
#                  #
#                  #
#                  #
#                  #
#                  #
#                  #
#                  #
####################`

	hero := &goscii.Sprite{
		Rep:  '⚇',
		Name: "Hero",
		Pos:  *goscii.NewPosition(1, 1),
	}

	spike := &goscii.Sprite{
		Rep:  '▲',
		Name: "Spike",
		Pos:  *goscii.NewPosition(5, 5),
	}

	level_1 := &goscii.Scene{
		Name:    "Level 1",
		Tilemap: goscii.TilemapFromString(tilemap),
		Sprites: []*goscii.Sprite{
			hero,
			spike,
		},
	}

	game.NewScene(level_1)
	game.CurrentScene = level_1

	for {
		game.Render()
		if hero.Pos.X < 10 {
			hero.Pos.X += 1
		} else if hero.Pos.X == spike.Pos.X && hero.Pos.Y == spike.Pos.Y {
			fmt.Println("Game Over!")
			break
		}
	}
}
