package mymap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashMap(t *testing.T) {
	h := &HashMap[int]{}

	h.Add("junbum", 100)
	val, ok := h.Get("junbum")
	assert.True(t, ok)
	assert.Equal(t, 100, val)

	h.Add("golang", 200)
	val, ok = h.Get("golang")
	assert.True(t, ok)
	assert.Equal(t, val, 200)

	h.Add("junbum", 100)
	val, ok = h.Get("junbum")
	assert.True(t, ok)
	assert.Equal(t, 100, val)

	h.Add("awesome", 300)
	val, ok = h.Get("awesome")
	assert.True(t, ok)
	assert.Equal(t, 300, val)
}

func TestGoBasicMap(t *testing.T) {
	m := make(map[string]int)
	// var m map[string]int
	// var m = map[string]int{
	// 	...
	// }

	m["junbum"] = 100
	m["golang"] = 200
	m["awesome"] = 300

	assert.Equal(t, 100, m["junbum"])
	assert.Equal(t, 200, m["golang"])
	assert.Equal(t, 300, m["awesome"])

	_, ok := m["aaa"]
	assert.False(t, ok)

	delete(m, "junbum")
}
