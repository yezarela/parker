package park

import (
	"container/heap"
	"errors"
	"strings"
	"sync"
)

var (
	instance *Park
	once     sync.Once
)

// Car represents car
type Car struct {
	ID    int
	RegNo string
	Color string
}

// Park represents park
type Park struct {
	Capacity   int
	AvailSlots IntHeap
	Slots      map[int]Car
}

// NewPark returns singleton park instance
func NewPark() *Park {

	once.Do(func() { instance = &Park{} })

	return instance
}

// InitSlots initialize slots with given capacity
func (p *Park) InitSlots(cap int) {

	p.AvailSlots = IntHeap{}
	p.Capacity = cap
	p.Slots = map[int]Car{}

	for i := 1; i <= cap; i++ {
		p.AvailSlots = append(p.AvailSlots, i)
	}

	heap.Init(&p.AvailSlots)
}

// In allocates a new car to a nearest slot
func (p *Park) In(car Car) (*int, error) {

	if p.AvailSlots.Len() == 0 {
		return nil, errors.New(ErrNotAvailable)
	}

	avail := heap.Pop(&p.AvailSlots)
	id := avail.(int)

	car.ID = id
	p.Slots[id] = car

	return &id, nil
}

// Out removes a car from slot by id
func (p *Park) Out(id int) (*int, error) {

	if slot := p.GetSlot(id); slot == nil {
		return nil, errors.New(ErrNotFound)
	}

	heap.Push(&p.AvailSlots, id)

	delete(p.Slots, id)

	return &id, nil
}

// GetSlot returns slot by slot id
func (p *Park) GetSlot(id int) *Car {

	if s, exist := p.Slots[id]; exist {
		return &s
	}

	return nil
}

// Search returns list of cars by specified term
func (p *Park) Search(term string) []*Car {

	term = strings.ToLower(term)
	cars := []*Car{}

	for _, v := range p.Slots {
		if v.RegNo == term || v.Color == term {
			cars = append(cars, &v)
		}
	}

	return cars
}
