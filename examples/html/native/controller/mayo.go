package controller

import (
	"strconv"
	"strings"

	"github.com/p9c/learngio/examples/html/native/model"
)

func getDefaultElementDisplay(element string) string {
	displayType := "block"

	switch element {
	case "script", "style", "meta", "link", "head", "title":
		displayType = "none"
	default:
		displayType = "block"
	}

	return displayType
}

func mapSizeValue(sizeValue string) float64 {
	valueString := sizeValue[0 : len(sizeValue)-2]
	value, err := strconv.ParseInt(valueString, 10, 0)

	if err != nil {
		return float64(14)
	}

	return float64(value)
}

func mapPropToStylesheet(parsedStyleSheet *model.Stylesheet, propSlice []string) *model.Stylesheet {
	propName := propSlice[0]
	propValue := propSlice[1]

	switch propName {
	case "color":
		parsedStyleSheet.Color = MapCSSColor(propValue)
	case "font-size":
		parsedStyleSheet.FontSize = mapSizeValue(propValue)
	case "display":
		parsedStyleSheet.Display = propValue
	case "postion":
		parsedStyleSheet.Position = propValue
	}

	return parsedStyleSheet
}

func parseInlineStylesheet(attributes []*model.Attribute) *model.Stylesheet {
	parsedStylesheet := &model.Stylesheet{}

	for i := 0; i < len(attributes); i++ {
		attributeName := attributes[i].Name
		if attributeName == "style" {

			styleString := attributes[i].Value
			styleProps := strings.Split(strings.Replace(styleString, " ", "", -1), ";")

			for i := 0; i < len(styleProps); i++ {
				styledProperty := strings.Split(styleProps[i], ":")
				if len(styledProperty) >= 2 {
					parsedStylesheet = mapPropToStylesheet(parsedStylesheet, styledProperty)
				}
			}
		}
	}

	return parsedStylesheet
}

func hasInlineStyle(attributes []*model.Attribute) bool {
	inlineStyle := false

	for i := 0; i < len(attributes); i++ {
		attributeName := attributes[i].Name
		if attributeName == "style" {
			inlineStyle = true
		}
	}

	return inlineStyle
}

func GetElementStylesheet(elementName string, attributes []*model.Attribute) *model.Stylesheet {
	elementStylesheet := &model.Stylesheet{
		Color:    &model.ColorRGBA{0, 0, 0, 0},
		FontSize: 0,
		Display:  "",
		Position: "Normal",
	}

	if hasInlineStyle(attributes) {
		elementStylesheet = parseInlineStylesheet(attributes)
	}

	if elementStylesheet.FontSize == float64(0) {
		fontSize := elementFontTable[elementName]

		if fontSize != float64(0) {
			elementStylesheet.FontSize = fontSize
		} else {
			elementStylesheet.FontSize = float64(14)
		}

		if elementStylesheet.Height == float64(0) {
			elementStylesheet.Height = elementStylesheet.FontSize
		}
	}

	if elementStylesheet.Display == "" {
		elementStylesheet.Display = getDefaultElementDisplay(elementName)
	}

	return elementStylesheet
}
