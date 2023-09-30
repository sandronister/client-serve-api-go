package pricerepository

import (
	"os"
)

func (r *repository) writeFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	_, err = file.Write([]byte(content))
	defer file.Close()
	if err != nil {
		return err
	}
	return nil

}
