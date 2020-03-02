package controller

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/p9c/learngio/examples/html/native/model"
	"io/ioutil"
	"net/http"
)

//RenderDocument "Renders the DOM to the screen"
func RenderDocument(gtx *layout.Context, th *material.Theme, document *model.NodeDOM) {
	html := document.Children[0]
	//runtime.LockOSThread()
	//glfw.Init()
	//widgets := []func(){}
	//
	//for _, dom := range document.Children {
	//		widgets = append(widgets, func() {
	//			ch := th.Caption(fmt.Sprint(dom.Element))
	//			ch.Layout(gtx)
	//		})
	//}
	//
	//list.Layout(gtx, len(widgets), func(i int) {
	//	layout.UniformInset(unit.Dp(16)).Layout(gtx, widgets[i])
	//})

	//browserWindow := createBrowserWindow(gtx, th, html)
	//attachBrowserWindowEvents(&browserWindow)
	//browserWindow.Viewport = canvas.New(browserWindow.ViewportBackend)
	//browserWindow.Addressbar = canvas.New(browserWindow.AddressbarBackend)
	browserWindowMainLoop(gtx, th, html)
}

// GetResource - Makes an http request and returns a resource struct
func GetResource(url string) *model.Resource {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return &model.Resource{
		Body: string(body),
	}
}
