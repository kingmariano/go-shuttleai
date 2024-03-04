package main
import (
	"context"
	"fmt"
    "os"
	shuttleai "github.com/charlesozo/go-shuttleai"
)

func main(){
    client := shuttleai.NewClient(os.getenv("OXYGEN_API_TOKEN"))
	response, err := client.TranscribeAudio(context.Background(), &shuttleai.AudioTranscribeRequest{
		File: []byte{"audio.mp3"},
	})
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(response.Text)
}