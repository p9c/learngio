package main

import (
	"os"
	"runtime/pprof"
	"runtime/trace"

	"github.com/gop9/olt/framework"
	nstyle "github.com/gop9/olt/framework/style"
)

const dotrace = false
const scaling = 1.6

var Wnd framework.MasterWindow

//var theme framework.Theme = framework.WhiteTheme
var theme nstyle.Theme = nstyle.DarkTheme

func main() {
	if dotrace {
		fh, _ := os.Create("overview.trace.out")
		if fh != nil {
			defer fh.Close()
			trace.Start(fh)
			defer trace.Stop()
		}
		f, _ := os.Create("overview.cpu.pprof")
		if f != nil {
			pprof.StartCPUProfile(f)
			defer pprof.StopCPUProfile()
		}
	}

	od := newOverviewDemo()
	od.Theme = theme

	Wnd = framework.NewMasterWindow(0, "Overview", od.overviewDemo)
	Wnd.SetStyle(nstyle.FromTheme(theme, scaling))
	Wnd.Main()
	if dotrace {
		fh, _ := os.Create("demo.heap.pprof")
		if fh != nil {
			defer fh.Close()
			pprof.WriteHeapProfile(fh)
		}
	}
}
