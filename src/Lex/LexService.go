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
    sess := session.Must(session.NewSession(&aws.Config {
        Region: aws.String(endpoints.UsEast1RegionID),
    }))
    svc := lexruntimeservice.New(sess)

    f, err := os.Open("MovePawnD2toF2.pcm")
    resp, err := svc.PostContent(&lexruntimeservice.PostContentInput{
        BotAlias: aws.String("ChessBot"),
        BotName: aws.String("ChessBot"),
        UserId: aws.String("VoiceRecChess8003"),
        ContentType: aws.String("audio/x-l16; sample-rate=16000; channel-count=1"),
        InputStream: aws.ReadSeekCloser(f),
        

    })

    // resp, err := svc.PostText(&lexruntimeservice.PostTextInput{
    //     BotAlias: aws.String("ChessBot"),
    //     BotName: aws.String("ChessBot"),
    //     UserId: aws.String("VoiceRecChess"),
    //     InputText: aws.String("Move Pawn from B2 to C2"),
        
    // })

    fmt.Println(resp)    
    fmt.Println(err) 
}