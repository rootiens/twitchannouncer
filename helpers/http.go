package helpers

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func HttpRequest(body interface{}, headers HttpReq) (*http.Response, error) {
	parsedBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(headers.Method, headers.Url, bytes.NewBuffer(parsedBody))

	for _, header := range headers.Headers {
		req.Header.Add(header.Key, header.Value)
	}

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
