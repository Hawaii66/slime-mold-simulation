package main

import (
	"math"
	"sync"
)

type World struct {
	image []float64
	rect  *Rect
}

func NewWorld(endX int, endY int) *World {
	rect := &Rect{
		originX: 0, originY: 0, endX: endX, endY: endY,
	}

	image := make([]float64, endX*endY)
	for x := 0; x < endX; x++ {
		for y := 0; y < endY; y++ {
			image[GetPos(x, y, rect)] = 0.0
		}
	}

	return &World{
		image: image, rect: rect,
	}
}

func GetPos(x, y int, rect *Rect) int {
	return x*(rect.endY-rect.originY) + y
}

func (world *World) GetValue(x, y int) float64 {
	return world.image[GetPos(x, y, world.rect)]
}

func (world *World) SetValue(x, y int, val float64) {
	if x > world.rect.endX-1 || y > world.rect.endY-1 || x < 0 || y < 0 {
		return
	}

	world.image[GetPos(x, y, world.rect)] = val
}

var offsets = []int{
	-1, 1, 0,
}

func (world *World) BlurSection(startX, startY, endX, endY int, rect *Rect, image *([]float64)) {
	for x := startX; x < endX; x++ {
		for y := startY; y < endY; y++ {
			total := 0.0
			sampled := 0.0
			for dxI := 0; dxI < len(offsets); dxI++ {
				for dyI := 0; dyI < len(offsets); dyI++ {
					dx := x + offsets[dxI]
					dy := y + offsets[dyI]

					if dx < 0 || dy < 0 || dx > rect.endX-1 || dy > rect.endY-1 {
						continue
					}

					total += world.GetValue(dx, dy)
					sampled += 1
				}
			}

			(*image)[GetPos(x, y, world.rect)] = total / sampled
		}
	}
}

func (world *World) Blur() {
	endX := world.rect.endX
	endY := world.rect.endY

	image := make([]float64, endX*endY)

	var wg sync.WaitGroup
	s := 50
	sizeX := SCREEN_SIZE_WIDTH / s
	sizeY := SCREEN_SIZE_HEIGHT / s

	for x := 0; x < s; x++ {
		for y := 0; y < s; y++ {
			wg.Add(1)

			go func(x, y int) {
				world.BlurSection(x*sizeX, y*sizeY, (x+1)*sizeX, (y+1)*sizeY, world.rect, &image)
				wg.Done()
			}(x, y)
		}
	}

	world.BlurSection(s*sizeX, 0, SCREEN_SIZE_WIDTH, SCREEN_SIZE_HEIGHT, world.rect, &image)
	world.BlurSection(0, s*sizeY, s*sizeX, SCREEN_SIZE_HEIGHT, world.rect, &image)

	wg.Wait()

	world.image = image
}

func (world *World) Decrement(d float64) {
	endX := world.rect.endX
	endY := world.rect.endY

	for x := 0; x < endX; x++ {
		for y := 0; y < endY; y++ {
			world.SetValue(x, y, math.Max(0, world.GetValue(x, y)-d))
		}
	}
}
