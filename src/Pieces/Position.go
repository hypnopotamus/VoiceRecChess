package pieces

type Position struct {
	Horizontal int
	Vertical   int
	Occupant   *Movable
	Color      Color
}

func (position Position) SetOccupant(piece Movable) {
	piece.InitializePlacement(position)
}

func (position Position) isOccupied() bool {
	return position.Occupant != nil
}
