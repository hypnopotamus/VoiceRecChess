package Pieces

type Queen struct {
	Piece
}

func (this Queen) IsMoveValid(newPosition Position) bool {
	return false
}
