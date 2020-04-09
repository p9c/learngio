// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"gioui.org/io/pointer"
	"github.com/gioapp/helpers"
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
)

type checkable struct {
	Label string
	Color color.RGBA
	Size  unit.Value
}

func (c *checkable) layout(gtx *layout.Context, checked bool) {
	var state layout.Direction

	if checked {
		state = layout.W
	} else {
		state = layout.E
	}

	layout.Center.Layout(gtx, func() {
		gtx.Constraints.Width.Min = 64
		gtx.Constraints.Height.Min = 32
		helpers.DrawRectangle(gtx, 64, 32, helpers.HexARGB("ff888888"), [4]float32{0, 0, 0, 0}, unit.Dp(0))
		layout.Center.Layout(gtx, func() {
			helpers.DrawRectangle(gtx, 48, 16, helpers.HexARGB("ffcf30cf"), [4]float32{8, 8, 8, 8}, unit.Dp(0))
		})

		state.Layout(gtx, func() {
			helpers.DrawRectangle(gtx, 24, 24, helpers.HexARGB("ffcf3030"), [4]float32{12, 12, 12, 12}, unit.Dp(0))
		})

		pointer.Rect(image.Rectangle{Max: gtx.Dimensions.Size}).Add(gtx.Ops)
	})
}
