package main

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/p9c/learngio/helpers"
)

func (t *Table) tableLayout(gtx *layout.Context, th *material.Theme) {

	//tableCellsData := []TableCell{
	//	TableCell{
	//		Id:      0,
	//		Title:   "Id",
	//		Content: func(){
	//			th.Label(th.TextSize.Scale(48.0/16.0), "dsdd").Layout(gtx)
	//		},
	//	},
	//}
	//
	cs := gtx.Constraints
	hmin := cs.Width.Max
	vmin := cs.Height.Max
	list := &layout.List{
		Axis: layout.Vertical,
	}
	layout.Stack{Alignment: layout.N}.Layout(gtx,
		layout.Expanded(func() {
			clip.Rect{
				Rect: f32.Rectangle{Max: f32.Point{
					X: float32(cs.Width.Min),
					Y: float32(cs.Height.Min),
				}},
			}.Op(gtx.Ops).Add(gtx.Ops)
		}),
		layout.Stacked(func() {
			cs.Width.Min = hmin
			cs.Height.Min = vmin
			layout.Center.Layout(gtx, func() {
				layout.Inset{Top: unit.Dp(t.Padding[0]), Bottom: unit.Dp(t.Padding[1]), Left: unit.Dp(t.Padding[2]), Right: unit.Dp(t.Padding[3])}.Layout(gtx, func() {

					layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func() {
							//rr := float32(gtx.Px(unit.Dp(4)))
							clip.Rect{
								Rect: f32.Rectangle{Max: f32.Point{
									X: float32(cs.Width.Min),
									Y: float32(cs.Height.Min),
								}},
								NE: t.Corners[0], NW: t.Corners[1], SE: t.Corners[2], SW: t.Corners[3],
							}.Op(gtx.Ops).Add(gtx.Ops)
							fill(gtx, helpers.HexARGB(t.BgColor))
						}),
						layout.Stacked(func() {
							cs.Width.Min = hmin
							cs.Height.Min = vmin
							layout.Center.Layout(gtx, func() {
								//layout.Inset{Top: unit.Dp(15), Bottom: unit.Dp(15), Left: unit.Dp(15), Right: unit.Dp(15)}.Layout(gtx, func() {

								// Content <<<
								layout.Flex{
									Axis: layout.Vertical,
								}.Layout(gtx,
									layout.Rigid(func() {
										th.Label(th.TextSize.Scale(48.0/16.0), t.Title).Layout(gtx)
									}),
									layout.Flexed(1, func() {

										//  content

										drawList(gtx, list, th)

										//  content

									}),
								)

								// Content >>>

								//})
							})
						}),
					)
				})
			})
		}),
	)
}
