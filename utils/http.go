// Copyright 2022 The mammoth Authors

// Package utils implement utils
package utils

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

// HttpRequest ...
func HttpRequest(httpMethod string, url string, data []byte) ([]byte, error) {
	client := &http.Client{}
	var err error
	request := new(http.Request)
	switch httpMethod {
	case http.MethodGet:
		request, err = http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}
	case http.MethodPost:
		request, err = http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
		if err != nil {
			return nil, err
		}
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if body == nil {
		return nil, errors.New("response body is nil ")
	}

	return body, nil
}
