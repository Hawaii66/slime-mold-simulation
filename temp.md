# Save rendering

// for i := 0; i < len(slimes); i++ {
// slime := slimes[i]
// screen.Set(int(slime.x), int(slime.y), color.White)
// }

    colors := [SCREEN_SIZE_HEIGHT * SCREEN_SIZE_WIDTH * 4]byte{}

    i := 0
    for y := 0; y < world.rect.endY; y++ {
    	for x := 0; x < world.rect.endX; x++ {
    		val := world.GetValue(x, y)
    		u := uint8(val * 255)
    		colors[i] = u
    		colors[i+1] = u
    		colors[i+2] = u
    		colors[i+3] = 0xff

    		i += 4
    	}
    }
    screen.WritePixels(colors[:])

    // for x := 0; x < world.rect.endX; x++ {
    // 	for y := 0; y < world.rect.endY; y++ {
    // 		val := world.GetValue(x, y)

    // 		scaled := uint8(val * 255)
    // 		c := color.RGBA{scaled, scaled, scaled, 0xff}
    // 		screen.Set(x, y, c)

    // 	}
    // }

    str := fmt.Sprintf("%f", ebiten.ActualFPS())
    ebitenutil.DebugPrint(screen, "FPS: "+str)
