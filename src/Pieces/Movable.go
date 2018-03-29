package Pieces

type Movable interface {
	IsMoveValid(newPosition Position) bool
}
