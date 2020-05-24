package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Movement struct {
	element *Element
}

func newMovement(renderer *sdl.Renderer) *Movement {
	return &Movement{}
}

func (m *Movement) SetElement(element *Element) {
	m.element = element
}

func (a *Movement) OnDraw(renderer *sdl.Renderer) error {
	return nil
}

func (m *Movement) OnUpdate() error {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 || keys[sdl.SCANCODE_A] == 1 {
		m.element.position.X -= playerSpeed
	}

	if keys[sdl.SCANCODE_RIGHT] == 1 || keys[sdl.SCANCODE_D] == 1 {
		m.element.position.X += playerSpeed
	}

	return nil
}
