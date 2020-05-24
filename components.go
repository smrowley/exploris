package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type SpriteRenderer struct {
	texture *sdl.Texture
	element *Element
}

func newSpriteRenderer(renderer *sdl.Renderer, fileName string) *SpriteRenderer {
	return &SpriteRenderer{
		texture: textureFromBitmap(renderer, fileName),
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
	err := renderer.Copy(sr.texture,
		&sdl.Rect{X: 0, Y: 0, W: spriteWidth, H: spriteHeight},
		&sdl.Rect{X: int32(sr.element.position.X /* - float64(playerWidth/2.0)*/), Y: int32(sr.element.position.Y /* - float64(playerHeight/2.0)*/), W: playerWidth, H: playerHeight})

	return err
}

func (sr *SpriteRenderer) OnUpdate() error {
	sr.element.position.X += sr.element.velocity.X
	sr.element.position.Y += sr.element.velocity.Y

	col := sr.element.collisionPoints[0]
	x := sr.element.position.X

	index := int(x / 10)

	//fmt.Printf("index %v, y left: %v, player pos: %v\n", index, ground[index].vy[0], sr.element.position.Y)

	if int(ground[index].vy[0]) < int(sr.element.position.Y+col.relativePosition.Y+col.radius) {
		sr.element.position.Y = float64(int32(ground[index].vy[0]) - playerHeight)
		//fmt.Println("ay")
	}

	return nil
}
