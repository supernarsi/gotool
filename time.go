package gotool

import (
	"time"
)

func LocTimestamp(date time.Time, timezone string) (sTimestamp, eTimestamp int64) {
	loc, dateInTimeZone := LocalTime(date, timezone)
	if loc == nil {
		return 0, 0
	}

	startOfDay := time.Date(dateInTimeZone.Year(), dateInTimeZone.Month(), dateInTimeZone.Day(), 0, 0, 0, 0, loc)
	endOfDay := time.Date(dateInTimeZone.Year(), dateInTimeZone.Month(), dateInTimeZone.Day(), 23, 59, 59, 999999999, loc)
	return startOfDay.Unix(), endOfDay.Unix()
}

func LocalTime(date time.Time, timezone string) (*time.Location, time.Time) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return nil, date
	}
	return loc, date.In(loc)
}

func TimeToYmdInt(localTime time.Time) int {
	formattedTime := localTime.Format("20060102")
	numericTime, err := time.Parse("20060102", formattedTime)
	if err != nil {
		return 0
	}
	return numericTime.Year()*10000 + int(numericTime.Month())*100 + numericTime.Day()
}
