package pieces

type Movable interface {
	InitializePlacement(initialPosition Position)

	IsMoveValid(newPosition Position) bool
	Move(newPosition Position) bool

	SetColor(color Color)
	GetColor() Color

	CurrentPosition() Position

	SetCaptured()
	IsCaptured() bool
}
