package main

import (
	"container/list"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"golang.org/x/exp/shiny/materialdesign/icons"
	"image"
)

var (
	widgetButtonUp   = new(widget.Button)
	widgetButtonDown = new(widget.Button)
)

type Slider struct {
	ColorBg      string
	BorderRadius [4]float32
	OperateValue interface{}
	Height       float32
	body         *SliderBody
	up           *SliderButton
	down         *SliderButton
}

type SliderBody struct {
	pressed      bool
	Do           func(interface{})
	ColorBg      string
	Position     float32
	Cursor       float32
	OperateValue interface{}
	Height       float32
	Icon         material.Icon
}

type SliderButton struct {
	icon        *material.Icon
	button      material.IconButton
	Height      float32
	iconColor   string
	iconBgColor string
	insetTop    float32
	insetRight  float32
	insetBottom float32
	insetLeft   float32
	iconSize    float32
	iconPadding float32
}

func (p *Panel) slider(th *material.Theme) *Slider {
	iconGrab, _ := material.NewIcon(icons.NavigationMenu)
	iconUp, _ := material.NewIcon(icons.NavigationArrowDropUp)
	iconDown, _ := material.NewIcon(icons.NavigationArrowDropDown)
	//itemValue := item{
	//	i: 0,
	//}
	up := &SliderButton{
		icon:        iconUp,
		button:      th.IconButton(iconUp),
		Height:      16,
		iconColor:   "ff445588",
		iconBgColor: "ffff882266",
		insetTop:    0,
		insetRight:  0,
		insetBottom: 0,
		insetLeft:   0,
		iconSize:    32,
		iconPadding: 0,
	}
	down := &SliderButton{
		icon:        iconUp,
		button:      th.IconButton(iconDown),
		Height:      16,
		iconSize:    16,
		iconColor:   "ff445588",
		iconBgColor: "ffff882266",
	}
	body := &SliderBody{
		pressed: false,
		ColorBg: "",
		//Position: 0,
		//Cursor:   0,
		Icon: *iconGrab,
		//Do: func(n interface{}) {
		//	itemValue.doSlide(n.(int))
		//},
		OperateValue: 1,
		Height:       30,
	}
	return &Slider{
		ColorBg:      "ff885566",
		BorderRadius: [4]float32{},
		OperateValue: 1,
		//ListPosition: 0,
		Height: 16,
		body:   body,
		up:     up,
		down:   down,
	}
}
func (s *SliderButton) sliderButton() *material.IconButton {
	button := s.button
	button.Inset.Top = unit.Dp(0)
	button.Inset.Bottom = unit.Dp(0)
	button.Inset.Right = unit.Dp(0)
	button.Inset.Left = unit.Dp(0)
	button.Size = unit.Dp(32)
	button.Padding = unit.Dp(0)
	return &button
}
func (s *Slider) Layout(gtx *layout.Context, list *list.List, listLenght int) {
	layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceBetween,
	}.Layout(gtx,
		layout.Rigid(func() {
			for widgetButtonUp.Clicked(gtx) {
				if list.Position.First > 0 {
					list.Position.First = list.Position.First - 1
				}
			}

			s.up.sliderButton().Layout(gtx, widgetButtonUp)
		}),
		layout.Flexed(1, func() {
			s.body.Layout(gtx, list, listLenght)
		}),
		layout.Rigid(func() {
			for widgetButtonDown.Clicked(gtx) {
				if list.Position.First < listLenght {
					list.Position.First = list.Position.First + 1
				}
			}
			s.down.sliderButton().Layout(gtx, widgetButtonDown)
		}),
	)
}

func (s *SliderBody) Layout(gtx *layout.Context, list *layout.List, ll int) {
	for _, e := range gtx.Events(s) {
		if e, ok := e.(pointer.Event); ok {
			s.Position = e.Position.Y
			switch e.Type {
			case pointer.Press:
				s.pressed = true
				s.Do(s.OperateValue)
				list.Position.First = int(s.Position)
			case pointer.Release:
				s.pressed = false
			}
		}
	}

	cs := gtx.Constraints
	//th := material.NewTheme()

	sliderBg := HexARGB("ff558899")
	colorBg := HexARGB("ff30cfcf")
	colorBorder := HexARGB("ffcf3030")
	border := unit.Dp(0)

	//listUnit := float32(cs.Height.Max) / float32(ll)
	//fmt.Println(listUnit)
	if s.pressed {
		if s.Position >= 0 && s.Position <= (float32(cs.Height.Max)-s.Height) {
			s.Cursor = s.Position
			//s.Cursor = s.Position * float32(list.Position.First)
			list.Position.First = int(s.Position)
		}
		colorBg = HexARGB("ffcf30cf")
		colorBorder = HexARGB("ff303030")
		border = unit.Dp(0)
	}
	pointer.Rect(
		image.Rectangle{Max: image.Point{X: cs.Width.Max, Y: cs.Height.Max}},
	).Add(gtx.Ops)
	pointer.InputOp{Key: s}.Add(gtx.Ops)
	DrawRectangle(gtx, float32(cs.Width.Max), float32(cs.Height.Max), colorBorder, [4]float32{0, 0, 0, 0}, unit.Dp(0))

	layout.UniformInset(border).Layout(gtx, func() {
		cs := gtx.Constraints
		DrawRectangle(gtx, float32(cs.Width.Max), float32(cs.Height.Max), colorBg, [4]float32{0, 0, 0, 0}, unit.Dp(0))
		//cs := gtx.Constraints
		layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func() {
				layout.Center.Layout(gtx, func() {

					layout.Inset{
						Top: unit.Dp(s.Cursor),
					}.Layout(gtx, func() {
						//cs := gtx.Constraints

						DrawRectangle(gtx, float32(30), s.Height, sliderBg, [4]float32{5, 5, 5, 5}, unit.Dp(0))

						s.Icon.Color = HexARGB("ff554499")
						s.Icon.Layout(gtx, unit.Px(float32(32)))
					})

				})
			}),
		)
	})
}
