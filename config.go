package main

import (
	"strconv"
	"strings"
)

type Config struct {
	gradient                   Gradient
	sizeX, sizeY               int
	numOfSlimes                int
	slimeDefaultSpeed          float64
	slimeSpeedRandomness       float64
	worldDecrementSpeed        float64
	trailSampleDirectionOffset float64
	trailSampleDirectionLength float64
	trailTurnForce             float64
	trailSampleSize            int
	trailRandomTurnForce       float64
	startCircleSize            float64
	startTowardsCenter         bool
}

func aspectVideo(a int) int {
	return a * 9 / 16
}

func fromHex(hex string, pos float64) GradientStop {
	if len(hex) != 7 {
		panic("Wrong size for hex: " + hex)
	}

	split := strings.Split(hex, "")
	if split[0] != "#" {
		panic("Wrong hex format: " + hex)
	}

	r, _ := strconv.ParseUint(split[1]+split[2], 16, 8)
	g, _ := strconv.ParseUint(split[3]+split[4], 16, 8)
	b, _ := strconv.ParseUint(split[5]+split[6], 16, 8)

	return GradientStop{
		R:        uint8(r),
		G:        uint8(g),
		B:        uint8(b),
		position: pos,
	}
}

var blueYellowGradient = Gradient{
	colors: []GradientStop{
		fromHex("#16325b", 0),
		fromHex("#227b94", 0.3),
		fromHex("#78b7d0", 0.85),
		fromHex("#ffdc7f", 1),
	},
}

var redSkyGradient = Gradient{
	colors: []GradientStop{
		fromHex("#d20062", 0),
		fromHex("#d6589f", 0.25),
		fromHex("#d895da", 0.75),
		fromHex("#c4e4ff", 1),
	},
}

var sunSetGradient = Gradient{
	colors: []GradientStop{
		fromHex("#f9ed69", 0),
		fromHex("#f08a5d", 0.25),
		fromHex("#b83b53", 0.75),
		fromHex("#6a2c70", 1),
	},
}

var pinkGradient = Gradient{
	colors: []GradientStop{
		fromHex("#ff00ff", 0),
		fromHex("#00ffff", 0.85),
		fromHex("#0000ff", 0.95),
		fromHex("#00ffff", 1),
	},
}

var creamPinkGradient = Gradient{
	colors: []GradientStop{
		fromHex("#fff5e4", 0),
		fromHex("#ffe3e1", 0.85),
		fromHex("#ffd1d1", 0.95),
		fromHex("#ff9494", 1),
	},
}

var darkBlueGradient = Gradient{
	colors: []GradientStop{
		fromHex("#1b262c", 0),
		fromHex("#0f4c75", 0.85),
		fromHex("#3282b8", 0.95),
		fromHex("#bbe1fa", 1),
	},
}

var config = Config{
	sizeX:                      720,
	sizeY:                      (720),
	numOfSlimes:                1_000_000,
	slimeDefaultSpeed:          2,
	slimeSpeedRandomness:       0.5,
	worldDecrementSpeed:        0.05,
	trailSampleDirectionOffset: 0.9,
	trailSampleDirectionLength: 40,
	trailTurnForce:             0.1,
	trailRandomTurnForce:       0.3,
	trailSampleSize:            1,
	startCircleSize:            50.0,
	startTowardsCenter:         false,
	gradient:                   blueYellowGradient,
}

/*





var config = Config{
	sizeX:                      720,
	sizeY:                      aspectVideo(720),
	numOfSlimes:                500_000,
	slimeDefaultSpeed:          3,
	slimeSpeedRandomness:       1,
	worldDecrementSpeed:        0.015,
	trailSampleDirectionOffset: 0.3,
	trailSampleDirectionLength: 10,
	trailTurnForce:             0.9,
	trailSampleSize:            1,
	startCircleSize:            75.0,
	startTowardsCenter: true,
	trailRandomTurnForce: 0.3,
	gradient: Gradient{
		colors: []GradientStop{
			fromHex("#ff00ff", 0),
			fromHex("#00ffff", 0.25),
			fromHex("#0000ff", 0.5),
			fromHex("#00ffff", 1),
		},
	},
}



*/
