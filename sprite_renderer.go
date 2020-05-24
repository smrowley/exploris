package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type SpriteRenderer struct {
	texture        *sdl.Texture
	element        *Element
	rowIndex       uint8
	animationIndex uint8
	flip           sdl.RendererFlip
}

func newSpriteRenderer(renderer *sdl.Renderer, fileName string) *SpriteRenderer {
	return &SpriteRenderer{
		texture: textureFromBitmap(renderer, fileName),
		flip:    sdl.FLIP_NONE,
	}
}

func textureFromBitmap(renderer *sdl.Renderer, fileName string) *sdl.Texture {
	img, err := sdl.LoadBMP(fileName)
	if err != nil {
		panic(fmt.Sprintf("Unable to load bitmap: %v. error: %v", fileName, err))
	}
	defer img.Free()
	texture, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Sprintf("error creating texture: %v", err))
	}

	return texture
}

func (sr *SpriteRenderer) SetElement(element *Element) {
	sr.element = element
}

func (sr *SpriteRenderer) OnDraw(renderer *sdl.Renderer) error {
	err := renderer.CopyEx(sr.texture,
		&sdl.Rect{X: int32(sr.animationIndex * spriteWidth), Y: int32(sr.rowIndex * spriteHeight), W: spriteWidth, H: spriteHeight},
		&sdl.Rect{X: int32(sr.element.position.X /* - float64(playerWidth/2.0)*/), Y: int32(sr.element.position.Y /* - float64(playerHeight/2.0)*/), W: playerWidth, H: playerHeight}, 0, nil, sr.flip)

	return err
}

func (sr *SpriteRenderer) OnUpdate() error {
	sr.element.grounded = false
	sr.element.position.X += sr.element.velocity.X
	sr.element.position.Y += sr.element.velocity.Y

	col := sr.element.collisionPoints[0]
	x := sr.element.position.X

	relativeX := (x + col.relativePosition.X) / 10 // ground width TODO read dynamically
	index := int(relativeX)
	relativeX -= float64(index)

	groundHeightAtPlayerX := float64(ground[index].vy[0]) - (float64(ground[index].vy[0]-ground[index].vy[1]) * relativeX) + 1

	//fmt.Printf("relativeX: %v, ground index: %v, v0: %v, v1: %v\n", relativeX, index, ground[index].vy[0], ground[index].vy[1])

	if groundHeightAtPlayerX < sr.element.position.Y+col.relativePosition.Y+col.radius {
		sr.element.position.Y = groundHeightAtPlayerX - float64(playerHeight)
		sr.element.velocity.Y = 0
		sr.element.grounded = true
	}

	fmt.Printf("grounded: %v\n", sr.element.grounded)

	fmt.Println(sr.element.velocity)

	if sr.element.velocity.X != 0 {
		fmt.Println("runnin")
		sr.rowIndex = 1
		sr.animationIndex++
		if sr.animationIndex > 3 {
			sr.animationIndex = 0
		}

		if sr.element.velocity.X < 0 {
			sr.flip = sdl.FLIP_HORIZONTAL
		} else {
			sr.flip = sdl.FLIP_NONE
		}
	} else {
		sr.rowIndex = 0
		sr.animationIndex = 0
	}

	return nil
}
