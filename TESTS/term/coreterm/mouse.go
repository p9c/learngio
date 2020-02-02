package coreterm

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"

	"github.com/gop9/olt/gio/pointers"
)

func Handler() {
	// When the "left mouse button" is pressed...
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		println("ebiten.MouseButtonLeft")
	}
	// When the "right mouse button" is pressed...
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		println("ebiten.MouseButtonRight")
	}
	// When the "middle mouse button" is pressed...
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle) {
		println("ebiten.MouseButtonMiddle")
	}

	cpx, cpy := ebiten.CursorPosition()
	//cpx -= cs.Border
	//cpy -= cs.Border
	fmt.Printf("X: %d, Y: %d\n", cpx, cpy)

}
