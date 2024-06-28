package gotool_test

import (
	"testing"

	"github.com/supernarsi/gotool"
)

func TestRandomReadableUniCode6Len(t *testing.T) {
	t.Run("test RandomReadableUniCode6Len", func(t *testing.T) {
		t.Cleanup(func() {})
		code := gotool.RandomReadableUniCode6Len()
		if len(code) != 6 {
			t.Error("build rand code error, code is:", code)
		}
	})
}

func TestRandomReadableUniCode8Len(t *testing.T) {
	t.Run("test RandomReadableUniCode8Len", func(t *testing.T) {
		t.Cleanup(func() {})
		code := gotool.RandomReadableUniCode8Len()
		if len(code) != 8 {
			t.Error("build rand code error, code is:", code)
		}
	})
}

func TestRandomString(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{name: "test 1", n: 1, want: 0},
		{name: "test 2", n: 2, want: 2},
		{name: "test 3", n: 3, want: 2},
		{name: "test 4", n: 4, want: 4},
		{name: "test 5", n: 5, want: 4},
		{name: "test 6", n: 6, want: 6},
		{name: "test 7", n: 7, want: 6},
		{name: "test 8", n: 8, want: 8},
		{name: "test 9", n: 9, want: 8},
		{name: "test 10", n: 10, want: 10},
		{name: "test 11", n: 11, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {})
			if code := gotool.RandomString(tt.n); len(code) != tt.want {
				t.Error("build rand code error, code is:", code)
			}
		})
	}
}

func BenchmarkRandomUniCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gotool.RandomUniCode(true, true)
		gotool.RandomUniCode(true, false)
		gotool.RandomUniCode(false, true)
		gotool.RandomUniCode(false, false)
	}
}

func TestSampleGenerateCode(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		wantLen int
	}{
		{"t1", 6, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gotool.SampleGenerateCode(tt.input); len(got) != tt.wantLen {
				t.Errorf("SampleGenerateCode() = %v, want %v", got, tt.wantLen)
			}
		})
	}
}

func TestUniInvCodeLen6ByUID(t *testing.T) {
	baseChars32 := []byte{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
		'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R',
		'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	}
	baseChars62 := []byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
		'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
		'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
		'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
		'U', 'V', 'W', 'X', 'Y', 'Z',
	}
	tests := []struct {
		name  string
		input uint64
		want  string
	}{
		{name: "t1", input: 0, want: "rflqBR"},
		{name: "t2", input: 1, want: "uuxzHU"},
		{name: "t3", input: 10000, want: "jBP2tx"},
		{name: "t4", input: 1000000, want: "xJJVno"},
		{name: "t5", input: 1234567, want: "yON1mt"},
		{name: "t6", input: 9999999, want: "mQ3dj9"},
		{name: "t7", input: 10000001, want: "skrvvf"},
		{name: "t8", input: 12332112, want: "xJLd1o"},
		{name: "t9", input: 12332113, want: "AYXm7r"},
		{name: "t10", input: 12332114, want: "Dd9vdu"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gotool.UniInvCodeLen6ByUID(tt.input, baseChars62); got != tt.want {
				t.Errorf("UniInvCodeLen6ByUID() = %v, want %v", got, tt.want)
			}
		})
	}

	tests2 := []struct {
		name  string
		input uint64
		want  string
	}{
		{name: "t1", input: 0, want: "DSQZFP"},
		{name: "t2", input: 1, want: "GHCJMS"},
		{name: "t3", input: 10000, want: "DSQCKR"},
		{name: "t4", input: 1000000, want: "DSZAFX"},
		{name: "t5", input: 1234567, want: "ABPUXM"},
		{name: "t6", input: 9999999, want: "AEXBCU"},
		{name: "t7", input: 10000001, want: "GLXVQA"},
		{name: "t8", input: 12332112, want: "DWGNMH"},
		{name: "t9", input: 12332113, want: "GMUXTL"},
		{name: "t10", input: 12332114, want: "KCGGZP"},
	}

	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			if got := gotool.UniInvCodeLen6ByUID(tt.input, baseChars32); got != tt.want {
				t.Errorf("UniInvCodeLen6ByUID() = %v, want %v", got, tt.want)
			}
		})
	}
}
