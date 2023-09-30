package pricerepository

import (
	"os"
)

func (r *repository) appendFile(filename string, content string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	data = append(data, []byte(content)...)

	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(data)
	return nil
}
