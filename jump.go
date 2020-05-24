package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Jump struct {
	element *Element
}

func newJump(renderer *sdl.Renderer) *Jump {
	return &Jump{}
}

func (j *Jump) SetElement(element *Element) {
	j.element = element
}

func (j *Jump) OnDraw(renderer *sdl.Renderer) error {
	return nil
}

func (j *Jump) OnUpdate() error {
	keys := sdl.GetKeyboardState()

	if j.element.grounded && (keys[sdl.SCANCODE_W] == 1 || keys[sdl.SCANCODE_UP] == 1) {
		fmt.Println("JUMP!")
		j.element.velocity.Y -= 3.0
	}

	return nil
}
