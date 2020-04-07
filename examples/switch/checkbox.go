// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

type DuoUIcheckBox struct {
	checkable
}

func NewDuoUIcheckBox() DuoUIcheckBox {
	return DuoUIcheckBox{
		checkable{
			//Font: text.Font{
			//	Typeface: t.Fonts["Primary"],
			//},
			//Label:              label,
			//Color:              helpers.HexARGB(color),
			//IconColor:          helpers.HexARGB(iconColor),
			//TextSize:           t.TextSize.Scale(14.0 / 16.0),
			Size: unit.Dp(48),
			//shaper:             t.Shaper,
			//checkedStateIcon:   t.Icons["Checked"],
			//uncheckedStateIcon: t.Icons["Unchecked"],
		},
	}
}

func (c DuoUIcheckBox) Layout(gtx *layout.Context, checkBox *CheckBox) {
	c.layout(gtx, checkBox.Checked(gtx))
	checkBox.Layout(gtx)
}
