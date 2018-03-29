package Pieces

type Pawn struct {
	Piece
	HasMoved bool
}

func (pawn Pawn) IsMoveValid(newPosition Position) bool {
	var horizontalDelta, verticalDelta = pawn.LocationDelta(newPosition)

	return (horizontalDelta == 0 && ((pawn.HasMoved && verticalDelta == 2) || verticalDelta == 1)) || (newPosition.Occupied && horizontalDelta == 1 && verticalDelta == 1)
}
