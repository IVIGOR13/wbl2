package pattern

/*
Паттерн Builder относится к порождающим паттернам уровня объекта.
Он определяет процесс поэтапного построения сложного продукта.
После того как будет построена последняя его часть, продукт можно использовать.
*/

type House struct {
	Square      float64
	FloorNumber uint32
}

type Builder interface {
	SetSquare(val float64)
	SetFloorNumber(val uint64)
}

type Director struct {
	builder Builder
}

func (d *Director) Build() {
	d.builder.SetSquare(54.3)
	d.builder.SetFloorNumber(2)
}

type HouseBuilder struct {
	house *House
}

func (h *HouseBuilder) SetSquare(val float64) {
	h.house.Square = val
}

func (h *HouseBuilder) SetFloorNumber(val uint32) {
	h.house.FloorNumber = val
}
