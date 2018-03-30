package pieces

type King struct {
	Piece
}

func (king King) IsMoveValid(newPosition Position) bool {
	var horizontalDelta, verticalDelta = king.LocationDelta(newPosition)

	return horizontalDelta <= 1 && verticalDelta <= 1
}

func (king King) GetPieceType() string {
	return "king"
}
