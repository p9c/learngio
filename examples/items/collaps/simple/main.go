package main

import (
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image"
)

var (
	button = new(widget.Button)
)

func main() {
	gofont.Register()
	th := material.NewTheme()
	//imgOp := paint.ImageOp{}
	//
	//existingImageFile, err := os.Open("sample.jpeg")
	//if err != nil {
	//	// Handle error
	//}
	//defer existingImageFile.Close()
	//
	//loadedImage, err := jpeg.Decode(existingImageFile)
	//if err != nil {
	//	// Handle error
	//}
	//sz := 555
	//if imgOp.Size().X != sz {
	//	imgRender := image.NewRGBA(image.Rectangle{Max: image.Point{X: sz, Y: sz}})
	//	draw.ApproxBiLinear.Scale(imgRender,
	//		imgRender.Bounds(),
	//		loadedImage,
	//		loadedImage.Bounds(),
	//		draw.Src, nil)
	//	imgOp = paint.NewImageOp(imgRender)
	//}

	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(400), unit.Dp(800)),
			app.Title("ParallelCoin"),
		)
		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)
				//imgRender := th.Image(imgOp)
				//imgRender.Scale = float32(sz) / float32(gtx.Px(unit.Dp(float32(sz))))

				layout.Flex{
					Axis: layout.Vertical,
				}.Layout(gtx,
					layout.Flexed(1, func() {

						th.Button("dada").Layout(gtx, button)
						th.Button("aaaaadada").Layout(gtx, button)
						//layout.Flex{
						//	Axis: layout.Horizontal,
						//}.Layout(gtx,
						//	layout.Rigid(func() {
						//		item(gtx, imgRender, sz)
						//	}),
						//	layout.Rigid(func() {
						//		item(gtx, imgRender, sz)
						//	}),
						//	layout.Rigid(func() {
						//		item(gtx, imgRender, sz)
						//	}),
						//	layout.Rigid(func() {
						//		item(gtx, imgRender, sz)
						//	}),
						//)
					}))
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

func toPointF(p image.Point) f32.Point {
	return f32.Point{X: float32(p.X), Y: float32(p.Y)}
}

func item(gtx *layout.Context, img material.Image, sz int) {
	img.Layout(gtx)
	gtx.Dimensions = layout.Dimensions{Size: image.Point{X: sz, Y: sz}}
}
