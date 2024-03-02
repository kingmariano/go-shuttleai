package shuttle

import (
	"bytes"
	"context"
	"mime/multipart"
	"fmt"
	"io"
	"encoding/json"
)

type AudioTranscribeRequest struct {
	Audio []byte 
}
type  AudioTranscribeResponse struct {
	Text string `json:"text"`
}
func (sh *ShuttleClient) TranscribeAudio(ctx context.Context, req *AudioTranscribeRequest) (*AudioTranscribeResponse, error){
	var requestBody bytes.Buffer
    writer := multipart.NewWriter(&requestBody)
	part, err := writer.CreateFormFile("file", "audio.mp3")
    if err != nil {
        return nil,  fmt.Errorf("Error creating form file: %v", err)
    }
	if _, err := io.Copy(part, bytes.NewReader(req.Audio)); err != nil {
        return nil, fmt.Errorf("Error copying file: %v", err)
    }
	if err := writer.Close(); err != nil {
        return nil, fmt.Errorf("Error closing writer: %v", err)
    }
	body, err := sh.post(ctx,"/v1/audio/transcriptions", "multipart/form-data", &requestBody)
	if err != nil{
		return nil, err
	}
	audioTranscribeResponse := AudioTranscribeResponse{}
	if err := json.Unmarshal(body, &audioTranscribeResponse); err != nil {
		return nil, err
	}
	return &audioTranscribeResponse, nil
}