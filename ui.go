package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

func CreateUI() *ebitenui.UI {
	face, _ := LoadFont(20)

	rootContainer := widget.NewContainer(

		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(2),
			widget.GridLayoutOpts.Spacing(20, 20),
			widget.GridLayoutOpts.Padding(widget.Insets{
				Top:  20,
				Left: 20,
			}),
		)),
	)

	label := widget.NewText(
		widget.TextOpts.Text("Turn force", face, color.White),
	)

	rootContainer.AddChild(label)

	slider := widget.NewSlider(
		widget.SliderOpts.Direction(widget.DirectionHorizontal),
		widget.SliderOpts.MinMax(0, 100),

		widget.SliderOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionStart,
				VerticalPosition:   widget.AnchorLayoutPositionStart,
			}),
			widget.WidgetOpts.MinSize(200, 6),
		),
		widget.SliderOpts.Images(
			&widget.SliderTrackImage{
				Idle:  image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
				Hover: image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
			},
			&widget.ButtonImage{
				Idle:    image.NewNineSliceColor(color.NRGBA{255, 100, 100, 255}),
				Hover:   image.NewNineSliceColor(color.NRGBA{255, 100, 100, 255}),
				Pressed: image.NewNineSliceColor(color.NRGBA{255, 50, 100, 255}),
			},
		),
		widget.SliderOpts.FixedHandleSize(6),
		widget.SliderOpts.TrackOffset(0),
		widget.SliderOpts.PageSizeFunc(func() int {
			return 1
		}),
		widget.SliderOpts.ChangedHandler(func(args *widget.SliderChangedEventArgs) {
			turnForceMax := 2 * math.Pi
			config.trailTurnForce = (float64(args.Current)) * (turnForceMax) / (100)
			fmt.Println("Turn force", config.trailTurnForce)

		}),
	)
	slider.Current = int((config.trailTurnForce) * (100 / (2.0 * math.Pi)))
	rootContainer.AddChild(slider)

	label2 := widget.NewText(
		widget.TextOpts.Text("Turn force random", face, color.White),
	)

	rootContainer.AddChild(label2)

	slider2 := widget.NewSlider(
		widget.SliderOpts.Direction(widget.DirectionHorizontal),
		widget.SliderOpts.MinMax(0, 100),

		widget.SliderOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionStart,
				VerticalPosition:   100,
			}),
			widget.WidgetOpts.MinSize(200, 6),
		),
		widget.SliderOpts.Images(
			&widget.SliderTrackImage{
				Idle:  image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
				Hover: image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
			},
			&widget.ButtonImage{
				Idle:    image.NewNineSliceColor(color.NRGBA{255, 100, 100, 255}),
				Hover:   image.NewNineSliceColor(color.NRGBA{255, 100, 100, 255}),
				Pressed: image.NewNineSliceColor(color.NRGBA{255, 50, 100, 255}),
			},
		),
		widget.SliderOpts.FixedHandleSize(6),
		widget.SliderOpts.TrackOffset(0),
		widget.SliderOpts.PageSizeFunc(func() int {
			return 1
		}),
		widget.SliderOpts.ChangedHandler(func(args *widget.SliderChangedEventArgs) {
			turnForceMax := 2 * math.Pi
			config.trailRandomTurnForce = (float64(args.Current)) * (turnForceMax) / (100)
			fmt.Println("Turn random force", config.trailRandomTurnForce)

		}),
	)
	slider2.Current = int(config.trailRandomTurnForce * (100 / (2 * math.Pi)))
	rootContainer.AddChild(slider2)

	label3 := widget.NewText(
		widget.TextOpts.Text("World decrement speed", face, color.White),
	)

	rootContainer.AddChild(label3)

	slider3 := widget.NewSlider(
		widget.SliderOpts.Direction(widget.DirectionHorizontal),
		widget.SliderOpts.MinMax(0, 100),

		widget.SliderOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionStart,
				VerticalPosition:   100,
			}),
			widget.WidgetOpts.MinSize(200, 6),
		),
		widget.SliderOpts.Images(
			&widget.SliderTrackImage{
				Idle:  image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
				Hover: image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
			},
			&widget.ButtonImage{
				Idle:    image.NewNineSliceColor(color.NRGBA{255, 100, 100, 255}),
				Hover:   image.NewNineSliceColor(color.NRGBA{255, 100, 100, 255}),
				Pressed: image.NewNineSliceColor(color.NRGBA{255, 50, 100, 255}),
			},
		),
		widget.SliderOpts.FixedHandleSize(6),
		widget.SliderOpts.TrackOffset(0),
		widget.SliderOpts.PageSizeFunc(func() int {
			return 1
		}),
		widget.SliderOpts.ChangedHandler(func(args *widget.SliderChangedEventArgs) {
			turnForceMax := 0.1
			config.worldDecrementSpeed = (float64(args.Current)) * (turnForceMax) / (100)
			fmt.Println("World decrement speed", config.worldDecrementSpeed)

		}),
	)
	slider3.Current = int(config.worldDecrementSpeed * (100 / (0.1)))
	rootContainer.AddChild(slider3)

	label4 := widget.NewText(
		widget.TextOpts.Text("Sample offset", face, color.White),
	)

	rootContainer.AddChild(label4)

	slider4 := widget.NewSlider(
		widget.SliderOpts.Direction(widget.DirectionHorizontal),
		widget.SliderOpts.MinMax(0, 100),

		widget.SliderOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionStart,
				VerticalPosition:   100,
			}),
			widget.WidgetOpts.MinSize(200, 6),
		),
		widget.SliderOpts.Images(
			&widget.SliderTrackImage{
				Idle:  image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
				Hover: image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
			},
			&widget.ButtonImage{
				Idle:    image.NewNineSliceColor(color.NRGBA{255, 100, 100, 255}),
				Hover:   image.NewNineSliceColor(color.NRGBA{255, 100, 100, 255}),
				Pressed: image.NewNineSliceColor(color.NRGBA{255, 50, 100, 255}),
			},
		),
		widget.SliderOpts.FixedHandleSize(6),
		widget.SliderOpts.TrackOffset(0),
		widget.SliderOpts.PageSizeFunc(func() int {
			return 1
		}),
		widget.SliderOpts.ChangedHandler(func(args *widget.SliderChangedEventArgs) {
			turnForceMax := 1.5
			config.trailSampleDirectionOffset = (float64(args.Current)) * (turnForceMax) / (100)
			fmt.Println("Sample direction offset", config.trailSampleDirectionOffset)

		}),
	)

	slider4.Current = int(config.trailSampleDirectionOffset * (100 / 1.5))
	rootContainer.AddChild(slider4)

	label5 := widget.NewText(
		widget.TextOpts.Text("Sample length", face, color.White),
	)

	rootContainer.AddChild(label5)

	slider5 := widget.NewSlider(
		widget.SliderOpts.Direction(widget.DirectionHorizontal),
		widget.SliderOpts.MinMax(0, 100),

		widget.SliderOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionStart,
				VerticalPosition:   100,
			}),
			widget.WidgetOpts.MinSize(200, 6),
		),
		widget.SliderOpts.Images(
			&widget.SliderTrackImage{
				Idle:  image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
				Hover: image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
			},
			&widget.ButtonImage{
				Idle:    image.NewNineSliceColor(color.NRGBA{255, 100, 100, 255}),
				Hover:   image.NewNineSliceColor(color.NRGBA{255, 100, 100, 255}),
				Pressed: image.NewNineSliceColor(color.NRGBA{255, 50, 100, 255}),
			},
		),
		widget.SliderOpts.FixedHandleSize(6),
		widget.SliderOpts.TrackOffset(0),
		widget.SliderOpts.PageSizeFunc(func() int {
			return 1
		}),
		widget.SliderOpts.ChangedHandler(func(args *widget.SliderChangedEventArgs) {
			turnForceMax := 20.0
			config.trailSampleDirectionLength = (float64(args.Current)) * (turnForceMax) / (100)
			fmt.Println("Sample direction length", config.trailSampleDirectionLength)
		}),
	)

	slider5.Current = int(config.trailSampleDirectionLength * (100 / 20.0))
	rootContainer.AddChild(slider5)

	ui := &ebitenui.UI{
		Container: rootContainer,
	}

	return ui
}

func LoadFont(size float64) (text.Face, error) {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &text.GoTextFace{
		Source: s,
		Size:   size,
	}, nil
}
