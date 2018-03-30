package lex

import (
	"strconv"
	"unicode"
)

type LocationRequest struct {
	Letter rune
	Number int
}

func (lr LocationRequest) toString() string {
	return string(lr.Letter) + strconv.Itoa(lr.Number)
}

func NewLocationRequest(letter rune, number int) *LocationRequest {
	if number < 1 || number > 8 {
		return nil
	}

	letter = unicode.ToUpper(letter)
	if letter < 'A' || letter > 'H' {
		return nil
	}

	bl := new(LocationRequest)
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
