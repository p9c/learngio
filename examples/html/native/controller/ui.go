package controller

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/p9c/learngio/examples/html/native/model"
	"github.com/tfriedel6/canvas"
)

//Input "Creates a new Input element"
func Input(id string, w float64, h float64, cv *canvas.Canvas, text string) model.UIElement {
	elementCursor := glfw.CreateStandardCursor(glfw.IBeamCursor)
	inputElement := model.UIElement{ID: id, WType: "input", X: w/2 - w/4, Y: 7.5, W: w / 2, H: 25, Cursor: elementCursor, Text: text}
	//inputElement := model.UIElement{ID: id, WType: "input", X: w/2 - w/4, Y: 7.5, W: w / 2, H: 25, Canvas: cv, Cursor: elementCursor, Text: text}
	model.DrawInput(inputElement)
	return inputElement
}

//Box "Creates a new Box element"
func Box(id string, x float64, y float64, w float64, h float64, cv *canvas.Canvas) model.UIElement {
	elementCursor := glfw.CreateStandardCursor(glfw.ArrowCursor)
	boxElement := model.UIElement{ID: id, WType: "box", X: x, Y: y, W: w, H: h, Cursor: elementCursor, Text: ""}
	//boxElement := model.UIElement{ID: id, WType: "box", X: x, Y: y, W: w, H: h, Canvas: cv, Cursor: elementCursor, Text: ""}
	model.DrawBox(boxElement)
	return boxElement
}

func getFocusedUIElement(eList []*model.UIElement, x float64, y float64) *model.UIElement {
	var focusedElement *model.UIElement

	for i := 0; i < len(eList); i++ {
		if x > eList[i].X && x < eList[i].X+eList[i].W && y > eList[i].Y && y < eList[i].Y+eList[i].H {
			focusedElement = eList[i]
		}
	}

	return focusedElement
}

func getSelectedUIElement(eList []*model.UIElement) *model.UIElement {
	var focusedElement *model.UIElement

	for i := 0; i < len(eList); i++ {
		if eList[i].Selected == true {
			focusedElement = eList[i]
		}
	}

	return focusedElement
}

func getUIElementByID(eList []*model.UIElement, id string) *model.UIElement {
	var element *model.UIElement

	for i := 0; i < len(eList); i++ {
		if eList[i].ID == id {
			element = eList[i]
		}
	}

	return element
}

func removeUIFocus(eList []*model.UIElement) {
	for i := 0; i < len(eList); i++ {
		eList[i].Focused = false
	}
}

func removeUISelection(eList []*model.UIElement) {
	for i := 0; i < len(eList); i++ {
		eList[i].Selected = false
	}
}
