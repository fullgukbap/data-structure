package buffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {
	buf := NewSliceBuffer[byte]()

	buf.Write([]byte{1, 2, 3, 4})

	assert.Equal(t, 4, buf.Readable())
}

func TestRead(t *testing.T) {
	buf := NewSliceBuffer[byte]()
	buf.Write([]byte{1, 2, 3, 4})

	readedData := buf.Read(4)
	for i := 0; i < 4; i++ {
		assert.Equal(t, byte(i+1), readedData[i])
	}

	assert.Equal(t, 0, buf.Readable())
}

func TestWriteAndRead(t *testing.T) {
	buf := NewSliceBuffer[byte]()

	buf.Write([]byte{1, 2, 3, 4})
	assert.Equal(t, 4, buf.Readable())

	buf.Write([]byte{5, 6})
	assert.Equal(t, 6, buf.Readable())

	readedData := buf.Read(4)
	for i := 0; i < 4; i++ {
		assert.Equal(t, byte(i+1), readedData[i])
	}
	assert.Equal(t, 2, buf.Readable())
}
