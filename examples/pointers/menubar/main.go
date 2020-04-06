package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/p9c/learngio/helpers"
)

var (
	listMenu = &layout.List{
		Axis: layout.Horizontal,
	}
	listFile = &layout.List{
		Axis: layout.Vertical,
	}
	listEdit = &layout.List{
		Axis: layout.Vertical,
	}
	listHelp = &layout.List{
		Axis: layout.Vertical,
	}
	fileButton = button("File")
	editButton = button("Edit")
	helpButton = button("Help")

	newButton  = button("New")
	openButton = button("Open")
	saveButton = button("Save")

	copyButton  = button("Copy")
	cutButton   = button("Cut")
	pasteButton = button("Paste")

	aboutButton   = button("About")
	subMenuSelect string
)

func main() {

	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(900), unit.Dp(556)),
			app.Title("Learn Gio"),
		)

		gofont.Register()

		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)
				layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func() {
						cs := gtx.Constraints
						helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, helpers.HexARGB("ff303030"), [4]float32{0, 0, 0, 0}, unit.Dp(0))
						in := layout.UniformInset(unit.Dp(0))
						in.Layout(gtx, func() {
							listMenu.Layout(gtx, len(menuItems(gtx)), func(i int) {
								layout.UniformInset(unit.Dp(0)).Layout(gtx, menuItems(gtx)[i])
							})
						})
					}),
					layout.Flexed(1, func() {
						cs := gtx.Constraints
						helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, helpers.HexARGB("ffcfcfcf"), [4]float32{0, 0, 0, 0}, unit.Dp(0))
						listFile.Layout(gtx, len(subMenuItems(gtx, subMenuSelect)), func(i int) {
							layout.UniformInset(unit.Dp(0)).Layout(gtx, subMenuItems(gtx, subMenuSelect)[i])
						})

					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

func subMenuItems(gtx *layout.Context, subMenu string) []func() {
	var sub []func()
	switch subMenu {
	case "file":
		sub = fileMenuItems(gtx)
	case "edit":
		sub = editMenuItems(gtx)
	case "help":
		sub = helpMenuItems(gtx)
	}
	return sub
}

func menuItems(gtx *layout.Context) []func() {
	return []func(){
		fileButton.Layout(gtx, func() {
			subMenuSelect = "file"
		}),
		editButton.Layout(gtx, func() {
			subMenuSelect = "edit"
		}),
		helpButton.Layout(gtx, func() {
			subMenuSelect = "help"
		}),
	}
}

func fileMenuItems(gtx *layout.Context) []func() {
	return []func(){
		newButton.Layout(gtx, func() {}),
		openButton.Layout(gtx, func() {}),
		saveButton.Layout(gtx, func() {}),
	}
}

func editMenuItems(gtx *layout.Context) []func() {
	return []func(){
		copyButton.Layout(gtx, func() {}),
		cutButton.Layout(gtx, func() {}),
		pasteButton.Layout(gtx, func() {}),
	}
}

func helpMenuItems(gtx *layout.Context) []func() {
	return []func(){
		helpButton.Layout(gtx, func() {}),
		aboutButton.Layout(gtx, func() {}),
	}
}
