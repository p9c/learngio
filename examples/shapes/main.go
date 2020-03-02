package main

import (
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/wrnrlr/shape"
	"image/color"
)

var (
	red   = color.RGBA{255, 0, 0, 125}
	green = color.RGBA{0, 255, 0, 125}
	blue  = color.RGBA{0, 0, 255, 125}
	black = color.RGBA{0, 0, 0, 255}
	grey  = color.RGBA{125, 125, 125, 125}
)

func main() {
	go func() {
		w := app.NewWindow()
		gtx := layout.NewContext(w.Queue())
		gofont.Register()
		for {
			e := <-w.Events()
			switch e := e.(type) {
			case system.DestroyEvent:
				return
			case system.FrameEvent:
				gtx.Reset(e.Config, e.Size)
				//painting(gtx)
				drawLine(gtx)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

func painting(gtx *layout.Context) {
	//defaultWidth := float32(gtx.Px(unit.Sp(2)))

	width := float32(gtx.Px(unit.Sp(10)))
	thin := float32(gtx.Px(unit.Sp(1)))
	//light := float32(gtx.Px(unit.Sp(2)))

	var a, b f32.Point

	w, h := float32(gtx.Constraints.Width.Max), float32(gtx.Constraints.Height.Max)

	//a, b = f32.Point{0, h / 2}, f32.Point{w, h / 2}
	//shape.Line{a, b}.Stroke(grey, width, gtx)
	//a, b = f32.Point{w / 2, 0}, f32.Point{w / 2, h}
	//shape.Line{a, b}.Stroke(grey, width, gtx)

	//c = f32.Point{w / 2, h / 2}
	//r := shape.Min(w, h) / 3
	//c1 := shape.Circle{c, r}
	//c1.Fill(grey, gtx)
	//c1.Stroke(grey, width, gtx)

	a, b = f32.Point{w / 2, 0}, f32.Point{w / 2, h}
	r1 := shape.Rectangle{a, b}
	r1.Fill(grey, gtx)
	r1.Stroke(grey, width, gtx)

	a = f32.Point{50, 50}
	shape.Circle{a, 100}.Stroke(red, thin, gtx)
	//shape.Circle{a, 100}.Fill(grey, gtx)

	//a, b = f32.Point{w / 2, 0}, f32.Point{w / 2, h}
	//shape.Rectangle{a, b}.Fill(grey, gtx)

	//a, b = f32.Point{100, 0}, f32.Point{100, 200}
	//shape.Line{a, b}.Stroke(red, light, gtx)
	//
	//a, b = f32.Point{0, 100}, f32.Point{200, 100}
	//shape.Line{a, b}.Stroke(red, light, gtx)
	//
	//a, b = f32.Point{10, 20}, f32.Point{210, 220}
	//shape.Line{a, b}.Stroke(red, light, gtx)

	//a = f32.Point{500, 500}
	//shape.Circle{a, 100}.Fill(red, gtx)
	//var stack op.StackOp
	//
	//stack.Push(gtx.Ops)
	//a = f32.Point{w/2, 0}
	//b = f32.Point{w/2, h}
	//line := shape.Line{a, b}
	//line.Stroke(red, float32(20), gtx)
	//paint.PaintOp{f32.Rectangle{Max:f32.Point{w,h}}}.Add(gtx.Ops)
	//stack.Pop()

	//stack.Push(gtx.Ops)
	//a = f32.Point{0, 0}
	//b = f32.Point{w, h}
	//line = shape.Line{a, b}
	//line.Stroke(unit.Sp(30), gtx)
	//paint.PaintOp{f32.Rectangle{Max:f32.Point{w,h}}}.Add(gtx.Ops)
	//stack.Pop()

	//stack.Push(gtx.Ops)
	//circle := shape.Circle{f32.Point{70, 70}, 40}
	//bbox = circle.Stroke(defaultWidth, gtx)
	//paint.ColorOp{blue}.Add(gtx.Ops)
	//paint.PaintOp{bbox}.Add(gtx.Ops)
	//stack.Pop()

	//stack.Push(gtx.Ops)
	//circle = shape.Circle{f32.Point{160, 70}, 40}
	//bbox = circle.Fill(gtx)
	//paint.ColorOp{blue}.Add(gtx.Ops)
	//paint.PaintOp{bbox}.Add(gtx.Ops)
	//stack.Pop()

	//stack.Push(gtx)
	//shape.StrokeRectangle(f32.Point{40, 160}, f32.Point{100, 60}, 10, gtx)
	//paint.ColorOp{green}.Add(gtx)
	//paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: 600, Y: 600}}}.Add(gtx)
	//stack.Pop()

	//stack.Push(gtx.Ops)
	//rect := shape.Rectangle{f32.Point{200, 160}, f32.Point{300, 400}}
	//bbox = rect.Stroke(10, gtx)
	//paint.ColorOp{green}.Add(gtx.Ops)
	//paint.PaintOp{bbox}.Add(gtx.Ops)
	//stack.Pop()

	//stack.Push(gtx)
	//shape.StrokeRectangle(f32.Point{40, 360}, f32.Point{100, 100}, 10, gtx)
	//paint.ColorOp{green}.Add(gtx)
	//paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: 600, Y: 600}}}.Add(gtx)
	//stack.Pop()
	//
	//stack.Push(gtx)
	//shape.FillRectangle(f32.Point{40, 240}, f32.Point{30, 60}, gtx)
	//paint.ColorOp{green}.Add(gtx)
	//paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: 600, Y: 600}}}.Add(gtx)
	//stack.Pop()

	//stack.Push(gtx.Ops)
	//shape.Triangle{f32.Point{300, 10}, f32.Point{360, 10}, f32.Point{340, 60}}.Fill(gtx)
	//paint.ColorOp{red}.Add(gtx.Ops)
	//paint.PaintOp{bbox}.Add(gtx.Ops)
	//stack.Pop()

	//stack.Push(gtx.Ops)
	//shape.Plus{f32.Point{300, 10}, f32.Point{360, 10}, f32.Point{340, 60}}.Fill(gtx)
	//paint.ColorOp{red}.Add(gtx.Ops)
	//paint.PaintOp{bbox}.Add(gtx.Ops)
	//stack.Pop()
}

func drawLine(gtx *layout.Context) {
	red := color.RGBA{255, 0, 0, 255}
	s := gtx.Constraints
	w, h := float32(s.Width.Max), float32(s.Height.Max)
	shape.Line{{0, h / 5}, {w, h / 3}}.Stroke(red, 1, gtx)

}
