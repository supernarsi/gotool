package gotool_test

import (
	"testing"
	"time"

	"github.com/supernarsi/gotool"
)

const (
	tzVN = "Asia/Ho_Chi_Minh"
	tzCN = "Asia/Shanghai"
	tzUS = "America/Los_Angeles"
)

var (
	locVN, _ = time.LoadLocation(tzVN)
	locCN, _ = time.LoadLocation(tzCN)
	locUS, _ = time.LoadLocation(tzUS)
)

func TestLocTimestamp(t *testing.T) {
	tests := []struct {
		name      string
		inputTime time.Time
		inputTz   string
		wantS     int64
		wantE     int64
	}{
		{"t1", time.Date(2024, 04, 8, 0, 0, 0, 0, locCN), tzCN, 1712505600, 1712591999},
		{"t2", time.Date(2000, 12, 31, 23, 59, 59, 0, locUS), tzCN, 978278400, 978364799},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if s, e := gotool.LocTimestamp(tt.inputTime, tt.inputTz); s != tt.wantS || e != tt.wantE {
				t.Errorf("LocTimestamp() = %v, %v, want %v, %v", s, e, tt.wantS, tt.wantE)
			}
		})
	}
}

func TestIsSameMonthDay(t *testing.T) {
	test := []struct {
		name          string
		inputDateStr  string
		inputTimezone string
		inputTime     time.Time
		want          bool
	}{
		{"t1", "2023-01-01", tzVN, time.Date(2023, 1, 1, 0, 0, 0, 0, locVN), true},
		{"t2", "2022-01-01", tzVN, time.Date(2021, 12, 31, 23, 59, 59, 0, locVN), false},
		{"t3", "1980-12-31", tzVN, time.Date(1980, 12, 31, 0, 0, 0, 0, locVN), true},
		{"t4", "2000-01-04", tzCN, time.Date(2000, 1, 4, 17, 59, 59, 0, locVN), true},
		{"t5", "2000-01-04", tzCN, time.Date(2000, 1, 4, 23, 59, 59, 0, locVN), false},
		{"t6", "2024-02-26", tzCN, time.Date(2024, 2, 26, 0, 0, 0, 0, locUS), true},
		{"t7", "2024-02-26", tzCN, time.Date(2024, 2, 26, 23, 0, 0, 0, locUS), false},
		{"t8", "2009-12-31", tzCN, time.Date(2009, 12, 31, 23, 59, 59, 0, locCN), true},
		{"t8", "2009-12-31", tzCN, time.Date(2009, 12, 31, 0, 0, 0, 0, locCN), true},
		{"t8", "2009-12-31", tzCN, time.Date(2009, 1, 1, 0, 0, 0, 0, locCN), false},
		{"t9", "2024-02-29", tzCN, time.Date(2023, 2, 28, 0, 0, 0, 0, locCN), false},
		{"t10", "2024-02-29", tzCN, time.Date(2020, 2, 29, 23, 59, 59, 0, locCN), true},
		{"t10", "2024-02-29", tzUS, time.Date(2020, 2, 28, 20, 59, 59, 0, locCN), false},
		{"t11", "2024-02-29", tzCN, time.Date(2020, 3, 1, 0, 59, 59, 0, locCN), false},
		{"t12", "2024-02-28", tzUS, time.Date(2023, 3, 1, 7, 59, 59, 0, locCN), true},
		{"t13", "2020-02-29", tzUS, time.Date(2024, 3, 1, 7, 59, 59, 0, locCN), true},
		{"t14", "2023-01-01", tzVN, time.Date(2000, 1, 1, 0, 0, 0, 0, locVN), true},
		{"t15", "1980-12-31", tzVN, time.Date(1993, 12, 31, 0, 0, 0, 0, locVN), true},
		{"t16", "2000-01-04", tzCN, time.Date(1999, 1, 4, 17, 59, 59, 0, locVN), true},
		{"t17", "2000-01-04", tzCN, time.Date(2020, 1, 4, 23, 59, 59, 0, locVN), false},
		{"t18", "2024-02-26", tzCN, time.Date(2030, 2, 26, 0, 0, 0, 0, locUS), true},
		{"t19", "2024-02-29", tzUS, time.Date(2023, 3, 1, 8, 0, 0, 0, locCN), false},
		{"t20", "2024-02-26", tzUS, time.Date(2024, 2, 26, 16, 0, 0, 0, locCN), true},
		{"t21", "2024-02-25", tzUS, time.Date(2024, 2, 26, 15, 23, 59, 0, locCN), true},
		{"t20", "2024-02-26", tzUS, time.Date(2024, 2, 26, 16, 0, 0, 1, locCN), true},
	}

	for _, v := range test {
		t.Run(v.name, func(t *testing.T) {
			if got := gotool.IsSameMonthDay(v.inputDateStr, v.inputTimezone, v.inputTime, false); got != v.want {
				t.Errorf("%s want %v, got %v", v.name, v.want, got)
			}
		})
	}
}

func TestTimeToYmdInt(t *testing.T) {
	tests := []struct {
		name  string
		input time.Time
		want  int
	}{
		{"t1", time.Date(2023, 1, 1, 0, 0, 0, 0, locUS), 20230101},
		{"t2", time.Date(2000, 12, 1, 23, 59, 0, 0, locUS), 20001201},
		{"t3", time.Date(1990, 2, 29, 23, 59, 0, 0, locUS), 19900301},
		{"t4", time.Date(1989, 2, 28, 0, 59, 0, 0, locCN), 19890228},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gotool.TimeToYmdInt(tt.input); got != tt.want {
				t.Errorf("TimeToYmdInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
