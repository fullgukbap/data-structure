package sort

import "golang.org/x/exp/constraints"

func QuickSort[T constraints.Ordered](arr []T) {
	if len(arr) <= 1 {
		return
	}

	idx := patition(arr)
	QuickSort(arr[:idx])
	QuickSort(arr[idx+1:])
}

func patition[T constraints.Ordered](arr []T) int {
	if len(arr) <= 1 {
		return 0
	}
	pivot := arr[0]
	i := 1
	j := len(arr) - 1

	for {
		// pivot 보다 큰 값 찾아감
		for i < len(arr) && arr[i] <= pivot {
			i++
		}

		// pivot 보다 작은 값 찾아감
		for j > 0 && arr[j] > pivot {
			j--
		}

		// 만약 서로 역전되었다면? -> pivot의 위치
		if i >= j {
			arr[0], arr[i-1] = arr[i-1], arr[0]
			return i - 1
		}

		arr[i], arr[j] = arr[j], arr[i]
	}
}

func IsSorted[T constraints.Ordered](arr []T) bool {
	if len(arr) <= 1 {
		return true
	}

	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}

	return true
}
