package shuttle

import (
	"context"
	"encoding/json"
)
//example of endpoint: /v1/images/generations
type ModelResponse []struct {
	Data []struct {
		Cost    int    `json:"cost"`
		Created int    `json:"created"`
		ID      string `json:"id"`
		Object  string `json:"object"`
		OwnedBy string `json:"owned_by"`
		Premium bool   `json:"premium"`
		Tokens  int    `json:"tokens"`
	} `json:"data"`
	Object string `json:"object"`
}

func (sh *ShuttleClient) GetModels() (*ModelResponse, error) {
	body, err := sh.post(context.Background(), "/v1/models", "application/json", nil)
	if err != nil {
		return nil, err
	}
	modelResponse := ModelResponse{}
	if err := json.Unmarshal(body, &modelResponse); err != nil {
		return nil, err
	}

	return &modelResponse, nil
}

func (sh *ShuttleClient) GetModelByEndpoint(endpoint string) (*ModelResponse, error) {
	body, err := sh.post(context.Background(), endpoint, "application/json", nil)
	if err != nil {
		return nil, err
	}
	modelResponse := ModelResponse{}
	if err := json.Unmarshal(body, &modelResponse); err != nil {
		return nil, err
	}

	return &modelResponse, nil
}
