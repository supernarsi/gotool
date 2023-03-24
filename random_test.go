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
