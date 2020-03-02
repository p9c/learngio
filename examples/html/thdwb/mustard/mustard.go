package mustard

import (
	"github.com/tfriedel6/canvas"
	"runtime"

	"github.com/p9c/learngio/examples/html/thdwb/structs"
)

//RenderDocument "Renders the DOM to the screen"
func RenderDocument(document *structs.NodeDOM, url string) {
	html := document.Children[0]
	runtime.LockOSThread()
	glfw.Init()

	browserWindow := createBrowserWindow(html, url)
	attachBrowserWindowEvents(&browserWindow)
	browserWindow.Viewport = canvas.New(browserWindow.ViewportBackend)
	browserWindow.Addressbar = canvas.New(browserWindow.AddressbarBackend)
	browserWindowMainLoop(&browserWindow)
}
