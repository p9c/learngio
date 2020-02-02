// SPDX-License-Identifier: Unlicense OR MIT

package framework

import (
	"image"
	"image/color"

	"github.com/gop9/olt/gio/io/pointer"
	"github.com/gop9/olt/gio/layout"
	"github.com/gop9/olt/gio/op/paint"
	"github.com/gop9/olt/gio/text"
	"github.com/gop9/olt/gio/unit"
	"github.com/gop9/olt/gio/widget"
)

type checkable struct {
	Label              string
	Color              color.RGBA
	Font               text.Font
	IconColor          color.RGBA
	Size               unit.Value
	shaper             *text.Shaper
	checkedStateIcon   *Icon
	uncheckedStateIcon *Icon
}

func (c *checkable) layout(gtx *layout.Context, checked bool) {

	var icon *Icon
	if checked {
		icon = c.checkedStateIcon
	} else {
		icon = c.uncheckedStateIcon
	}

	hmin := gtx.Constraints.Width.Min
	vmin := gtx.Constraints.Height.Min
	layout.Flex{Alignment: layout.Middle}.Layout(gtx,
		layout.Rigid(func() {
			layout.Align(layout.Center).Layout(gtx, func() {
				layout.UniformInset(unit.Dp(2)).Layout(gtx, func() {
					size := gtx.Px(c.Size)
					icon.Color = c.IconColor
					icon.Layout(gtx, unit.Px(float32(size)))
					gtx.Dimensions = layout.Dimensions{
						Size: image.Point{X: size, Y: size},
					}
				})
			})
		}),

		layout.Rigid(func() {
			gtx.Constraints.Width.Min = hmin
			gtx.Constraints.Height.Min = vmin
			layout.Align(layout.Start).Layout(gtx, func() {
				layout.UniformInset(unit.Dp(2)).Layout(gtx, func() {
					paint.ColorOp{Color: c.Color}.Add(gtx.Ops)
					widget.Label{}.Layout(gtx, c.shaper, c.Font, c.Label)
				})
			})
		}),
	)
	pointer.Rect(image.Rectangle{Max: gtx.Dimensions.Size}).Add(gtx.Ops)
}
