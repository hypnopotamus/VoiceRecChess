package lex

type MoveRequest struct {
	Piece       string
	Source      *LocationRequest
	Destination *LocationRequest
}

func NewMoveRequest(piece string, sourceLetter string, sourceNumber int,
	destinationLetter string, destinationNumber int) *MoveRequest {
	mr := new(MoveRequest)

	mr.Piece = piece
	mr.Source = NewLocationRequest(rune(sourceLetter[0]), sourceNumber)
	mr.Destination = NewLocationRequest(rune(destinationLetter[0]), destinationNumber)
	return mr
}

func (mr MoveRequest) ToString() string {
	return mr.Piece + ":" + mr.Source.toString() + ":" + mr.Destination.toString()
}

// func main() {
// 	x := NewMoveRequest("a2", "b3")

// 	fmt.Println(string(x.Source.Letter))
// 	fmt.Println(x.Source.Number)
// 	fmt.Println(string(x.Destination.Letter))
// 	fmt.Println(x.Destination.Number)
// }
