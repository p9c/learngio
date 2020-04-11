package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/helpers"
)

type Data struct {
	Title       string
	Content     string
	ShowContent bool
}

var (
	list = &layout.List{
		Axis: layout.Vertical,
	}
	data = []*Data{
		&Data{
			Title:   "First title",
			Content: "Content of the first item Content of the first item Content of the first item Content of the first item Content of the first item",
		},
		&Data{
			Title:   "Second title",
			Content: "Content of the second item Content of the second item Content of the second item Content of the second item Content of the second item",
		},
		&Data{
			Title:   "Third title",
			Content: "Content of the third item Content of the third item Content of the third item Content of the third item Content of the third item",
		},
		&Data{
			Title:   "Fourth title",
			Content: "Content of the fourth item Content of the fourth item Content of the fourth item Content of the fourth item Content of the fourth item",
		},
		&Data{
			Title:   "Fifth title",
			Content: "Content of the fifth item Content of the fifth item Content of the fifth item Content of the fifth item Content of the fifth item",
		},
	}
)

func main() {
	gofont.Register()
	th := material.NewTheme()
	buttons := createButtons(data)
	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(400), unit.Dp(800)),
			app.Title("ParallelCoin"),
		)
		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)

				list.Layout(gtx, len(data), func(i int) {
					if buttons[data[i].Title].Clicked(gtx) {
						hideContent(data)
						data[i].ShowContent = true
					}
					layout.Flex{
						Axis: layout.Vertical,
					}.Layout(gtx,
						layout.Rigid(func() {
							cs := gtx.Constraints
							helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, helpers.HexARGB("ffcf30cf"), [4]float32{0, 0, 0, 0}, unit.Dp(0))
							th.Button(data[i].Title).Layout(gtx, buttons[data[i].Title])
						}),
						layout.Rigid(func() {
							if data[i].ShowContent != false {
								cs := gtx.Constraints
								helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, helpers.HexARGB("ff3030cf"), [4]float32{0, 0, 0, 0}, unit.Dp(0))
								layout.UniformInset(unit.Dp(16)).Layout(gtx, func() {
									th.Body1(data[i].Content).Layout(gtx)
								})
							}
						}),
					)
				})
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

func createButtons(data []*Data) map[string]*widget.Button {
	buttons := make(map[string]*widget.Button)
	for _, val := range data {
		buttons[val.Title] = new(widget.Button)
	}
	return buttons
}

func hideContent(data []*Data) {
	for _, val := range data {
		val.ShowContent = false
	}
}
