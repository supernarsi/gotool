package encrypt_test

import (
	"testing"

	"github.com/supernarsi/gotool/encrypt"
)

func TestIpEncrypt(t *testing.T) {
	type nT struct {
		in   string
		want string
		err  error
	}
	tests := []nT{
		{"none", "", encrypt.ErrIPBeyondIPv4},
		{"127.0.0.1", "724e6433d351713a", nil},
		{"127.0.0.2", "724e6433d351713b", nil},
		{"127.0.0.3", "724e6433d351713c", nil},
		{"91.24.21.3", "724e6433af69863c", nil},
		{"113.250.150.170", "724e6433c64c07e3", nil},
		{"113.248.169.17", "724e6433c64a1a4a", nil},
		{"113.250.150.168", "724e6433c64c07e1", nil},
		{"125.80.202.253", "724e6433d1a23c36", nil},
		{"255.255.255.255", "724e643454517138", nil},
		{"255.255.255.254", "724e643454517137", nil},
		{"255.255.255.250", "724e643454517133", nil},
		{"255.255.255.999", "", encrypt.ErrIPBeyondIPv4},
		{"101.23.31.999", "", encrypt.ErrIPBeyondIPv4},
		{"91.202.11.64", "724e6433b01b7c79", nil},
		{"2.0.0.0", "724e643356517139", nil},
		{"0.0.0.0", "724e643354517139", nil},
		{"0.0.0.1", "724e64335451713a", nil},
		{"0.0.0.1000", "", encrypt.ErrIPBeyondIPv4},
	}
	for _, tt := range tests {
		t.Run("test for "+tt.in, func(t *testing.T) {
			t.Cleanup(func() {})
			if got, err := encrypt.NewIPEncrypter().Encrypt(tt.in); err != nil {
				if tt.err == nil || err.Error() != tt.err.Error() {
					t.Error("got err", err)
				}
				if tt.want != "" {
					t.Error("want should be empty, got", tt.want)
				}
			} else if got != tt.want {
				t.Errorf("got result %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIpDecrypt2(t *testing.T) {
	type nT struct {
		want  string
		input string
		err   error
	}
	tests := []nT{
		{"", "", nil},
		{"127.0.0.1", "724e6433d351713a", nil},
		{"127.0.0.2", "724e6433d351713b", nil},
		{"127.0.0.3", "724e6433d351713c", nil},
		{"91.24.21.3", "724e6433af69863c", nil},
		{"255.255.255.255", "724e643454517138", nil},
		{"255.255.255.254", "724e643454517137", nil},
		{"255.255.255.250", "724e643454517133", nil},
		{"91.202.11.64", "724e6433b01b7c79", nil},
		{"2.0.0.0", "724e643356517139", nil},
		{"0.0.0.0", "724e643354517139", nil},
		{"0.0.0.1", "724e64335451713a", nil},
		{"", "123", encrypt.ErrInvalidIPHex},
		{"", "dafkfo1h23", encrypt.ErrInvalidIPHex},
		{"", " ", encrypt.ErrInvalidIPHex},
	}
	for _, tt := range tests {
		t.Run("test for "+tt.input, func(t *testing.T) {
			t.Cleanup(func() {})
			if got, err := encrypt.NewIPEncrypter().Decrypt(tt.input); err != nil {
				if tt.err == nil || err.Error() != tt.err.Error() {
					t.Error("got err", err)
				}
				if tt.want != "" {
					t.Error("want should be empty, got", tt.want)
				}
			} else if got != tt.want || tt.err != nil {
				t.Errorf("got result %v, want %v", got, tt.want)
			}
		})
	}
}
