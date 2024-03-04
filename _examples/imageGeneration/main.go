package main

import (
	"context"
	"fmt"
    "os"
	shuttleai "github.com/charlesozo/go-shuttleai"
)

func main(){
    client := shuttleai.NewClient(os.getenv("OXYGEN_API_TOKEN"))
	response, err := client.ImageGeneration(context.Background(), &shuttleai.ImageRequest{
		Prompt: "two ladies holding and talking to each other while crossing the road",
	})
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(response.Data[0].URL)
}