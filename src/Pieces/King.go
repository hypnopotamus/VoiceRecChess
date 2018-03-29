package Pieces

type King struct {
	Piece
}

func (king King) IsMoveValid(newPosition Position) bool {
	if king.Location.Horizontal-1 == newPosition.Horizontal && king.Location.Vertical == newPosition.Vertical {
		return true
	}
	if king.Location.Horizontal+1 == newPosition.Horizontal && king.Location.Vertical == newPosition.Vertical {
		return true
	}
	if king.Location.Horizontal == newPosition.Horizontal && king.Location.Vertical-1 == newPosition.Vertical {
		return true
	}
	if king.Location.Horizontal == newPosition.Horizontal && king.Location.Vertical+1 == newPosition.Vertical {
		return true
	}

	return false
}
