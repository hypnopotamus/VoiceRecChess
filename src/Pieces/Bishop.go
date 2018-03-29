package Pieces

type Bishop struct {
	Piece
}

func (this Bishop) IsMoveValid(newPosition Position) bool {
	return false
}
