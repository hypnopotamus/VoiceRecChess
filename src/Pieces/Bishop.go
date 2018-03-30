package pieces

type Bishop struct {
	Piece
}

func (bishop Bishop) IsMoveValid(newPosition Position) bool {
	var horizontalDelta, verticalDelta = bishop.LocationDelta(newPosition)

	return horizontalDelta == verticalDelta
}

func (bishop Bishop) GetPieceType() string {
	return "bishop"
}
