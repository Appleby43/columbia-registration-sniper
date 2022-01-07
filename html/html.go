//html handles the parsing of strings into a tree o
package html

//Element represents an HTML Element - composed of starting and
//ending tags as well as 'content' between. Content may contain more html.
type Element struct {
	html *string
	startTag Tag //<p>
	endTag Tag //</p>
}

type Tag struct {
	startIndex int // <
	endIndex int // >
}

//ElementAround gathers indicies of carats in html tags with referece to some starting index within the tags
func ElementAround(inpt string, index int) Element {
	startTagEnd := findAt('>', inpt, index, true)
	startTagStart := findAt('<', inpt, startTagEnd, true)

	endTagStart := findAt('<', inpt, index, false)
	endTagEnd := findAt('>', inpt, endTagStart, false)

	return Element{
		html: &inpt,
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

//FindNextElement looks for the next set of HTML tags directly after this one
//Note that this approach is entirely unsound except for the special case where the next adjacent
//element has no children, and only a text body.
//Since all the data we're parsing exists in adjacent <td> elements, we're (hopefully) safe to use this here. 
//Largely, this approach is just a workaround to avoid parsing the whole html document.
func (e *Element) FindNextElement() Element {
	//find an index in the next element and run ElementAround
	startTagEnd := findAt('>', *e.html, e.endTag.endIndex + 1, false)
	return ElementAround(*e.html, startTagEnd + 1)
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