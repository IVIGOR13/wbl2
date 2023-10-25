package pattern

import (
	"strings"
)

/*
Паттерн Facade относится к структурным паттернам уровня объекта.
Паттерн Facade предоставляет высокоуровневый унифицированный интерфейс в виде набора имен методов к набору
взаимосвязанных классов или объектов некоторой подсистемы, что облегчает ее использование.
Разбиение сложной системы на подсистемы позволяет упростить процесс разработки, а также помогает максимально снизить
зависимости одной подсистемы от другой. Однако использовать такие подсистемы становиться довольно сложно.
Один из способов решения этой проблемы является паттерн Facade. Наша задача, сделать простой, единый интерфейс,
через который можно было бы взаимодействовать с подсистемами.
*/

func NewFacade() *Facade {
	return &Facade{
		a: &A{},
		b: &B{},
		c: &C{},
	}
}

type Facade struct {
	a *A
	b *B
	c *C
}

func (m *Facade) execute() string {
	result := []string{
		m.a.TaskA(),
		m.b.TaskB(),
		m.c.TaskC(),
	}
	return strings.Join(result, " + ")
}

type A struct{}

func (this *A) TaskA() string {
	return "executed task A"
}

type B struct{}

func (this *B) TaskB() string {
	return "executed task B"
}

type C struct{}

func (this *C) TaskC() string {
	return "executed task C"
}
