package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lexruntimeservice"
)

func main() {
	processVoiceCommand()
}

func processVoiceCommand() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(endpoints.UsEast1RegionID),
	}))
	svc := lexruntimeservice.New(sess)

	f, err := os.Open("PlayPawnD2.pcm")
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

	fmt.Println(resp)
}
