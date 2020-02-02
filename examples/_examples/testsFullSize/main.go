package main

import (
	"fmt"
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/label"
	"github.com/aarzilli/nucular/style"
)

var (
	outerWindowFlags         = nucular.WindowNoScrollbar // also try framework.WindowNoScrollbar
	forceVerticalScrollbar   = false
	forceHorizontalScrollbar = 0
	beginningOfRow           = true
)

func main() {
	wnd := nucular.NewMasterWindow(outerWindowFlags, "Counter", updatefn)
	wnd.SetStyle(style.FromTheme(style.DarkTheme, 1))
	wnd.Main()
}

func updatefn(w *nucular.Window) {
	//w.Row(20).Dynamic(1)
	//if w.TreePush(framework.TreeNode, "Simple", false) {

	w.Row(0).Dynamic(1)
	//if sw := w.GroupBegin("Group Without Border", 0); sw != nil {
	//	sw.Row(18).Static(150)
	//	for i := 0; i < 64; i++ {
	//		sw.Label(fmt.Sprintf("%#02x", i), "LC")
	//	}
	//	sw.GroupEnd()
	//}
	if sw := w.GroupBegin("Group With Border", 0); sw != nil {
		sw.Row(50).Dynamic(2)
		for i := 0; i < 64; i++ {
			sw.Button(label.T(fmt.Sprintf("%08d", (((i%7)*10)^32)+(64+(i%2)*2))), false)
		}
		sw.GroupEnd()
	}
	//w.TreePop()
	//}
}
