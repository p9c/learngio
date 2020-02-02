// SPDX-License-Identifier: Unlicense OR MIT

package framework

import (
	"github.com/gop9/olt/gio/layout"
	"github.com/gop9/olt/gio/text"
	"github.com/gop9/olt/gio/unit"
	"github.com/gop9/olt/gio/widget"
)

type RadioButton struct {
	checkable
	Key string
}

// RadioButton returns a RadioButton with a label. The key specifies
// the value for the Enum.
func (t *Framework) RadioButton(key, label string) RadioButton {
	return RadioButton{
		checkable: checkable{
			Label: label,

			Color:     t.Color.Text,
			IconColor: t.Color.Primary,
			Font: text.Font{
				Size: t.TextSize.Scale(14.0 / 16.0),
			},
			Size:               unit.Dp(26),
			shaper:             t.Shaper,
			checkedStateIcon:   t.radioCheckedIcon,
			uncheckedStateIcon: t.radioUncheckedIcon,
		},
		Key: key,
	}
}

func (r RadioButton) Layout(gtx *layout.Context, enum *widget.Enum) {
	r.layout(gtx, enum.Value(gtx) == r.Key)
	enum.Layout(gtx, r.Key)
}
