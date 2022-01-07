package course

import "testing"

func TestFull(t *testing.T) {
	c := course{
		capacity:   10,
		enrollment: 10,
	}

	if !c.isFull() {
		t.Fail()
	}
}

func TestBelowCapactiy(t *testing.T){
	c := course{
		capacity:   50,
		enrollment: 25,
	}

	if c.isFull() {
		t.Fail()
	}
}

func TestAboveCapacity(t *testing.T){
	c := course{
		capacity:   100,
		enrollment: 101,
	}

	if !c.isFull() {
		t.Fail()
	}
}