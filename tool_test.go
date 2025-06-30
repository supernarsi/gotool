package gotool_test

import (
	"context"
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/supernarsi/gotool"
	"github.com/supernarsi/gotool/util"
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

func TestCoroutineForeach(t *testing.T) {
	// 测试 int 类型
	t.Run("test int", func(t *testing.T) {
		process1 := func(ctx context.Context, i int) (int, error) {
			return i * 2, nil
		}
		process2 := func(ctx context.Context, i int) (int, error) {
			return i - 1, nil
		}
		tests := []struct {
			name    string
			input   []int
			process func(ctx context.Context, i int) (int, error)
			want    []int
		}{
			{"t1", []int{1, 2, 3}, process1, []int{2, 4, 6}},
			{"t2", []int{1, 1, 2, 0}, process1, []int{2, 2, 4, 0}},
			{"t3", []int{4, 3, 2, 1, 0}, process2, []int{3, 2, 1, 0, -1}},
		}

		for _, v := range tests {
			got := util.GoForeach(context.Background(), v.input, v.process, 0, 100)
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
			{"s3", []string{"a", " ", "+", "B", "dBa"}, p2, []string{"A", " ", "+", "B", "DBA"}},
		}
		for _, v := range tests {
			got := util.GoForeach(context.Background(), v.input, v.process, "", 100)
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
			{"s1", []int{1, 2, 3}, p1, []string{"1_", "2_", "3_"}},
			{"s2", []int{}, p1, []string{}},
			{"s3", []int{0, 2}, p1, []string{"0_", "2_"}},
		}
		for _, v := range tests {
			got := util.GoForeach(context.Background(), v.input, v.process, "", 100)
			for i, val := range got {
				if val != v.want[i] {
					t.Errorf("Got %v, want %v", got, v.want)
				}
			}
		}
	})
}

func TestRandTasks(t *testing.T) {
	tests := []struct {
		name     string
		inputArr []int
		inputNum uint
		want     []int
	}{
		{name: "t1", inputArr: []int{1, 2, 3, 4, 5, 6, 7}, inputNum: 3, want: []int{2, 1, 3}},
		{name: "t2", inputArr: []int{1, 2, 3, 4, 5, 6, 7}, inputNum: 0, want: []int{3, 5, 6}},
		{name: "t3", inputArr: []int{1, 2, 3, 4, 5, 6, 7}, inputNum: 10, want: []int{1, 6, 7}},
		{name: "t4", inputArr: []int{1, 2, 3, 4, 5, 6, 7}, inputNum: 999, want: []int{7, 4, 3}},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := gotool.RandInt(v.inputArr, v.inputNum, 3)
			for k, gV := range got {
				if gV != v.want[k] {
					t.Errorf("%s, want %v, got %v", v.name, v.want, got)

				}
			}
		})
	}
}

func TestDifference(t *testing.T) {
	tests := []struct {
		name   string
		inputA []uint
		inputB []uint
		want   []uint
	}{
		{"t1", []uint{1, 2, 3}, []uint{1, 2, 3}, []uint{}},
		{"t2", []uint{2, 3}, []uint{1, 2, 3}, []uint{}},
		{"t3", []uint{1, 2, 3}, []uint{1, 2}, []uint{3}},
		{"t4", []uint{2, 3}, []uint{1, 2, 4}, []uint{3}},
		{"t5", []uint{}, []uint{1, 2, 4}, []uint{}},
		{"t6", []uint{1, 2, 3, 4}, []uint{}, []uint{1, 2, 3, 4}},
		{"t7", []uint{1, 2, 3, 4}, []uint{}, []uint{1, 2, 3, 4}},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := gotool.Difference(v.inputA, v.inputB)
			for k, gV := range got {
				if gV != v.want[k] {
					t.Errorf("%s, want %v, got %v", v.name, v.want, got)
				}
			}
		})
	}
}

func TestAssignGroup(t *testing.T) {
	tests := []struct {
		name      string
		inputUid  uint32
		inputSeed uint32
		want      uint32
	}{
		{"t0_1", 10001, 0, 38},
		{"t0_2", 10002, 0, 82},
		{"t0_3", 0, 0, 54},
		{"t0_4", 1, 0, 18},
		{"t0_5", 2, 0, 63},
		{"t0_6", 10, 0, 17},
		{"t0_7", 999, 0, 87},
		{"t0_8", 99999999, 0, 52},
		{"t0_9", 399999999, 0, 71},

		{"t1_1", 10001, 1, 2},
		{"t1_2", 10002, 1, 97},
		{"t1_3", 0, 1, 45},
		{"t1_4", 1, 1, 56},
		{"t1_5", 2, 1, 44},
		{"t1_6", 10, 1, 70},
		{"t1_7", 999, 1, 13},
		{"t1_8", 99999999, 1, 56},
		{"t1_9", 399999999, 1, 80},

		{"t100_1", 10001, 100, 26},
		{"t100_2", 10002, 100, 64},
		{"t100_3", 0, 100, 87},
		{"t100_4", 1, 100, 97},
		{"t100_5", 2, 100, 22},
		{"t100_6", 10, 100, 35},
		{"t100_7", 999, 100, 66},
		{"t100_8", 99999999, 100, 78},
		{"t100_9", 399999999, 100, 19},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := gotool.AssignGroup(v.inputUid, v.inputSeed); got != v.want {
				t.Errorf("AssignGroup() = %v, want %v", got, v.want)
			}
		})
	}
}

func TestFloatRatioToInt(t *testing.T) {
	tests := []struct {
		name  string
		input []float32
		want  []int
	}{
		{name: "t1", input: []float32{1, 2, 3}, want: []int{17, 33, 50}},
		{name: "t2", input: []float32{1, 1, 1, 1}, want: []int{25, 25, 25, 25}},
		{name: "t3", input: []float32{0.1, 0.1, 0.1, 0.1}, want: []int{25, 25, 25, 25}},
		{name: "t4", input: []float32{0.5, 0.5}, want: []int{50, 50}},
		{name: "t5", input: []float32{0.5, 0, 5}, want: []int{9, 0, 91}},
		{name: "t6", input: []float32{0.5, 1, 2, 3, 4, 5, 10, 0, 0.1, 0.2, 0.3}, want: []int{2, 4, 8, 11, 15, 19, 38, 0, 0, 1, 2}},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := gotool.FloatRatioToInt(v.input)
			for k, gotV := range got {
				if gotV != v.want[k] {
					t.Errorf("%s, want %v, got %v", v.name, v.want, got)
					break
				}
			}
		})
	}
}

func TestLottery(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"t1", []int{3}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gotool.Lottery(tt.input); got != tt.want {
				t.Errorf("Lottery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaskNickname(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"空字符串", "", ""},
		{"单个ASCII字符", "a", "a"},
		{"两个ASCII字符", "ab", "ab"},
		{"三个ASCII字符", "abc", "ab*"},
		{"四个ASCII字符", "abcd", "ab**"},
		{"五个ASCII字符", "abcde", "ab***"},
		{"单个中文字符", "中", "中"},
		{"两个中文字符", "中文", "中*"},
		{"三个中文字符", "中文名", "中**"},
		{"四个中文字符", "中文名字", "中***"},
		{"五个中文字符", "中文名字啊", "中***"},
		{"中英混合-中英", "中a", "中*"},
		{"中英混合-英中", "a中", "a中"},
		{"中英混合-中英中", "中a中", "中**"},
		{"中英混合-英中英", "a中a", "a中*"},
		{"全角字符", "ＡＢＣ", "Ａ**"},
		{"全角字符混合", "ＡＢＣＤ", "Ａ***"},
		{"emoji表情", "😀😃😄", "😀**"},
		{"emoji和文字混合", "😀中文", "😀**"},
		{"特殊符号", "!@#$%", "!@***"},
		{"数字", "12345", "12***"},
		{"空格", "a b c", "a ***"},
		{"制表符", "a\tb\tc", "a\t***"},
		{"换行符", "a\nb\nc", "a\n***"},
		{"日文字符", "あいう", "あ**"},
		{"韩文字符", "가나다", "가**"},
		{"混合字符-英中日", "a中b", "a中*"},
		{"混合字符-中日英", "中a日", "中**"},
		{"全角数字", "１２３", "１**"},
		{"全角字母", "ＡＢＣＤＥ", "Ａ***"},
		{"全角符号", "！＠＃", "！**"},
		{"边界情况-单字节两个字符", "ab", "ab"},
		{"边界情况-多字节两个字符", "中文", "中*"},
		{"边界情况-单字节三个字符", "abc", "ab*"},
		{"边界情况-多字节三个字符", "中文名", "中**"},
		{"长字符串-单字节", "abcdefghijklmnop", "ab***"},
		{"长字符串-多字节", "中文名字很长很长很长", "中***"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gotool.MaskNickname(tt.input)
			if result != tt.expected {
				t.Errorf("MaskNickname(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsMultibyte(t *testing.T) {
	tests := []struct {
		name     string
		input    rune
		expected bool
	}{
		{
			name:     "ASCII字符",
			input:    'a',
			expected: false,
		},
		{
			name:     "ASCII数字",
			input:    '1',
			expected: false,
		},
		{
			name:     "ASCII符号",
			input:    '!',
			expected: false,
		},
		{
			name:     "ASCII空格",
			input:    ' ',
			expected: false,
		},
		{
			name:     "ASCII制表符",
			input:    '\t',
			expected: false,
		},
		{
			name:     "ASCII换行符",
			input:    '\n',
			expected: false,
		},
		{
			name:     "中文字符",
			input:    '中',
			expected: true,
		},
		{
			name:     "日文字符",
			input:    'あ',
			expected: true,
		},
		{
			name:     "韩文字符",
			input:    '가',
			expected: true,
		},
		{
			name:     "emoji表情",
			input:    '😀',
			expected: true,
		},
		{
			name:     "全角字符-字母",
			input:    'Ａ',
			expected: true,
		},
		{
			name:     "全角字符-数字",
			input:    '１',
			expected: true,
		},
		{
			name:     "全角字符-符号",
			input:    '！',
			expected: true,
		},
		{
			name:     "全角字符-空格",
			input:    '　',
			expected: true,
		},
		{
			name:     "边界值-ASCII最大值",
			input:    127,
			expected: false,
		},
		{
			name:     "边界值-ASCII最大值+1",
			input:    128,
			expected: true,
		},
		{
			name:     "边界值-全角字符开始",
			input:    0xFF01,
			expected: true,
		},
		{
			name:     "边界值-全角字符结束",
			input:    0xFF5E,
			expected: true,
		},
		{
			name:     "边界值-全角字符结束+1",
			input:    0xFF5F,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gotool.IsMultibyte(tt.input)
			if result != tt.expected {
				t.Errorf("isMultibyte(%q) = %v, want %v", string(tt.input), result, tt.expected)
			}
		})
	}
}

// BenchmarkMaskNickname 性能测试
func BenchmarkMaskNickname(b *testing.B) {
	testCases := []string{
		"",
		"a",
		"ab",
		"abc",
		"abcd",
		"abcde",
		"中",
		"中文",
		"中文名",
		"中文名字",
		"中a",
		"a中",
		"中a中",
		"a中a",
		"ＡＢＣ",
		"ＡＢＣＤ",
		"😀😃😄",
		"😀中文",
		"!@#$%",
		"12345",
		"a b c",
		"a\tb\tc",
		"a\nb\nc",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, input := range testCases {
			gotool.MaskNickname(input)
		}
	}
}

// BenchmarkIsMultibyte 性能测试
func BenchmarkIsMultibyte(b *testing.B) {
	testCases := []rune{
		'a', '1', '!', ' ', '\t', '\n',
		'中', 'あ', '가', '😀',
		'Ａ', '１', '！', '　',
		127, 128, 0xFF01, 0xFF5E, 0xFF5F,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, input := range testCases {
			gotool.IsMultibyte(input)
		}
	}
}
