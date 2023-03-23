package gotool_test

import (
	"math"
	"testing"

	"github.com/supernarsi/gotool"
)

type uniqueElementsTestCase[T gotool.ElementType] struct {
	name  string
	input []T
	want  []T
}

type inArrayTestCase[T gotool.ElementType] struct {
	name   string
	target T
	arr    []T
	want   bool
}

func runUniqueElementsTestCases[T gotool.ElementType](t *testing.T, cases []uniqueElementsTestCase[T]) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := gotool.UniqueElements(tt.input); len(got) != len(tt.want) {
				t.Errorf("got result %v, want %v", got, tt.want)
			} else {
				for i, v := range got {
					if v != tt.want[i] {
						t.Errorf("got result %v, want %v", got, tt.want)
					}
				}
			}
		})
	}
}

func runInArrayTestCases[T gotool.ElementType](t *testing.T, cases []inArrayTestCase[T]) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := gotool.InArray(tt.target, tt.arr); got != tt.want {
				t.Errorf("got result %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInArray(t *testing.T) {
	// 测试 uint array
	uintCases := []inArrayTestCase[uint]{
		{name: "uint 1 not in", target: 1, arr: []uint{0, 11, 10, 0, 111}, want: false},
		{name: "uint 1 in", target: 1, arr: []uint{0, 11, 10, 0, 1, 111}, want: true},
		{name: "uint 0 in", target: 0, arr: []uint{0, 00, 1, 2, 3}, want: true},
		{name: "uint 0 not in", target: 0, arr: []uint{10, 20, 1, 2, 3}, want: false},
	}
	// 测试 int array
	intCases := []inArrayTestCase[int]{
		{name: "int 1 not in", target: 1, arr: []int{0, -1, 11, 10, 0, 111}, want: false},
		{name: "int 1 in", target: 1, arr: []int{0, 11, -10, 0, 1, 111, 11, -11}, want: true},
		{name: "int 0 in", target: 0, arr: []int{0, 00, -1, 2, 3}, want: true},
		{name: "int 0 not in", target: 0, arr: []int{10, 20, 1, 2, 3}, want: false},
		{name: "int 9999999 in", target: 9999999, arr: []int{10, 20, 9999999, 2, 3}, want: true},
		{name: "int 999999999999 in", target: 999999999999, arr: []int{10, 20, 999999999999, 2, 3}, want: true},
		{name: "int max in", target: math.MaxInt, arr: []int{10, 20, math.MaxInt, 2, 3}, want: true},
		{name: "int max not in", target: math.MaxInt, arr: []int{10, 20, math.MaxInt - 1, 2, 3}, want: false},
		{name: "int min in", target: math.MinInt, arr: []int{10, 20, math.MinInt, 2, 3}, want: true},
		{name: "int min not in", target: math.MinInt, arr: []int{10, 20, math.MinInt + 1, 2, 3}, want: false},
	}
	// 测试 float array
	floatCases := []inArrayTestCase[float64]{
		{name: "float 1 not in", target: 1, arr: []float64{0, -1, 1.1, 1.01, 0.0, 111, 00.000000000001}, want: false},
		{name: "float 1 in", target: 1, arr: []float64{0, 11, -10, 0, 1.00000000000000000, 111, 11, -11}, want: true},
		{name: "float 0 in", target: 0, arr: []float64{0.00, 0.001, 0001, -1, 2, 3}, want: true},
		{name: "float 0 not in", target: 0, arr: []float64{10, 2.0, 0.000000000001, 1, 2, 3}, want: false},
		{name: "float 99999999 in", target: 99999999, arr: []float64{10, 20, 99999999.0000, 2, 3}, want: true},
		{name: "float 999999999999 not in", target: 999999999999, arr: []float64{10, 20, 999999999999.00000000001, 2, 3}, want: true},
		{name: "float 99999999999.00001 in", target: 99999999999.00001, arr: []float64{10, 20, 99999999999.00001, 2, 3}, want: true},
		{name: "float 99999999999.00001 in", target: 99999999999.0001, arr: []float64{10, 20, 99999999999.00001, 2, 3}, want: false},
		{name: "float max in", target: math.MaxFloat64, arr: []float64{10, 20, math.MaxFloat64, 2, 3}, want: true},
	}
	// 测试 string array
	stringCases := []inArrayTestCase[string]{
		{name: "str 1 not in", target: "1", arr: []string{" ", "1 ", " 1", " 1 ", "11", "01", "1.0", "1a"}, want: false},
		{name: "str 1 in", target: "1", arr: []string{" ", "1 ", " 1", " 1 ", "1", "01"}, want: true},
		{name: "str empty in", target: "", arr: []string{" ", "", "0"}, want: true},
		{name: "str empty not in", target: "", arr: []string{" ", "  ", ".", "0"}, want: false},
		{name: "str + in", target: "+", arr: []string{"=", "++", "+", " ", "十"}, want: true},
		{name: "str + not in", target: "+", arr: []string{"=", "++", "-", " ", "十"}, want: false},
		{name: "str not in", target: "123", arr: []string{" 123", "12", "123 ", "123a", "1230", "0123", "123."}, want: false},
	}

	runInArrayTestCases(t, uintCases)
	runInArrayTestCases(t, intCases)
	runInArrayTestCases(t, floatCases)
	runInArrayTestCases(t, stringCases)
}

func TestUniqueElements(t *testing.T) {
	// 测试 uint
	uintCases := []uniqueElementsTestCase[uint]{
		{name: "test1", input: []uint{1, 2, 3, 1, 2}, want: []uint{1, 2, 3}},
		{name: "test2", input: []uint{0, 2, 3, 1, 2}, want: []uint{0, 2, 3, 1}},
		{name: "test3", input: []uint{0, 0, 3, 1, 1, 3, 1, 2}, want: []uint{0, 3, 1, 2}},
		{name: "test4", input: []uint{0, 0}, want: []uint{0}},
	}
	// 测试 int
	intCases := []uniqueElementsTestCase[int]{
		{name: "test1", input: []int{1, 2, -3, 1, 2}, want: []int{1, 2, -3}},
		{name: "test2", input: []int{0, 2, 3, -1, 1, -1, 2}, want: []int{0, 2, 3, -1, 1}},
		{name: "test3", input: []int{0, 0, 3, 1, 1, -3, 1, 2}, want: []int{0, 3, 1, -3, 2}},
		{name: "test4", input: []int{0, 0}, want: []int{0}},
	}
	// 测试 float
	floatCases := []uniqueElementsTestCase[float64]{
		{name: "test1", input: []float64{1, 2, -3, 1, 2, 2.0, 0, 0.0, -3.0, -3.1}, want: []float64{1, 2, -3, 0, -3.1}},
		{name: "test2", input: []float64{1.1, 1.0, 1, 0, 0.0, 1.01, 0.00, 0.01, 0}, want: []float64{1.1, 1.0, 0, 1.01, 0.01}},
		{name: "test3", input: []float64{1.0, 1.00, 1}, want: []float64{1}},
		{name: "test4", input: []float64{1, -1.00, 1.0}, want: []float64{1, -1}},
	}
	// 测试 string
	stringCases := []uniqueElementsTestCase[string]{
		{name: "test1", input: []string{"", " ", "", "0", "123", "123abc", "a 1"}, want: []string{"", " ", "0", "123", "123abc", "a 1"}},
		{name: "test2", input: []string{"aaa", "AAA", "aAa", "a aa"}, want: []string{"aaa", "AAA", "aAa", "a aa"}},
		{name: "test3", input: []string{"1.0", "1", "a1", "a1.0", "a ", "a"}, want: []string{"1.0", "1", "a1", "a1.0", "a ", "a"}},
		{name: "test4", input: []string{"1.0", "1", "1.0", "1.00"}, want: []string{"1.0", "1", "1.00"}},
		{name: "test5", input: []string{"a", "a", "a ", " a", " a ", " a", "  a", "a  "}, want: []string{"a", "a ", " a", " a ", "  a", "a  "}},
	}

	runUniqueElementsTestCases(t, uintCases)
	runUniqueElementsTestCases(t, intCases)
	runUniqueElementsTestCases(t, floatCases)
	runUniqueElementsTestCases(t, stringCases)
}
