package gotool

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"
)

var (
	ErrInvalidBase64 = errors.New("invalid base64 string")
	ErrInvalidRegex  = errors.New("invalid regular expression")
)

// StringIsEmpty checks if a string is empty or contains only whitespace
func StringIsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// StringIsNotEmpty checks if a string is not empty and contains non-whitespace characters
func StringIsNotEmpty(s string) bool {
	return !StringIsEmpty(s)
}

// StringTruncate truncates a string to the specified maximum length and adds ellipsis
// If maxLen is less than or equal to 3, it will simply truncate without ellipsis
func StringTruncate(s string, maxLen int) string {
	if maxLen <= 0 {
		return ""
	}

	runeCount := utf8.RuneCountInString(s)
	if runeCount <= maxLen {
		return s
	}

	runes := []rune(s)
	if maxLen <= 3 {
		return string(runes[:maxLen])
	}
	return string(runes[:maxLen-3]) + "..."
}

// StringContainsAny checks if a string contains any of the provided substrings
func StringContainsAny(s string, substrings ...string) bool {
	if len(substrings) == 0 {
		return false
	}

	for _, sub := range substrings {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}

// StringContainsAll checks if a string contains all of the provided substrings
func StringContainsAll(s string, substrings ...string) bool {
	if len(substrings) == 0 {
		return true
	}

	for _, sub := range substrings {
		if !strings.Contains(s, sub) {
			return false
		}
	}
	return true
}

// StringToMD5 calculates the MD5 hash of a string
func StringToMD5(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

// StringToSHA1 calculates the SHA1 hash of a string
func StringToSHA1(s string) string {
	hash := sha1.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

// StringToSHA256 calculates the SHA256 hash of a string
func StringToSHA256(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}

// StringToBase64 encodes a string to Base64
func StringToBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Base64ToString decodes a Base64 string
func Base64ToString(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	bytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrInvalidBase64, err)
	}
	return string(bytes), nil
}

// StringReplaceAll replaces multiple substrings in a string according to the provided map
func StringReplaceAll(s string, replacements map[string]string) string {
	if len(replacements) == 0 {
		return s
	}

	result := s
	for old, new := range replacements {
		result = strings.ReplaceAll(result, old, new)
	}
	return result
}

// StringIsMatch checks if a string matches the provided regular expression pattern
func StringIsMatch(s, pattern string) (bool, error) {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return false, fmt.Errorf("%w: %v", ErrInvalidRegex, err)
	}
	return reg.MatchString(s), nil
}

// StringExtractByRegex extracts content from a string using a regular expression
func StringExtractByRegex(s, pattern string) ([]string, error) {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidRegex, err)
	}
	return reg.FindStringSubmatch(s), nil
}

// StringExtractAllByRegex extracts all matches from a string using a regular expression
func StringExtractAllByRegex(s, pattern string) ([][]string, error) {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidRegex, err)
	}
	return reg.FindAllStringSubmatch(s, -1), nil
}

// StringCamelToSnake converts a camelCase string to snake_case
func StringCamelToSnake(s string) string {
	if s == "" {
		return s
	}

	var result bytes.Buffer
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteRune('_')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// StringSnakeToCamel converts a snake_case string to camelCase
func StringSnakeToCamel(s string) string {
	if s == "" {
		return s
	}

	parts := strings.Split(s, "_")
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}

// StringPascalToCamel converts a PascalCase string to camelCase
func StringPascalToCamel(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// StringCamelToPascal converts a camelCase string to PascalCase
func StringCamelToPascal(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// StringTemplate renders a string using the provided template and data
func StringTemplate(templateStr string, data interface{}) (string, error) {
	if templateStr == "" {
		return "", nil
	}

	tmpl, err := template.New("string_template").Parse(templateStr)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %v", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %v", err)
	}

	return buf.String(), nil
}

// StringMaskEmail masks an email address for privacy
func StringMaskEmail(email string) string {
	if email == "" {
		return email
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return email
	}

	name := parts[0]
	domain := parts[1]

	if len(name) <= 2 {
		return name + "@" + domain
	}

	maskedName := name[:2] + strings.Repeat("*", len(name)-2)
	return maskedName + "@" + domain
}

// StringMaskPhone masks a phone number for privacy
func StringMaskPhone(phone string) string {
	if phone == "" {
		return phone
	}

	cleanPhone := strings.ReplaceAll(strings.ReplaceAll(phone, "-", ""), " ", "")
	if len(cleanPhone) < 7 {
		return phone
	}

	prefix := cleanPhone[:3]
	suffix := cleanPhone[len(cleanPhone)-4:]
	return prefix + strings.Repeat("*", len(cleanPhone)-7) + suffix
}

// StringFormatByteSize formats a byte size into a human-readable string
func StringFormatByteSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}

	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// StringSplitAndTrim splits a string by the separator and trims whitespace from each part
func StringSplitAndTrim(s, sep string) []string {
	if s == "" {
		return nil
	}

	parts := strings.Split(s, sep)
	result := make([]string, 0, len(parts))

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}

	return result
}

// StringJoinWithSeparator joins strings with a separator, ignoring empty strings
func StringJoinWithSeparator(sep string, parts ...string) string {
	if len(parts) == 0 {
		return ""
	}

	nonEmpty := make([]string, 0, len(parts))
	for _, part := range parts {
		if part != "" {
			nonEmpty = append(nonEmpty, part)
		}
	}
	return strings.Join(nonEmpty, sep)
}

// StringReverse reverses a string while preserving UTF-8 encoding
func StringReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// StringCountWords counts the number of words in a string
func StringCountWords(s string) int {
	if s == "" {
		return 0
	}

	words := strings.Fields(s)
	return len(words)
}

// StringRemoveDuplicates removes duplicate strings from a slice while preserving order
func StringRemoveDuplicates(slice []string) []string {
	if len(slice) == 0 {
		return nil
	}

	seen := make(map[string]struct{})
	result := make([]string, 0, len(slice))

	for _, s := range slice {
		if _, ok := seen[s]; !ok {
			seen[s] = struct{}{}
			result = append(result, s)
		}
	}

	return result
}
