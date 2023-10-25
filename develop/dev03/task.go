package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {

	var (
		filename string = os.Args[len(os.Args)-1]
		k        int
		n        bool
		r        bool
		u        bool
	)

	flag.CommandLine.SetOutput(os.Stdout)
	flag.IntVar(&k, "k", 0, "")
	flag.BoolVar(&n, "n", false, "")
	flag.BoolVar(&r, "r", false, "")
	flag.BoolVar(&u, "u", false, "")
	flag.Parse()

	data := readFile(filename)

	if u {
		data = removeDuplicate(data)
	}

	sort.Slice(data, func(i, j int) bool {
		row1, row2 := data[i], data[j]

		if k > 0 {
			row1 = strings.Fields(data[i])[k-1]
			row2 = strings.Fields(data[j])[k-1]
		}
		if n {
			n1, _ := strconv.Atoi(row1)
			n2, _ := strconv.Atoi(row2)
			return n1 < n2
		}

		return row1 < row2
	})

	if r {
		slices.Reverse(data)
	}

	printOutput(data)

}

func removeDuplicate(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func readFile(filename string) []string {
	var data []string

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Unable to open file:", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
		}
		line = strings.Trim(line, "\n")
		line = strings.Trim(line, "\r")
		data = append(data, line)
	}

	return data
}

func printOutput(data []string) {
	for _, line := range data {
		fmt.Fprintln(os.Stdout, line)
	}
}
