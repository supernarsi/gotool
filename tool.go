package gotool

import (
	"encoding/binary"
	"math/rand"
	"reflect"
	"time"
	_ "time/tzdata" // use built-in timezone database

	"github.com/spaolacci/murmur3"
)

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

// MergeMaps 合并 map
func MergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// MergeMapsAny 合并多个 map，适用于所有类型的 key 和 value
func MergeMapsAny[K comparable, V any](maps ...map[K]V) map[K]V {
	result := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// InArray 判断元素是否在数组切片中
func InArray[T ElementType](target T, arr []T) bool {
	for _, item := range arr {
		if target == item {
			return true
		}
	}
	return false
}

// UniqueElements 过滤掉切片中个重复元素
func UniqueElements[T ElementType](elements []T) []T {
	uniqueMap := make(map[T]struct{})
	res := make([]T, 0, len(elements))
	for _, ele := range elements {
		if _, ok := uniqueMap[ele]; !ok {
			uniqueMap[ele] = struct{}{}
			res = append(res, ele)
		}
	}
	return res
}

func ListToMap[T any](list []T, key string) map[int]T {
	result := make(map[int]T)
	for _, item := range list {
		// 使用反射来获取结构体字段的值
		val := reflect.ValueOf(item)
		switch val.Kind() {
		case reflect.Ptr:
			val = val.Elem()
		default:
			continue
		}

		idx := 0
		val = val.FieldByName(key)
		switch val.Type().Kind() {
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
			idx = int(val.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			idx = int(val.Uint())
		default:
			continue
		}
		result[idx] = item
	}
	return result
}

func RandInt(nums []int, dayN uint, needNum int) []int {
	rng := rand.New(rand.NewSource(int64(dayN)))
	rng.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	if len(nums) >= needNum {
		return nums[:needNum]
	} else {
		return nums
	}
}

// Lottery 抽奖算法
// probabilities 概率数组
// return 返回概率数组中奖下标
func Lottery(probabilities []int) int {
	totalProbability := 0
	for _, probability := range probabilities {
		totalProbability += probability
	}

	randomNumber := int(rand.Float64() * float64(totalProbability))
	for idx, probability := range probabilities {
		if randomNumber < probability {
			return idx
		}
		randomNumber -= probability
	}
	return 0
}

func RandomIdx(num int) int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return r.Intn(num)
}

func Difference[T ElementType](slice1, slice2 []T) []T {
	m := make(map[T]struct{})
	for _, item := range slice2 {
		m[item] = struct{}{}
	}

	var diff []T
	for _, item := range slice1 {
		if _, exists := m[item]; !exists {
			diff = append(diff, item)
		}
	}

	return diff
}

func AssignGroup(id uint32, seed uint32) uint32 {
	// 将用户 ID 转换为字节数组
	userIDBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(userIDBytes, id)

	// 创建带 seed 的 Murmur3 哈希函数
	hashDealer := murmur3.New32WithSeed(seed)
	// 写入 id 字节数据
	_, err := hashDealer.Write(userIDBytes)
	if err != nil {
		return 0
	}

	// 计算哈希值
	hashValue := hashDealer.Sum32()

	// 将哈希值映射到 0 到 99 之间的一个组
	group := hashValue % 100

	return group
}

func FloatRatioToInt(input []float32) []int {
	var total float32
	for _, value := range input {
		total += value
	}

	// 创建一个结果数组用于存储百分比值
	result := make([]int, len(input))
	var sum int
	for i, value := range input {
		percentage := (value / total) * 100
		result[i] = int(percentage)
		sum += result[i]
	}
	// 调整最后一个元素以确保总和为100
	if sum != 100 {
		result[len(result)-1] += 100 - sum
	}
	return result
}
