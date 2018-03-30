package lex

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lexruntimeservice"
)

func main() {
	moveIntent := processVoiceCommand("MovePawnD2toF2.pcm")
	processIntentResponse(moveIntent)
}

func processIntentResponse(intentResponse *lexruntimeservice.PostContentOutput) *MoveRequest {
	var piece string
	if intentResponse.Slots["Piece"] != nil {
		piece = intentResponse.Slots["Piece"].(string)
	}
	sourceLetter := string(intentResponse.Slots["SourceLetter"].(string)[0])
	sourceNumber := processIntentNumber(intentResponse.Slots["SourceNumber"].(string))

	destinationLetter := string(intentResponse.Slots["DestinationLetter"].(string)[0])
	destinationNumber := processIntentNumber(intentResponse.Slots["DestinationNumber"].(string))

	mr := NewMoveRequest(piece, sourceLetter, sourceNumber, destinationLetter, destinationNumber)
	return mr
}

func processIntentNumber(numberAsString string) int {
	switch numberAsString {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	default:
		return 0
	}
}

func processVoiceCommand(voiceFileName string) *lexruntimeservice.PostContentOutput {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(endpoints.UsEast1RegionID),
	}))
	svc := lexruntimeservice.New(sess)

	f, err := os.Open(voiceFileName)
	resp, err := svc.PostContent(&lexruntimeservice.PostContentInput{
		BotAlias:    aws.String("ChessBot"),
		BotName:     aws.String("ChessBot"),
		UserId:      aws.String("VoiceRecChess800177"),
		ContentType: aws.String("audio/x-l16; sample-rate=16000; channel-count=1"),
		InputStream: aws.ReadSeekCloser(f),
	})

	// resp, err := svc.PostText(&lexruntimeservice.PostTextInput{
	//     BotAlias: aws.String("ChessBot"),
	//     BotName: aws.String("ChessBot"),
	//     UserId: aws.String("VoiceRecChess"),
	//     InputText: aws.String("Move Pawn from B2 to C2"),

	// })

	if err != nil {
		panic(err)
	}

	return resp
}
