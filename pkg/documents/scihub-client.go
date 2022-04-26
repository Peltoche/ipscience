package documents

import (
	"fmt"
	"io"
	"net/http"
)

type ScihubClient struct {
}

func NewScihubClient() *ScihubClient {
	return &ScihubClient{}
}

func (c *ScihubClient) FetchDocument(doi string) ([]byte, error) {
	req, _ := http.NewRequest("GET", "https://www.sci-hub.st/"+doi, nil)

	req.Header.Set("User-Agent", "HTTPie/2.6.0")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call the remote server: %w", err)
	}

	defer res.Body.Close()
	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read the body")
	}

	return raw, nil
}
