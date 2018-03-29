package pieces

type Movable interface {
	InitializePlacement(initialPosition Position)

	IsMoveValid(newPosition Position) bool
	Move(newPosition Position) bool

	CurrentPosition() Position

	SetCaptured()
	IsCaptured() bool
}
