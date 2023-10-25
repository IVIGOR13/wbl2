package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func printRow(row string, n bool, lineNumber int) {
	if n {
		fmt.Printf("%d: ", lineNumber)
	}
	fmt.Println(row)
}

func main() {

	var (
		A, B, C       int
		c, i, v, F, n bool
	)

	flag.CommandLine.SetOutput(os.Stdout)
	flag.IntVar(&A, "A", 0, "")
	flag.IntVar(&B, "B", 0, "")
	flag.IntVar(&C, "C", 0, "")
	flag.BoolVar(&c, "c", false, "")
	flag.BoolVar(&i, "i", false, "")
	flag.BoolVar(&v, "v", false, "")
	flag.BoolVar(&F, "F", false, "")
	flag.BoolVar(&n, "n", false, "")
	flag.Parse()

	if C != 0 {
		A = C
		B = C
	}

	reader := bufio.NewScanner(os.Stdin)

	pattern := os.Args[len(os.Args)-1]
	pattern = strings.Trim(pattern, "\n")
	pattern = strings.Trim(pattern, "\r")
	if i {
		pattern = strings.ToLower(pattern)
	}

	regex, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal(err)
	}

	before := make([]string, B)
	lastCoincidence := -1
	printedLines := []int{}

	coincidenceNumber := 0
	lineNumber := 0
	for reader.Scan() {

		lineNumber += 1

		srcLine := reader.Text()

		line := srcLine
		line = strings.Trim(line, "\n")
		line = strings.Trim(line, "\r")
		if i {
			line = strings.ToLower(line)
		}

		coincidence := regex.MatchString(line)
		if F {
			coincidence = line == pattern
		}

		if v {
			coincidence = !coincidence
		}

		if coincidence {
			coincidenceNumber += 1
			lastCoincidence = lineNumber
			if c {
				continue
			}

			for i := 0; i < len(before); i++ {
				lineN := lineNumber - len(before) + i
				if lineN < 1 || slices.Contains(printedLines, lineN) {
					continue
				}
				printRow(before[i], n, lineN)
			}
			before = []string{}

			if n {
				fmt.Printf("%d: ", lineNumber)
			}
			fmt.Println(srcLine)
		} else if A > 0 && lastCoincidence > 0 {
			if lineNumber-lastCoincidence <= A {
				printRow(srcLine, n, lineNumber)
				printedLines = append(printedLines, lineNumber)
			}
		}
		if !coincidence {
			before = append(before, srcLine)
			if len(before) > B {
				before = before[1:]
			}
		}
	}

	if c {
		fmt.Println(coincidenceNumber)
	}

	if err := reader.Err(); err != nil {
		log.Fatal(err)
	}
}
