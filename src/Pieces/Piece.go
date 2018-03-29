package pieces

type Piece struct {
	Movable
	Location *Position
	Captured bool
}

func (piece Piece) InitializePlacement(initialPosition Position) {
	piece.Location = &initialPosition
	piece.Location.Occupant = &piece.Movable
}

func (piece Piece) Move(newPosition Position) bool {
	if piece.IsMoveValid(newPosition) {
		piece.Location.Occupant = nil
		if newPosition.isOccupied() {
			var occupant = *newPosition.Occupant
			occupant.SetCaptured()
		}
		piece.Location = &newPosition
		piece.Location.Occupant = &piece.Movable

		return true
	}

	return false
}

func (piece Piece) LocationDelta(newPosition Position) (int, int) {
	var horizontalChange = abs(piece.Location.Horizontal - newPosition.Horizontal)
	var verticalChange = abs(piece.Location.Horizontal - newPosition.Horizontal)

	return horizontalChange, verticalChange
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func (piece Piece) CurrentPosition() Position {
	return *piece.Location
}

func (piece Piece) SetCaptured() {
	piece.Captured = true
}

func (piece Piece) IsCaptured() bool {
	return piece.Captured
}
