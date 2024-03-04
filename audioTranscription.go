package shuttle

import (
	"context"
	"encoding/json"
)

type AudioTranscribeRequest struct {
	File []byte `json:"file"`
}
type  AudioTranscribeResponse struct {
	Text string `json:"text"`
}
func (sh *ShuttleClient) TranscribeAudio(ctx context.Context, req *AudioTranscribeRequest) (*AudioTranscribeResponse, error){

	body, err := sh.post(ctx, "/v1/audio/transcriptions", "multipart/form-data", req.File)
	if err != nil {
		return nil, err
	}
	audioTranscribeResponse := AudioTranscribeResponse{}
	if err := json.Unmarshal(body, &audioTranscribeResponse); err != nil {
		return nil, err
	}
	return &audioTranscribeResponse, nil
}