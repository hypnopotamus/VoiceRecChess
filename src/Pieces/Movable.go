package Pieces

type Movable interface {
	IsMoveValid(newPosition Position) bool
	Move(newPosition Position) bool
}
