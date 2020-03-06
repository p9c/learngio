package main

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

//
//func (p *Panel)mainLayout(gtx *layout.Context) []func() {
//	return []func(){
//		func() {
//			p.drawList(gtx)
//		},
//	}
//}

func (p *Panel) panelLayout(gtx *layout.Context, th *material.Theme, c func()) func() {
	return func() {
		content := []func(){
			c,
		}
		p.panelContent.Layout(gtx, len(content), func(i int) {
			content[i]()
			p.totalHeight = gtx.Dimensions.Size.Y
		})
		p.visibleHeight = gtx.Constraints.Height.Max
	}
}

func testContent(gtx *layout.Context, th *material.Theme) func() {
	return func() {
		th.H6("txt more txt and more txt, so more more more txt and more txt, so more txt and more txt, so txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so txt xt more txt and more txt, so more more more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so txt xt more txt and more txt, so more more more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so txt xt more txt and more txt, so more more more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so txt xt more txt and more txt, so more more more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so txt xt more txt and more txt, so more more more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so txt xt more txt and more txt, so more more more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so txt xt more txt and more txt, so more more more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so txt xt more txt and more txt, so more more more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so txt xt more txt and more txt, so more more more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so txt xt more txt and more txt, so more more more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so txt xt more txt and more txt, so more more more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so txt xt more txt and more txt, so more more more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so txt xt more txt and more txt, so more more more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so txt xt more txt and more txt, so more more more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so txt xt more txt and more txt, so more more more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so txt xt more txt and more txt, so more more more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so txt xt more txt and more txt, so more more more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so txt xt more txt and more txt, so more more more txt and mo so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so txt 600000").Layout(gtx)
	}
}

func listObjects1(gtx *layout.Context, th *material.Theme) []func() {
	return []func(){
		func() {
			th.H6("txt more txt and more txt, so more more txt ").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 1").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 2").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 3").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 4").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 5").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so txt 6").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 7").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 8").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 9").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so 10").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 11").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 12").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 13").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 14").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 15").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 16").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 17").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 18").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 19").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 20").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 21").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 22").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 23").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 24").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 25").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 26").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 27").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 28").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 29").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 30").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 31").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 32").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 33").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 34").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more more txt 35").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 36").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 37").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 38").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more  more txt and more txt, so  more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, so more txt and more txt, somore txt 39").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 40").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 41").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 42").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 43").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 44").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 45").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 46").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 47").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 48").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 49").Layout(gtx)
		},
		func() {
			th.H6("txt more txt and more txt, so more more txt 50").Layout(gtx)
		},
	}
}
