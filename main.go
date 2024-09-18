package main

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	ui *ebitenui.UI
}

var SCREEN_SIZE_WIDTH = config.sizeX
var SCREEN_SIZE_HEIGHT = config.sizeY

var world World

var slimes []*Slime
var rect = &Rect{
	originX: 0, originY: 0, endX: SCREEN_SIZE_WIDTH, endY: SCREEN_SIZE_HEIGHT,
}

func (g *Game) Update() error {
	g.ui.Update()

	var wg sync.WaitGroup
	per := 1_000

	for j := 0; j < len(slimes)/per; j++ {
		wg.Add(1)

		go func(j int) {
			for i := j * per; i < (j+1)*per; i++ {
				slime := slimes[i]
				slime.Update(1.0, &world)
				world.SetValue(int(slime.x), int(slime.y), 1)
			}
			wg.Done()
		}(j)
	}
	wg.Wait()

	world.Blur()
	world.Decrement(config.worldDecrementSpeed)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	colors := []byte{}

	for y := 0; y < world.rect.endY; y++ {
		for x := 0; x < world.rect.endX; x++ {
			val := world.GetValue(x, y)

			color := config.gradient.GetColorFromGradient(val)

			colors = append(colors, color.R)
			colors = append(colors, color.G)
			colors = append(colors, color.B)
			colors = append(colors, 0xff)
		}
	}

	screen.WritePixels(colors[:])

	// slimes[0].DrawDebug(screen)
	// slimes[1].DrawDebug(screen)
	// slimes[2].DrawDebug(screen)
	// slimes[3].DrawDebug(screen)
	// slimes[4].DrawDebug(screen)

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.ui.Draw(screen)
	}

	if ebiten.IsKeyPressed(ebiten.KeyF) {
		str := fmt.Sprintf("%f", ebiten.ActualFPS())
		ebitenutil.DebugPrint(screen, "FPS: "+str)
	}
}

func (g *Game) Layout(outsideWidth, outsideHieght int) (screenWidth, screenHeight int) {
	return SCREEN_SIZE_WIDTH, SCREEN_SIZE_HEIGHT
}

func main() {
	fromHex("#ffddcc", 0)

	world = *NewWorld(SCREEN_SIZE_WIDTH, SCREEN_SIZE_HEIGHT)

	for i := 0; i < config.numOfSlimes; i++ {
		slimes = append(slimes, NewSlime(config.slimeDefaultSpeed+(rand.Float64()*2-1)*config.slimeSpeedRandomness))
	}

	ebiten.SetWindowSize(SCREEN_SIZE_WIDTH, SCREEN_SIZE_HEIGHT)
	ebiten.SetWindowTitle("Slime Simulation")
	// ebiten.SetVsyncEnabled(false)

	game := &Game{
		ui: CreateUI(),
	}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
