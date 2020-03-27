package main

import (
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"golang.org/x/image/draw"
	"image"
	"image/jpeg"
	"os"
)

func main() {
	th := material.NewTheme()
	imgOp := paint.ImageOp{}

	existingImageFile, err := os.Open("sample.jpeg")
	if err != nil {
		// Handle error
	}
	defer existingImageFile.Close()

	loadedImage, err := jpeg.Decode(existingImageFile)
	if err != nil {
		// Handle error
	}

	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(400), unit.Dp(800)),
			app.Title("ParallelCoin"),
		)
		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)

				layout.Flex{
					Axis: layout.Horizontal,
				}.Layout(gtx,
					layout.Flexed(1, func() {

						sz := gtx.Constraints.Width.Min
						if imgOp.Size().X != sz {
							imgRender := image.NewRGBA(image.Rectangle{Max: image.Point{X: sz, Y: sz}})
							draw.ApproxBiLinear.Scale(imgRender,
								imgRender.Bounds(),
								loadedImage,
								loadedImage.Bounds(),
								draw.Src, nil)
							imgOp = paint.NewImageOp(imgRender)
						}
						imgRender := th.Image(imgOp)
						imgRender.Scale = float32(sz) / float32(gtx.Px(unit.Dp(float32(sz))))
						imgRender.Layout(gtx)

					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

func toPointF(p image.Point) f32.Point {
	return f32.Point{X: float32(p.X), Y: float32(p.Y)}
}
