package Pieces

type Pawn struct {
	Piece
}

func (pawn Pawn) IsMoveValid(newPosition Position) bool {
	var horizontalDelta, verticalDelta = pawn.LocationDelta(newPosition)

	return horizontalDelta == 0 || (newPosition.Occupied && horizontalDelta == 1 && verticalDelta == 1)
}
