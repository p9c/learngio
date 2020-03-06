package main

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type item struct {
	i int
}

func (it *item) doSlide(n int) {
	it.i = it.i + n
}

type Panel struct {
	Name              string
	totalHeight       int
	visibleHeight     int
	totalOffset       int
	panelContent      *layout.List
	panelObject       []func()
	panelObjectHeight int
	scrollBar         *ScrollBar
	scrollUnit        float32
}

func onePanel(th *material.Theme) *Panel {
	return &Panel{
		Name: "OnePanel",
		panelContent: &layout.List{
			Axis:        layout.Vertical,
			ScrollToEnd: false,
		},
		scrollBar: scrollBar(th),
	}
}

func (p *Panel) Layout(gtx *layout.Context, th *material.Theme) {
	layout.Flex{
		Axis:    layout.Horizontal,
		Spacing: layout.SpaceBetween,
	}.Layout(gtx,
		layout.Flexed(1, p.panelLayout(gtx, th, testContent(gtx, th))),
		layout.Rigid(func() {
			if p.totalOffset > 0 {
				p.SliderLayout(gtx)
			}
		}),
	)

	p.totalOffset = p.totalHeight - p.visibleHeight

	p.scrollUnit = float32(p.scrollBar.body.Height) / float32(p.totalHeight)
	p.scrollBar.body.CursorHeight = p.scrollUnit * float32(p.visibleHeight)
	p.scrollBar.body.Cursor = float32(p.panelContent.Position.Offset) * p.scrollUnit
	//p.scrollBar.body.Cursor = float32(p.panelContent.Position.Offset)
	fmt.Println("bodyHeight:", p.scrollBar.body.Height)

	fmt.Println("totalOffset:", p.totalOffset)
	fmt.Println("scrollUnit:", p.scrollUnit)

	fmt.Println("cursor:", p.scrollBar.body.Cursor)
	fmt.Println("visibleHeight:", p.visibleHeight)

	fmt.Println("total:", p.totalHeight)
	fmt.Println("offset:", p.panelContent.Position.Offset)
}

//func (p *Panel) drawList(gtx *layout.Context) {
//	p.listLayout.Layout(gtx, len(p.listObjects), func(i int) {
//		p.listObjects[i]()
//	})
//}

func (p *Panel) calculateListHeight(gtx *layout.Context) {

	//for _, listItem := range p.listObjects{
	//	//p.listHeight = p.listHeight + listItem()
	//}

}
