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
	PRIME1 = 3          // 与字符集长度 62 互质
	PRIME2 = 5          // 与邀请码长度 6 互质
	SALT   = 4213098675 // 随意取一个数值
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
	return RandomUniCode(false, true)
}

// RandomReadableUniCode8Len 返回可读性高的 8 位唯一码
func RandomReadableUniCode8Len() string {
	return RandomUniCode(true, true)
}

// RandomStrictUniCode8Len 返回严格不重复的 8 位唯一码
func RandomStrictUniCode8Len() string {
	return RandomUniCode(true, false)
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
func RandomUniCode(longCode bool, readability bool) string {
	var (
		ans        []byte
		appendChar byte
	)
	// 32 个 16 进制的字符
	pureUuid := strings.ReplaceAll(uuid.New().String(), "-", "")
	length, step := 6, 5
	if longCode {
		length, step = 8, 4
	}
	for i := 0; i < length; i++ {
		// 如果非 longCode 模式，将摒弃 32 位 uuid 的最后两位，可能会导致生成结果不唯一
		if hexNum, err := strconv.ParseUint(pureUuid[i*step:i*step+step], 16, 32); err == nil {
			if readability {
				appendChar = baseChars32[hexNum%32]
			} else {
				appendChar = baseChars62[hexNum%62]
			}
			ans = append(ans, appendChar)
		} else {
			return ""
		}
	}
	return string(ans)
}

// RandomString 随机生成 n 位字符串
// 重复概率较高，参考数据：
// 每连续生成 10 万次 6 位字符串，约产生 300 次重复
// 每连续生成 10 万次 8 位字符串，约产生 3 次重复
func RandomString(n int) string {
	randBytes := make([]byte, n/2)
	_, _ = rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

// UniInvCodeLen6ByUID 根据 uid 生成 6 位的唯一码（可逆）
func UniInvCodeLen6ByUID(uid uint64, baseChars []byte) string {
	l := 6
	// 放大 + 加盐
	uid = uid*PRIME1 + SALT

	var code []rune
	slIdx := make([]byte, l)

	// 扩散
	for i := 0; i < l; i++ {
		// 获取 62 进制的每一位值
		slIdx[i] = byte(uid % uint64(len(baseChars)))
		// 其他位与个位加和再取余（让个位的变化影响到所有位）
		slIdx[i] = (slIdx[i] + byte(i)*slIdx[0]) % byte(len(baseChars))
		// 相当于右移一位（62进制）
		uid = uid / uint64(len(baseChars))
	}

	// 混淆
	for i := 0; i < l; i++ {
		idx := (byte(i) * PRIME2) % byte(l)
		code = append(code, rune(baseChars[slIdx[idx]]))
	}
	return string(code)
}

func SampleGenerateCode(codeLength int) string {
	var letters = []rune("ABCDEFGHJKMNPQRSTUVWXYZ")
	code := make([]rune, codeLength)
	for i := range code {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		code[i] = letters[index.Int64()]
	}
	return string(code)
}
