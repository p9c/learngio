package model

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"github.com/p9c/learngio/examples/icons/ico"
)

type Screen struct {
	TextSize          float32
	IconSize          float32
	BgColor           string
	AccentColor       string
	GroupsIco         []string
	IcoBank           ico.Ico
	NavList           *layout.List
	List              *layout.List
	NavButtonsGroup   *widget.Enum
	BgColorLight      *widget.Button
	BgColorYellow     *widget.Button
	BgColorPurple     *widget.Button
	BgColorRed        *widget.Button
	BgColorBlue       *widget.Button
	BgColorDark       *widget.Button
	BgColorOrange     *widget.Button
	BgColorGreen      *widget.Button
	AccentColorLight  *widget.Button
	AccentColorYellow *widget.Button
	AccentColorPurple *widget.Button
	AccentColorRed    *widget.Button
	AccentColorBlue   *widget.Button
	AccentColorDark   *widget.Button
	AccentColorOrange *widget.Button
	AccentColorGreen  *widget.Button
	Colors            Colors
}
type Colors map[string]string
