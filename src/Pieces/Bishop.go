package Pieces

type Bishop struct {
	Piece
}

func (bishop Bishop) IsMoveValid(newPosition Position) bool {
	var horizontalDelta, verticalDelta = bishop.LocationDelta(newPosition)

	return horizontalDelta == verticalDelta
}
