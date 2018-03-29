package Pieces

type Pawn struct {
	Piece
}

func (this Pawn) IsMoveValid(newPosition Position) bool {
	return false
}
