package main

import (
	"github.com/zerenadam/goscii"
)

func main() {
	// Create a new game
	game := goscii.NewGame(10, 5)

	// Create a tilemap
	tilemapStr := `
..........\n+..........\n+..........\n+..........\n+..........`
	tilemap := goscii.TilemapFromString(tilemapStr)

	// Create a scene
	scene := &goscii.Scene{
		Name:    "DemoScene",
		Tilemap: tilemap,
		Sprites: []*goscii.Sprite{},
	}

	// Create a sprite
	sprite := &goscii.Sprite{
		Rep:  '@',
		Name: "Hero",
		Pos:  *goscii.NewPosition(2, 2),
	}
	scene.Sprites = append(scene.Sprites, sprite)

	// Set current scene
	game.CurrentScene = scene
	game.NewScene(scene)

	// Render the game
	game.Render()
}
