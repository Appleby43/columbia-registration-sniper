package course

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/Appleby43/columbia-registration-sniper/htmlutil"
)

type Course struct {
	CallNum int
	Professor string
	Capacity int
	Enrollment int
}

func (c *Course) isFull() bool {
	return c.Enrollment >= c.Capacity
}

func StripFrom(html string) (Course, error) {
	html = strings.ReplaceAll(html, "\n", "")
	html = htmlutil.StripMeta(html)
	callNumString := strings.TrimSpace(pullFromNextCell(html, "Call Number"))
	callNum, err := strconv.Atoi(callNumString)

	professor := pullFromNextCell(html, "Instructor")
	enrollmentDetails := pullFromNextCell(html, "Enrollment")

	enrollment, capacity := parseEnrollmentDetails(enrollmentDetails)

	return Course{
		CallNum: callNum,
		Professor: professor,
		Enrollment: enrollment,
		Capacity: capacity,
	}, err
}

//pullFromNextCell pulls data from HTML in the general form of
//<td>CellName</td><td>some data you want</td>.
func pullFromNextCell(html string, cellName string) string {
	cellIndex := strings.Index(html, cellName)
	cellElement := htmlutil.ElementAround(html, cellIndex)
	nextElement := cellElement.FindNextElement()
	return nextElement.Contents()
}

//parseEnrollmentDetails parses the strings displayed in the form:
//36 students (72 max) as of 8:55PM Friday, January 7, 2022
func parseEnrollmentDetails(details string) (enrollment int, capacity int){
	details = strings.TrimSpace(details)

	parenthesisIndex := strings.Index(details, "(")

	enrollment = parseIntAt(details, 0)
	capacity = parseIntAt(details, parenthesisIndex + 1)
	return
}

//parseIntAt tries to parse the longest integer possible starting from the given index (inclusive)
func parseIntAt(input string, index int) int{
	retVal := 0
	runeInput := []rune(input)

	for index < len(runeInput) {
		if unicode.IsDigit(runeInput[index]) {
			retVal *= 10
			retVal += int(runeInput[index] - '0');
		} else {
			return retVal
		}
		index++
	}
	return retVal
}