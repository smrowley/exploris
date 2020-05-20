package main

import (
	"fmt"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed  = 0.10 //0.05
	spriteWidth  = 24
	spriteHeight = 32
	spriteScale  = 5.0
)

var (
	playerWidth  = int32(math.Round(spriteWidth * spriteScale))
	playerHeight = int32(math.Round(spriteHeight * spriteScale))
)

// Player struct
type Player struct {
	Texture *sdl.Texture
	X, Y    float64
}

// NewPlayer create a new player
func NewPlayer(renderer *sdl.Renderer, x, y int32) (*Player, error) {
	img, err := sdl.LoadBMP("sprites/sprite.bmp")
	if err != nil {
		return &Player{}, err
	}
	defer img.Free()
	texture, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		fmt.Println("error creating texture: ", err)
	}

	return &Player{Texture: texture, X: float64(x), Y: float64(y)}, nil
}

func (p *Player) OnDraw(renderer *sdl.Renderer) {
	renderer.Copy(p.Texture,
		&sdl.Rect{X: 0, Y: 0, W: spriteWidth, H: spriteHeight},
		&sdl.Rect{X: int32(p.X - float64(playerWidth/2.0)), Y: int32(p.Y-float64(playerHeight/2.0)) - 100, W: playerWidth, H: playerHeight})
}

func (p *Player) OnUpdate() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 || keys[sdl.SCANCODE_A] == 1 {
		p.X -= playerSpeed
	}

	if keys[sdl.SCANCODE_RIGHT] == 1 || keys[sdl.SCANCODE_D] == 1 {
		p.X += playerSpeed
	}

}
