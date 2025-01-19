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

func (u *utilityTime) LocTimestamp(date time.Time, timezone string) (sTimestamp, eTimestamp int64) {
	loc, dateInTimeZone := u.LocalTime(date, timezone)
	if loc == nil {
		return 0, 0
	}

	// 获取当天开始时间
	startOfDay := time.Date(dateInTimeZone.Year(), dateInTimeZone.Month(), dateInTimeZone.Day(), 0, 0, 0, 0, loc)

	// 获取当天结束时间
	endOfDay := time.Date(dateInTimeZone.Year(), dateInTimeZone.Month(), dateInTimeZone.Day(), 23, 59, 59, 999999999, loc)

	// 将时间转换为 Unix 时间戳
	sTimestamp = startOfDay.Unix()
	eTimestamp = endOfDay.Unix()

	return sTimestamp, eTimestamp
}

func (u *utilityTime) GetUTCOffset(timezone string) (int, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return 0, err
	}

	now := time.Now().In(loc)
	_, offset := now.Zone()

	return offset / 3600, nil
}

func (u *utilityTime) LocalTime(date time.Time, timezone string) (*time.Location, time.Time) {
	// 通过时区字符串加载时区
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return nil, date
	}
	// 将时间转换为指定时区的时间
	return loc, date.In(loc)
}

func (u *utilityTime) TimeToYmdInt(localTime time.Time) int {
	// 将日期格式化为数字型
	formattedTime := localTime.Format("20060102")
	numericTime, err := time.Parse("20060102", formattedTime)
	if err != nil {
		return 0
	}
	return numericTime.Year()*10000 + int(numericTime.Month())*100 + numericTime.Day()
}

func (u *utilityTime) IsSameMonthDay(dateStr, timezone string, datetime time.Time, excludeSameYear bool) bool {
	date, locTime, err := u.parseTimeToSameLoc(dateStr, datetime, timezone)
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

func (u *utilityTime) TimeIsAfterDateEnd(dateStr, timezone string, datetime time.Time) bool {
	date, locTime, err := u.parseTimeToSameLoc(dateStr, datetime, timezone)
	if err != nil {
		return false
	}
	return locTime.After(date.Add(24*time.Hour - 1*time.Second))
}

func (u *utilityTime) TimeIsBeforeDateBegin(dateStr, timezone string, datetime time.Time) bool {
	date, locTime, err := u.parseTimeToSameLoc(dateStr, datetime, timezone)
	if err != nil {
		return false
	}
	return locTime.Before(*date)
}

func (u *utilityTime) CheckTimeIsInPeriod(checkTime int64, momentMin uint, momentMax int, now int64) bool {
	// 如果限定了展示时机，判断是否符合展示时机
	if momentMin < 0 || (momentMax >= 0 && momentMax < int(momentMin)) {
		return false
	}
	// 注册时间需要在 showSTime ~ showETime 范围内
	showSTime := checkTime + int64(momentMin)*86400
	var showETime int64 = math.MaxInt64
	if momentMax >= 0 {
		showETime = checkTime + int64(momentMax)*86400
	}
	if now < showSTime || now > showETime {
		return false
	}
	return true
}

func (u *utilityTime) GetDiffDays(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), t1.Hour(), t1.Minute(), t1.Second(), 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), t2.Hour(), t2.Minute(), t2.Second(), 0, time.Local)
	return int(t1.Sub(t2).Hours() / 24)
}

func (u *utilityTime) GetDiffDaysBySecond(t1, t2 int64) int {
	time1 := time.Unix(t1, 0)
	time2 := time.Unix(t2, 0)
	// 调用上面的函数
	return u.GetDiffDays(time1, time2)
}

func (u *utilityTime) GetTimeWithZone(dateStr, timezone string) *time.Time {
	// 解析时区
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return nil
	}
	// 解析给定的日期字符串
	date, err := time.ParseInLocation("2006-01-02", dateStr, loc)
	if err != nil {
		return nil
	}
	return &date
}

func (u *utilityTime) GetDateTimeWithZone(dateTimeStr string, timezone string) *time.Time {
	// 解析时区
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return nil
	}
	// 解析给定的日期字符串
	date, err := time.ParseInLocation("2006-01-02 15:04:05", dateTimeStr, loc)
	if err != nil {
		return nil
	}
	return &date
}

func (u *utilityTime) T1IsEarlier(clockTime1, clockTime2 string) bool {
	return u.DurationSecondsInClocks(clockTime1, clockTime2) > 0
}

// DurationSecondsInClockwise 计算 time1、time2 两个时钟间按顺时针相差的秒数（time1 晚于 time2 时为跨天情况，返回整数）
func (u *utilityTime) DurationSecondsInClockwise(time1, time2 string) int {
	durationInSeconds := u.DurationSecondsInClocks(time1, time2)
	if durationInSeconds < 0 {
		durationInSeconds = 86400 + durationInSeconds
	}
	return durationInSeconds
}

// DurationSecondsInClocks 计算 time1、time2 两个时钟间相差的秒数（time1 晚于 time2 时返回负数）
func (u *utilityTime) DurationSecondsInClocks(time1, time2 string) int {
	const timeFormat = "15:04:05" // 使用 24 小时制 "HH:MM:SS" 格式
	// 解析时间
	t1, err := time.Parse(timeFormat, time1)
	if err != nil {
		return 0
	}
	t2, err := time.Parse(timeFormat, time2)
	if err != nil {
		return 0
	}
	// 计算时间差，返回的是秒数
	durationInSeconds := int(t2.Sub(t1).Seconds())
	return durationInSeconds
}

// LocalDailyEndTime 当地时区当天结束时间戳
func (u *utilityTime) LocalDailyEndTime(now time.Time, timezone string) int64 {
	// 当前时区该日的起止时间戳
	loc, localTime := u.LocalTime(now, timezone)
	eTimestamp := time.Date(localTime.Year(), localTime.Month(), localTime.Day(), 23, 59, 59, 999999999, loc).Unix()
	return eTimestamp
}

// DailyLeftTime 距离当地时区当天结束时间还剩多少秒
func (u *utilityTime) DailyLeftTime(now time.Time, timezone string) int {
	eTimestamp := u.LocalDailyEndTime(now, timezone)
	leftTime := int(eTimestamp - now.Unix())
	if leftTime < 0 {
		leftTime = 0
	}
	return leftTime
}

func (u *utilityTime) TimeToStamp(dateTime *time.Time, timezone string) (int64, error) {
	local, err := time.LoadLocation(timezone)
	if err != nil {
		return 0, err
	}
	t := time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), dateTime.Hour(), dateTime.Minute(), dateTime.Second(), 0, local)
	return t.Unix(), nil
}

func (u *utilityTime) parseTimeToSameLoc(dateStr string, datetime time.Time, timezone string) (time1, time2 *time.Time, err error) {
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
