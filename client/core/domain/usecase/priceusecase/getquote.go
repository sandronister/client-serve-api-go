package priceusecase

import (
	"context"
	"io"
	"net/http"
)

func (u *usecase) GetQuote(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", nil
	}

	content := u.formatContent(body)
	return content, nil
}
