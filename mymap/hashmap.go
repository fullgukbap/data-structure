package mymap

import "hash/crc32"

const arraySize = 3571

type hashData[T any] struct {
	key   string
	value T
}

type HashMap[T any] struct {
	// arr [3571]{key string, value any(T)}
	arr [arraySize][]hashData[T]
}

func hashfn(key string) uint32 {
	// 데이터 검증에서 사용되는 해쉬 알고리즘, 매우 단순한 알고리즘이다. CRC-32
	return crc32.ChecksumIEEE([]byte(key))
}

func (h *HashMap[T]) Add(key string, value T) {
	hash := hashfn(key)
	// 해쉬 충돌이 일어날 수 있다.
	// -> 그러면 같은 위치에 있는 데이터를 덮어쓸 수 있다.
	// -> 대비책 ==> slice로 충돌나면 append 해주기
	var hashData = hashData[T]{
		key:   key,
		value: value,
	}

	h.arr[hash%arraySize] = append(h.arr[hash%arraySize], hashData)
}

func (h *HashMap[T]) Get(key string) (T, bool) {
	hash := hashfn(key)
	values := h.arr[hash%arraySize]
	for _, hashData := range values {
		if hashData.key == key {
			return hashData.value, true
		}
	}

	return *new(T), false
}
