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
	"github.com/gioapp/helpers"
	"github.com/p9c/learngio/examples/icons/controller"
	"github.com/p9c/learngio/examples/icons/model"
	"image"
	"image/color"
	"sort"
)

func main() {
	gofont.Register()
	th := material.NewTheme()

	scr := controller.NewScreen()

	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(400), unit.Dp(800)),
			app.Title("ParallelCoin"),
		)
		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)
				DrawRectangle(gtx, gtx.Constraints.Width.Max, gtx.Constraints.Height.Max, helpers.HexARGB(scr.BgColor))
				layout.Flex{}.Layout(gtx,
					layout.Rigid(func() {
						layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func() {
								th.Body1("Background").Layout(gtx)
							}),
							layout.Rigid(func() {
								layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
									renderColors(gtx, th, scr.BgColorLight, scr, "bg", scr.Colors["light"]),
									renderColors(gtx, th, scr.BgColorYellow, scr, "bg", scr.Colors["yellow"]),
									renderColors(gtx, th, scr.BgColorPurple, scr, "bg", scr.Colors["purple"]),
									renderColors(gtx, th, scr.BgColorRed, scr, "bg", scr.Colors["red"]),
									renderColors(gtx, th, scr.BgColorBlue, scr, "bg", scr.Colors["blue"]),
									renderColors(gtx, th, scr.BgColorDark, scr, "bg", scr.Colors["dark"]),
									renderColors(gtx, th, scr.BgColorOrange, scr, "bg", scr.Colors["orange"]),
									renderColors(gtx, th, scr.BgColorGreen, scr, "bg", scr.Colors["green"]))
							}),
							layout.Rigid(func() {
								th.Body1("Accent").Layout(gtx)
							}),
							layout.Rigid(func() {
								layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
									renderColors(gtx, th, scr.AccentColorLight, scr, "accent", scr.Colors["light"]),
									renderColors(gtx, th, scr.AccentColorYellow, scr, "accent", scr.Colors["yellow"]),
									renderColors(gtx, th, scr.AccentColorPurple, scr, "accent", scr.Colors["purple"]),
									renderColors(gtx, th, scr.AccentColorRed, scr, "accent", scr.Colors["red"]),
									renderColors(gtx, th, scr.AccentColorBlue, scr, "accent", scr.Colors["blue"]),
									renderColors(gtx, th, scr.AccentColorDark, scr, "accent", scr.Colors["dark"]),
									renderColors(gtx, th, scr.AccentColorOrange, scr, "accent", scr.Colors["orange"]),
									renderColors(gtx, th, scr.AccentColorGreen, scr, "accent", scr.Colors["green"]))
							}),
							layout.Rigid(func() {
								DrawRectangle(gtx, 160, gtx.Constraints.Height.Max, helpers.HexARGB(scr.AccentColor))
								scr.NavList.Layout(gtx, len(scr.GroupsIco), func(i int) {
									layout.UniformInset(unit.Dp(16)).Layout(gtx, func() {
										th.RadioButton(scr.GroupsIco[i], scr.GroupsIco[i]).Layout(gtx, scr.NavButtonsGroup)
									})
								})
							}))
					}),
					layout.Flexed(1, func() {
						icoS := make([]string, 0, len(scr.IcoBank[scr.NavButtonsGroup.Value(gtx)]))
						for k := range scr.IcoBank[scr.NavButtonsGroup.Value(gtx)] {
							icoS = append(icoS, k)
						}
						sort.Strings(icoS)
						scr.List.Layout(gtx, len(scr.IcoBank[scr.NavButtonsGroup.Value(gtx)]), func(i int) {
							layout.UniformInset(unit.Dp(16)).Layout(gtx, renderIcon(gtx, th, scr.IcoBank[scr.NavButtonsGroup.Value(gtx)][icoS[i]], scr, icoS[i]))
						})
					}))
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

func renderIcon(gtx *layout.Context, th *material.Theme, icon *material.Icon, scr *model.Screen, iconLabel string) func() {
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
						icon.Color = helpers.HexARGB(scr.AccentColor)
						icon.Layout(gtx, unit.Dp(float32(scr.IconSize)))
					}),
					layout.Rigid(func() {
						th.Label(th.TextSize.Scale(14.0/float32(scr.TextSize)), iconLabel).Layout(gtx)
					}),
				)
			}),
			layout.Rigid(func() {
				DrawRectangle(gtx, gtx.Constraints.Width.Max, 1, color.RGBA{A: 0xff, R: 0xcf, G: 0x30, B: 0x30})
			}))
	}
}

func renderColors(gtx *layout.Context, th *material.Theme, controllerButton *widget.Button, scr *model.Screen, part, color string) layout.FlexChild {
	return layout.Rigid(func() {
		layout.Flex{
			Spacing: layout.SpaceBetween,
		}.Layout(gtx,
			layout.Rigid(func() {
				var linkButton material.Button
				linkButton = th.Button("")
				linkButton.Background = helpers.HexARGB(color)
				for controllerButton.Clicked(gtx) {
					switch part {
					case "bg":
						scr.BgColor = color
					case "accent":
						scr.AccentColor = color
					}
				}
				linkButton.Layout(gtx, controllerButton)
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
