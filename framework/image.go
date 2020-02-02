// SPDX-License-Identifier: Unlicense OR MIT

package framework

import (
	"image"

	"github.com/gop9/olt/gio/f32"
	"github.com/gop9/olt/gio/layout"
	"github.com/gop9/olt/gio/op"
	"github.com/gop9/olt/gio/op/clip"
	"github.com/gop9/olt/gio/op/paint"
	"github.com/gop9/olt/gio/unit"
)

// Image is a widget that displays an image.
type Image struct {
	// Src is the image to display.
	Src paint.ImageOp
	// Scale is the ratio of image pixels to
	// dps.
	Scale float32
}

func (t *Framework) Image(img paint.ImageOp) Image {
	return Image{
		Src:   img,
		Scale: 160 / 72, // About 72 DPI.
	}
}

func (im Image) Layout(gtx *layout.Context) {
	size := im.Src.Size()
	wf, hf := float32(size.X), float32(size.Y)
	w, h := gtx.Px(unit.Dp(wf*im.Scale)), gtx.Px(unit.Dp(hf*im.Scale))
	cs := gtx.Constraints
	d := image.Point{X: cs.Width.Constrain(w), Y: cs.Height.Constrain(h)}
	var s op.StackOp
	s.Push(gtx.Ops)
	clip.Rect{Rect: f32.Rectangle{Max: toPointF(d)}}.Op(gtx.Ops).Add(gtx.Ops)
	im.Src.Add(gtx.Ops)
	paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: float32(w), Y: float32(h)}}}.Add(gtx.Ops)
	s.Pop()
	gtx.Dimensions = layout.Dimensions{Size: d}
}
