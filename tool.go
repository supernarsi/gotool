package gotool

import "reflect"

type eleInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
type eleUint interface {
	~uint | ~uint8 | ~uint16 | ~uint32
}
type eleFloat interface{ ~float32 | ~float64 }
type eleNum interface{ eleInt | eleUint | eleFloat }
type ElementType interface{ eleNum | string }

// MergeMaps 合并 map
func MergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{}, len(maps)*10)
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
	uniqueMap := make(map[T]struct{}, len(elements))
	res := make([]T, 0, cap(elements))
	for _, ele := range elements {
		if _, ok := uniqueMap[ele]; !ok {
			uniqueMap[ele] = struct{}{}
			res = append(res, ele)
		}
	}
	return res
}

type reflectHelper struct{}

func (rh reflectHelper) getKey(row interface{}, key string) (int, bool) {
	immutable := reflect.ValueOf(row)
	var val reflect.Value
	switch immutable.Kind() {
	case reflect.Ptr:
		val = immutable.Elem().FieldByName(key)
	case reflect.Struct:
		val = immutable.FieldByName(key)
	default:
		return 0, false
	}
	switch val.Type().Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		return int(val.Int()), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int(val.Uint()), true
	default:
		return 0, false
	}
}

func ListToMap(list interface{}, key string) map[int]interface{} {
	res := make(map[int]interface{})
	arr := toSlice(list)
	rh := reflectHelper{}
	for _, row := range arr {
		if k, ok := rh.getKey(row, key); ok {
			res[k] = row
		}
	}
	return res
}

func toSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		return []interface{}{arr}
	}
	l := v.Len()
	res := make([]interface{}, 0, l)
	for i := 0; i < l; i++ {
		res = append(res, v.Index(i).Interface())
	}
	return res
}
