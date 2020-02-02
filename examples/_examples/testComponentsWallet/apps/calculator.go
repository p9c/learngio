package apps

import (
	"fmt"
	"strconv"

	"github.com/gop9/olt/framework"
	"github.com/gop9/olt/framework/label"
)

type CalcDemo struct {
	A, b    float64
	Current *float64

	set      bool
	op, prev string

	editor framework.TextEditor
}

var calcBtns = []string{
	"7", "8", "9", "+",
	"4", "5", "6", "-",
	"1", "2", "3", "*",
	"C", "0", "=", "/",
}

func (c *CalcDemo) CalculatorDemo(w *framework.Window) {
	w.Row(35).Dynamic(1)
	c.editor.Flags = framework.EditSimple
	c.editor.Filter = framework.FilterFloat
	c.editor.Maxlen = 255
	c.editor.Buffer = []rune(fmt.Sprintf("%.2f", *c.Current))
	c.editor.Edit(w)
	*c.Current, _ = strconv.ParseFloat(string(c.editor.Buffer), 64)

	w.Row(35).Dynamic(4)
	solve := false
	for _, btn := range calcBtns {
		if w.Button(label.T(btn), false) {
			switch btn {
			case "+", "-", "*", "/":
				if !c.set {
					if c.Current != &c.b {
						c.Current = &c.b
					} else {
						c.prev = c.op
						solve = true
					}
				}
				c.op = btn
				c.set = true
			case "C":
				c.A = 0.0
				c.b = 0.0
				c.op = ""
				c.Current = &c.A
				c.set = false
			case "=":
				solve = true
				c.prev = c.op
				c.op = ""
				c.set = false
			default:
				*c.Current = *c.Current*10 + float64(btn[0]-'0')
			}
		}
	}
	if solve {
		switch c.prev {
		case "+":
			c.A = c.A + c.b
		case "-":
			c.A = c.A - c.b
		case "*":
			c.A = c.A * c.b
		case "/":
			c.A = c.A / c.b
		}
		c.Current = &c.A
		if c.set {
			c.Current = &c.b
		}
		c.b = 0
		c.set = false
	}

}
