package coreterm

import (
	"fmt"
	"image"
	"strings"

	"gioui.org/op/paint"
)

const (
	rows    = 25
	columns = 80
)

type Instance struct {
	videoTextMemory  [rows * columns * 2]byte
	Border           int
	Height           int
	Width            int
	Scale            float64
	CurrentColor     byte
	uTime            uint64
	updateScreen     bool
	tmpScreen        *paint.ImageOp
	img              *image.RGBA
	ScreenHandler    func(*Instance) error
	Title            string
	cursor           int
	cursorBlinkTimer int
	cursorSetBlink   bool
	machine          int
	cpx, cpy         int
	Font             struct {
		Height int
		Width  int
		Bitmap []byte
	}
	lastKey struct {
		Time uint64
		Char byte
	}
	noKey bool
}

func New() *Instance {
	var i *Instance
	i = &Instance{}
	i.Width = columns * 9
	i.Height = rows * 16
	i.Scale = 1
	i.ScreenHandler = i.update
	i.Title = "term"
	i.CurrentColor = 0x0F
	i.cursorSetBlink = true
	return i
}

var Colors = []struct {
	R byte
	G byte
	B byte
}{
	{0, 0, 0},
	{0, 0, 170},
	{0, 170, 0},
	{0, 170, 170},
	{170, 0, 0},
	{170, 0, 170},
	{170, 85, 0},
	{170, 170, 170},
	{85, 85, 85},
	{85, 85, 255},
	{85, 255, 85},
	{85, 255, 255},
	{255, 85, 85},
	{255, 85, 255},
	{255, 255, 85},
	{255, 255, 255},
}

func MergeColorCode(b, f byte) byte {
	return (f & 0xff) | (b << 4)
}

func (i *Instance) updateTermScreen(screen *paint.ImageOp) error {
	if i.ScreenHandler != nil {
		err := i.ScreenHandler(i)
		if err != nil {
			return err
		}
	}
	if i.updateScreen {
		//i.tmpScreen.ReplacePixels(i.img.Pix)
		i.updateScreen = false
	}
	//screen.DrawImage(i.tmpScreen, nil)
	i.uTime++
	return nil
}

func (i *Instance) Run() {

	i.Font.Bitmap = bitmap
	i.Font.Height = 16
	i.Font.Width = 9
	i.img = image.NewRGBA(image.Rect(0, 0, i.Width, i.Height))
	//i.tmpScreen, _ = paint.NewImageOp(i.Width, i.Height, ebiten.FilterNearest)

	i.Clear()
	i.updateScreen = true
	i.clearVideoTextMode()

	//if err := ebiten.Run(i.updateTermScreen, i.Width, i.Height, i.Scale, i.Title); err != nil {
	//	log.Fatal(err)
	//}
}

func (i *Instance) DrawPix(x, y int, color byte) {
	x += i.Border
	y += i.Border
	if x < i.Border ||
		y < i.Border ||
		x >= i.Width-i.Border ||
		y >= i.Height-i.Border {
		return
	}
	pos := 4*y*i.Width + 4*x
	i.img.Pix[pos] = Colors[color].R
	i.img.Pix[pos+1] = Colors[color].G
	i.img.Pix[pos+2] = Colors[color].B
	i.img.Pix[pos+3] = 0xff
	i.updateScreen = true
}

func (i *Instance) DrawChar(index, fgColor, bgColor byte, x, y int) {
	var a uint
	var b uint
	var lColor byte
	for b = 0; b < 16; b++ {
		for a = 0; a < 9; a++ {
			if a == 8 {
				c := bgColor
				if index >= 192 && index <= 223 {
					c = lColor
				}
				i.DrawPix(int(a)+x, int(b)+y, c)
				continue
			}
			idx := uint(index)*16 + b
			if bitmap[idx]&(0x80>>a) != 0 {
				lColor = fgColor
				i.DrawPix(int(a)+x, int(b)+y, lColor)
				continue
			}
			lColor = bgColor
			i.DrawPix(int(a)+x, int(b)+y, lColor)
		}
	}
}

func (i *Instance) Clear() {
	for idx := 0; idx < i.Height*i.Width*4; idx += 4 {
		i.img.Pix[idx] = Colors[i.CurrentColor].R
		i.img.Pix[idx+1] = Colors[i.CurrentColor].G
		i.img.Pix[idx+2] = Colors[i.CurrentColor].B
		i.img.Pix[idx+3] = 0xff
	}
}

func (i *Instance) DrawCursor(index, fgColor, bgColor byte, x, y int) {
	if i.cursorSetBlink {
		if i.cursorBlinkTimer < 15 {
			fgColor, bgColor = bgColor, fgColor
		}
		i.DrawChar(index, fgColor, bgColor, x, y)
		i.cursorBlinkTimer++
		if i.cursorBlinkTimer > 30 {
			i.cursorBlinkTimer = 0
		}
		return
	}
	i.DrawChar(index, bgColor, fgColor, x, y)
}

func (i *Instance) DrawVideoTextMode() {
	idx := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			color := i.videoTextMemory[idx]
			f := color & 0x0f
			b := color & 0xf0 >> 4
			if idx == i.cursor {
				idx++
				i.DrawCursor(i.videoTextMemory[idx], f, b, c*9, r*16)
			} else {
				idx++
				i.DrawChar(i.videoTextMemory[idx], f, b, c*9, r*16)
			}
			idx++
		}
	}
}

func (i *Instance) clearVideoTextMode() {
	copy(i.videoTextMemory[:], make([]byte, len(i.videoTextMemory)))
	for idx := 0; idx < len(i.videoTextMemory); idx += 2 {
		i.videoTextMemory[idx] = i.CurrentColor
	}
}

func (i *Instance) moveLineUp() {
	copy(i.videoTextMemory[0:], i.videoTextMemory[columns*2:])
	copy(i.videoTextMemory[len(i.videoTextMemory)-columns*2:], make([]byte, columns*2))
	for idx := len(i.videoTextMemory) - columns*2; idx < len(i.videoTextMemory); idx += 2 {
		i.videoTextMemory[idx] = i.CurrentColor
	}
}

func (i *Instance) correctVideoCursor() {
	if i.cursor < 0 {
		i.cursor = 0
	}
	for i.cursor >= rows*columns*2 {
		i.cursor -= columns * 2
		i.moveLineUp()
	}
}

func (i *Instance) putChar(c byte) {
	i.correctVideoCursor()
	i.videoTextMemory[i.cursor] = i.CurrentColor
	i.cursor++
	i.correctVideoCursor()
	i.videoTextMemory[i.cursor] = c
	i.cursor++
	i.correctVideoCursor()
}

func (i *Instance) bPrint(msg string) {
	for idx := 0; idx < len(msg); idx++ {
		c := msg[idx]

		switch c {
		case 13:
			i.cursor += columns * 2
			continue
		case 10:
			aux := i.cursor / (columns * 2)
			aux = aux * (columns * 2)
			i.cursor = aux
			continue
		}
		i.putChar(msg[idx])
	}
}

func (i *Instance) bPrintln(msg string) {
	msg += "\r\n"
	i.bPrint(msg)
}

func (i *Instance) keyTreatment(c byte, f func(c byte)) {
	if i.noKey || i.lastKey.Char != c || i.lastKey.Time+20 < i.uTime {
		f(c)
		i.noKey = false
		i.lastKey.Char = c
		i.lastKey.Time = i.uTime
	}
}

func (i *Instance) getLine() string {
	aux := i.cursor / (columns * 2)
	var ret string
	for idx := aux*(columns*2) + 1; idx < aux*(columns*2)+columns*2; idx += 2 {
		c := i.videoTextMemory[idx]
		if c == 0 {
			break
		}
		ret += string(i.videoTextMemory[idx])
	}

	ret = strings.TrimSpace(ret)
	return ret
}

func (i *Instance) eval(cmd string) {
	fmt.Println("eval:", cmd)
}

//
//func (i *Instance) input() {
//	for c := 'A'; c <= 'Z'; c++ {
//		if pointer.IsKeyPressed(gkey.Key(c) - 'A' + gkey.KeyA) {
//			i.keyTreatment(byte(c), func(c byte) {
//				if pointer.IsKeyPressed(gkey.ModShift) {
//					i.putChar(c)
//					return
//				}
//				i.putChar(c + 32) // convert to lowercase
//			})
//			return
//		}
//	}
//
//	for c := '0'; c <= '9'; c++ {
//		if pointer.IsKeyPressed(gkey.Key(c) - '0' + gkey.Key0) {
//			i.keyTreatment(byte(c), func(c byte) {
//				i.putChar(c)
//			})
//			return
//		}
//	}
//
//	if pointer.IsKeyPressed(gkey.CodeSpacebar) {
//		i.keyTreatment(byte(' '), func(c byte) {
//			i.putChar(c)
//		})
//		return
//	}
//
//	if pointer.IsKeyPressed(gkey.CodeComma) {
//		i.keyTreatment(byte(','), func(c byte) {
//			i.putChar(c)
//		})
//		return
//	}
//
//	if pointer.IsKeyPressed(gkey.NameReturn) {
//		i.keyTreatment(0, func(c byte) {
//			i.eval(i.getLine())
//			i.cursor += columns * 2
//			aux := i.cursor / (columns * 2)
//			aux = aux * (columns * 2)
//			i.cursor = aux
//			i.correctVideoCursor()
//		})
//		return
//	}
//
//	if pointer.IsKeyPressed(gkey.CodeDeleteBackspace) {
//		i.keyTreatment(0, func(c byte) {
//			i.cursor -= 2
//			line := i.cursor / (columns * 2)
//			lineEnd := line*columns*2 + columns*2
//			if i.cursor < 0 {
//				i.cursor = 0
//			}
//
//			copy(i.videoTextMemory[i.cursor:lineEnd], i.videoTextMemory[i.cursor+2:lineEnd])
//			i.videoTextMemory[lineEnd-2] = i.CurrentColor
//			i.videoTextMemory[lineEnd-1] = 0
//
//			i.correctVideoCursor()
//		})
//		return
//	}

/*
   KeyMinus: -
   KeyEqual: =
   KeyLeftBracket: [
   KeyRightBracket: ]
   KeyBackslash:
   KeySemicolon: ;
   KeyApostrophe: '
   KeySlash: /
   KeyGraveAccent: `
*/

//shift := gkey.IsKeyPressed(gkey.ModShift)
//
//	if gkey.IsKeyPressed(gkey.CodeEqualSign) {
//		if shift {
//			i.keyTreatment('+', func(c byte) {
//				i.putChar(c)
//				println("+")
//			})
//			return
//		} else {
//			i.keyTreatment('=', func(c byte) {
//				i.putChar(c)
//				println("=")
//			})
//			return
//		}
//	}
//
//	if pointer.IsKeyPressed(gkey.NameUpArrow) {
//		i.keyTreatment(0, func(c byte) {
//			i.cursor -= columns * 2
//			i.correctVideoCursor()
//		})
//		return
//	}
//	if pointer.IsKeyPressed(gkey.NameDownArrow) {
//		i.keyTreatment(0, func(c byte) {
//			i.cursor += columns * 2
//			i.correctVideoCursor()
//		})
//		return
//	}
//	if pointer.IsKeyPressed(gkey.NameLeftArrow) {
//		i.keyTreatment(0, func(c byte) {
//			i.cursor -= 2
//			i.correctVideoCursor()
//		})
//		return
//	}
//	if pointer.IsKeyPressed(gkey.NameRightArrow) {
//		i.keyTreatment(0, func(c byte) {
//			i.cursor += 2
//			i.correctVideoCursor()
//		})
//		return
//	}
//
//	// When the "left mouse button" is pressed...
//	if pointer.IsMouseButtonPressed(pointer.ButtonLeft) {
//		//ebitenutil.DebugPrint(screen, "You're pressing the 'LEFT' mouse button.")
//	}
//	// When the "right mouse button" is pressed...
//	if pointer.IsMouseButtonPressed(pointer.ButtonRight) {
//		//ebitenutil.DebugPrint(screen, "\nYou're pressing the 'RIGHT' mouse button.")
//	}
//	// When the "middle mouse button" is pressed...
//	if pointer.IsMouseButtonPressed(pointer.ButtonMiddle) {
//		//ebitenutil.DebugPrint(screen, "\n\nYou're pressing the 'MIDDLE' mouse button.")
//	}
//
//	i.cpx, i.cpy = pointer.CursorPosition()
//	//fmt.Printf("X: %d, Y: %d\n", x, y)
//
//	// Display the information with "X: xx, Y: xx" format
//	//ebitenutil.DebugPrint(screen, fmt.Sprintf("X: %d, Y: %d", x, y))
//
//	i.noKey = true
//
//}

func (i *Instance) update(screen *Instance) error {
	i.uTime++

	if i.machine == 0 {
		i.machine++
		i.bPrintln("          1         2         3         4         5         6         7")
		i.bPrintln("01234567890123456789012345678901234567890123456789012345678901234567890123456789")
		i.bPrintln("terminal v0.01")
		i.bPrintln("https://crg.eti.br")
		i.bPrintln(fmt.Sprintf("Width: %v, Height: %v", screen.Width, screen.Height))
		i.bPrintln("")

		var c byte
		for ; c < 246; c++ {
			i.putChar(c)
		}

	}

	i.DrawVideoTextMode()

	//i.input()
	return nil
}
