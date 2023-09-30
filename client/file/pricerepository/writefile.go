package pricerepository

import (
	"os"
)

func (r *repository) writeFile(filename string, body []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	content := r.formatContent(body)
	_, err = file.Write([]byte(content))
	defer file.Close()
	if err != nil {
		return err
	}
	return nil

}
