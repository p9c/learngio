package main

import (
	"flag"
	"fmt"
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	qrcode "github.com/skip2/go-qrcode"
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/font/gofont"
)

type fill color.RGBA

var (
	theme *material.Theme
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
func init() {
	gofont.Register()
	theme = material.NewTheme()
	theme.Color.Primary = rgb(0x1b5e20)
	//theme.TextSize = unit.Sp(20)
}

type App struct {
	transList *layout.List
	qrBtn     widget.Button
	showQR    bool
	addrQR    paint.ImageOp
}

func main() {
	flag.Parse()
	go func() {
		runUI()
	}()
	app.Main()
}

func NewApp() *App {
	var err error
	var q *qrcode.QRCode
	q, err = qrcode.New("inputinputindfginputine645645wytergrfgsddfh45g45putinput", qrcode.Highest)
	checkError(err)
	//
	//if *negative {
	//	q.ForegroundColor, q.BackgroundColor = q.BackgroundColor, q.ForegroundColor
	//}

	//var png []byte
	q.BackgroundColor = rgb(0xe8f5e9)
	q.ForegroundColor = rgb(0xe11133)
	q.Level = 3
	checkError(err)
	return &App{
		transList: &layout.List{
			Axis: layout.Vertical,
		},
		addrQR: paint.NewImageOp(q.Image(500)),
	}
}

func runUI() {
	a := NewApp()
	w := app.NewWindow()
	gtx := layout.NewContext(w.Queue())
	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				if err := e.Err; err != nil {
					log.Fatal(err)
				}
				return
			case system.FrameEvent:
				gtx.Reset(e.Config, e.Size)
				a.Layout(gtx)
				e.Frame(gtx.Ops)
			}
		}
	}
}

func (a *App) Layout(gtx *layout.Context) {
	layout.Center.Layout(gtx, func() {
		sz := gtx.Constraints.Width.Constrain(gtx.Px(unit.Dp(500)))
		a.addrQR.Add(gtx.Ops)
		paint.PaintOp{
			Rect: f32.Rectangle{
				Max: f32.Point{
					X: float32(sz), Y: float32(sz),
				},
			},
		}.Add(gtx.Ops)
		gtx.Dimensions.Size = image.Point{X: sz, Y: sz}
	})
}

func rgb(c uint32) color.RGBA {
	return argb((0xff << 24) | c)
}

func argb(c uint32) color.RGBA {
	return color.RGBA{A: uint8(c >> 24), R: uint8(c >> 16), G: uint8(c >> 8), B: uint8(c)}
}
