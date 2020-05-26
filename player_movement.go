package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Movement struct {
	element      *Element
	maxVelocity  float32
	acceleration float32
}

func newMovement(renderer *sdl.Renderer) *Movement {
	return &Movement{maxVelocity: 1.0, acceleration: .02}
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
		if float32(m.element.velocity.X) > -m.maxVelocity {
			m.element.velocity.X -= float64(m.acceleration)
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 || keys[sdl.SCANCODE_D] == 1 {
		if float32(m.element.velocity.X) < m.maxVelocity {
			m.element.velocity.X += float64(m.acceleration)
		}
	} else if m.element.velocity.X != 0 {
		if m.element.velocity.X > 0 {
			m.element.velocity.X -= float64(m.acceleration)
		} else {
			m.element.velocity.X += float64(m.acceleration)
		}
	}

	return nil
}
