package priceusecase

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/gofor-little/env"
	"sandroni.fullcycle.server/core/domain"
)

func getRequest() (*http.Request, error) {
	url := env.Get("URL", "FAIL")

	if url == "FAIL" {
		return nil, errors.New("URL does not defined")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func Get() (*domain.Price, error) {
	result := domain.Price{}
	req, err := getRequest()

	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
