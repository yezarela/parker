package park

import (
	"testing"
)

func TestInitSlots(t *testing.T) {
	p := NewPark()
	p.InitSlots(6)

	if slots := p.AvailSlots.Len(); slots != 6 {
		t.Errorf("expect slot capacity to be 6, got: %d", slots)
	}
}

func TestParkIn(t *testing.T) {
	p := NewPark()
	p.InitSlots(6)

	id, err := p.In(Car{RegNo: "B1", Color: "White"})
	if err != nil {
		t.Errorf("expect no error, got: %s", err.Error())
	}

	if *id != 1 {
		t.Errorf("expect slot number to be 1, got: %d", *id)
	}
}

func TestParkOut(t *testing.T) {
	p := NewPark()
	p.InitSlots(6)
	p.In(Car{RegNo: "B2", Color: "Black"})
	p.In(Car{RegNo: "B2", Color: "Black"})
	p.In(Car{RegNo: "B3", Color: "White"})

	id, err := p.Out(3)
	if err != nil {
		t.Errorf("expect no error, got: %s", err.Error())
	}

	if *id != 3 {
		t.Errorf("expect slot number to be 3, got: %d", *id)
	}
}

func TestGetSlot(t *testing.T) {
	p := NewPark()
	p.InitSlots(6)
	p.In(Car{RegNo: "B1", Color: "White"})

	s := p.GetSlot(1)
	if s == nil {
		t.Errorf("expect slot to be exist, got: nil")
	}
}
