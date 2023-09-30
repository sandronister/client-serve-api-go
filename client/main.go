package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !errors.Is(err, os.ErrNotExist)
}

func appendFile(filename string, body []byte) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	content := fmt.Sprintf("Data: %v - Cotacao: %v \n", time.Now(), string(body))
	data = append(data, []byte(content)...)

	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(data)
	return nil
}

func writeFile(filename string, body []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	_, err = file.Write([]byte(fmt.Sprintf("Data: %v - Cotacao: %v \n", time.Now(), string(body))))
	defer file.Close()
	if err != nil {
		return err
	}
	return nil

}

func saveFile(filename string, content []byte) error {
	if fileExists(filename) {
		return appendFile(filename, content)
	}

	return writeFile(filename, content)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	saveFile("cotacao.txt", body)
	defer res.Body.Close()

}
