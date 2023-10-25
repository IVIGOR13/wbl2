package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func removeEmptyStrings(input []string) []string {
	result := []string{}
	for _, str := range input {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}

func main() {

	var (
		f, d string
		s    bool
	)

	flag.CommandLine.SetOutput(os.Stdout)
	flag.StringVar(&f, "f", "", "")
	flag.StringVar(&d, "d", " ", "")
	flag.BoolVar(&s, "s", false, "")
	flag.Parse()

	fieldsStr := strings.Split(f, ",")
	fields := []int{}
	fromToEnd := -1
	for _, v := range fieldsStr {
		if strings.Contains(v, "-") {
			if v[0] == '-' {
				val, _ := strconv.Atoi(v[1:])
				for i := 0; i < val; i++ {
					fields = append(fields, i)
				}
			} else if v[len(v)-1] == '-' {
				val, _ := strconv.Atoi(v[:len(v)-1])
				fromToEnd = val - 1
			} else {
				vArr := strings.Split(v, "-")
				a, _ := strconv.Atoi(vArr[0])
				b, _ := strconv.Atoi(vArr[1])
				for ; a <= b; a++ {
					fields = append(fields, a)
				}
			}
		} else {
			val, _ := strconv.Atoi(v)
			fields = append(fields, val-1)
		}
	}

	reader := bufio.NewScanner(os.Stdin)

	for reader.Scan() {
		srcLine := reader.Text()
		lineArr := removeEmptyStrings(strings.Split(srcLine, d))
		if len(lineArr) == 1 && s {
			continue
		}

		var lineOut string

		if f == "" && fromToEnd == -1 {
			lineOut = strings.Join(lineArr, "\t")
		} else {

			builder := strings.Builder{}
			for i, v := range lineArr {
				if slices.Contains(fields, i) || (fromToEnd <= i && fromToEnd != -1) {
					builder.WriteString(v)
					builder.WriteString(d)
				}
			}

			lineOut = builder.String()
			if len(lineOut) != 0 {
				lineOut = lineOut[:len(lineOut)-len(d)]
			}
		}

		if lineOut != "" {
			fmt.Println(lineOut)
		}
	}

}
