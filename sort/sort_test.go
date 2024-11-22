package sort

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	arr := make([]int, 1000)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(100)
	}

	assert.False(t, IsSorted(arr))
	QuickSort(arr)
	assert.True(t, IsSorted(arr), arr)
}

func TestMergeSort(t *testing.T) {
	arr := make([]int, 1000)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(100)
	}

	assert.False(t, IsSorted(arr))
	sorted := MergeSort(arr)
	assert.True(t, IsSorted(sorted), arr)
}

func BenchmarkQuickSort(b *testing.B) {
	arr := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		arr[i] = rand.Intn(b.N)
	}

	QuickSort(arr)
}

func BenchmarkMergeSort(b *testing.B) {
	arr := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		arr[i] = rand.Intn(b.N)
	}

	MergeSort(arr)
}

func BenchmarkBinaryInsertSort(b *testing.B) {
	arr := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		arr = BinaryInsertSort(arr, rand.Intn(100))
	}
}

func TestBinaryInsertSort(t *testing.T) {
	arr := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		arr = BinaryInsertSort(arr, rand.Intn(100))
	}

	assert.True(t, IsSorted(arr), arr)
}
