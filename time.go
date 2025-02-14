package gotool

import (
	"errors"
	"math"
	"time"
	_ "time/tzdata" // use built-in timezone database
)

type utilityTime struct{}

func GetUtilityTime() *utilityTime {
	return &utilityTime{}
}

func (u *utilityTime) LocTimestamp(date time.Time, timezone string) (int64, int64) {
	loc, dateInLoc := u.LocalTime(date, timezone)
	if loc == nil {
		return 0, 0
	}
	startOfDay := time.Date(dateInLoc.Year(), dateInLoc.Month(), dateInLoc.Day(), 0, 0, 0, 0, loc)
	endOfDay := startOfDay.Add(24*time.Hour - time.Nanosecond)
	return startOfDay.Unix(), endOfDay.Unix()
}

func (u *utilityTime) GetUTCOffset(timezone string) (int, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return 0, err
	}

	_, offset := time.Now().In(loc).Zone()
	return offset / 3600, nil
}

func (u *utilityTime) LocalTime(date time.Time, timezone string) (*time.Location, time.Time) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return nil, date
	}
	return loc, date.In(loc)
}

func (u *utilityTime) TimeToYmdInt(localTime time.Time) int {
	return localTime.Year()*10000 + int(localTime.Month())*100 + localTime.Day()
}

func (u *utilityTime) IsSameMonthDay(dateStr, timezone string, datetime time.Time, excludeSameYear bool) bool {
	date, locTime, err := u.parseTimeToSameLoc(dateStr, datetime, timezone)
	if err != nil {
		return false
	}
	return locTime.Month() == date.Month() && locTime.Day() == date.Day() && (!excludeSameYear || locTime.Year() != date.Year())
}

func (u *utilityTime) TimeIsAfterDateEnd(dateStr, timezone string, datetime time.Time) bool {
	date, locTime, err := u.parseTimeToSameLoc(dateStr, datetime, timezone)
	return err == nil && locTime.After(date.Add(24*time.Hour-time.Second))
}

func (u *utilityTime) TimeIsBeforeDateBegin(dateStr, timezone string, datetime time.Time) bool {
	date, locTime, err := u.parseTimeToSameLoc(dateStr, datetime, timezone)
	return err == nil && locTime.Before(*date)
}

func (u *utilityTime) CheckTimeIsInPeriod(checkTime int64, momentMin, momentMax int, now int64) bool {
	if momentMin < 0 || (momentMax >= 0 && momentMax < momentMin) {
		return false
	}
	showSTime := checkTime + int64(momentMin)*86400
	showETime := int64(math.MaxInt64)
	if momentMax >= 0 {
		showETime = checkTime + int64(momentMax)*86400
	}
	return now >= showSTime && now <= showETime
}

func (u *utilityTime) GetDiffDays(t1, t2 time.Time) int {
	t1, t2 = t1.Truncate(24*time.Hour), t2.Truncate(24*time.Hour)
	return int(t1.Sub(t2).Hours() / 24)
}

func (u *utilityTime) GetDiffDaysBySecond(t1, t2 int64) int {
	return u.GetDiffDays(time.Unix(t1, 0), time.Unix(t2, 0))
}

func (u *utilityTime) GetTimeWithZone(dateStr, timezone string) *time.Time {
	return u.parseTimeInLocation("2006-01-02", dateStr, timezone)
}

func (u *utilityTime) GetDateTimeWithZone(dateTimeStr, timezone string) *time.Time {
	return u.parseTimeInLocation("2006-01-02 15:04:05", dateTimeStr, timezone)
}

func (u *utilityTime) T1IsEarlier(clockTime1, clockTime2 string) bool {
	return u.DurationSecondsInClocks(clockTime1, clockTime2) > 0
}

// DurationSecondsInClockwise 计算 time1、time2 两个时钟间按顺时针相差的秒数（time1 晚于 time2 时为跨天情况，返回整数）
func (u *utilityTime) DurationSecondsInClockwise(time1, time2 string) int {
	duration := u.DurationSecondsInClocks(time1, time2)
	if duration < 0 {
		return 86400 + duration
	}
	return duration
}

// DurationSecondsInClocks 计算 time1、time2 两个时钟间相差的秒数（time1 晚于 time2 时返回负数）
func (u *utilityTime) DurationSecondsInClocks(time1, time2 string) int {
	const format = "15:04:05"
	t1, err1 := time.Parse(format, time1)
	t2, err2 := time.Parse(format, time2)
	if err1 != nil || err2 != nil {
		return 0
	}
	return int(t2.Sub(t1).Seconds())
}

// LocalDailyEndTime 当地时区当天结束时间戳
func (u *utilityTime) LocalDailyEndTime(now time.Time, timezone string) int64 {
	_, localTime := u.LocalTime(now, timezone)
	return time.Date(localTime.Year(), localTime.Month(), localTime.Day(), 23, 59, 59, 999999999, localTime.Location()).Unix()
}

// DailyLeftTime 距离当地时区当天结束时间还剩多少秒
func (u *utilityTime) DailyLeftTime(now time.Time, timezone string) int {
	left := int(u.LocalDailyEndTime(now, timezone) - now.Unix())
	if left < 0 {
		left = 0
	}
	return left
}

func (u *utilityTime) TimeToStamp(dateTime *time.Time, timezone string) (int64, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return 0, err
	}
	t := time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), dateTime.Hour(), dateTime.Minute(), dateTime.Second(), 0, loc)
	return t.Unix(), nil
}

func (u *utilityTime) parseTimeToSameLoc(dateStr string, datetime time.Time, timezone string) (*time.Time, *time.Time, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return nil, nil, errors.New("invalid timezone")
	}
	date, err := time.ParseInLocation("2006-01-02", dateStr, loc)
	if err != nil {
		return nil, nil, errors.New("invalid dateStr format")
	}
	locTime := datetime.In(loc)
	return &date, &locTime, nil
}

func (u *utilityTime) parseTimeInLocation(format, dateStr, timezone string) *time.Time {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return nil
	}
	t, err := time.ParseInLocation(format, dateStr, loc)
	if err != nil {
		return nil
	}
	return &t
}
