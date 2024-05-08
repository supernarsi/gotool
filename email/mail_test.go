package email_test

import (
	"testing"

	"github.com/supernarsi/gotool/email"
)

func TestIsEmailValid(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"t1", "xxx@abc.com", true},
		{"t2", "hhhh@example.com", true},
		{"t3", "@sm070102", false},
		{"t4", ".com", false},
		{"t5", "cccc@gmail.com", true},
		{"t6", "cccc+123@gmail.com", true},
		{"t7", "@gmail.com", false},
		{"t8", "123@gmail", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := email.IsEmailValid(tt.input); got != tt.want {
				t.Errorf("IsEmailValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
