package pattern

import (
	"fmt"
	"math"
)

/*
Паттерн Visitor относится к поведенческим паттернам уровня объекта.
Паттерн Visitor позволяет обойти набор элементов (объектов) с разнородными интерфейсами,
а также позволяет добавить новый метод в класс объекта, при этом, не изменяя сам класс этого объекта.
*/

type Shape interface {
	getType() string
	accept(Visitor)
}

type Triangle struct {
	a float64
	b float64
	c float64
}

func (this *Triangle) getType() string {
	return "Triangle"
}

func (this *Triangle) accept(v Visitor) {
	v.visitForTriangle(this)
}

type Circle struct {
	radius float64
}

func (this *Circle) getType() string {
	return "Circle"
}

func (this *Circle) accept(v Visitor) {
	v.visitForCircle(this)
}

type Visitor interface {
	visitForTriangle(*Triangle)
	visitForCircle(*Circle)
}

type SquareCalc struct{}

func (this *SquareCalc) visitForTriangle(t *Triangle) {
	p := (t.a + t.b + t.c) / 2
	square := math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
	fmt.Println(square)
}

func (this *SquareCalc) visitForCircle(c *Circle) {
	square := math.Pi * math.Pow(c.radius, 2)
	fmt.Println(square)
}
