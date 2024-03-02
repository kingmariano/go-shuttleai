package shuttle
import (
	"context"
	"encoding/json"
	"errors"
)
type AudioGenRequest struct {
	//required
	Input string `json:"input"`

	//required
	Model string `json:"model"`

	//The voice to use for the generation > Voices can be found https://api.shuttleai.app/v1/models
	Voice *string `json:"voice,omitempty"`
}

type AudioGenResponse struct {
	Chars int `json:"chars"`
	Data  []struct {
		URL string `json:"url"`
	} `json:"data"`
	ExpiresIn int    `json:"expiresIn"`
	Model     string `json:"model"`
}

func (sh *ShuttleClient) AudioGeneration(ctx context.Context, req *AudioGenRequest) (*AudioGenResponse, error) {
	if len(req.Input) == 0 {
		return nil, errors.New("input text is required")
	}
	if req.Model == "" {
		req.Model = "eleven-labs-2"
	}
	body, err := sh.post(ctx, "/v1/audio/generations", "application/json", req)
	if err != nil {
		return nil, err
	}
	audioGenResponse := AudioGenResponse{}
	if err := json.Unmarshal(body, &audioGenResponse); err != nil {
		return nil, err
	}

	return &audioGenResponse, nil
}