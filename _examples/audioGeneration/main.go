package main
import (
	"context"
	"fmt"
    "os"
	shuttleai "github.com/charlesozo/go-shuttleai"
)

func main(){
    client := shuttleai.NewClient(os.getenv("OXYGEN_API_TOKEN"))
	response, err := client.AudioGeneration(context.Background(), &shuttleai.AudioGenRequest{
		Input: "I have been running around till she called me",
		Model: "eleven-labs-2",
		Voice: "dave",
	})
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(response.Data[0].URL)
}	