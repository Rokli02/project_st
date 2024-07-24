package utils

import (
	"time"
)

func IsDateExpired(expireDate *string) bool {
	return expireDate != nil && *expireDate != "" && ToTime(*expireDate).Before(time.Now())
}
