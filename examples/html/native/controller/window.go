package controller

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/davecgh/go-spew/spew"
	"github.com/p9c/learngio/examples/html/native/model"
)

var (
	list = &layout.List{
		Axis: layout.Vertical,
	}
)

func createBrowserWindow(gtx *layout.Context, th *material.Theme, document *model.NodeDOM) model.AppWindow {
	//defaultCursor := glfw.CreateStandardCursor(glfw.ArrowCursor)
	pageTile := getPageTitle(document) + " - thdwb"
	browserWindow := model.AppWindow{
		Width:          900,
		Height:         600,
		Title:          pageTile,
		ViewportOffset: 0,
		DOM:            document,
	}
	return browserWindow
}

func browserWindowMainLoop(gtx *layout.Context, th *material.Theme, document *model.NodeDOM) {
	widgets := []func(){}
	nodeChildren := getNodeChildren(document)
	for i := 0; i < len(nodeChildren); i++ {
		ReflowNode(nodeChildren[i], nodeChildren[i], 0)
		///
		widgets = append(widgets, func() {
			ch := th.Caption(fmt.Sprint(nodeChildren[i-1].Element))
			ch.Layout(gtx)
		})
		spew.Dump(nodeChildren[i])

		//}

		////
	}

	//
	//	browserWindow.Reflow = false
	//}

	//updateViewport(browserWindow)

	//glfw.WaitEvents()
	//}
	list.Layout(gtx, len(widgets), func(i int) {
		layout.UniformInset(unit.Dp(16)).Layout(gtx, widgets[i])
	})

}

func updateAddressBar(browserWindow *model.AppWindow) {
	//oldAddressBarInput := getUIElementByID(browserWindow.UIElements, "addressbarInput")

	if browserWindow.Resize {
		browserWindow.Initialized = false
		browserWindow.UIElements = nil
	}

	if browserWindow.Initialized {
		for i := 0; i < len(browserWindow.UIElements); i++ {
			browserWindow.UIElements[i].Redraw()
		}
	} else {

		//addressbarText := browserWindow.Location
		//if oldAddressBarInput != nil {
		//	addressbarText = oldAddressBarInput.Text
		//}

		//addressbarInput := Input("addressbarInput", w, h, browserWindow.Addressbar, addressbarText)

		//browserWindow.UIElements = append(browserWindow.UIElements, &addressbarBackground, &addressbarInput)
		browserWindow.Initialized = true
	}
}

func updateViewport(browserWindow *model.AppWindow) {
	//viewport := browserWindow.Viewport
	vO := float64(browserWindow.ViewportOffset)

	//w := float64(browserWindow.ViewportWidth)
	//h := float64(browserWindow.ViewportHeight)

	//viewport.SetFillStyle("#FFF")
	//viewport.FillRect(0, 0, w, h)
	renderNode(browserWindow.DOM, browserWindow, vO)
}
