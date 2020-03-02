package controller

import (
	"regexp"

	"github.com/p9c/learngio/examples/html/native/model"
)

var rgba = regexp.MustCompile(`rgba?\([\.?\d?\.?\d?%?\s?,?]+\)`)
var rgbaParams = regexp.MustCompile(`\([\.?\d?\.?\d?%?\s?,?]+\)`)

var colorTable = map[string]*model.ColorRGBA{
	"maroon":  &model.ColorRGBA{R: 0.5, G: 0.0, B: 0.0, A: 1.0},
	"red":     &model.ColorRGBA{R: 1.0, G: 0.0, B: 0.0, A: 1.0},
	"orange":  &model.ColorRGBA{R: 1.0, G: 0.6, B: 0.0, A: 1.0},
	"yellow":  &model.ColorRGBA{R: 1.0, G: 1.0, B: 0.0, A: 1.0},
	"olive":   &model.ColorRGBA{R: 0.5, G: 0.5, B: 0.0, A: 1.0},
	"green":   &model.ColorRGBA{R: 0.0, G: 1.0, B: 0.0, A: 1.0},
	"purple":  &model.ColorRGBA{R: 0.5, G: 0.0, B: 0.5, A: 1.0},
	"fuchsia": &model.ColorRGBA{R: 1.0, G: 0.0, B: 1.0, A: 1.0},
	"lime":    &model.ColorRGBA{R: 0.0, G: 1.0, B: 0.0, A: 1.0},
	"teal":    &model.ColorRGBA{R: 0.0, G: 0.5, B: 0.5, A: 1.0},
	"aqua":    &model.ColorRGBA{R: 0.0, G: 1.0, B: 1.0, A: 1.0},
	"blue":    &model.ColorRGBA{R: 0.0, G: 0.0, B: 1.0, A: 1.0},
	"navy":    &model.ColorRGBA{R: 0.0, G: 0.0, B: 0.5, A: 1.0},
	"black":   &model.ColorRGBA{R: 0.0, G: 0.0, B: 0.0, A: 1.0},
	"gray":    &model.ColorRGBA{R: 0.5, G: 0.5, B: 0.5, A: 1.0},
	"silver":  &model.ColorRGBA{R: 0.7, G: 0.7, B: 0.7, A: 1.0},
	"white":   &model.ColorRGBA{R: 1.0, G: 1.0, B: 1.0, A: 1.0},
}

var elementFontTable = map[string]float64{
	"h1": float64(32),
	"h2": float64(28),
	"p":  float64(14),
}
