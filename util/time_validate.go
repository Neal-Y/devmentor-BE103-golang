package util

import (
	"time"
)

func ConvertDateRangeToUTC(startDateStr, endDateStr, timezone string) (time.Time, time.Time, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	startDate, err := time.ParseInLocation("2006-01-02", startDateStr, loc)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	endDate, err := time.ParseInLocation("2006-01-02", endDateStr, loc)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	startDateUTC := startDate.UTC()
	endDateUTC := endDate.Add(23*time.Hour + 59*time.Minute + 59*time.Second).UTC()

	return startDateUTC, endDateUTC, nil
}
