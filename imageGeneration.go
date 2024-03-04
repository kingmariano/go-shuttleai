package shuttle

import (
	"context"
	"encoding/json"
	"errors"
)

type ImageRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}
type ImageResponse struct {
	Created int `json:"created"`
	Data    []struct {
		URL string `json:"url"`
	} `json:"data"`
}

func (sh *ShuttleClient) ImageGeneration(ctx context.Context, req *ImageRequest) (*ImageResponse, error) {
	if len(req.Prompt) == 0 {
		return nil, errors.New("input text is required")
	}
	if req.Model == "" {
		req.Model = "turbovision-xl"
	}
	body, err := sh.post(ctx, "/v1/images/generations", "application/json", req)
	if err != nil {
		return nil, err
	}
	imageResponse := ImageResponse{}
	if err := json.Unmarshal(body, &imageResponse); err != nil {
		return nil, err
	}
	return &imageResponse, nil
}
