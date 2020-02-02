package main

import (
	"github.com/gop9/olt/gio/f32"
	"github.com/gop9/olt/gio/op/paint"
	"golang.org/x/exp/shiny/iconvg"
	"image"
	"image/color"
	"image/draw"
)

func rgb(c uint32) color.RGBA {
	return argb(0xff000000 | c)
}
func argb(c uint32) color.RGBA {
	return color.RGBA{A: uint8(c >> 24), R: uint8(c >> 16), G: uint8(c >> 8), B: uint8(c)}
}
func toPointF(p image.Point) f32.Point {
	return f32.Point{X: float32(p.X), Y: float32(p.Y)}
}

func (ic *Icon) image(sz int) paint.ImageOp {
	if sz == ic.imgSize && ic.Color == ic.imgColor {
		return ic.op
	}
	m, _ := iconvg.DecodeMetadata(ic.src)
	dx, dy := m.ViewBox.AspectRatio()
	img := image.NewRGBA(image.Rectangle{Max: image.Point{X: sz, Y: int(float32(sz) * dy / dx)}})
	var ico iconvg.Rasterizer
	ico.SetDstImage(img, img.Bounds(), draw.Src)
	m.Palette[0] = ic.Color
	iconvg.Decode(&ico, ic.src, &iconvg.DecodeOptions{
		Palette: &m.Palette,
	})
	ic.op = paint.NewImageOp(img)
	ic.imgSize = sz
	ic.imgColor = ic.Color
	return ic.op
}
