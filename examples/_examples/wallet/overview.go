package main

import (
	"github.com/gop9/olt/framework"
	"github.com/gop9/olt/framework/label"
	"github.com/gop9/olt/framework/rect"
	nstyle "github.com/gop9/olt/framework/style"
	"image"
	"image/color"
)

type OptionEnum int

const (
	OptionA = OptionEnum(iota)
	OptionB
	OptionC
)

type overviewDemo struct {
	ShowMenu    bool
	Border      bool
	Resize      bool
	Movable     bool
	NoScrollbar bool
	Minimizable bool
	Close       bool

	HeaderAlign nstyle.HeaderAlign

	// Menu status
	Mprog   int
	Mslider int
	Mcheck  bool
	Prog    int
	Slider  int
	Check   bool

	// Basic widgets status
	IntSlider                          int
	FloatSlider                        float64
	ProgValue                          int
	PropertyFloat                      float64
	PropertyInt                        int
	PropertyNeg                        int
	RangeMin, RangeValue, RangeMax     float64
	RangeIntMin, RangeInt, RangeIntMax int
	Checkbox                           bool
	Option                             OptionEnum

	// Selectable
	Selected  []bool
	Selected2 []bool

	// Combo
	CurrentWeapon              int
	CheckValues                []bool
	Position                   []float64
	ComboColor                 color.RGBA
	ProgA, ProgB, ProgC, ProgD int
	Weapons                    []string
	ColMode                    int
	TimeSelected               int
	Text0Editor, Text1Editor   framework.TextEditor
	FieldEditor                framework.TextEditor
	PasswordEditor             framework.TextEditor
	BoxEditor                  framework.TextEditor

	// Popup
	PSelect []bool
	PProg   int
	PSlider int

	// Layout
	GroupBorder             bool
	GroupNoScrollbar        bool
	GroupWidth, GroupHeight int

	GroupSelected []bool

	// Vertical Split
	A, B, C    int
	HA, HB, HC int

	Img *image.RGBA

	Resizing1, Resizing2, Resizing3, Resizing4 bool

	edEntry1, edEntry2, edEntry3 framework.TextEditor

	Theme nstyle.Theme

	FitWidthEditor framework.TextEditor
	FitWidthIdx    int
}

func newOverviewDemo() (od *overviewDemo) {
	od = &overviewDemo{}
	od.ShowMenu = true
	od.Border = true
	od.Resize = true
	od.Movable = true
	od.NoScrollbar = false
	od.Close = true
	od.HeaderAlign = nstyle.HeaderRight
	od.Mprog = 60
	od.Mslider = 8
	od.Mcheck = true
	od.Prog = 40
	od.Slider = 10
	od.Check = true
	od.IntSlider = 5
	od.FloatSlider = 2.5
	od.ProgValue = 40
	od.Option = OptionA
	od.Selected = []bool{false, false, true, false}
	od.Selected2 = []bool{true, false, false, false, false, true, false, false, false, false, true, false, false, false, false, true}
	od.CurrentWeapon = 0
	od.CheckValues = make([]bool, 5)
	od.Position = make([]float64, 3)
	od.ComboColor = color.RGBA{130, 150, 150, 255}
	od.ProgA = 20
	od.ProgB = 40
	od.ProgC = 10
	od.ProgD = 90
	od.Weapons = []string{"Fist", "Pistol", "Shotgun", "Plasma", "BFG"}
	od.PSelect = make([]bool, 4)
	od.PProg = 0
	od.PSlider = 10

	//Layout
	od.GroupBorder = true
	od.GroupNoScrollbar = false

	od.GroupSelected = make([]bool, 16)

	od.A = 100
	od.B = 100
	od.C = 100

	od.HA = 100
	od.HB = 100
	od.HC = 100

	od.PropertyFloat = 2
	od.PropertyInt = 10
	od.PropertyNeg = 10

	od.RangeMin = 0
	od.RangeValue = 50
	od.RangeMax = 100

	od.RangeIntMin = 0
	od.RangeInt = 2048
	od.RangeIntMax = 4096

	od.GroupWidth = 320
	od.GroupHeight = 200

	return od
}

func (od *overviewDemo) overviewDemo(w *framework.Window) {
	w.Row(0).SpaceBegin(0)
	//w.Row(0).Dynamic(1)
	w.LayoutSpacePush(rect.Rect{0, 0, 150, 500})
	if sw := w.GroupBegin("Group Left", framework.WindowBorder); sw != nil {
		sw.Row(18).Static(100)
		for i := range od.GroupSelected {
			txt := "Unselected"
			if od.GroupSelected[i] {
				txt = "Selected"
			}
			sw.SelectableLabel(txt, "CC", &od.GroupSelected[i])
		}
		sw.GroupEnd()
	}

	w.LayoutSpacePush(rect.Rect{160, 0, 150, 240})
	if sw := w.GroupBegin("Group top", framework.WindowBorder); sw != nil {
		sw.Row(25).Dynamic(1)
		sw.Button(label.T("#FFAA"), false)
		sw.Button(label.T("#FFBB"), false)
		sw.Button(label.T("#FFCC"), false)
		sw.Button(label.T("#FFDD"), false)
		sw.Button(label.T("#FFEE"), false)
		sw.Button(label.T("#FFFF"), false)
		sw.GroupEnd()
	}

	w.LayoutSpacePush(rect.Rect{160, 250, 150, 250})
	if sw := w.GroupBegin("Group bottom", framework.WindowBorder); sw != nil {
		sw.Row(25).Dynamic(1)
		sw.Button(label.T("#FFAA"), false)
		sw.Button(label.T("#FFBB"), false)
		sw.Button(label.T("#FFCC"), false)
		sw.Button(label.T("#FFDD"), false)
		sw.Button(label.T("#FFEE"), false)
		sw.Button(label.T("#FFFF"), false)
		sw.GroupEnd()
	}

}
