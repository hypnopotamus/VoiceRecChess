package game

import "pieces"

type Space struct {
	pieces.Position
	Occupant *pieces.Movable
}

func (space Space) setOccupant(piece pieces.Movable) {
	space.Position.Occupied = true
	space.Occupant = &piece
	piece.InitializePlacement(space.Position)
}
