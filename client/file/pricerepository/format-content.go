package pricerepository

import (
	"fmt"
	"time"
)

func (r *repository) formatContent(body []byte) string {
	current := time.Now()
	return fmt.Sprintf("Data: %v - Cotacao: %v \n", current.Format("07-09-2006 15:06"), string(body))
}
