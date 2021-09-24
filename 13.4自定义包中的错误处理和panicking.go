package main

import (
	"fmt"
	"strconv"
	"strings"
)

//这是所有自定义包实现者应该遵守的最佳实践：
//1)在包内部，总是应该从panic中recover: 不允许显示的超出包范围的panic()。
//2)向包的调用者返回错误值（而不是panic）。
//在包内部，特别是在非导出函数中有很深层次的嵌套调用时，对主调函数来说用panic来表示应该被翻译成错误的错误场景是很有用的（并且提高了代码可读性）。

//这在下面的代码中被很好地阐述了。我们有一个简单的parse包用来把输入的字符串解析为整数切片；这个包有自己特殊的ParseError。
// A ParseError indicates an error in converting a word into an integer.
type ParseError struct {
	Index int      // The index into the space-separated list of words.
	Word  string   // The word that generated the parse error.
	Err error // The raw error that precipitated this error, if any.
}

// String returns a human-readable error message.
func (e *ParseError) String() string {
	return fmt.Sprintf("pkg parse: error parsing %q as int", e.Word)
}

// Parse parses the space-separated words in in put as integers.
func Parse(input string) (numbers []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()

	fields := strings.Fields(input)
	numbers = fields2numbers(fields)
	return
}

func fields2numbers(fields []string) (numbers []int) {
	if len(fields) == 0 {
		panic("no words to parse")
	}
	for idx, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(&ParseError{idx, field, err})
		}
		numbers = append(numbers, num)
	}
	return
}
//当没有东西需要转换或者转换成整数失败时，这个包会panic（在函数fields2numbers中）。但是可导出的Parse函数会从panic中recover并用所有这些信息返回一个错误给调用者。为了演示这个过程，调用了parse包；不可解析的字符串会导致错误并被打印出来。

func main() {
	var examples = []string {
		"1 2 3 4 5",
		"100 50 25 12.5 6.25",
		"2 + 2 = 4",
		"1st class",
		"",
	}

	for _, ex := range examples {
		fmt.Printf("parsing %q\n", ex)
		nums, err := Parse(ex)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(nums)
	}
	/**
	parsing "1 2 3 4 5"
	[1 2 3 4 5]
	parsing "100 50 25 12.5 6.25"
	pkg: pkg parse: error parsing "12.5" as int
	parsing "2 + 2 = 4"
	pkg: pkg parse: error parsing "+" as int
	parsing "1st class"
	pkg: pkg parse: error parsing "1st" as int
	parsing ""
	pkg: no words to parse
	*/
}
