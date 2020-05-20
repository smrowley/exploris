package main

import (
	"exploris/player"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("Exploris", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_OPENGL)

	if err != nil {
		panic(err)
	}

	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	if err != nil {
		panic(err)
	}

	defer renderer.Destroy()

	player, err := player.NewPlayer(renderer, screenWidth/2, screenHeight-32)

	if err != nil {
		panic(err)
	}

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		keys := sdl.GetKeyboardState()

		if (keys[sdl.SCANCODE_LCTRL] == 1 || keys[sdl.SCANCODE_RCTRL] == 1) && (keys[sdl.SCANCODE_Q] == 1 || keys[sdl.SCANCODE_W] == 1) {
			return
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		player.Draw(renderer)
		player.Update()

		renderer.Present()
	}
}
