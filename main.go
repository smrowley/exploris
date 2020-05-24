package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

var (
	elements = make([]*Element, 0, 0)
	ground   = make([]*GroundPiece, screenWidth/10)
)

type GroundPiece struct {
	vx []int16
	vy []int16
}

func main() {
	rand.Seed(time.Now().UnixNano())

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

	elements = append(elements, NewPlayer(renderer, 100.0, 100.0))

	y := int16(400)
	offset := int16(0)
	for i := int16(0); int(i) < len(ground); i++ {
		x := i * 10
		num := rand.Intn(5) - 2
		offset += int16(num)
		offset = int16(math.Min(float64(offset), 5))
		offset = int16(math.Max(float64(offset), -5))
		ground[i] = &GroundPiece{[]int16{x, x + 9, x + 9, x}, []int16{y, y + offset, y + offset + 9, y + 9}}
		y += offset
	}

	groundColor := sdl.Color{R: 0, G: 153, B: 51, A: 255}

	fpsManager := &gfx.FPSmanager{}

	gfx.InitFramerate(fpsManager)
	gfx.SetFramerate(fpsManager, gfx.FPS_UPPER_LIMIT)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		keys := sdl.GetKeyboardState()

		if (keys[sdl.SCANCODE_LCTRL] == 1 || keys[sdl.SCANCODE_RCTRL] == 1) && (keys[sdl.SCANCODE_Q] == 1 || keys[sdl.SCANCODE_W] == 1 || keys[sdl.SCANCODE_C] == 1) {
			return
		}

		renderer.SetDrawColor(200, 200, 200, 255)
		renderer.Clear()

		for _, g := range ground {
			gfx.FilledPolygonColor(renderer, g.vx, g.vy, groundColor)
		}

		for _, el := range elements {
			for _, comp := range el.components {
				comp.OnDraw(renderer)
				comp.OnUpdate()
			}
		}

		renderer.Present()

		gfx.FramerateDelay(fpsManager)
	}
}
