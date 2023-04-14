package ninja

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/igilgyrg/arbitrage/use/integration/ninja/response"
)

func (c *client) CryptoSymbols(ctx context.Context) (symbols []string, err error) {
	strUrl := fmt.Sprintf("%s/v1/cryptosymbols", c.cfg.Endpoint)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, strUrl, http.NoBody)
	if err != nil {
		err = fmt.Errorf("ninja: failed to make request: %v", err)

		return
	}

	request.Header.Set("x-api-key", c.cfg.ApiKey)

	resp, err := c.httpClient.Do(request)
	if err != nil {
		err = fmt.Errorf("ninja: failed to do request: %v", err)

		return
	}

	if resp.StatusCode >= 400 {
		c.log.Errorf("ninja: response status %d %s", resp.StatusCode, resp.Status)

		return
	}

	var responseBody response.Symbols

	if err = json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		err = fmt.Errorf("ninja: failed to decode request: %v", err)

		return
	}

	if len(responseBody.Symbols) > 0 {
		c.log.Infof("ninja: number of crypto symbols: %d", len(responseBody.Symbols))
	}

	symbols = responseBody.Symbols

	return
}
