package pattern

import "fmt"

/*
Паттерн Strategy относится к поведенческим паттернам уровня объекта.
Он определяет набор алгоритмов схожих по роду деятельности, инкапсулирует их в отдельный класс и делает их подменяемыми.
Паттерн Strategy позволяет подменять алгоритмы без участия клиентов, которые используют эти алгоритмы.
*/

type FindStrategy interface {
	findRoute(int, int)
}

type DFS struct{}

func (this *DFS) findRoute(start, end int) {
	fmt.Printf("DFS find route from %d to %d\n", start, end)
}

type BFS struct{}

func (this *BFS) findRoute(start, end int) {
	fmt.Printf("BFS find route from %d to %d\n", start, end)
}

type ASharp struct{}

func (this *ASharp) findRoute(start, end int) {
	fmt.Printf("ASharp find route from %d to %d\n", start, end)
}

type Map struct {
	findStrategy FindStrategy
}

func (this *Map) setFindAlgoritm(findStrategy *FindStrategy) {
	this.findStrategy = *findStrategy
}
