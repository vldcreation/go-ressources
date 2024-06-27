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
