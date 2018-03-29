package pieces

type Rook struct {
	Piece
}

func (rook Rook) IsMoveValid(newPosition Position) bool {
	var horizontalDelta, verticalDelta = rook.LocationDelta(newPosition)

	return (horizontalDelta != 0 && verticalDelta == 0) || (horizontalDelta == 0 && verticalDelta != 0)
}
