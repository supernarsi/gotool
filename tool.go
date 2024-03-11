package gotool

import (
	"errors"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

const emptyIdfa = "00000000-0000-0000-0000-000000000000"

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

func ListGroup[T any](list []T, key string) (values map[string][]T, err error) {
	values = make(map[string][]T)
	reflectValue := reflect.ValueOf(list)
	if reflectValue.Kind() != reflect.Slice && reflectValue.Kind() != reflect.Array {
		err = errors.New("参数必须为切片或数组")
		return
	}

	t := reflect.TypeOf((*T)(nil)).Elem()
	isPtr := t.Kind() == reflect.Ptr

	for i := 0; i < reflectValue.Len(); i++ {
		item := reflectValue.Index(i)
		if item.Kind() == reflect.Ptr {
			item = item.Elem()
		}

		val := item.FieldByName(key)
		if !val.IsValid() {
			err = errors.New("未找到键对应的值")
			return
		}

		var keyVal string
		switch val.Type().Kind() {
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
			keyVal = strconv.Itoa(int(val.Int()))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			keyVal = strconv.Itoa(int(val.Uint()))
			break
		case reflect.Float64, reflect.Float32:
			keyVal = strconv.Itoa(int(val.Float()))
			break
		case reflect.String:
			keyVal = val.String()
			break
		default:
			err = errors.New("键值类型不正确！")
			return
		}

		if _, ok := values[keyVal]; !ok {
			values[keyVal] = make([]T, 0)
		}
		arr := values[keyVal]
		itemVal := item.Interface()
		if isPtr {
			itemVal = reflect.New(reflect.TypeOf(item.Interface())).Interface()
			// 将 item 的值赋值给这个指针
			reflect.ValueOf(itemVal).Elem().Set(item)
		}

		values[keyVal] = append(arr, itemVal.(T))
	}
	return
}

func IdfaAvailable(idfa string) bool {
	return idfa != "" && idfa != emptyIdfa
}

// TimeToStamp 时间根据时区转时间戳
func TimeToStamp(dateTime *time.Time, timezone string) (int64, error) {
	local, err := time.LoadLocation(timezone)
	if err != nil {
		return 0, err
	}
	t := time.Date(dateTime.Year(), time.Month(dateTime.Month()), dateTime.Day(), dateTime.Hour(), dateTime.Minute(), dateTime.Second(), 0, local)
	return t.Unix(), nil
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
