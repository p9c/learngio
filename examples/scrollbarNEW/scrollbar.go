package main

import (
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

type ScrollBar struct {
	ColorBg      string
	BorderRadius [4]float32
	OperateValue interface{}
	//Height       float32
	body *ScrollBarBody
	up   *ScrollBarButton
	down *ScrollBarButton
}

type ScrollBarBody struct {
	pressed      bool
	Do           func(interface{})
	ColorBg      string
	Position     float32
	Cursor       float32
	OperateValue interface{}
	Height       int
	CursorHeight float32
	Icon         material.Icon
}

type ScrollBarButton struct {
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

func scrollBar(th *material.Theme) *ScrollBar {
	iconGrab, _ := material.NewIcon(icons.NavigationMenu)
	iconUp, _ := material.NewIcon(icons.NavigationArrowDropUp)
	iconDown, _ := material.NewIcon(icons.NavigationArrowDropDown)
	itemValue := item{
		i: 0,
	}
	up := &ScrollBarButton{
		icon:        iconUp,
		button:      th.IconButton(iconUp),
		Height:      16,
		iconColor:   "ff445588",
		iconBgColor: "ff882266",
		insetTop:    0,
		insetRight:  0,
		insetBottom: 0,
		insetLeft:   0,
		iconSize:    32,
		iconPadding: 0,
	}
	down := &ScrollBarButton{
		icon:        iconUp,
		button:      th.IconButton(iconDown),
		Height:      16,
		iconSize:    16,
		iconColor:   "ff445588",
		iconBgColor: "ff882266",
	}
	body := &ScrollBarBody{
		pressed:  false,
		ColorBg:  "",
		Position: 0,
		Cursor:   0,
		Icon:     *iconGrab,
		Do: func(n interface{}) {
			itemValue.doSlide(n.(int))
		},
		OperateValue: 1,
		CursorHeight: 30,
	}
	return &ScrollBar{
		ColorBg:      "ff885566",
		BorderRadius: [4]float32{},
		OperateValue: 1,
		//ListPosition: 0,
		//Height: 16,
		body: body,
		up:   up,
		down: down,
	}
}
func (s *ScrollBarButton) scrollBarButton() *material.IconButton {
	button := s.button
	button.Inset.Top = unit.Dp(0)
	button.Inset.Bottom = unit.Dp(0)
	button.Inset.Right = unit.Dp(0)
	button.Inset.Left = unit.Dp(0)
	button.Size = unit.Dp(32)
	button.Padding = unit.Dp(0)
	return &button
}
func (p *Panel) SliderLayout(gtx *layout.Context) {
	layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func() {
			for widgetButtonUp.Clicked(gtx) {
				p.panelContent.Position.Offset = p.panelContent.Position.Offset - int(p.scrollBar.body.CursorHeight)
			}
			p.scrollBar.up.scrollBarButton().Layout(gtx, widgetButtonUp)
		}),
		layout.Flexed(1, func() {
			p.bodyLayout(gtx)
		}),
		layout.Rigid(func() {
			for widgetButtonDown.Clicked(gtx) {
				p.panelContent.Position.Offset = p.panelContent.Position.Offset + int(p.scrollBar.body.CursorHeight)
			}
			p.scrollBar.down.scrollBarButton().Layout(gtx, widgetButtonDown)
		}),
	)
}

func (p *Panel) bodyLayout(gtx *layout.Context) {
	for _, e := range gtx.Events(p.scrollBar.body) {
		if e, ok := e.(pointer.Event); ok {
			p.scrollBar.body.Position = e.Position.Y - (p.scrollBar.body.CursorHeight / 2)
			switch e.Type {
			case pointer.Press:
				p.scrollBar.body.pressed = true
				p.scrollBar.body.Do(p.scrollBar.body.OperateValue)
				//list.Position.First = int(s.Position)
			case pointer.Release:
				p.scrollBar.body.pressed = false
			}
		}
	}
	cs := gtx.Constraints
	p.scrollBar.body.Height = cs.Height.Max
	sliderBg := HexARGB("ff558899")
	colorBg := HexARGB("ff30cfcf")
	colorBorder := HexARGB("ffcf3030")
	border := unit.Dp(0)
	if p.scrollBar.body.pressed {
		if p.scrollBar.body.Position >= 0 && p.scrollBar.body.Position <= (float32(cs.Height.Max)-p.scrollBar.body.CursorHeight) {
			p.scrollBar.body.Cursor = p.scrollBar.body.Position
			p.panelContent.Position.Offset = int(p.scrollBar.body.Cursor / p.scrollUnit)
		}
		colorBg = HexARGB("ffcf30cf")
		colorBorder = HexARGB("ff303030")
		border = unit.Dp(0)
	}
	pointer.Rect(
		image.Rectangle{Max: image.Point{X: cs.Width.Max, Y: cs.Height.Max}},
	).Add(gtx.Ops)
	pointer.InputOp{Key: p.scrollBar.body}.Add(gtx.Ops)
	DrawRectangle(gtx, float32(cs.Width.Max), float32(cs.Height.Max), colorBorder, [4]float32{0, 0, 0, 0}, unit.Dp(0))
	layout.UniformInset(border).Layout(gtx, func() {
		cs := gtx.Constraints
		DrawRectangle(gtx, float32(cs.Width.Max), float32(cs.Height.Max), colorBg, [4]float32{0, 0, 0, 0}, unit.Dp(0))
		//cs := gtx.Constraints
		layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func() {
				layout.Center.Layout(gtx, func() {
					layout.Inset{
						Top: unit.Dp(p.scrollBar.body.Cursor),
					}.Layout(gtx, func() {
						//cs := gtx.Constraints
						DrawRectangle(gtx, float32(30), p.scrollBar.body.CursorHeight, sliderBg, [4]float32{5, 5, 5, 5}, unit.Dp(0))
						layout.Center.Layout(gtx, func() {
							p.scrollBar.body.Icon.Color = HexARGB("ff554499")
							p.scrollBar.body.Icon.Layout(gtx, unit.Px(float32(32)))
						})
					})
				})
			}),
		)
	})
}
