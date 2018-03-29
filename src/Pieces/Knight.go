package Pieces

type Knight struct {
	Piece
}

func (this Knight) IsMoveValid(newPosition Position) bool {
	return false
}
