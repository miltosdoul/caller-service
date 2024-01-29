package client

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

type CounterClient struct {
	url string
}

func NewCounterClient() *CounterClient {
	counterServiceUrl := os.Getenv("COUNTER_SERVICE_URL")

	if counterServiceUrl == "" {
		log.Fatalln(errors.New("counter service URL was not found"))
	}

	return &CounterClient{
		url: counterServiceUrl,
	}
}

func (c *CounterClient) GetRequestCount() string {
	resp, err := http.Get(c.url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}
