package html

import (
	"strings"
	"testing"
)

func TestFind(t *testing.T) {
	result := find('k', "hike", false)

	if result != 2 {
		t.Fail()
	}
}

func TestReverse(t *testing.T) {
	result := find('w', "hello world", false)
	resultReverse := find('w', "hello world", true)

	if result != resultReverse {
		t.Fail()
	}
}

func TestNoTarget(t *testing.T) {
	result := find('u', "hello world", false)

	if result != -1 {
		t.Fail()
	}
}

func TestEmptyString(t *testing.T) {
	result := find('a', "", false)

	if result != -1 {
		t.Fail()
	}
}

func TestTwoOccurrances(t *testing.T) {
	leftR := find('r', "racecar", false)
	rightR := find('r', "racecar", true)

	if leftR != 0 {
		t.Fail()
	}
	if rightR != 6 {
		t.Fail()
	}
}

func TestUnicode(t *testing.T) {
	result := find('🙏', "✨ab🙏cd😂", false)

	if result != 3 {
		t.Fail()
	}
}

func TestFindAt(t *testing.T) {
	//start at index 8 (space) and go right
	uInUniversity := findAt('u', "columbia university", 8, false)
	uInColumbia := findAt('u', "columbia university", 8, true) //go left

	if uInUniversity != 9 {
		t.Fail()
	}
	if uInColumbia != 3 {
		t.Fail()
	}
}

func TestElementAround(t *testing.T) { 
	inpt := "<p> Hello World! </p>"
	helloIndex := strings.Index(inpt, "Hello World!");

	result := ElementAround(inpt, helloIndex)

	if result.startTag.startIndex != 0 {
		t.Fail()
	}

	if result.startTag.endIndex != 2 {
		t.Fail()
	}

	if result.endTag.endIndex != len(inpt) - 1 {
		t.Fail()
	}

	if result.endTag.startIndex != len(inpt) - 4 {
		t.Fail()
	}
}