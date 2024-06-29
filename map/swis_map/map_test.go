package swis_map

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMap(t *testing.T) {
	btSample := []byte(`{"key1": "value1", "key2": "value2"}`)
	m, err := NewMap(btSample)
	assert.Nil(t, err)

	valKey1, ok := m.Get("key1")
	assert.True(t, ok)
	assert.Equal(t, "value1", valKey1)
}

func TestEncodeMap(t *testing.T) {
	btSample := []byte(`{"key1":"value1","key2":"value2"}`)
	m, err := NewMap(btSample)
	assert.Nil(t, err)

	valKey1, ok := m.Get("key1")
	assert.True(t, ok)
	assert.Equal(t, "value1", valKey1)

	// encode full map
	t.Log("Encode Full map")
	btEncode, err := Encode(m, m.Count())
	assert.Nil(t, err)
	assert.Equal(t, string(btEncode), string(btSample))

	// encode partial map
	t.Log("Encode Partial map")
	btEncode, err = Encode(m, 1)
	assert.Nil(t, err)
	assert.Equal(t, string(btEncode), `{"key1":"value1"}`)
}
