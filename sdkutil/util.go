package sdkutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type BaseSDK struct {
	BaseURL    string
	HTTPClient *http.Client
}

func (baseSDK *BaseSDK) SendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := baseSDK.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errMsg string
		if err = json.NewDecoder(res.Body).Decode(&errMsg); err == nil {
			return errors.New(errMsg)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	return json.NewDecoder(res.Body).Decode(&v)
}
