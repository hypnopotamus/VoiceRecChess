package ui

import (
	"game"

	"github.com/lxn/walk/declarative"
)

type Board struct {
	gameBoard game.Board
}

func (board Board) Render() declarative.Widget {
	var rowWidgets declarative.HSplitter

	for _, rows := range board.gameBoard.Spaces {
		var row declarative.VSplitter
		for _, space := range rows {
			spaceWidget := SpaceWidget{Space: space}
			row.Children = append(row.Children, spaceWidget)
		}
		rowWidgets.Children = append(rowWidgets.Children, row)
	}

	return rowWidgets
}
