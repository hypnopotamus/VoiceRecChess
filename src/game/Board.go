package game

import (
	"pieces"
	"reflect"
)

type Board struct {
	BlackPieces [16]pieces.Movable
	WhitePieces [16]pieces.Movable
	Spaces      [8][8]Space
}

func InitializeBoard() Board {
	var board Board

	board.initialize(board.BlackPieces, board.Spaces[0], board.Spaces[1])
	board.initialize(board.WhitePieces, board.Spaces[7], board.Spaces[6])
	swapKingAndQueenPosition(board.BlackPieces)

	return board
}

func (board Board) initialize(piecesArray [16]pieces.Movable, notPawnRow [8]Space, pawnRow [8]Space) {
	var pieceIndex = 0

	for _, space := range pawnRow {
		var pawn pieces.Pawn

		space.setOccupant(pawn.Piece)
		piecesArray[pieceIndex] = pawn.Piece

		pieceIndex++
	}

	for index, space := range pawnRow {
		var piece = makePieceForSpace(index)

		space.setOccupant(piece)
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
