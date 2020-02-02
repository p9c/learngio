package main

import (
	"github.com/gop9/olt/gio/app"
	"github.com/gop9/olt/gio/io/system"
	"github.com/gop9/olt/gio/layout"
	"github.com/gop9/olt/gio/unit"
	"github.com/p9c/learngio/helpers"
	"image"
	"image/color"
)

func main() {
	go func() {
		w := app.NewWindow()
		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)
				layout.Flex{
					Axis:    layout.Horizontal,
					Spacing: layout.SpaceSides,
				}.Layout(gtx,
					layout.Flexed(0.5, func() {
						layout.Flex{
							Axis:    layout.Vertical,
							Spacing: layout.SpaceSides,
						}.Layout(gtx,
							layout.Flexed(0.5, func() {
								cs := gtx.Constraints
								helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, helpers.HexARGB("30303030"), [4]float32{0, 0, 0, 0}, unit.Dp(0))

								m := image.NewRGBA(image.Rect(0, 0, cs.Width.Max, cs.Height.Max))
								drawGradient(*m)
							}),
						)
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

func drawGradient(m image.RGBA) {
	size := m.Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			color := color.RGBA{
				uint8(255 * x / size.X),
				uint8(255 * y / size.Y),
				55,
				255}
			m.Set(x, y, color)
		}
	}
}
