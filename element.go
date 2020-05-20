package main

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type Vector struct {
	X, Y float64
}

type Component interface {
	onUpdate() error
	onDraw(renderer *sdl.Renderer) error
}

type Element struct {
	position   Vector
	rotation   float64
	active     bool
	components []Component
}

func (el *Element) AddComponent(new Component) {
	for _, existing := range el.components {
		if reflect.TypeOf(new) == reflect.TypeOf(existing) {
			panic(fmt.Sprintf("component of type %v already exists on element", reflect.TypeOf(new)))
		}
	}

	el.components = append(el.components, new)
}

func (el *Element) GetComponent(withType Component) Component {
	lookupType := reflect.TypeOf(withType)

	for _, existing := range el.components {
		if reflect.TypeOf(lookupType) == reflect.TypeOf(existing) {
			return existing
		}
	}

	panic(fmt.Sprintf("component of type %v already doesn't exist on element", lookupType))
}