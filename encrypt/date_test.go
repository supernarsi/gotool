package encrypt_test

import (
	"testing"
	"time"

	"github.com/supernarsi/gotool/encrypt"
)

func TestDateEncrypt(t *testing.T) {
	tests := []struct {
		name string
		in   time.Time
		want string
	}{
		{name: "test date", in: time.Date(2002, 2, 13, 0, 0, 0, 0, time.Local), want: "1DBBD3BD3CE"},
		{name: "test date", in: time.Date(2000, 12, 31, 0, 0, 0, 0, time.Local), want: "2ECCC1DE1FD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {})
			if got := encrypt.DateEncrypt(tt.in); got != tt.want {
				t.Errorf("got result %v, want %v", got, tt.want)
			}
		})
	}
}
