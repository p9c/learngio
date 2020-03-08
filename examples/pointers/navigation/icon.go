package main

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"golang.org/x/exp/shiny/iconvg"
	"image/color"
)

type Icon struct {
	Color color.RGBA
	src   []byte
	size  unit.Value
	// Cached values.
	op       paint.ImageOp
	imgSize  int
	imgColor color.RGBA
}

// NewDuoUIicon returns a new DuoUIicon from DuoUIiconVG data.
func NewIcon(data []byte) (*Icon, error) {
	_, err := iconvg.DecodeMetadata(data)
	if err != nil {
		return nil, err
	}
	return &Icon{src: data, Color: rgb(0x000000)}, nil
}

func (ic *Icon) Layout(gtx *layout.Context, sz unit.Value) {
	ico := ic.image(gtx.Px(sz))
	ico.Add(gtx.Ops)
	paint.PaintOp{
		Rect: f32.Rectangle{
			Max: toPointF(ico.Size()),
		},
	}.Add(gtx.Ops)
}
