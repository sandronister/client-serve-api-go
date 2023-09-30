package pricerepository

import (
	"fmt"
	"os"
	"time"
)

func (r *repository) appendFile(filename string, body []byte) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	current := time.Now()
	content := fmt.Sprintf("Data: %v - Cotacao: %v \n", current.Format("07-09-2006 15:06"), string(body))
	data = append(data, []byte(content)...)

	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(data)
	return nil
}
