package exchangers

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func DoRequest(ctx context.Context, client *http.Client, method string, hosts []string, query string, body io.Reader) (resp *http.Response, err error) {
	if len(hosts) == 0 {
		err = ErrUnavailable

		return
	}

	url := fmt.Sprintf("%s/%s", hosts[0], query)
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		err = fmt.Errorf("error make http request: %v", err)

		return
	}

	resp, err = client.Do(req)
	if err != nil || resp.StatusCode >= 500 {
		resp, err = DoRequest(ctx, client, method, hosts[1:], url, body)
	}

	return
}
