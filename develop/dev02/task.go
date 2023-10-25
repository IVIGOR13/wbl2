package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	fmt.Println(fmt.Sprintf("[%s]", Unpacking("a4bc2d5e")))
	fmt.Println(fmt.Sprintf("[%s]", Unpacking("abcd")))
	fmt.Println(fmt.Sprintf("[%s]", Unpacking("45")))
	fmt.Println(fmt.Sprintf("[%s]", Unpacking("")))

	fmt.Println(fmt.Sprintf("[%s]", Unpacking(`qwe\4\5`)))
	fmt.Println(fmt.Sprintf("[%s]", Unpacking(`qwe\45`)))
	fmt.Println(fmt.Sprintf("[%s]", Unpacking(`qwe\\5`)))
}

func Unpacking(str string) string {
	result := strings.Builder{}

	var escape bool = false
	var prev rune

	for _, v := range []rune(str) {

		if escape {
			result.WriteRune(v)
			prev = v
			escape = false
			continue
		}

		if v == '\\' {
			escape = true
		} else if unicode.IsLetter(v) {
			result.WriteRune(v)
			prev = v
		} else if unicode.IsDigit(v) && (prev != 0) {
			count := int(v) - '0'
			for ; count > 1; count-- {
				result.WriteRune(prev)
			}
		}
	}

	return result.String()
}
