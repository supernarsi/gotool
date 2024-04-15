package gotool_test

import (
	"testing"

	"github.com/supernarsi/gotool"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name string
		pass string
	}{
		{name: "test empty", pass: ""},
		{name: "test number", pass: "123"},
		{name: "test number 0", pass: "0"},
		{name: "test random str", pass: "jfieh91=-+23"},
		{name: "test random symbol", pass: "(&^&^$&%()*#_+!|~"},
		{name: "test simple pass 000000", pass: "000000"},
		{name: "test simple pass 123456", pass: "123456"},
		{name: "test simple pass 88888888", pass: "88888888"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
			})

			if hash, err := gotool.HashPassword(tt.pass); err != nil {
				t.Errorf("create password hash failed")
			} else {
				if !gotool.CheckPasswordHash(tt.pass, hash) {
					t.Errorf("hash error pass is: %v , hash is: %v", tt.pass, hash)
				}
			}
		})
	}
}

func TestCheckPasswordHash(t *testing.T) {
	tests := []struct {
		name string
		pass string
		hash string
		want bool
	}{
		{name: "test empty", pass: "", hash: "$2a$14$eq4nqzOmpSIo9qB2SQJkbukQ/D7M87IC5rcy3V2hsylRDCOxMbMKS", want: true},
		{name: "test number", pass: "123", hash: "$2a$14$7nDWKWfAUiBv.VGwRt/90ebQhzw5bOGjyc8/dq6gh9PqAQB71lyIC", want: true},
		{name: "test number 0", pass: "0", hash: "$2a$14$ej2AVUXfWwrx/U5R6IQ7M.ai1oNZYAhbK9w.96maqkOXvbsDIMNrK", want: true},
		{name: "test random str", pass: "jfieh91=-+23", hash: "$2a$14$YK1B8A90KUq18QYUG4hBkubWnvPsabBS.ynK33/9KLeAjxmZGjfRy", want: true},
		{name: "test random symbol", pass: "(&^&^$&%()*#_+!|~", hash: "$2a$14$eO9iERvmjJYNypXzKrxaduIDpTgMf8KDXf0qqG1rLPGGLhyZJQ4yS", want: true},
		{name: "test simple pass 000000", pass: "000000", hash: "$2a$14$2Fuhbbn/FfO.WV5Jn4J9x.Asj.QVonN/wqttgXE23ymqjD1xTqg6q", want: true},
		{name: "test simple pass 123456", pass: "123456", hash: "$2a$14$2cAkx9kEF6lPZ1GCDoES5ORn0Jf3GtHiY5o8CDRVapftPHu7h8Znu", want: true},
		{name: "test simple pass 88888888", pass: "88888888", hash: "$2a$14$a9A1Gn/WxTFbrKtEVX92Je9mWaX3M6f1zletgAiSeaSVlNOXE6/c.", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
			})

			if got := gotool.CheckPasswordHash(tt.pass, tt.hash); got != tt.want {
				t.Errorf("got result %v, want %v", got, tt.want)
			}
		})
	}
}
