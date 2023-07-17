package util_test

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/supernarsi/gotool/util"
)

func TestCoroutineForeach(t *testing.T) {
	// 测试 int 类型
	t.Run("test int", func(t *testing.T) {
		process1 := func(ctx context.Context, i int) (int, error) {
			if i == 0 {
				return 0, errors.New("can not be 0")
			}
			return i * 2, nil
		}
		process2 := func(ctx context.Context, i int) (int, error) {
			time.Sleep(1 * time.Second)
			return i - 1, nil
		}
		tests := []struct {
			name    string
			input   []int
			process func(ctx context.Context, i int) (int, error)
			want    []int
		}{
			{"t0", []int{0}, process1, []int{999}},
			{"t1", []int{1, 2, 3}, process1, []int{2, 4, 6}},
			{"t2", []int{1, 1, 2, 0, -2}, process1, []int{2, 2, 4, 999, -4}},
			{"t3", []int{3, 2, 1, 0}, process2, []int{2, 1, 0, -1}},
		}

		for _, v := range tests {
			got := util.GoForeach(context.Background(), v.input, v.process, 999, 2)
			for i, val := range got {
				if val != v.want[i] {
					t.Errorf("Got %v, want %v", got, v.want)
				}
			}
		}
	})

	// 测试 string 类型
	t.Run("test string", func(t *testing.T) {
		p1 := func(ctx context.Context, s string) (string, error) {
			return s + "_", nil
		}
		p2 := func(ctx context.Context, s string) (string, error) {
			if s == "err" {
				return "", errors.New("got error")
			}
			return strings.ToUpper(s), nil
		}

		tests := []struct {
			name    string
			input   []string
			process func(ctx context.Context, s string) (string, error)
			want    []string
		}{
			{"s1", []string{"a", "b"}, p1, []string{"a_", "b_"}},
			{"s2", []string{}, p1, []string{}},
			{"s3", []string{"a", " ", "+", "B", "dBa", "err"}, p2, []string{"A", " ", "+", "B", "DBA", "[ERROR]"}},
		}
		for _, v := range tests {
			got := util.GoForeach(context.Background(), v.input, v.process, "[ERROR]", 100)
			for i, val := range got {
				if val != v.want[i] {
					t.Errorf("Got %v, want %v", got, v.want)
				}
			}
		}
	})

	// 测试类型互转
	t.Run("test int to string", func(t *testing.T) {
		p1 := func(ctx context.Context, s int) (string, error) {
			return strconv.Itoa(s) + "_", nil
		}

		tests := []struct {
			name    string
			input   []int
			process func(ctx context.Context, s int) (string, error)
			want    []string
		}{
			{"s1", []int{1, 2, 3, 4, 5, 0, -1}, p1, []string{"1_", "2_", "3_", "4_", "5_", "0_", "-1_"}},
			{"s2", []int{}, p1, []string{}},
			{"s3", []int{0, 2}, p1, []string{"0_", "2_"}},
		}
		for _, v := range tests {
			got := util.GoForeach(context.Background(), v.input, v.process, "", 3)
			for i, val := range got {
				if val != v.want[i] {
					t.Errorf("Got %v, want %v", got, v.want)
				}
			}
		}
	})
}
