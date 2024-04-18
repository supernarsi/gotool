package gotool_test

import (
	"testing"

	"github.com/supernarsi/gotool"
)

func TestVerCompare(t *testing.T) {
	tests := []struct {
		name string
		v1   string
		v2   string
		op   string
		want bool
	}{
		{name: "test > true", v1: "1.1.1", v2: "1.1.0", op: ">", want: true},
		{name: "test > true", v1: "3.2.1", v2: "1.2.3", op: ">", want: true},
		{name: "test > true", v1: "3.2.1", v2: "1.9.8", op: ">", want: true},
		{name: "test > true", v1: "3.11.1", v2: "3.1.1", op: ">", want: true},
		{name: "test > false", v1: "1.1.1", v2: "1.1.0", op: "<", want: false},
		{name: "test >= true", v1: "1.1.1", v2: "1.1.0", op: ">=", want: true},
		{name: "test >= true", v1: "1.1.1", v2: "1.1.1", op: ">=", want: true},
		{name: "test >= false", v1: "1.0.1", v2: "1.1.0", op: ">=", want: false},
		{name: "test >= false", v1: "1.12.1", v2: "10.1.0", op: ">=", want: false},
		{name: "test >= false", v1: "1.12.1", v2: "2.1.0", op: ">=", want: false},
		{name: "test < true", v1: "2.1.0", v2: "2.10.0", op: "<", want: true},
		{name: "test < false", v1: "21.0.1", v2: "2.99.99", op: "<", want: false},
		{name: "test <= true", v1: "3.10.99", v2: "21.1.0", op: "<=", want: true},
		{name: "test <= false", v1: "10.101.1", v2: "10.29.99", op: "<=", want: false},
		{name: "test == true", v1: "101.10.1", v2: "101.10.1", op: "==", want: true},
		{name: "test == false", v1: "101.10.1", v2: "10.110.1", op: "==", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {})
			if got := gotool.VerCompare(tt.v1, tt.v2, tt.op); got != tt.want {
				t.Errorf("got result %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	tests := []struct {
		name string
		v1   string
		v2   string
		want int
	}{
		{name: "1.1.1 > 1.1.0", v1: "1.1.1", v2: "1.1.0", want: 1},
		{name: "3.2.1 > 1.2.3", v1: "3.2.1", v2: "1.2.3", want: 1},
		{name: "3.2.1 > 1.9.8", v1: "3.2.1", v2: "1.9.8", want: 1},
		{name: "3.11.1 > 3.1.1", v1: "3.11.1", v2: "3.1.1", want: 1},
		{name: "1.1.1 > 1.1.0", v1: "1.1.1", v2: "1.1.0", want: 1},
		{name: "1.11.1 > 1.10.1", v1: "1.11.1", v2: "1.10.1", want: 1},
		{name: "1.11.3 > 1.10.11", v1: "1.11.3", v2: "1.10.11", want: 1},
		{name: "1.1.1 == 1.1.1", v1: "1.1.1", v2: "1.1.1", want: 0},
		{name: "1.0.1 < 1.1.0", v1: "1.0.1", v2: "1.1.0", want: 2},
		{name: "1.12.1 < 10.1.0", v1: "1.12.1", v2: "10.1.0", want: 2},
		{name: "1.12.1 < 2.1.0", v1: "1.12.1", v2: "2.1.0", want: 2},
		{name: "2.1.0 < 2.10.0", v1: "2.1.0", v2: "2.10.0", want: 2},
		{name: "21.0.1 > 2.99.99", v1: "21.0.1", v2: "2.99.99", want: 1},
		{name: "3.10.99 < 21.1.0", v1: "3.10.99", v2: "21.1.0", want: 2},
		{name: "10.101.1 > 10.29.99", v1: "10.101.1", v2: "10.29.99", want: 1},
		{name: "101.10.1 == 101.10.1", v1: "101.10.1", v2: "101.10.1", want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {})
			if got := gotool.Compare(tt.v1, tt.v2); got != tt.want {
				t.Errorf("got result %v, want %v", got, tt.want)
			}
		})
	}
}
