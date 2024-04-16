package gotool

import "strings"

// VerCompare 比较 v1、v2 的版本号大小
func VerCompare(v1, v2, operator string) bool {
	com := Compare(v1, v2)
	switch operator {
	case "==":
		if com == 0 {
			return true
		}
	case "<":
		if com == 2 {
			return true
		}
	case ">":
		if com == 1 {
			return true
		}
	case "<=":
		if com == 0 || com == 2 {
			return true
		}
	case ">=":
		if com == 0 || com == 1 {
			return true
		}
	}
	return false
}

// Compare 比较格式为 X.Y.Z 的版本号的大小关系（X、Y、Z 为纯数字）
// 返回值：0 表示v1与v2相等；1 表示v1大于v2；2 表示v1小于v2；-1 表示版本号格式错误
func Compare(v1, v2 string) int {
	ver1 := strings.Split(v1, ".")
	ver2 := strings.Split(v2, ".")
	if len(ver1) != len(ver2) || len(ver1) != 3 {
		return -1
	}
	// 循环比较
	for i := 0; i < 3; i++ {
		ver1I, ver2I := ver1[i], ver2[i]
		if len(ver1I) > len(ver2I) {
			return 1
		} else if len(ver1I) < len(ver2I) {
			return 2
		} else {
			if ver1I > ver2I {
				return 1
			} else if ver1I < ver2I {
				return 2
			}
		}
	}
	return 0
}
