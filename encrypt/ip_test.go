package encrypt_test

import (
	"strconv"
	"testing"

	"github.com/supernarsi/gotool/encrypt"
)

func TestIPString2Long(t *testing.T) {
	tests := []struct {
		in   string
		want uint
		err  error
	}{
		{"255.255.255.255", 4294967295, nil},
		{"0.0.0.0", 0, nil},
		{"0.0.0.0", 0, nil},
	}
	for _, tt := range tests {
		t.Run("test for "+tt.in, func(t *testing.T) {
			t.Cleanup(func() {})
			if got, err := encrypt.IPString2Long(tt.in); err != nil {
				t.Error("got err", err)
			} else if got != tt.want {
				t.Errorf("got result %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLong2IPString(t *testing.T) {
	type nT struct {
		in   uint
		want string
		err  error
	}
	tests := []nT{
		{100000029, "5.245.225.29", nil},
		{99999987, "5.245.224.243", nil},
		{0, "0.0.0.0", nil},
		{4294967295, "255.255.255.255", nil},
	}
	for _, tt := range tests {
		t.Run("test for "+strconv.Itoa(int(tt.in)), func(t *testing.T) {
			t.Cleanup(func() {})
			if got, err := encrypt.Long2IPString(tt.in); err != nil {
				t.Error("got err", err)
			} else if got != tt.want {
				t.Errorf("got result %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIpEncrypt(t *testing.T) {
	type nT struct {
		in   string
		want string
	}
	tests := []nT{
		{"", ""},
		{"127.0.0.1", "7f00002b"},
		{"127.0.0.2", "7f00002c"},
		{"127.0.0.3", "7f00002d"},
		{"91.24.21.3", "5b18152d"},
		{"255.255.255.255", "100000029"},
		{"255.255.255.254", "100000028"},
		{"255.255.255.250", "100000024"},
		{"255.255.255.999", ""},
		{"101.23.31.999", ""},
		{"91.202.11.64", "5bca0b6a"},
		{"2.0.0.0", "200002a"},
		{"0.0.0.0", "2a"},
		{"0.0.0.1", "2b"},
		{"0.0.0.1000", ""},
	}
	for _, tt := range tests {
		t.Run("test for "+tt.in, func(t *testing.T) {
			t.Cleanup(func() {})
			if got := encrypt.IpEncrypt(tt.in); got != tt.want {
				t.Errorf("got result %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIpDecrypt2(t *testing.T) {
	type nT struct {
		want  string
		input string
	}
	tests := []nT{
		{"", ""},
		{"127.0.0.1", "7f00002b"},
		{"127.0.0.2", "7f00002c"},
		{"127.0.0.3", "7f00002d"},
		{"91.24.21.3", "5b18152d"},
		{"91.202.11.64", "5bca0b6a"},
		{"2.0.0.0", "200002a"},
		{"0.0.0.0", "2a"},
		{"0.0.0.1", "2b"},
		{"255.255.255.255", "100000029"},
		{"255.255.255.254", "100000028"},
		{"255.255.255.250", "100000024"},
		{"0.0.0.249", "123"},
		{"0.0.10.146", "abc"},
		{"", "zzzabc"},
		{"", "a.c.d"},
	}
	for _, tt := range tests {
		t.Run("test for "+tt.input, func(t *testing.T) {
			t.Cleanup(func() {})
			if got := encrypt.IpDecrypt(tt.input); got != tt.want {
				t.Errorf("got result %v, want %v", got, tt.want)
			}
		})
	}
}
