package benchvalue

type item struct {
	id int64
}

func NewItemPointer() *item {
	return &item{}
}

func NewItemPointerSlice(size int) []*item {
	s := make([]*item, size)
	for i := 0; i < size; i++ {
		s[i] = &item{}
	}
	return s
}

func NewItem() item {
	return item{}
}

func NewItemSlice(size int) []item {
	s := make([]item, size)
	for i := 0; i < size; i++ {
		s[i] = item{}
	}
	return s
}
