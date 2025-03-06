package gotool

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

const (
	InvCodePrime1 = 3          // 与字符集长度 62 互质
	InvCodePrime2 = 5          // 与邀请码长度 6 互质
	InvCodeSalt   = 4213098675 // 随意取一个数值
)

var (
	baseChars32 = [32]byte{
		'2', '3', '4', '5', '6', '7', '8', '9',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
		'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R',
		'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	}
	baseChars62 = [62]byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
		'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
		'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
		'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
		'U', 'V', 'W', 'X', 'Y', 'Z',
	}
)

// RandomReadableUniCode6Len 返回可读性高的 6 位唯一码
func RandomReadableUniCode6Len() string {
	code, err := RandomUniCode(false, true)
	if err != nil {
		return ""
	}
	return code
}

// RandomReadableUniCode8Len 返回可读性高的 8 位唯一码
func RandomReadableUniCode8Len() string {
	code, err := RandomUniCode(true, true)
	if err != nil {
		return ""
	}
	return code
}

// RandomStrictUniCode8Len 返回严格不重复的 8 位唯一码
func RandomStrictUniCode8Len() string {
	code, err := RandomUniCode(true, false)
	if err != nil {
		return ""
	}
	return code
}

// RandomUniCode 随机生成 6 或 8 位不重复（非严格意义）字符串
// longCode 为 true 时返回 8 位长度；false 时返回 6 位
// readability 为 true 时返回只包含数字和大写字母（排除 0、O、1、I）；为 false 时返回可能包含所有数字和大小写字母
//
// 重复概率参考数据：
// [longCode == false && readability == true]
//   - 连续生成 10 万次，约 5 次重复；
//   - 连续生成 100 万次，约 500 次重复；
//   - 连续生成 1000 万次，约 50000 次重复
//
// [longCode == false && readability == false]
//   - 连续生成 10 万次，约 <1 次重复；
//   - 连续生成 100 万次，约 10 次重复；
//   - 连续生成 1000 万次，约 1000 次重复
//
// [longCode == true && readability == true]
//   - 连续生成 10 万次，约 <1 次重复；
//   - 连续生成 100 万次，约 <1 次重复；
//   - 连续生成 1000 万次，约 <50 次重复
//
// [longCode == true && readability == false]
//   - 理论上不重复（与 uuid 一一对应）
func RandomUniCode(longCode bool, readability bool) (string, error) {
	var (
		ans        strings.Builder
		appendChar byte
	)
	pureUuid := strings.ReplaceAll(uuid.New().String(), "-", "")
	length, step := 6, 5
	if longCode {
		length, step = 8, 4
	}
	ans.Grow(length)

	for i := 0; i < length; i++ {
		hexNum, err := strconv.ParseUint(pureUuid[i*step:i*step+step], 16, 32)
		if err != nil {
			return "", fmt.Errorf("parse uuid failed: %w", err)
		}
		if readability {
			appendChar = baseChars32[hexNum%32]
		} else {
			appendChar = baseChars62[hexNum%62]
		}
		ans.WriteByte(appendChar)
	}
	return ans.String(), nil
}

func UniInvCodeLen6ByUID(uid uint64, baseChars []byte) (string, error) {
	if len(baseChars) == 0 {
		return "", fmt.Errorf("baseChars cannot be empty")
	}

	l := 6
	uid = uid*InvCodePrime1 + InvCodeSalt

	code := make([]rune, l)
	slIdx := make([]byte, l)

	for i := 0; i < l; i++ {
		slIdx[i] = byte(uid % uint64(len(baseChars)))
		slIdx[i] = (slIdx[i] + byte(i)*slIdx[0]) % byte(len(baseChars))
		uid = uid / uint64(len(baseChars))
	}

	for i := 0; i < l; i++ {
		idx := (byte(i) * InvCodePrime2) % byte(l)
		code[i] = rune(baseChars[slIdx[idx]])
	}
	return string(code), nil
}

func SampleGenerateCode(codeLength int) (string, error) {
	if codeLength <= 0 {
		return "", fmt.Errorf("code length must be positive")
	}

	letters := []rune("ABCDEFGHJKMNPQRSTUVWXYZ")
	code := make([]rune, codeLength)

	for i := range code {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", fmt.Errorf("generate random number failed: %w", err)
		}
		code[i] = letters[index.Int64()]
	}
	return string(code), nil
}

// RandomString 随机生成指定长度的十六进制字符串
// 输入长度 n 会被转换为最接近的偶数（向下取整）
// 例如：输入1或2返回2位字符，输入3或4返回4位字符
// RandomString 随机生成 n 位字符串
func RandomString(n int) string {
    if n <= 0 {
        return ""
    }
    randBytes := make([]byte, n)
    _, _ = rand.Read(randBytes)
    return fmt.Sprintf("%x", randBytes)[:n]
}
