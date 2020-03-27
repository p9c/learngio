package main

import (
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/p9c/learngio/examples/image/wav/waveform"
	"github.com/p9c/learngio/examples/image/wav/wavreader"
	"image"
	"image/color"
	"os"
)

var (
	ww = new(wav)
)

type wav struct {
	w *wavreader.Reader
}

func main() {
	th := material.NewTheme()
	imgOp := paint.ImageOp{}
	r, err := os.Open("muz.wav")
	if err != nil {
	}
	defer r.Close()

	wr, err := wavreader.New(r)
	if err != nil {
	}

	ww.w = wr

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
						cs := gtx.Constraints
						sz := gtx.Constraints.Width.Min
						if imgOp.Size().X != sz {

							imgOp = paint.NewImageOp(ww.loadWav(cs.Width.Max, cs.Height.Max, 1.6))
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

func (w *wav) loadWav(width, height int, zoom float32) *image.NRGBA {
	return waveform.MinMax(w.w, &waveform.Options{
		Width:   width,
		Height:  height,
		Zoom:    zoom,
		Half:    false,
		MarginL: 0,
		MarginR: 0,
		MarginT: 0,
		MarginB: 0,
		Front: &color.NRGBA{
			R: 255,
			G: 128,
			B: 0,
			A: 150,
		},
		Back: &color.NRGBA{
			A: 0,
		},
	})

}
