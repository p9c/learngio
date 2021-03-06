package main

import (
	"fmt"
	"gioui.org/layout"
)

type item struct {
	i int
}

func (it *item) doSlide(n int) {
	it.i = it.i + n
}

type Panel struct {
	Name                 string
	size                 int
	visibleObjectsNumber int
	//totalOffset       int
	panelContent       *layout.List
	panelObject        []func()
	panelObjectsNumber int
	scrollBar          *ScrollBar
	scrollUnit         int
}

func onePanel() *Panel {
	p := &Panel{
		Name: "OnePanel",
		panelContent: &layout.List{
			Axis:        layout.Vertical,
			ScrollToEnd: false,
		},
		size: 16,
	}
	p.ScrollBar()
	return p
}

func (p *Panel) panelLayout(gtx *layout.Context) func() {
	return func() {
		visibleObjectsNumber := 0
		p.panelContent.Layout(gtx, p.panelObjectsNumber, func(i int) {
			p.panelObject[i]()
			visibleObjectsNumber = visibleObjectsNumber + 1
			p.visibleObjectsNumber = visibleObjectsNumber
			//p.totalHeight = p.totalHeight + content[i]()
		})
		//p.visibleHeight = gtx.Constraints.Height.Max
	}
}

func (p *Panel) Layout(gtx *layout.Context, content []func()) {
	p.panelObject = content
	p.panelObjectsNumber = len(p.panelObject)
	layout.Flex{
		Axis:    layout.Horizontal,
		Spacing: layout.SpaceBetween,
	}.Layout(gtx,
		layout.Flexed(1, p.panelLayout(gtx)),
		layout.Rigid(func() {
			//if p.totalOffset > 0 {
			p.SliderLayout(gtx)
			//}
		}),
	)

	//p.totalOffset = p.totalHeight - p.visibleHeight

	p.scrollUnit = p.scrollBar.body.Height / p.panelObjectsNumber
	cursorHeight := p.visibleObjectsNumber * p.scrollUnit
	//p.scrollBar.body.CursorHeight = p.scrollUnit * float32(p.visibleHeight)
	if cursorHeight > 30 {
		p.scrollBar.body.CursorHeight = cursorHeight
	}
	//p.scrollBar.body.Cursor = float32(p.panelContent.Position.First) * p.scrollUnit
	//p.scrollBar.body.Cursor = float32(p.panelContent.Position.Offset)
	fmt.Println("bodyHeight:", p.scrollBar.body.Height)
	fmt.Println("visibleObjectsNumber:", p.visibleObjectsNumber)

	fmt.Println("scrollBarbodyPosition:", p.scrollBar.body.Position)
	fmt.Println("scrollUnit:", p.scrollUnit)

	fmt.Println("cursor:", p.panelContent.Position.Offset)
	//fmt.Println("Daf:",int(p.scrollBar.body.Position/float32(len(p.panelObject))))

	fmt.Println("First:", p.panelContent.Position.First)
	fmt.Println("BeforeEnd:", p.panelContent.Position.BeforeEnd)
}
