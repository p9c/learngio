// SPDX-License-Identifier: Unlicense OR MIT

package framework

import (
	"github.com/gop9/olt/gio/layout"
	"github.com/gop9/olt/gio/text"
	"github.com/gop9/olt/gio/unit"
	"github.com/gop9/olt/gio/widget"
)

type CheckBox struct {
	checkable
}

func (t *Framework) CheckBox(label string) CheckBox {
	return CheckBox{
		checkable{
			Label:     label,
			Color:     t.Color.Text,
			IconColor: t.Color.Primary,
			Font: text.Font{
				Size: t.TextSize.Scale(14.0 / 16.0),
			},
			Size:               unit.Dp(26),
			shaper:             t.Shaper,
			checkedStateIcon:   t.checkBoxCheckedIcon,
			uncheckedStateIcon: t.checkBoxUncheckedIcon,
		},
	}
}

func (c CheckBox) Layout(gtx *layout.Context, checkBox *widget.CheckBox) {
	c.layout(gtx, checkBox.Checked(gtx))
	checkBox.Layout(gtx)
}
