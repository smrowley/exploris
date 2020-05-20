package player

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed  = 0.05
	playerWidth  = 24
	playerHeight = 32
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

func (p *Player) Draw(renderer *sdl.Renderer) {
	renderer.Copy(p.Texture,
		&sdl.Rect{X: 0, Y: 0, W: 24, H: 32},
		&sdl.Rect{X: int32(p.X) - playerWidth/2, Y: int32(p.Y) - playerHeight/2, W: 24, H: 32})
}

func (p *Player) Update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 || keys[sdl.SCANCODE_A] == 1 {
		p.X -= playerSpeed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 || keys[sdl.SCANCODE_D] == 1 {
		p.X += playerSpeed
	}
}
