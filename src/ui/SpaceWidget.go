package ui

import (
	"pieces"

	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
)

type SpaceWidget struct {
	declarative.Widget
	declarative.CustomWidget
	Space pieces.Position
}

func (space SpaceWidget) Create(builder *declarative.Builder) error {
	colorByte := getColorByte(space.Space.Color)
	brush, err := walk.NewSolidColorBrush(walk.RGB(colorByte, colorByte, colorByte))
	defer brush.Dispose()
	declarativeBrush := declarative.SolidColorBrush{Color: brush.Color()}
	space.CustomWidget.Background = declarativeBrush

	if err != nil {
		panic(err)
	}
	return space.CustomWidget.Create(builder)
}

func getColorByte(color pieces.Color) byte {
	switch color {
	case pieces.Black:
		return 0
	case pieces.White:
		return 255
	default:
		return 255
	}
}
