package pricerepository

import (
	"errors"
	"os"
)

func (r *repository) fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !errors.Is(err, os.ErrNotExist)
}
