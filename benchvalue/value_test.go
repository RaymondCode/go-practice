package benchvalue

import "testing"

func BenchmarkNewItemPointer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewItemPointer()
	}
}

func BenchmarkNewItem(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewItem()
	}
}

func BenchmarkNewItemPointerSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewItemPointerSlice(1000)
	}
}

func BenchmarkNewItemSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewItemSlice(1000)
	}
}
