package game

import (
	"pieces"
	"reflect"
)

type Board struct {
	BlackPieces [16]pieces.Movable
	WhitePieces [16]pieces.Movable
	Spaces      [8][8]pieces.Position
}

func InitializeBoard() Board {
	var board Board

	for rowIndex, rows := range board.Spaces {
		for columnIndex, space := range rows {
			space.Horizontal = rowIndex
			space.Vertical = columnIndex
			space.Color = pieces.ToColor(rowIndex + columnIndex)
		}
	}

	board.initialize(board.BlackPieces, board.Spaces[0], board.Spaces[1], pieces.Black)
	board.initialize(board.WhitePieces, board.Spaces[7], board.Spaces[6], pieces.White)
	swapKingAndQueenPosition(board.BlackPieces)

	return board
}

func (board Board) initialize(piecesArray [16]pieces.Movable, notPawnRow [8]pieces.Position, pawnRow [8]pieces.Position, color pieces.Color) {
	var pieceIndex = 0

	for _, space := range pawnRow {
		var pawn pieces.Pawn

		pawn.SetColor(color)
		space.SetOccupant(pawn.Piece.Movable)
		piecesArray[pieceIndex] = pawn.Piece

		pieceIndex++
	}

	for index, space := range pawnRow {
		var piece = makePieceForSpace(index)

		piece.SetColor(color)
		space.SetOccupant(piece)
		piecesArray[pieceIndex] = piece

		pieceIndex++
	}

	return
}

func makePieceForSpace(index int) pieces.Movable {
	var piece pieces.Movable

	if index == 0 || index == 7 {
		piece = new(pieces.Rook)
	}

	if index == 1 || index == 6 {
		piece = new(pieces.Knight)
	}

	if index == 2 || index == 5 {
		piece = new(pieces.Bishop)
	}

	if index == 3 {
		piece = new(pieces.Queen)
	}

	if index == 4 {
		piece = new(pieces.King)
	}

	return piece
}

func swapKingAndQueenPosition(pieceArray [16]pieces.Movable) {
	var king pieces.Movable
	var queen pieces.Movable

	for _, piece := range pieceArray {
		t := reflect.TypeOf(piece)
		if reflect.TypeOf(pieces.King{}) == t {
			king = piece
		}
		if reflect.TypeOf(pieces.King{}) == t {
			queen = piece
		}
	}

	var kingLocation = king.CurrentPosition()
	king.InitializePlacement(queen.CurrentPosition())
	queen.InitializePlacement(kingLocation)

	return
}

func (board Board) RemoveCapturedPieces() {
	removeCaptured(board.WhitePieces)
	removeCaptured(board.BlackPieces)
}

func removeCaptured(pieceArray [16]pieces.Movable) {
	for index, piece := range pieceArray {
		if piece.IsCaptured() {
			pieceArray[index] = nil
		}
	}
}
