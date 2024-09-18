package main

import "math"

type Color struct {
	R, G, B uint8
}

type GradientStop struct {
	R, G, B  uint8
	position float64
}

type Gradient struct {
	colors []GradientStop
}

func Lerp(a, b float64, t float64) float64 {
	return a + t*(b-a)
}

func InterpolateColor(c1, c2 GradientStop, t float64) GradientStop {
	return GradientStop{
		R: uint8(math.Round(Lerp(float64(c1.R), float64(c2.R), t))),
		G: uint8(math.Round(Lerp(float64(c1.G), float64(c2.G), t))),
		B: uint8(math.Round(Lerp(float64(c1.B), float64(c2.B), t))),
	}
}

func (gradient *Gradient) GetColorFromGradient(t float64) GradientStop {

	for i := 0; i < len(gradient.colors)-1; i++ {
		posI := (gradient.colors[i].position)
		posI2 := (gradient.colors[i+1].position)
		if t >= posI && t <= posI2 {
			segmentT := (t - posI) / (posI2 - posI)
			return InterpolateColor(gradient.colors[i], gradient.colors[i+1], segmentT)
		}
	}

	return gradient.colors[len(gradient.colors)-1]
}
