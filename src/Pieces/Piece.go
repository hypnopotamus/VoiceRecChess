package pieces

type Piece struct {
	Movable
	Location *Position
}

func (piece Piece) InitializePlacement(initialPosition Position) {
	piece.Location = &initialPosition
}

func (piece Piece) Move(newPosition Position) bool {
	if piece.IsMoveValid(newPosition) {
		piece.Location.Occupied = false
		//if newPosition.Occupied capture that piece
		piece.Location = &newPosition
		piece.Location.Occupied = true

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
