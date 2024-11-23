package client

import (
	"io"
	"net/http"
)

func NewHttpClient() HttpClient {
	client := http.Client{}
	return &HttpClientImpl{
		client: client,
	}
}

type HttpClientImpl struct {
	client http.Client
}

func (impl *HttpClientImpl) Get(url string) (string, error) {
	res, err := impl.client.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
