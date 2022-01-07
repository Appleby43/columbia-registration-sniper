package course

type course struct {
	callNum int
	professor string
	capacity int
	enrollment int
}

func (c *course) isFull() bool {
	return c.enrollment >= c.capacity
}

func stripFrom() course {
	return course{
		callNum: 0,
		professor: "",
		capacity: 0,
		enrollment: 0,
	}
}