package benchpool

import "sync"

var slicePool sync.Pool

const size = 10000

func init() {
	slicePool.New = func() interface{} {
		return make([]int, 0, size)
	}
}

func AppendMany() []int {
	s := make([]int, 0, size)
	for i := 0; i < size; i++ {
		s = append(s, i)
	}
	return s
}

func PoolAppendMany() []int {
	s := slicePool.Get().([]int)
	for i := 0; i < size; i++ {
		s = append(s, i)
	}
	return s
}
