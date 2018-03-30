package main

import (
	"strconv"
)

type MoveRequest struct {
	Source      *BoardLocation
	Destination *BoardLocation
}

func NewMoveRequest(source string, destination string) *MoveRequest {
	mr := new(MoveRequest)

	sourceNumber, _ := strconv.Atoi(string(source[1]))
	mr.Source = NewBoardLocation(rune(source[0]), sourceNumber)

	destinationNumber, _ := strconv.Atoi(string(destination[1]))
	mr.Destination = NewBoardLocation(rune(destination[0]), destinationNumber)
	return mr
}

// func main() {
// 	x := NewMoveRequest("a2", "b3")

// 	fmt.Println(string(x.Source.Letter))
// 	fmt.Println(x.Source.Number)
// 	fmt.Println(string(x.Destination.Letter))
// 	fmt.Println(x.Destination.Number)
// }
