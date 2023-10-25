package pattern

import "fmt"

/*
Паттерн Factory Method относится к порождающим паттернам уровня класса и сфокусирован только на отношениях между классами.
Он полезен, когда система должна оставаться легко расширяемой путем добавления объектов новых типов.
Этот паттерн является основой для всех порождающих паттернов и может легко трансформироваться под нужды системы.
По этому, если перед разработчиком стоят не четкие требования для продукта или не ясен способ организации взаимодействия
между продуктами, то для начала можно воспользоваться паттерном Factory Method, пока полностью не сформируются все требования.
Он применяется для создания объектов с определенным интерфейсом, реализации которого предоставляются потомками.
Другими словами, есть базовый абстрактный класс фабрики, который говорит, что каждая его наследующая фабрика должна реализовать такой-то метод для создания своих продуктов.
*/

type Storage interface {
	Open(string)
}

type StorageType int32

const (
	PgStorageType StorageType = iota
	MongoStorageType
	MemoryStorageType
)

func NewStorage(t StorageType) Storage {
	switch t {
	case PgStorageType:
		return &PgStorage{}
	case MongoStorageType:
		return &MongoStorage{}
	default:
		return &MemoryStorage{}
	}
}

type PgStorage struct{}

func (this *PgStorage) Open(str string) {
	fmt.Println("Open PG with:", str)
}

type MongoStorage struct{}

func (this *MongoStorage) Open(str string) {
	fmt.Println("Open MongoDB with:", str)
}

type MemoryStorage struct{}

func (this *MemoryStorage) Open(str string) {
	fmt.Println("Open Memory storage with:", str)
}
