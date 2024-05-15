package gotool

import (
	"errors"
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

func IsSameMonthDay(dateStr, timezone string, datetime time.Time, excludeSameYear bool) bool {
	date, locTime, err := parseTimeToSameLoc(dateStr, datetime, timezone)
	if err != nil {
		return false
	}

	// 比较月份和日期
	isSameMonth := locTime.Month() == date.Month()
	isSameDay := locTime.Day() == date.Day()

	yearOk := true
	if excludeSameYear {
		yearOk = locTime.Year() != date.Year()
	}

	return isSameMonth && isSameDay && yearOk
}

func TimeIsAfterDateEnd(dateStr, timezone string, datetime time.Time) bool {
	date, locTime, err := parseTimeToSameLoc(dateStr, datetime, timezone)
	if err != nil {
		return false
	}
	return locTime.After(date.Add(24*time.Hour - 1*time.Second))
}

func TimeIsBeforeDateBegin(dateStr, timezone string, datetime time.Time) bool {
	date, locTime, err := parseTimeToSameLoc(dateStr, datetime, timezone)
	if err != nil {
		return false
	}
	return locTime.Before(*date)
}

func TimeToStamp(dateTime *time.Time, timezone string) (int64, error) {
	local, err := time.LoadLocation(timezone)
	if err != nil {
		return 0, err
	}
	t := time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), dateTime.Hour(), dateTime.Minute(), dateTime.Second(), 0, local)
	return t.Unix(), nil
}

func GetUTCOffset(timezone string) (int, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return 0, err
	}

	now := time.Now().In(loc)
	_, offset := now.Zone()

	return offset / 3600, nil
}

func parseTimeToSameLoc(dateStr string, datetime time.Time, timezone string) (time1, time2 *time.Time, err error) {
	// 解析时区
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return nil, nil, errors.New("")
	}

	// 解析给定的日期字符串
	date, err := time.ParseInLocation("2006-01-02", dateStr, loc)
	if err != nil {
		return nil, nil, errors.New("dateStr is invalidate")
	}
	// 获取时间戳在指定时区的时间
	locTime := datetime.In(loc)

	return &date, &locTime, nil
}
