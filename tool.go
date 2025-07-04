package gotool

import (
	"encoding/binary"
	"math"
	"math/rand"
	"reflect"
	"strings"
	"time"

	"github.com/spaolacci/murmur3"
)

// 泛型定义
type eleInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
type eleUint interface {
	~uint | ~uint8 | ~uint16 | ~uint32
}
type eleFloat interface{ ~float32 | ~float64 }
type eleNum interface{ eleInt | eleUint | eleFloat }
type ElementType interface{ eleNum | string }

type IpDetail struct {
	Ip           string
	CountryShort string
	CountryLong  string
	Province     string
	City         string
	Isp          string
	Latitude     float32
	Longitude    float32
	Zipcode      string
	Timezone     string
}

// MergeMaps 泛型合并 map，适用于所有类型的 key 和 value
func MergeMaps[K comparable, V any](maps ...map[K]V) map[K]V {
	result := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// InArray 判断元素是否在切片中
func InArray[T ElementType](target T, arr []T) bool {
	for _, item := range arr {
		if target == item {
			return true
		}
	}
	return false
}

// UniqueElements 过滤切片中的重复元素
func UniqueElements[T ElementType](elements []T) []T {
	uniqueMap := make(map[T]struct{}, len(elements))
	res := make([]T, 0, len(elements))
	for _, ele := range elements {
		if _, ok := uniqueMap[ele]; !ok {
			uniqueMap[ele] = struct{}{}
			res = append(res, ele)
		}
	}
	return res
}

// ListToMap 泛型方法，允许自定义 key 选择
func ListToMap[T any, K comparable](list []T, keySelector func(T) K) map[K]T {
	result := make(map[K]T, len(list))
	for _, item := range list {
		result[keySelector(item)] = item
	}
	return result
}

// ListToMapByField 反射版 ListToMap，支持结构体字段作为 key
func ListToMapByField[T any](list []T, key string) map[int]T {
	result := make(map[int]T, len(list))
	for _, item := range list {
		val := reflect.ValueOf(item)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}

		field := val.FieldByName(key)
		if !field.IsValid() {
			continue
		}

		var idx int
		switch field.Kind() {
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
			idx = int(field.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			idx = int(field.Uint())
		default:
			continue
		}
		result[idx] = item
	}
	return result
}

// RandInt 返回打乱后的前 needNum 个元素
func RandInt(nums []int, dayN uint, needNum int) []int {
	rng := rand.New(rand.NewSource(int64(dayN)))
	shuffled := make([]int, len(nums))
	copy(shuffled, nums)
	rng.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	if len(shuffled) > needNum {
		return shuffled[:needNum]
	}
	return shuffled
}

// Lottery 抽奖算法
func Lottery(probabilities []int) int {
	totalProbability := 0
	for _, probability := range probabilities {
		totalProbability += probability
	}

	randomNumber := rand.Intn(totalProbability)
	for idx, probability := range probabilities {
		if randomNumber < probability {
			return idx
		}
		randomNumber -= probability
	}
	return 0
}

// RandomIdx 生成 0 到 num-1 的随机索引
func RandomIdx(num int) int {
	var globalRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	return globalRand.Intn(num)
}

// Difference 计算两个切片的差集
func Difference[T ElementType](slice1, slice2 []T) []T {
	m := make(map[T]struct{}, len(slice2))
	for _, item := range slice2 {
		m[item] = struct{}{}
	}

	diff := make([]T, 0, len(slice1))
	for _, item := range slice1 {
		if _, exists := m[item]; !exists {
			diff = append(diff, item)
		}
	}
	return diff
}

// AssignGroup 使用 Murmur3 哈希算法分配组
func AssignGroup(id uint32, seed uint32) uint32 {
	userIDBytes := [4]byte{}
	binary.LittleEndian.PutUint32(userIDBytes[:], id)

	hashValue := murmur3.Sum32WithSeed(userIDBytes[:], seed)
	return hashValue % 100
}

// FloatRatioToInt 将浮点数数组转换为整数百分比（确保总和 100）
func FloatRatioToInt(input []float32) []int {
	var total float32
	for _, value := range input {
		total += value
	}

	result := make([]int, len(input))
	var sum int
	for i, value := range input {
		percentage := math.Round((float64(value) / float64(total)) * 100)
		result[i] = int(percentage)
		sum += result[i]
	}

	// 确保总和为 100
	if sum != 100 {
		diff := 100 - sum
		result[len(result)-1] += diff
	}
	return result
}

// MaskNickname 隐藏昵称
// - 单字节字符（如英文、数字）显示前两位，其余用星号替换
// - 多字节字符（如中文、日文、韩文）显示第一位，其余用星号替换
func MaskNickname(nickname string) string {
	if nickname == "" {
		return ""
	}

	runes := []rune(nickname)
	if len(runes) == 0 {
		return ""
	}

	// 如果只有一个字符，直接返回
	if len(runes) == 1 {
		return string(runes)
	}

	// 检查第一个字符的类型
	firstCharIsMultibyte := IsMultibyte(runes[0])

	var visible int
	if firstCharIsMultibyte {
		// 多字节字符：显示第一位
		visible = 1
	} else {
		// 单字节字符：显示前两位
		visible = 2
		if visible > len(runes) {
			visible = len(runes)
		}
	}

	// 计算需要隐藏的字符数
	hiddenLength := len(runes) - visible
	if hiddenLength > 3 {
		hiddenLength = 3
	}

	// 如果不需要隐藏任何字符，直接返回原字符串
	if hiddenLength <= 0 {
		return string(runes[:visible])
	}

	return string(runes[:visible]) + strings.Repeat("*", hiddenLength)
}

// IsMultibyte 判断字符是否为多字节字符
func IsMultibyte(r rune) bool {
	// ASCII字符（单字节）
	if r <= 127 {
		return false
	}
	// 其他字符（包括全角字符、中文、日文、韩文、emoji等）视为多字节
	return true
}

func CalSwitchRemain(lastTime int, cd int) int {
	now := int(time.Now().Unix())
	remain := lastTime + cd - now
	if remain < 0 {
		return 0
	}
	return remain
}

func GetConstellation(timestamp int64) uint {
	t := time.Unix(timestamp, 0)
	month := int(t.Month())
	day := t.Day()

	// 星座开始日期
	constellations := [][]int{
		{12, 22}, {1, 20}, {2, 19}, {3, 21}, {4, 20}, {5, 21},
		{6, 22}, {7, 23}, {8, 23}, {9, 23}, {10, 24}, {11, 23},
	}

	for i := 0; i < 12; i++ {
		nextIndex := (i + 1) % 12
		if (month > constellations[i][0] || (month == constellations[i][0] && day >= constellations[i][1])) &&
			(month < constellations[nextIndex][0] || (month == constellations[nextIndex][0] && day < constellations[nextIndex][1])) {
			return uint(i)
		}
	}
	return 0
}
