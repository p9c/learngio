package main

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/p9c/learngio/examples/list"
)

type item struct {
	i int
}

func (it *item) doSlide(n int) {
	it.i = it.i + n
}

type Panel struct {
	Name        string
	listLayout  *layout.List
	listObjects []func()
	listHeight  int
}

func onePanel() *Panel {
	return &Panel{
		Name: "OnePanel",
		listLayout: &list.List{
			Axis:        layout.Vertical,
			ScrollToEnd: false,
		},
	}
}

func (p *Panel) Layout(gtx *layout.Context, th *material.Theme) {

	//for _, lo := range p.listObjects{
	//	fmt.Println("Normal:")
	//	fmt.Println(gtx.Dimensions.Size.Y)
	//}

	layout.Flex{
		Axis:    layout.Horizontal,
		Spacing: layout.SpaceBetween,
	}.Layout(gtx,
		layout.Flexed(1, func() {
			p.drawList(gtx)
		}),
		layout.Rigid(func() {
			p.slider(th).Layout(gtx, p.listLayout, len(p.listObjects))
		}),
	)
}

func (p *Panel) drawList(gtx *layout.Context) {
	p.listHeight = 0
	p.listLayout.Layout(gtx, len(p.listObjects), func(i int) {
		p.listObjects[i]()
		p.listHeight = p.listHeight + gtx.Dimensions.Size.Y
		fmt.Println(p.listHeight)
		fmt.Println(gtx.Dimensions.Size.Y)
	})
}

func (p *Panel) calculateListHeight(gtx *layout.Context) {

	//for _, listItem := range p.listObjects{
	//	//p.listHeight = p.listHeight + listItem()
	//}

}
