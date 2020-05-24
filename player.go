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

// Player struct
/*type Player struct {
	Texture *sdl.Texture
	X, Y    float64
}*/

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

// NewPlayer create a new player
/*func NewPlayer(renderer *sdl.Renderer, x, y int32) *Player {
	img, err := sdl.LoadBMP("sprites/sprite.bmp")
	if err != nil {
		return &Player{}
	}
	defer img.Free()
	texture, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		fmt.Println("error creating texture: ", err)
	}

	return &Player{Texture: texture, X: float64(x), Y: float64(y)}
}*/

/*
func (p *Player) OnDraw(renderer *sdl.Renderer) {
	renderer.Copy(p.Texture,
		&sdl.Rect{X: 0, Y: 0, W: spriteWidth, H: spriteHeight},
		&sdl.Rect{X: int32(p.X - float64(playerWidth/2.0)), Y: int32(p.Y-float64(playerHeight/2.0)) - 100, W: playerWidth, H: playerHeight})
}*/
/*
func (p *Player) OnUpdate() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 || keys[sdl.SCANCODE_A] == 1 {
		p.X -= playerSpeed
	}

	if keys[sdl.SCANCODE_RIGHT] == 1 || keys[sdl.SCANCODE_D] == 1 {
		p.X += playerSpeed
	}

}*/
