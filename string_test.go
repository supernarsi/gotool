package gotool

import (
	"testing"
)

func TestStringIsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", true},
		{"whitespace only", "   ", true},
		{"non-empty string", "hello", false},
		{"string with spaces", " hello ", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringIsEmpty(tt.input); got != tt.expected {
				t.Errorf("StringIsEmpty() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestStringTruncate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		maxLen   int
		expected string
	}{
		{"empty string", "", 5, ""},
		{"shorter than max", "hello", 10, "hello"},
		{"exact length", "hello", 5, "hello"},
		{"longer than max", "hello world", 5, "he..."},
		{"max len 3", "hello", 3, "hel"},
		{"max len 0", "hello", 0, ""},
		{"unicode string", "你好世界", 3, "你好世"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringTruncate(tt.input, tt.maxLen); got != tt.expected {
				t.Errorf("StringTruncate() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestStringContainsAny(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		substrs  []string
		expected bool
	}{
		{"empty string", "", []string{}, false},
		{"no substrings", "hello", []string{}, false},
		{"contains one", "hello world", []string{"world"}, true},
		{"contains none", "hello world", []string{"foo", "bar"}, false},
		{"contains multiple", "hello world", []string{"foo", "world", "bar"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringContainsAny(tt.input, tt.substrs...); got != tt.expected {
				t.Errorf("StringContainsAny() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestStringToMD5(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", "d41d8cd98f00b204e9800998ecf8427e"},
		{"hello world", "hello world", "5eb63bbbe01eeed093cb22bb8f5acdc3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToMD5(tt.input); got != tt.expected {
				t.Errorf("StringToMD5() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestBase64ToString(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		expectError bool
	}{
		{"empty string", "", "", false},
		{"valid base64", "SGVsbG8gV29ybGQ=", "Hello World", false},
		{"invalid base64", "invalid", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Base64ToString(tt.input)
			if (err != nil) != tt.expectError {
				t.Errorf("Base64ToString() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if got != tt.expected {
				t.Errorf("Base64ToString() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestStringIsMatch(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		pattern     string
		expected    bool
		expectError bool
	}{
		{"valid pattern match", "hello123", `\d+`, true, false},
		{"valid pattern no match", "hello", `\d+`, false, false},
		{"invalid pattern", "hello", `[`, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringIsMatch(tt.input, tt.pattern)
			if (err != nil) != tt.expectError {
				t.Errorf("StringIsMatch() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if got != tt.expected {
				t.Errorf("StringIsMatch() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestStringCamelToSnake(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single word", "hello", "hello"},
		{"camelCase", "helloWorld", "hello_world"},
		{"PascalCase", "HelloWorld", "hello_world"},
		{"multiple words", "helloWorldTest", "hello_world_test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringCamelToSnake(tt.input); got != tt.expected {
				t.Errorf("StringCamelToSnake() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestStringMaskEmail(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"short email", "a@b.com", "a@b.com"},
		{"normal email", "john.doe@example.com", "jo******@example.com"},
		{"invalid email", "notanemail", "notanemail"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringMaskEmail(tt.input); got != tt.expected {
				t.Errorf("StringMaskEmail() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestStringMaskPhone(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"short number", "123", "123"},
		{"normal number", "1234567890", "123***7890"},
		{"formatted number", "123-456-7890", "123***7890"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringMaskPhone(tt.input); got != tt.expected {
				t.Errorf("StringMaskPhone() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestStringFormatByteSize(t *testing.T) {
	tests := []struct {
		name     string
		input    int64
		expected string
	}{
		{"bytes", 500, "500 B"},
		{"kilobytes", 1500, "1.5 KB"},
		{"megabytes", 1500000, "1.4 MB"},
		{"gigabytes", 1500000000, "1.4 GB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringFormatByteSize(tt.input); got != tt.expected {
				t.Errorf("StringFormatByteSize() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestStringReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single char", "a", "a"},
		{"ascii string", "hello", "olleh"},
		{"unicode string", "你好世界", "界世好你"},
		{"mixed string", "hello世界", "界世olleh"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringReverse(tt.input); got != tt.expected {
				t.Errorf("StringReverse() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestStringCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"empty string", "", 0},
		{"single word", "hello", 1},
		{"multiple words", "hello world", 2},
		{"extra spaces", "  hello   world  ", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringCountWords(tt.input); got != tt.expected {
				t.Errorf("StringCountWords() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestStringRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{"empty slice", []string{}, nil},
		{"no duplicates", []string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{"with duplicates", []string{"a", "b", "a", "c", "b"}, []string{"a", "b", "c"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringRemoveDuplicates(tt.input)
			if len(got) != len(tt.expected) {
				t.Errorf("StringRemoveDuplicates() length = %v, want %v", len(got), len(tt.expected))
				return
			}
			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("StringRemoveDuplicates()[%d] = %v, want %v", i, got[i], tt.expected[i])
				}
			}
		})
	}
}
