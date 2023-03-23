package encrypt_date_test

import (
	"testing"

	"github.com/supernarsi/gotool/encrypt/encrypt_date"
)

func TestDateStringEncrypt(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "test date", in: "2000-12-", want: ""},
		{name: "test date", in: "20001201", want: ""},
		{name: "test date", in: "2000-12-31", want: "2ECCC1DE1FD"},
		{name: "test date", in: "2022-02-31", want: "2ECEE1CE1FD"},
		{name: "test date", in: "1990-01-01", want: "2DLLC1CD1CD"},
		{name: "test date", in: "1720-12-31", want: "2DJEC1DE1FD"},
		{name: "test date", in: "2023-03-03", want: "1DBDE3BE3BE"},
		{name: "test date", in: "2199-06-10", want: "1DCKK0BH0CB"},
		{name: "test date", in: "2199-06-20", want: "1DCKK0BH0DB"},
		{name: "test date", in: "2000-11-27", want: "2ECCC7DD7EJ"},
		{name: "test date", in: "2000-08-19", want: "1DBBB9BJ9CK"},
		{name: "test date", in: "1904-04-24", want: "2DLCG4CG4EG"},
	}
	for _, tt := range tests {
		t.Run(tt.name+" "+tt.in, func(t *testing.T) {
			t.Cleanup(func() {})
			if got := encrypt_date.DateStringEncrypt(tt.in); got != tt.want {
				t.Errorf("got result %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateDecrypt(t *testing.T) {
	type testS struct {
		name string
		in   string
		want string
	}
	tests := []testS{
		{name: "test date", in: "2ECEE1CE1FD", want: "2022-02-31"},
		{name: "test date", in: "0CACC3AC3BD", want: "2022-02-13"},
		{name: "test date", in: "1DBDD3BD3CE", want: "2022-02-13"},
		{name: "test date", in: "1DBDD1BD1EC", want: "2022-02-31"},
		{name: "test date", in: "0CAAA1BC1DB", want: "2000-12-31"},
		{name: "test date", in: "2DLCG4CG4EG", want: "1904-04-24"},
		{name: "test date", in: "2DLCG4CG4E^", want: ""},
		{name: "test date", in: "9DL]G4CG411", want: ""},
		{name: "test date", in: "0CAAA1BC1D", want: ""},
		{name: "test date", in: "11111111111", want: ""},
		{name: "test date", in: "234u82j", want: ""},
		{name: "test date", in: "", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {})
			if got := encrypt_date.DateDecrypt(tt.in); got != tt.want {
				t.Errorf("got result %v, want %v", got, tt.want)
			}
		})
	}
}
