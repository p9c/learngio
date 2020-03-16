package main

import (
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/samples"
	"github.com/llgcode/ps"
	"image"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//th := material.NewTheme()
	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, 555, 555))
	gc := draw2dimg.NewGraphicContext(dest)
	// Draw Android logo
	_, err := Main(gc, "png")
	if err != nil {
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
						addrQR := paint.NewImageOp(dest)
						sz := gtx.Constraints.Width.Constrain(gtx.Px(unit.Dp(100)))
						addrQR.Add(gtx.Ops)
						paint.PaintOp{
							Rect: f32.Rectangle{
								Max: f32.Point{
									X: float32(sz), Y: float32(sz),
								},
							},
						}.Add(gtx.Ops)
						gtx.Dimensions.Size = image.Point{X: sz, Y: sz}
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

func Main(gc draw2d.GraphicContext, ext string) (string, error) {
	gc.Save()

	// flip the image
	gc.Translate(0, 200)
	gc.Scale(0.35, -0.35)
	gc.Translate(70, -200)

	// Tiger postscript drawing
	tiger := samples.Resource("image", "tiger.ps", ext)

	// Draw tiger
	Draw(gc, tiger)
	gc.Restore()

	// Return the output filename
	return samples.Output("postscript", ext), nil
}

// Draw a tiger
func Draw(gc draw2d.GraphicContext, filename string) {
	// Open the postscript
	src, err := os.OpenFile(filename, 0, 0)
	if err != nil {
		panic(err)
	}
	defer src.Close()
	bytes, err := ioutil.ReadAll(src)
	reader := strings.NewReader(string(bytes))

	// Initialize and interpret the postscript
	interpreter := ps.NewInterpreter(gc)
	interpreter.Execute(reader)
}
