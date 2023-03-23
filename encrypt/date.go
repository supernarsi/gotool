package encrypt

import (
	"time"

	"github.com/supernarsi/gotool/encrypt/encrypt_date"
)

func DateEncrypt(date time.Time) string {
	dateString := date.Format("2006-01-02")
	return DateStringEncrypt(dateString)
}

func DateStringEncrypt(date string) string {
	return encrypt_date.DateStringEncrypt(date)
}

func DateDecrypt(encryptStr string) string {
	return encrypt_date.DateDecrypt(encryptStr)
}
