package priceusecase

import (
	"fmt"
	"time"
)

func (u *usecase) formatContent(body []byte) string {
	current := time.Now()
	return fmt.Sprintf("{Data: %v - Dólar: %v }\n", current.Format("07-09-2006 15:06"), string(body))
}
