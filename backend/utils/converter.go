package utils

import (
	"regexp"
	"st/backend/logger"
	"st/backend/settings"
	"strconv"
	"strings"
	"time"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// Converts a 'string' date (that is in `yyyy[?]MM[?]dd[?]hh[?]mm[?]ss`) to Time struct.
//
// In place of [?] any character can be used.
func ToTime(date string) *time.Time {
	if len(date) != len(settings.Database.DateFormat) {
		logger.WarningF("[%s] doesn't match format: %s", date, settings.Database.DateFormat)

		return nil
	}

	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:10])
	hour, _ := strconv.Atoi(date[11:13])
	min, _ := strconv.Atoi(date[14:16])
	sec, _ := strconv.Atoi(date[17:])

	timeDate := time.Date(year, time.Month(month), day, hour, min, sec, 0, time.Local)

	return &timeDate
}

func ToRef[T any](param T) *T {
	return &param
}
