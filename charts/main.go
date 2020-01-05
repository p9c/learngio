package main

import (
	"bytes"
	"fmt"
	"gioui.org/app"
	"gioui.org/io/pointer"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/ajstarks/svgo"
	"github.com/p9c/learngio/helpers"
	"github.com/vdobler/chart"
	"image"
	"image/color"
)

//
// Pie Charts
//
func pieChart() []byte {
	var b bytes.Buffer
	//_ = io.Writer(&b)
	svgOut := new(svg.SVG)

	//pie := chart.PieChart{Title: "Some Pies"}
	//data := []chart.CatValue{{"D", 10, false}, {"GB", 20, true}, {"CH", 30, false}, {"F", 60, false}}
	//lw := 4
	//red := chart.Style{LineColor: color.NRGBA{0xcc, 0x00, 0x00, 0xff}, FillColor: color.NRGBA{0xff, 0x80, 0x80, 0xff},
	//	LineStyle: chart.SolidLine, LineWidth: lw}
	//green := chart.Style{LineColor: color.NRGBA{0x00, 0xcc, 0x00, 0xff}, FillColor: color.NRGBA{0x80, 0xff, 0x80, 0xff},
	//	LineStyle: chart.SolidLine, LineWidth: lw}
	//blue := chart.Style{LineColor: color.NRGBA{0x00, 0x00, 0xcc, 0xff}, LineWidth: lw, LineStyle: chart.SolidLine, FillColor: color.NRGBA{0x80, 0x80, 0xff, 0xff}}
	//pink := chart.Style{LineColor: color.NRGBA{0x99, 0x00, 0x99, 0xff}, LineWidth: lw, LineStyle: chart.SolidLine, FillColor: color.NRGBA{0xaa, 0x60, 0xaa, 0xff}}
	//
	//styles := []chart.Style{red, green, blue, pink}
	//pie.FmtKey = chart.IntegerValue
	//pie.AddData("Data1", data, styles)
	//pie.Inner = 0
	//pie.Key.Cols = 2
	//pie.Key.Pos = "ibr"
	////dumper2.Plot(&pie)

	pie := chart.PieChart{Title: "Some Rings"}
	data2 := []chart.CatValue{{"D", 15, false}, {"GB", 25, false}, {"CH", 30, false}, {"F", 50, false}}

	lw := 2
	lightred := chart.Style{LineColor: color.NRGBA{0xcc, 0x40, 0x40, 0xff}, FillColor: color.NRGBA{0xff, 0xc0, 0xc0, 0xff},
		LineStyle: chart.SolidLine, LineWidth: lw}
	lightgreen := chart.Style{LineColor: color.NRGBA{0x40, 0xcc, 0x40, 0xff}, FillColor: color.NRGBA{0xc0, 0xff, 0xc0, 0xff},
		LineStyle: chart.SolidLine, LineWidth: lw}
	lightblue := chart.Style{LineColor: color.NRGBA{0x40, 0x40, 0xcc, 0xff}, FillColor: color.NRGBA{0xc0, 0xc0, 0xff, 0xff},
		LineWidth: lw, LineStyle: chart.SolidLine}
	lightpink := chart.Style{LineColor: color.NRGBA{0xaa, 0x00, 0xaa, 0xff}, FillColor: color.NRGBA{0xff, 0x80, 0xff, 0xff},
		LineWidth: lw, LineStyle: chart.SolidLine}
	lightstyles := []chart.Style{lightred, lightgreen, lightblue, lightpink}

	pie.Inner = 0.3
	pie.Key.Cols = 2
	pie.Key.Pos = "ibr"
	pie.FmtVal = chart.PercentValue
	chart.PieChartShrinkage = 0.55
	pie.FmtKey = chart.IntegerValue

	pie.AddData("2010", data2, lightstyles)
	//svgOut.Plot(&pie)pie
	//row, col := svgOut.Cnt/svgOut.N, svgOut.Cnt%svgOut.N
	//
	//sgr := svgg.AddTo(svgOut.S, col*svgOut.W, row*svgOut.H, svgOut.W, svgOut.H, "", 12, color.RGBA{0xff, 0xff, 0xff, 0xff})
	//c.Plot(sgr)
	//
	////tgr := txtg.New(100, 30)
	////c.Plot(tgr)
	////d.txtFile.Write([]byte(tgr.String() + "\n\n\n"))
	//
	//d.Cnt++

	svgOut.Start(2*500, 1*400)
	svgOut.Title("test chart")
	svgOut.Rect(0, 0, 2*500, 1*400, "fill: #cfcfcf")
	svgOut.End()

	//dumper.I = image.NewRGBA(image.Rect(0, 0, 2*500, 1*400))
	//bg := image.NewUniform(color.RGBA{0xff, 0xff, 0xff, 0xff})
	//draw.Draw(dumper.I, dumper.I.Bounds(), bg, image.ZP, draw.Src)
	svgOut = &svg.SVG{
		Writer: &b,
	}
	fmt.Println("dadada", b.String())

	return helpers.LoadSvgFile(b.Bytes(), 1)
}

//
//  Main
//
func main() {

	//svgOOO := ico.ParallelCoin
	//svgOOO := pieChart()

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
								helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, helpers.HexARGB("ff303030"), [4]float32{0, 0, 0, 0}, unit.Dp(0))

								layout.Align(layout.Center).Layout(gtx, func() {
									layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(4), Left: unit.Dp(5), Right: unit.Dp(4)}.Layout(gtx, func() {

										icsLogo, err := material.NewIcon(pieChart())
										if err != nil {
											//log.FATAL(err)
										}

										icsLogo.Color = color.RGBA{A: 0xff, R: 0xcf, G: 0xcf, B: 0xcf}
										icsLogo.Layout(gtx, unit.Dp(64))
									})
								})

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
