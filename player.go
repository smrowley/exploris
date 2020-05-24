package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed  = 1.0
	spriteWidth  = 24
	spriteHeight = 32
	spriteScale  = 1.0
)

var (
	playerWidth  = int32(math.Round(spriteWidth * spriteScale))
	playerHeight = int32(math.Round(spriteHeight * spriteScale))
)

func NewPlayer(renderer *sdl.Renderer, x, y float64) *Element {
	player := &Element{}

	player.position = Vector{
		X: x,
		Y: y,
	}

	player.active = true

	player.AddComponent(newSpriteRenderer(renderer, "sprites/sprite.bmp"))
	player.AddComponent(newAccelerator(renderer, Vector{0, .05}))
	player.AddComponent(newMovement(renderer))

	player.collisionPoints = append(player.collisionPoints, CollisionPoint{Vector{11, 24}, 8})

	return player
}
