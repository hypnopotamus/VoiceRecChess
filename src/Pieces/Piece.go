package Pieces

type Piece struct {
	Movable
	Location Position
}

func (piece Piece) Move(newPosition Position) bool {
	if piece.IsMoveValid(newPosition) {
		piece.Location.Occupied = false
		//if newPosition.Occupied capture that piece
		piece.Location = newPosition
		piece.Location.Occupied = true

		return true
	}

	return false
}

func (piece Piece) LocationDelta(newPosition Position) (int, int) {
	var horizontalChange = Abs(piece.Location.Horizontal - newPosition.Horizontal)
	var verticalChange = Abs(piece.Location.Horizontal - newPosition.Horizontal)

	return horizontalChange, verticalChange
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
