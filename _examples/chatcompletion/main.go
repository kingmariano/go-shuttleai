package main

import (
	"context"
	"fmt"
    "os"
	shuttleai "github.com/charlesozo/go-shuttleai"
)

func main(){
    client := shuttleai.NewClient(os.getenv("OXYGEN_API_TOKEN"))
	response, err := client.ChatCompletion(context.Background(), &oxygenai.ChatRequest{
		Messages: []oxygenai.ChatMessage{
		  {
			Role: "user",
			Content: "write an essay on global warming",
		  },
		},
	})

	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(response.Choices[0].Message.Content)
}