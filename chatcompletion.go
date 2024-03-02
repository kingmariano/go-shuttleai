package shuttle
import (
	"encoding/json"
	"context"
	"errors"
)
type ChatRequest struct {
	//Required
	Messages []ChatMessage `json:"messages"`

	//The id of the model to use for the generation
	Model string `json:"model"`

	//Whether to include citations in the generation > For use in owned_by: bing/openai models only, defaults to true.
	Citations *bool `json:"citations,omitempty"`

	//The url of the image to use for the generation > Model limitations apply

	Image *string `json:"image,omitempty"`

	//Whether to use the internet for the generation > Model limitations apply, defaults to true for bing models.
	Internet *bool `json:"internet,omitempty"`

	//The maximum number of tokens to generate
	MaxToken *int `json:"max_token,omitempty"`

	//Whether to return the raw response > For use in owned_by: bing/openai models only! Returns Bing AI Suggestions and Search Results, defaults to false.
	Raw *bool `json:"raw,omitempty"`

	//The temperature of the sampling distribution
	Temperature *int `json:"temperature,omitempty"`

	//The tone of the generation > For use in owned_by: bing/openai models only, defaults to precise.
	Tone *string `json:"tone,omitempty"`

	//The cumulative probability of the top tokens to keep in the nucleus of the distribution
	TopP *int `json:"top_p,omitempty"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	Choices []struct {
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
		Logprobs     struct {
		} `json:"logprobs"`
		Message struct {
			Content string `json:"content"`
			Role    string `json:"role"`
		} `json:"message"`
	} `json:"choices"`
	Created           int    `json:"created"`
	ID                string `json:"id"`
	Model             string `json:"model"`
	Object            string `json:"object"`
	SystemFingerprint string `json:"system_fingerprint"`
	Usage             struct {
		CompletionTokens int `json:"completion_tokens"`
		PromptTokens     int `json:"prompt_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}
func (sh *ShuttleClient) ChatCompletion(ctx context.Context, req *ChatRequest) (*ChatResponse, error) {
	if len(req.Messages) == 0 {
		return nil, errors.New("input text is required")
	}
	if req.Model == "" {
		req.Model = "gpt-3.5-turbo-1106"
	}
	body, err := sh.post(ctx, "/v1/chat/completions", "application/json", req)
	if err != nil {
		return nil, err
	}
	chatResponse := ChatResponse{}
	if err := json.Unmarshal(body, &chatResponse); err != nil {
		return nil, err
	}

	return &chatResponse, nil
}

