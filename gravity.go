package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Accelerator struct {
	vector  Vector
	element *Element
}

func newAccelerator(renderer *sdl.Renderer, vector Vector) *Accelerator {
	return &Accelerator{
		vector: vector,
	}
}

func (a *Accelerator) SetElement(element *Element) {
	a.element = element
}

func (a *Accelerator) OnDraw(renderer *sdl.Renderer) error {
	return nil
}

func (a *Accelerator) OnUpdate() error {
	a.element.velocity.X += a.vector.X
	a.element.velocity.Y += a.vector.Y
	return nil
}
