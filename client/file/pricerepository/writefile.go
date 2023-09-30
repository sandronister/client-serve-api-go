package pricerepository

import (
	"fmt"
	"os"
	"time"
)

func (r *repository) writeFile(filename string, body []byte) error {
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
