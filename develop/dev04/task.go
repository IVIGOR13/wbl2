package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	fmt.Println(len(*findAnagram(&[]string{"пятка", "пятак", "листок", "столик", "Тяпка", "Слиток", "арваы"})))
}

func findAnagram(list *[]string) *map[string][]string {
	m := make(map[string][]string)
	mk := make(map[string]string)

	for _, word := range *list {
		word = strings.ToLower(word)
		sortWord := sortString(word)
		if v, ok := m[mk[sortWord]]; ok {
			m[mk[sortWord]] = append(v, word)
		} else {
			mk[sortWord] = word
			m[word] = []string{word}
		}
	}

	for k := range m {
		arr := m[k]

		if len(arr) == 1 {
			delete(m, k)
			continue
		}

		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
		m[k] = arr
	}

	return &m
}

func sortString(text string) string {
	runes := []rune(text)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
