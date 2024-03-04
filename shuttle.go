package shuttle
import (
	"net/http"
	"fmt"
	"encoding/json"
	"context"
	"bytes"
	"io"
)

const APIBaseURL = "https://api.shuttleai.app"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type ShuttleClient struct {
	Httpclient HTTPClient
	ApiKey      string
	Baseurl    string
}
type ErrorResponse struct {
	Error string `json:"error"`
}

func NewClient(apikey string) *ShuttleClient {
	return &ShuttleClient{
		Httpclient: http.DefaultClient,
		ApiKey:     apikey,
		Baseurl:    APIBaseURL,
	}
}

func (sh *ShuttleClient) post(ctx context.Context, task string, contentType string, payload interface{}) ([]byte, error) {
	url := sh.resolveURL(task)
	var body io.Reader

	switch v := payload.(type) {
	case []byte:
		body = bytes.NewReader(v)
	default:
		jsonBody, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	req.Header.Set("Accept", "application/json")
	if sh.ApiKey != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", sh.ApiKey))
	}

	res, err := sh.Httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		errResp := ErrorResponse{}
		if err := json.Unmarshal(resBody, &errResp); err != nil {
			return nil, fmt.Errorf("shuttleAI error: %s", resBody)
		}

		return nil, fmt.Errorf("shuttleAI error: %s", errResp.Error)
	}

	return resBody, nil
}


func (oc *ShuttleClient) resolveURL(task string) string {

	return fmt.Sprintf("%s/%s", oc.Baseurl, task)
}