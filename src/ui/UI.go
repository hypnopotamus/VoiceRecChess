package ui

import (
	"game"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type UI struct {
	board game.Board
}

func (ui UI) Render() {
	var pieceIn, columnIn, rowIn *walk.TextEdit
	var pieceOut, columnOut, rowOut *walk.TextEdit

	MainWindow{
		Title:   "Chess",
		MinSize: Size{800, 600},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{Text: "the board should go here", ReadOnly: true},
					VSplitter{
						Children: []Widget{
							VSplitter{
								Children: []Widget{
									TextEdit{AssignTo: &pieceIn},
									TextEdit{AssignTo: &pieceOut, ReadOnly: true},
									PushButton{
										Text: "Piece",
										OnClicked: func() {
											pieceOut.SetText(pieceIn.Text())
										},
									},
								},
							},
							VSplitter{
								Children: []Widget{
									TextEdit{AssignTo: &columnIn},
									TextEdit{AssignTo: &columnOut, ReadOnly: true},
									PushButton{
										Text: "Column",
										OnClicked: func() {
											columnOut.SetText(columnIn.Text())
										},
									},
								},
							},
							VSplitter{
								Children: []Widget{
									TextEdit{AssignTo: &rowIn},
									TextEdit{AssignTo: &rowOut, ReadOnly: true},
									PushButton{
										Text: "Row",
										OnClicked: func() {
											rowOut.SetText(rowIn.Text())
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}.Run()
}
