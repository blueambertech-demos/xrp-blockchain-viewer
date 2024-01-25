package httputil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// PostJson takes a URL, content type and object. It serialises the object into json then posts it. Any response is
// read into a byte slice and returned
func PostJson(url string, obj interface{}) ([]byte, error) {
	body, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Non-OK status received: %d", resp.StatusCode)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
