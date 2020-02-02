package main

import (
	"fmt"
	"github.com/gop9/olt/gio/app"
	"github.com/gop9/olt/gio/f32"
	"github.com/gop9/olt/gio/io/system"
	"github.com/gop9/olt/gio/layout"
	"github.com/gop9/olt/gio/op/paint"
	"github.com/gop9/olt/gio/unit"
	"github.com/gop9/olt/gio/widget/material"
	"github.com/p9c/learngio/pkg/fonts"

	"image"
	"image/color"
)

func main() {
	fonts.Register()

	th := material.NewTheme()
	go func() {
		w := app.NewWindow()
		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)
				t := Table{
					Title:      "Test Table",
					TitleSize:  0,
					TitleColor: "ffcfcf30",
					BgColor:    "ffcf3030",
					CellNames:  false,
					Sorting:    false,
					Paging:     false,
					Filtering:  false,
					PageSize:   0,
					Padding:    []float32{0, 0, 0, 0},
					Border:     []float32{0, 0, 0, 0},
					Corners:    []float32{0, 0, 0, 0},
					Cells:      nil,
				}
				t.tableLayout(gtx, th)

				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

type Table struct {
	Title      string
	TitleSize  int
	TitleColor string
	BgColor    string
	CellNames  bool
	Sorting    bool
	Paging     bool
	Filtering  bool
	PageSize   int
	Padding    []float32
	Border     []float32
	Corners    []float32
	Cells      []TableCell
}

type TableCell struct {
	Id      int
	Title   string
	Content func()
}

func fill(gtx *layout.Context, col color.RGBA) {
	cs := gtx.Constraints
	d := image.Point{X: cs.Width.Min, Y: cs.Height.Min}
	dr := f32.Rectangle{
		Max: f32.Point{X: float32(d.X), Y: float32(d.Y)},
	}
	paint.ColorOp{Color: col}.Add(gtx.Ops)
	paint.PaintOp{Rect: dr}.Add(gtx.Ops)
	gtx.Dimensions = layout.Dimensions{Size: d}
}

func drawList(gtx *layout.Context, list *layout.List, th *material.Theme) {
	widgets := []func(){
		func() {
			layout.Rigid(func() {
				txt := fmt.Sprintf("List element #%d")

				th.H3(txt).Layout(gtx)
			})
		},
		func() {

		},
	}
	list.Layout(gtx, len(widgets), func(i int) {
		layout.UniformInset(unit.Dp(16)).Layout(gtx, widgets[i])
	})

}
