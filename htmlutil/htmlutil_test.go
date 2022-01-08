package htmlutil

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

func TestNextElement(t *testing.T) { 
	inpt := "<td> Hello </td><td> World! </td>"
	helloIndex := strings.Index(inpt, "Hello");

	firstElement := ElementAround(inpt, helloIndex)
	secondElement := firstElement.FindNextElement()

	if secondElement.startTag.startIndex != 16 {
		t.Fail()
	}

	if secondElement.startTag.endIndex != 19 {
		t.Fail()
	}

	if secondElement.endTag.endIndex != len(inpt) - 1 {
		t.Fail()
	}

	if secondElement.endTag.startIndex != len(inpt) - 5 {
		t.Fail()
	}
}

func TestContents(t *testing.T) {
	inpt := "<html>Hello World!</html>"
	element := ElementAround(inpt, 10) // 10 is fairly arbitrary

	if element.Contents() != "Hello World!" {
		t.Fail()
	}
}

func TestSplice(t *testing.T) {
	output := splice("foobar", 1, 3)
	
	if output != "far" {
		t.Fail()
	}
}

func TestSpliceFull(t *testing.T){
	output := splice("hello", 0, len("hello") - 1)

	if output != "" {
		t.Fail()
	}
}

func TestStripMeta(t *testing.T) {
	inpt := "<meta aslkdjslkafjlkasdjfldskajf><meta><meta foo=bar>hello world<meta>"
	outpt := StripMeta(inpt)

	if outpt != "hello world" {
		t.Fail()
	}
}