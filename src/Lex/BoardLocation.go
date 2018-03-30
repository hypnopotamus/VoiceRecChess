package main

import (
	"unicode"
)

type BoardLocation struct {
	Letter rune
	Number int
}

func NewBoardLocation(letter rune, number int) *BoardLocation {
	if number < 1 || number > 8 {
		return nil
	}

	letter = unicode.ToUpper(letter)
	if letter < 'A' || letter > 'H' {
		return nil
	}

	bl := new(BoardLocation)
	bl.Letter = letter
	bl.Number = number
	return bl
}

// func main(){
// 	x := NewBoardLocation('a',8)

// 	if(x == nil) {
// 		fmt.Println("Invalid Location")
// 	} else {
// 		fmt.Println(string(x.Letter))
// 		fmt.Println(x.Number)
// 	}
// }
