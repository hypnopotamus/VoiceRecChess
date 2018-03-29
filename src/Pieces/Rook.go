package Pieces

type Rook struct {
	Piece
}

func (this Rook) IsMoveValid(newPosition Position) bool {
	return false
}
