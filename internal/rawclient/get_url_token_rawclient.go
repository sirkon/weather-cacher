package rawclient

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// GetURLTokenRawClient конструктор
func GetURLTokenRawClient(formatter func(lat, lon float64) string, client *http.Client) RawClient {
	return &getURLTokenRawClient{
		urlFormatter: formatter,
		client:       client,
	}
}

type getURLTokenRawClient struct {
	urlFormatter func(lat, log float64) string
	client       *http.Client
}

// Get ...
func (rc *getURLTokenRawClient) Get(ctx context.Context, lat, log float64) (io.ReadCloser, error) {
	url := rc.urlFormatter(lat, log)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to formed HTTP GET request: %s", err)
	}
	req = req.WithContext(ctx)

	resp, err := rc.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve data via `%s`: %s", url, err)
	}

	if resp.StatusCode != http.StatusOK {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf("failed to retrieve data via `%s`: unexpected status code %d", url, resp.StatusCode)
	}

	return resp.Body, nil
}
