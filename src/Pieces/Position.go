package pieces

type Position struct {
	Horizontal int
	Vertical   int
	Occupant   *Movable
}

func (position Position) SetOccupant(piece Movable) {
	position.Occupant = &piece
	piece.InitializePlacement(position)
}

func (position Position) isOccupied() bool {
	return position.Occupant != nil
}
