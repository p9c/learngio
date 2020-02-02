package main

import (
	"fmt"
	"github.com/p9c/learngio/examples/_examples/testComponentsWallet/apps"
	"image"
	"image/color"
	"io/ioutil"
	"os"
	"runtime/pprof"
	"runtime/trace"
	"time"

	"github.com/gop9/olt/framework"
	"github.com/gop9/olt/framework/font"
	"github.com/gop9/olt/framework/label"
	"github.com/gop9/olt/framework/rect"
	nstyle "github.com/gop9/olt/framework/style"
	gkey "github.com/gop9/olt/gio/io/key"
)

const dotrace = false
const scaling = 1.62

//var theme framework.Theme = framework.WhiteTheme
var theme nstyle.Theme = nstyle.DarkTheme

func id(fn func(*framework.Window)) func() func(*framework.Window) {
	return func() func(*framework.Window) {
		return fn
	}
}

type Demo struct {
	Name     string
	Title    string
	Flags    framework.WindowFlags
	UpdateFn func() func(*framework.Window)
}

var demos = []Demo{
	{"calc", "Calculator Demo", 0, func() func(*framework.Window) {
		var cd apps.CalcDemo
		cd.Current = &cd.A
		return cd.CalculatorDemo
	}},
}

var Wnd framework.MasterWindow

func main() {

	if dotrace {
		fh, _ := os.Create("demo.trace.out")
		if fh != nil {
			defer fh.Close()
			trace.Start(fh)
			defer trace.Stop()
		}
		f, _ := os.Create("demo.cpu.pprof")
		if f != nil {
			pprof.StartCPUProfile(f)
			defer pprof.StopCPUProfile()
		}
	}

	whichdemo := ""
	if len(os.Args) > 1 {
		whichdemo = os.Args[1]
	}

	switch whichdemo {
	case "multi", "":
		Wnd = framework.NewMasterWindow(0, "Multiwindow Demo", func(w *framework.Window) {})
		Wnd.PopupOpen("Multiwindow Demo", framework.WindowTitle|framework.WindowBorder|framework.WindowMovable|framework.WindowScalable|framework.WindowNonmodal, rect.Rect{0, 0, 400, 300}, true, multiDemo)
	default:
		for i := range demos {
			if demos[i].Name == whichdemo {
				Wnd = framework.NewMasterWindow(demos[i].Flags, demos[i].Title, demos[i].UpdateFn())
				break
			}
		}
	}
	if Wnd == nil {
		fmt.Fprintf(os.Stderr, "unknown demo %q\n", whichdemo)
		fmt.Fprintf(os.Stderr, "known demos:\n")
		for i := range demos {
			fmt.Fprintf(os.Stderr, "\t%s\n", demos[i].Name)
		}
		os.Exit(1)
	}

	Wnd.SetStyle(nstyle.FromTheme(theme, scaling))

	normalFontData, normerr := ioutil.ReadFile("demofont.ttf")
	if normerr == nil {
		szf := 12 * scaling
		face, normerr := font.NewFace(normalFontData, int(szf))
		if normerr == nil {
			style := Wnd.Style()
			style.Font = face
		}
	}

	Wnd.Main()
	if dotrace {
		fh, _ := os.Create("demo.heap.pprof")
		if fh != nil {
			defer fh.Close()
			pprof.WriteHeapProfile(fh)
		}
	}
}

func buttonDemo(w *framework.Window) {
	w.Row(20).Static(60, 60)
	if w.Button(label.T("button1"), false) {
		fmt.Printf("button pressed!\n")
	}
	if w.Button(label.T("button2"), false) {
		fmt.Printf("button 2 pressed!\n")
	}
}

type difficulty int

const (
	easy = difficulty(iota)
	hard
)

var op difficulty = easy
var compression int

func basicDemo(w *framework.Window) {
	w.Row(30).Dynamic(1)
	w.Label(time.Now().Format("15:04:05"), "RT")

	w.Row(30).Static(80)
	if w.Button(label.T("button"), false) {
		fmt.Printf("button pressed! difficulty: %v compression: %d\n", op, compression)
	}
	w.Row(30).Dynamic(2)
	if w.OptionText("easy", op == easy) {
		op = easy
	}
	if w.OptionText("hard", op == hard) {
		op = hard
	}
	w.Row(25).Dynamic(1)
	w.PropertyInt("Compression:", 0, &compression, 100, 10, 1)
}

func textEditorDemo() func(w *framework.Window) {
	var textEditorEditor framework.TextEditor
	textEditorEditor.Flags = framework.EditSelectable
	textEditorEditor.Buffer = []rune("prova")
	return func(w *framework.Window) {
		w.Row(30).Dynamic(1)
		textEditorEditor.Maxlen = 30
		textEditorEditor.Edit(w)
	}
}

func multilineTextEditorDemo() func(w *framework.Window) {
	var multilineTextEditor framework.TextEditor
	bs, _ := ioutil.ReadFile("overview.go")
	multilineTextEditor.Buffer = []rune(string(bs))
	var ibeam bool
	return func(w *framework.Window) {
		w.Row(20).Dynamic(1)
		w.CheckboxText("I-Beam cursor", &ibeam)
		w.Row(0).Dynamic(1)
		multilineTextEditor.Flags = framework.EditMultiline | framework.EditSelectable | framework.EditClipboard
		if ibeam {
			multilineTextEditor.Flags |= framework.EditIbeamCursor
		}
		multilineTextEditor.Edit(w)
	}
}

type panelDebug struct {
	splitv          framework.ScalableSplit
	splith          framework.ScalableSplit
	showblocks      bool
	showsingleblock bool
	showtabs        bool
}

func (pd *panelDebug) Init() {
	pd.splitv.MinSize = 80
	pd.splitv.Size = 120
	pd.splitv.Spacing = 5
	pd.splith.MinSize = 100
	pd.splith.Size = 300
	pd.splith.Spacing = 5
	pd.showtabs = true
	pd.showsingleblock = true
}

func (pd *panelDebug) Update(w *framework.Window) {
	for _, k := range w.Input().Keyboard.Keys {
		if k.Rune == 'b' {
			pd.showsingleblock = false
			pd.showblocks = !pd.showblocks
		}
		if k.Rune == 'B' {
			pd.showsingleblock = !pd.showsingleblock
		}
		if k.Rune == 't' {
			pd.showtabs = !pd.showtabs
		}
	}

	if pd.showtabs {
		w.Row(20).Dynamic(2)
		w.Label("A", "LC")
		w.Label("B", "LC")
	}

	area := w.Row(0).SpaceBegin(0)

	if pd.showsingleblock {
		w.LayoutSpacePushScaled(area)
		bounds, out := w.Custom(nstyle.WidgetStateInactive)
		if out != nil {
			out.FillRect(bounds, 10, color.RGBA{0x00, 0x00, 0xff, 0xff})
		}
	} else {
		leftbounds, rightbounds := pd.splitv.Vertical(w, area)
		viewbounds, commitbounds := pd.splith.Horizontal(w, rightbounds)

		w.LayoutSpacePushScaled(leftbounds)
		pd.groupOrBlock(w, "index-files", framework.WindowBorder)

		w.LayoutSpacePushScaled(viewbounds)
		pd.groupOrBlock(w, "index-diff", framework.WindowBorder)

		w.LayoutSpacePushScaled(commitbounds)
		pd.groupOrBlock(w, "index-right-column", framework.WindowNoScrollbar|framework.WindowBorder)
	}
}

func (pd *panelDebug) groupOrBlock(w *framework.Window, name string, flags framework.WindowFlags) {
	if pd.showblocks {
		bounds, out := w.Custom(nstyle.WidgetStateInactive)
		if out != nil {
			out.FillRect(bounds, 10, color.RGBA{0x00, 0x00, 0xff, 0xff})
		}
	} else {
		if sw := w.GroupBegin(name, flags); sw != nil {
			sw.GroupEnd()
		}
	}
}

func nestedMenu(w *framework.Window) {
	w.Row(20).Static(180)
	w.Label("Test", "CC")
	w.ContextualOpen(0, image.Point{0, 0}, w.LastWidgetBounds, func(w *framework.Window) {
		w.Row(20).Dynamic(1)
		if w.MenuItem(label.TA("Submenu", "CC")) {
			w.ContextualOpen(0, image.Point{0, 0}, rect.Rect{0, 0, 0, 0}, func(w *framework.Window) {
				w.Row(20).Dynamic(1)
				if w.MenuItem(label.TA("Done", "CC")) {
					fmt.Printf("done\n")
				}
			})
		}
	})
}

var listDemoSelected = -1
var listDemoCnt = 0

func listDemo(w *framework.Window) {
	const N = 100
	recenter := false
	for _, e := range w.Input().Keyboard.Keys {
		switch e.Code {
		case gkey.CodeDownArrow:
			listDemoSelected++
			if listDemoSelected >= N {
				listDemoSelected = N - 1
			}
			recenter = true
		case gkey.CodeUpArrow:
			listDemoSelected--
			if listDemoSelected < -1 {
				listDemoSelected = -1
			}
			recenter = true
		}
	}
	w.Row(0).Dynamic(1)
	if gl, w := framework.GroupListStart(w, N, "list", framework.WindowNoHScrollbar); w != nil {
		if !recenter {
			gl.SkipToVisible(20)
		}
		w.Row(20).Dynamic(1)
		cnt := 0
		for gl.Next() {
			cnt++
			i := gl.Index()
			selected := i == listDemoSelected
			w.SelectableLabel(fmt.Sprintf("label %d", i), "LC", &selected)
			if selected {
				listDemoSelected = i
				if recenter {
					gl.Center()
				}
			}
		}
		if cnt != listDemoCnt {
			listDemoCnt = cnt
			fmt.Printf("called %d times\n", listDemoCnt)
		}
	}
}

func keybindings(w *framework.Window) {
	mw := w.Master()
	if in := w.Input(); in != nil {
		k := in.Keyboard
		for _, e := range k.Keys {
			scaling := mw.Style().Scaling
			switch {
			case (e.Modifiers == gkey.ModControl || e.Modifiers == gkey.ModControl|gkey.ModShift) && (e.Code == gkey.CodeEqualSign):
				mw.Style().Scale(scaling + 0.1)
			case (e.Modifiers == gkey.ModControl || e.Modifiers == gkey.ModControl|gkey.ModShift) && (e.Code == gkey.CodeHyphenMinus):
				mw.Style().Scale(scaling - 0.1)
			case (e.Modifiers == gkey.ModControl) && (e.Code == gkey.CodeF):
				mw.SetPerf(!mw.GetPerf())
			}
		}
	}
}

func multiDemo(w *framework.Window) {
	keybindings(w)

	w.Row(20).Dynamic(1)
	w.Label("Welcome to the multi-window demo.", "LC")
	w.Label("Open any demo window by clicking on the buttons.", "LC")
	w.Label("To run a demo as a stand-alone window use:", "LC")
	w.Label("     \"./demo <demo-name>\"", "LC")
	w.Row(30).Static(100, 100, 100)
	for i := range demos {
		if w.ButtonText(demos[i].Name) {
			w.Master().PopupOpen(demos[i].Title, framework.WindowDefaultFlags|framework.WindowNonmodal|demos[i].Flags, rect.Rect{0, 0, 200, 200}, true, demos[i].UpdateFn())
		}
	}
}
