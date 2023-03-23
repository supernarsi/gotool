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

func ListToMap(list interface{}, key string) map[int]interface{} {
	res := make(map[int]interface{})
	arr := toSlice(list)
	var val reflect.Value
	for _, row := range arr {
		immutable := reflect.ValueOf(row)
		switch immutable.Kind() {
		case reflect.Ptr:
			val = immutable.Elem().FieldByName(key)
		case reflect.Struct:
			val = immutable.FieldByName(key)
		default:
			continue
		}
		key := 0
		switch val.Type().Kind() {
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
			key = int(val.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			key = int(val.Uint())
		default:
			continue
		}
		res[key] = row
	}
	return res
}

func toSlice(arr interface{}) []interface{} {
	res := make([]interface{}, 0)
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		res = append(res, arr)
		return res
	}

	l := v.Len()
	for i := 0; i < l; i++ {
		res = append(res, v.Index(i).Interface())
	}
	return res
}
