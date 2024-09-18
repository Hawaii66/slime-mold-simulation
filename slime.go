package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Slime struct {
	x         float64
	y         float64
	direction float64
	speed     float64
}

func NewSlime(speed float64) *Slime {
	centerx := float64(SCREEN_SIZE_WIDTH) / 2.0
	centery := float64(SCREEN_SIZE_HEIGHT) / 2.0

	dir := RandomDirection()
	offset := config.startCircleSize
	spawnx := centerx + (math.Cos(dir))*offset
	spawny := centery + (math.Sin(dir))*offset

	var dirToCenter float64
	if config.startTowardsCenter {
		dirToCenter = math.Atan2(centery-spawny, centerx-spawnx)
	} else {
		dirToCenter = RandomDirection()
	}

	return &Slime{
		x:         (spawnx),
		y:         (spawny),
		direction: dirToCenter,
		speed:     speed,
	}

}

func RandomDirection() float64 {
	return rand.Float64() * 2 * math.Pi
}

func (slime *Slime) Update(dt float64, world *World) {
	slime.MoveForward(dt)
	slime.BoundPosition(world.rect)
	slime.SampleAndSteer(world)
}

func (slime *Slime) SampleAndSteer(world *World) {
	forward := slime.Sample(0, config.trailSampleDirectionLength, world)
	left := slime.Sample(-config.trailSampleDirectionOffset, config.trailSampleDirectionLength, world)
	right := slime.Sample(config.trailSampleDirectionOffset, config.trailSampleDirectionLength, world)

	if forward == -1 {
		return
	}

	turnForce := config.trailTurnForce + (rand.Float64()*2-1)*config.trailRandomTurnForce

	if right == -1 && left != -1 {
		slime.direction -= turnForce
		return
	}
	if left == -1 && right != -1 {
		slime.direction += turnForce
		return
	}

	if forward > left && forward > right {
		return
	}

	if left > right {
		slime.direction -= turnForce
		return
	}

	if right > left {
		slime.direction += turnForce
		return
	}
}

func (slime *Slime) Sample(directionOffset float64, offset float64, world *World) float64 {
	x := slime.x + math.Cos(slime.direction+directionOffset)*offset
	y := slime.y + math.Sin(slime.direction+directionOffset)*offset

	if !world.rect.IsInside(int(x), int(y)) {
		return -1
	}

	val := 0.0
	sampleRegionSize := config.trailSampleSize
	for dx := -sampleRegionSize; dx <= sampleRegionSize; dx++ {
		for dy := -sampleRegionSize; dy <= sampleRegionSize; dy++ {
			a := int(x) + (dx)
			b := int(y) + (dy)

			if world.rect.IsInside((a), (b)) {
				val += (world.GetValue((a), (b)))
			} else {
				return -1
			}
		}
	}

	return val
}

type Coord struct {
	x, y int
}

func (slime *Slime) SamplePositions(directionOffset float64, offset float64, world *World) []Coord {
	x := slime.x + math.Cos(slime.direction+directionOffset)*offset
	y := slime.y + math.Sin(slime.direction+directionOffset)*offset

	if x < float64(world.rect.originX) || y < float64(world.rect.originY) || x > float64(world.rect.endX-1) || y > float64(world.rect.endY-1) {
		return []Coord{}
	}

	test := []Coord{}
	sampleRegionSize := 2
	for dx := -sampleRegionSize; dx <= sampleRegionSize; dx++ {
		for dy := -sampleRegionSize; dy <= sampleRegionSize; dy++ {
			a := x + float64(dx)
			b := y + float64(dy)
			if world.rect.IsInside(int(a), int(b)) {
				test = append(test, Coord{x: int(a), y: int(b)})
			}
		}
	}

	return test
}

func (slime *Slime) MoveForward(dt float64) {
	dX := math.Cos(slime.direction)
	dY := math.Sin(slime.direction)

	slime.x += dX * dt * slime.speed
	slime.y += dY * dt * slime.speed
}

func (slime *Slime) BoundPosition(rect *Rect) {
	if !rect.IsInside(int(slime.x), int(slime.y)) {
		clampedX, clampedY := rect.ClampInside(int(slime.x), int(slime.y))
		slime.x = float64(clampedX)
		slime.y = float64(clampedY)
		slime.direction = RandomDirection()
	}
}

func (slime *Slime) Log() {
	s := fmt.Sprintf("X: %f, Y: %f, dir: %f, spd: %f", slime.x, slime.y, slime.direction, slime.speed)
	fmt.Println(s)
}

func (slime *Slime) DrawDebug(screen *ebiten.Image) {
	screen.Set(int(slime.x), int(slime.y), color.RGBA{0xff, 0xff, 0x00, 0xff})

	samled := slime.SamplePositions(0, 12, &world)
	for i := 0; i < len(samled); i++ {
		screen.Set(samled[i].x, samled[i].y, color.RGBA{0xff, 0x00, 0xff, 0xff})
	}
	screen.Set(int(slime.x), int(slime.y), color.RGBA{0x00, 0xff, 0xff, 0xff})

	samled2 := slime.SamplePositions(0.9, 12, &world)
	for i := 0; i < len(samled2); i++ {
		screen.Set(samled2[i].x, samled2[i].y, color.RGBA{0xff, 0x00, 0xff, 0xff})
	}

	samled3 := slime.SamplePositions(-0.9, 12, &world)
	for i := 0; i < len(samled3); i++ {
		screen.Set(samled3[i].x, samled3[i].y, color.RGBA{0xff, 0x00, 0xff, 0xff})
	}
}
