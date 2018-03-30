package pieces

type Knight struct {
	Piece
}

func (knight Knight) IsMoveValid(newPosition Position) bool {
	var horizontalDelta, verticalDelta = knight.LocationDelta(newPosition)

	return (horizontalDelta == 2 || verticalDelta == 2) && (horizontalDelta == 1 || verticalDelta == 1)
}

func (knight Knight) GetPieceType() string {
	return "knight"
}
