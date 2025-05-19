package gotool

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"
)

// StringIsEmpty 检查字符串是否为空
func StringIsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// StringIsNotEmpty 检查字符串是否非空
func StringIsNotEmpty(s string) bool {
	return !StringIsEmpty(s)
}

// StringTruncate 截断字符串到指定长度，并添加省略号
func StringTruncate(s string, maxLen int) string {
	if utf8.RuneCountInString(s) <= maxLen {
		return s
	}
	runes := []rune(s)
	if maxLen <= 3 {
		return string(runes[:maxLen])
	}
	return string(runes[:maxLen-3]) + "..."
}

// StringContainsAny 检查字符串是否包含任意一个子串
func StringContainsAny(s string, substrings ...string) bool {
	for _, sub := range substrings {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}

// StringContainsAll 检查字符串是否包含所有子串
func StringContainsAll(s string, substrings ...string) bool {
	for _, sub := range substrings {
		if !strings.Contains(s, sub) {
			return false
		}
	}
	return true
}

// StringToMD5 计算字符串的MD5哈希值
func StringToMD5(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

// StringToSHA1 计算字符串的SHA1哈希值
func StringToSHA1(s string) string {
	hash := sha1.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

// StringToSHA256 计算字符串的SHA256哈希值
func StringToSHA256(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}

// StringToBase64 将字符串编码为Base64
func StringToBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Base64ToString 将Base64解码为字符串
func Base64ToString(s string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// StringReplaceAll 替换字符串中的多个子串
func StringReplaceAll(s string, replacements map[string]string) string {
	result := s
	for old, new := range replacements {
		result = strings.ReplaceAll(result, old, new)
	}
	return result
}

// StringIsMatch 检查字符串是否匹配正则表达式
func StringIsMatch(s, pattern string) bool {
	match, err := regexp.MatchString(pattern, s)
	return err == nil && match
}

// StringExtractByRegex 使用正则表达式提取字符串中的内容
func StringExtractByRegex(s, pattern string) []string {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return nil
	}
	return reg.FindStringSubmatch(s)
}

// StringExtractAllByRegex 使用正则表达式提取字符串中的所有匹配内容
func StringExtractAllByRegex(s, pattern string) [][]string {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return nil
	}
	return reg.FindAllStringSubmatch(s, -1)
}

// StringCamelToSnake 将驼峰命名转换为蛇形命名
func StringCamelToSnake(s string) string {
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

// StringSnakeToCamel 将蛇形命名转换为驼峰命名
func StringSnakeToCamel(s string) string {
	parts := strings.Split(s, "_")
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}

// StringPascalToCamel 将帕斯卡命名转换为驼峰命名
func StringPascalToCamel(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// StringCamelToPascal 将驼峰命名转换为帕斯卡命名
func StringCamelToPascal(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// StringTemplate 使用模板渲染字符串
func StringTemplate(templateStr string, data interface{}) (string, error) {
	tmpl, err := template.New("string_template").Parse(templateStr)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// StringMaskEmail 对邮箱地址进行部分遮蔽
func StringMaskEmail(email string) string {
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

// StringMaskPhone 对手机号码进行部分遮蔽
func StringMaskPhone(phone string) string {
	cleanPhone := strings.ReplaceAll(strings.ReplaceAll(phone, "-", ""), " ", "")
	if len(cleanPhone) < 7 {
		return phone
	}

	prefix := cleanPhone[:3]
	suffix := cleanPhone[len(cleanPhone)-4:]
	return prefix + strings.Repeat("*", len(cleanPhone)-7) + suffix
}

// StringFormatByteSize 格式化字节大小为人类可读格式
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

// StringSplitAndTrim 分割字符串并去除每部分的空白
func StringSplitAndTrim(s, sep string) []string {
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

// StringJoinWithSeparator 使用分隔符连接字符串，忽略空字符串
func StringJoinWithSeparator(sep string, parts ...string) string {
	nonEmpty := make([]string, 0, len(parts))
	for _, part := range parts {
		if part != "" {
			nonEmpty = append(nonEmpty, part)
		}
	}
	return strings.Join(nonEmpty, sep)
}