package main

import (
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/p9c/learngio/examples/icons/ico"
	"image"
	"image/color"
	"sort"
)

var (
	navList = &layout.List{
		Axis:      layout.Vertical,
		Alignment: layout.Start,
	}
	list = &layout.List{
		Axis:      layout.Vertical,
		Alignment: layout.Start,
	}
	navButtonsGroup = new(widget.Enum)
	icoBank         = ico.NewIco()
	groupsIco       = icoBank.GroupsIco()
)

func main() {
	gofont.Register()
	th := material.NewTheme()
	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(400), unit.Dp(800)),
			app.Title("ParallelCoin"),
		)
		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)
				layout.Flex{}.Layout(gtx,
					layout.Rigid(func() {
						DrawRectangle(gtx, 160, gtx.Constraints.Height.Max, color.RGBA{A: 0xff, R: 0xcf, G: 0x30, B: 0x30})
						navList.Layout(gtx, len(groupsIco), func(i int) {
							layout.UniformInset(unit.Dp(16)).Layout(gtx, func() {
								th.RadioButton(groupsIco[i], groupsIco[i]).Layout(gtx, navButtonsGroup)
							})
						})
					}),
					layout.Flexed(1, func() {
						icoS := make([]string, 0, len(icoBank[navButtonsGroup.Value(gtx)]))
						for k := range icoBank[navButtonsGroup.Value(gtx)] {
							icoS = append(icoS, k)
						}
						sort.Strings(icoS)
						list.Layout(gtx, len(icoBank[navButtonsGroup.Value(gtx)]), func(i int) {
							layout.UniformInset(unit.Dp(16)).Layout(gtx, renderIcon(gtx, th, icoBank[navButtonsGroup.Value(gtx)][icoS[i]], icoS[i]))
						})
					}))
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

func renderIcon(gtx *layout.Context, th *material.Theme, icon *material.Icon, iconLabel string) func() {
	return func() {
		layout.Flex{
			Axis:    layout.Vertical,
			Spacing: layout.SpaceBetween,
		}.Layout(gtx,
			layout.Rigid(func() {

				layout.Flex{
					Spacing: layout.SpaceBetween,
				}.Layout(gtx,
					layout.Rigid(func() {
						icon.Color = color.RGBA{A: 0xff, R: 0xcf, G: 0x30, B: 0x30}
						icon.Layout(gtx, unit.Dp(float32(48)))
					}),
					layout.Rigid(func() {
						th.H2(iconLabel).Layout(gtx)
					}),
				)
			}),
			layout.Rigid(func() {
				DrawRectangle(gtx, gtx.Constraints.Width.Max, 1, color.RGBA{A: 0xff, R: 0xcf, G: 0x30, B: 0x30})
			}))
	}
}

func renderIconFlex(gtx *layout.Context, th *material.Theme, icon *material.Icon, iconLabel string) layout.FlexChild {
	return layout.Rigid(func() {
		layout.Flex{
			Spacing: layout.SpaceBetween,
		}.Layout(gtx,
			layout.Rigid(func() {
				icon.Color = color.RGBA{A: 0xff, R: 0xcf, G: 0x30, B: 0x30}
				icon.Layout(gtx, unit.Dp(float32(32)))
			}),
			layout.Rigid(func() {
				th.H6(iconLabel).Layout(gtx)
			}),
		)
	})
}

func DrawRectangle(gtx *layout.Context, w, h int, color color.RGBA) {
	square := f32.Rectangle{
		Max: f32.Point{
			X: float32(w),
			Y: float32(h),
		},
	}
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	clip.Rect{Rect: square}.Op(gtx.Ops).Add(gtx.Ops)
	paint.PaintOp{Rect: square}.Add(gtx.Ops)
	gtx.Dimensions = layout.Dimensions{Size: image.Point{X: w, Y: h}}
}
