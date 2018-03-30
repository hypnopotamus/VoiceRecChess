package pieces

type Queen struct {
	Piece
}

func (queen Queen) IsMoveValid(newPosition Position) bool {
	bishop := Bishop{Piece: queen.Piece}
	rook := Rook{Piece: queen.Piece}

	return bishop.IsMoveValid(newPosition) || rook.IsMoveValid(newPosition)
}

func (queen Queen) GetPieceType() string {
	return "queen"
}
