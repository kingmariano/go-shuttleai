# ⚗️ go-shuttleai
![Build Status](https://github.com/hupe1980/go-huggingface/workflows/build/badge.svg) 
[![Go Reference](https://pkg.go.dev/badge/github.com/hupe1980/go-huggingface.svg)](https://pkg.go.dev/github.com/charlesozo/go-oxygenai)
> The ShuttleAI  Client in Golang is a modul designed to interact with the openai model repository for free and perform inference tasks using state-of-the-art natural language processing models. Developed in Golang, it provides a seamless and efficient way to integrate openai  models into your Golang applications.

## Installation
```
go get github.com/charlesozo/go-shuttleai
```
## Shuttle AI documentation
> Visit [Discord server](https://discord.gg/shuttleai "Visit ShuttleAi discord server") to get the $OXYGEN_API_KEY and  Run `/getkey` command in the Discord server.

___[shuttle api docs](https://docs.shuttleai.app/ "Visit ShuttleAi")___

## How to use
```golang
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
```
Output:
```text
response goes here
```

For more example usage, see [_examples](./_examples).

## License
[MIT](LICENCE)