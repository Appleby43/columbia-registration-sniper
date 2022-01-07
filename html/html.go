//html handles the parsing of strings into a tree o
package html

//Element represents an HTML Element - composed of starting and
//ending tags as well as 'content' between. Content may contain more html.
type Element struct {
	startTag Tag //<p>
	endTag Tag //</p>
}

type Tag struct {
	startIndex int // <
	endIndex int // >
}

//ElementAround gathers indicies of carats in html tags
func ElementAround(inpt string, index int) Element {
	startTagEnd := findAt('>', inpt, index, true)
	startTagStart := findAt('<', inpt, startTagEnd, true)

	endTagStart := findAt('<', inpt, index, false)
	endTagEnd := findAt('>', inpt, endTagStart, false)

	return Element{
		startTag: Tag{
			startIndex: startTagStart,
			endIndex: startTagEnd,
		},
		endTag: Tag{
			startIndex: endTagStart,
			endIndex: endTagEnd,
		},
	}
}


func find(target rune, inpt string, reverse bool) int {
	i := 0;

	if reverse {
		i = len(inpt) - 1
	}

	return findAt(target, inpt, i, reverse)
}

func findAt(target rune, inpt string, startingPoint int, reverse bool) int {
	if len(inpt) == 0 {
		return -1
	}

	i := startingPoint;

	runeInpt := []rune(inpt)
	increment := 1;
	finalIndex := len(inpt) - 1

	if reverse {
		increment = -1
		finalIndex = 0;
	}

	for {
		if runeInpt[i] == target {
			return i
		}

		if i == finalIndex {
			break
		}

		i += increment
	}
	return -1
}