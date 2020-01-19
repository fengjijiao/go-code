package main
import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str1 :="asSASA ddd dsjkdsjs dk"
	str2 :="asSASA ddd dsjkdsjsこん dk"
	fmt.Println("str1 len():",len(str1))
	fmt.Println("str1 utf8.RunneCountInString:",utf8.RuneCountInString(str1))
	fmt.Println("str2 len():",len(str2))
	fmt.Println("str2 utf8.RuneCountInString:",utf8.RuneCountInString(str2))
}
