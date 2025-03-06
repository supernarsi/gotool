package gotool_test

import (
	"testing"

	"github.com/supernarsi/gotool"
)

func TestRandomReadableUniCode6Len(t *testing.T) {
	t.Run("test RandomReadableUniCode6Len", func(t *testing.T) {
		code := gotool.RandomReadableUniCode6Len()
		if len(code) != 6 {
			t.Error("build rand code error, code is:", code)
		}
	})
}

func TestRandomReadableUniCode8Len(t *testing.T) {
	t.Run("test RandomReadableUniCode8Len", func(t *testing.T) {
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
		{name: "test 1", n: 1, want: 1},
		{name: "test 2", n: 2, want: 2},
		{name: "test 3", n: 3, want: 3},
		{name: "test 4", n: 4, want: 4},
		{name: "test 5", n: 5, want: 5},
		{name: "test 6", n: 6, want: 6},
		{name: "test 7", n: 7, want: 7},
		{name: "test 8", n: 8, want: 8},
		{name: "test 9", n: 9, want: 9},
		{name: "test 10", n: 10, want: 10},
		{name: "test 11", n: 11, want: 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code := gotool.RandomString(tt.n)
			if len(code) != tt.want {
				t.Errorf("RandomString(%d) = %s, want length %d, got length %d",
					tt.n, code, tt.want, len(code))
			}
		})
	}
}

func TestSampleGenerateCode(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		wantErr bool
	}{
		{"valid length", 6, false},
		{"zero length", 0, true},
		{"negative length", -1, true},
		{"large length", 20, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := gotool.SampleGenerateCode(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SampleGenerateCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(got) != tt.input {
				t.Errorf("SampleGenerateCode() = %v, want length %v", got, tt.input)
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
		name      string
		input     uint64
		baseChars []byte
		want      string
		wantErr   bool
	}{
		{name: "base62 test1", input: 0, baseChars: baseChars62, want: "rflqBR", wantErr: false},
		{name: "base32 test1", input: 0, baseChars: baseChars32, want: "DSQZFP", wantErr: false},
		// ... 可以保留其他现有的测试用例 ...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := gotool.UniInvCodeLen6ByUID(tt.input, tt.baseChars)
			if (err != nil) != tt.wantErr {
				t.Errorf("UniInvCodeLen6ByUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UniInvCodeLen6ByUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomUniCode(t *testing.T) {
	tests := []struct {
		name        string
		longCode    bool
		readability bool
		wantLen     int
		wantErr     bool
	}{
		{"6位可读性高", false, true, 6, false},
		{"8位可读性高", true, true, 8, false},
		{"8位严格不重复", true, false, 8, false},
		{"6位全字符", false, false, 6, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := gotool.RandomUniCode(tt.longCode, tt.readability)
			if (err != nil) != tt.wantErr {
				t.Errorf("RandomUniCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(got) != tt.wantLen {
				t.Errorf("RandomUniCode() = %v, want length %v", got, tt.wantLen)
			}
		})
	}
}

func BenchmarkRandomUniCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := gotool.RandomUniCode(true, true); err != nil {
			b.Fatal(err)
		}
		if _, err := gotool.RandomUniCode(true, false); err != nil {
			b.Fatal(err)
		}
		if _, err := gotool.RandomUniCode(false, true); err != nil {
			b.Fatal(err)
		}
		if _, err := gotool.RandomUniCode(false, false); err != nil {
			b.Fatal(err)
		}
	}
}
