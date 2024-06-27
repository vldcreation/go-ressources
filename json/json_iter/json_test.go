package json_iter

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var pathTestData = "test_data.json"

func TestJsonIterUnMarshall(t *testing.T) {
	bt, err := os.ReadFile(pathTestData)
	assert.NoError(t, err)
	assert.NotNil(t, bt)

	var data map[string]interface{}
	err = Unmarshal(bt, &data)
	assert.NoError(t, err)

	assert.Equal(t, float64(2), data["_v"])
}

type Adf struct {
	Version  float64   `json:"_v"`
	Sections []Section `json:"sections"`
}

type Section struct {
	ID     string  `json:"_id"`
	Tittle string  `json:"tittle"`
	Fields []Field `json:"fields"`
}

type Field struct {
	ID         string `json:"_id"`
	Key        string `json:"key"`
	Type       string `json:"type"`
	Directive  string `json:"directive"`
	IsRequired bool   `json:"isRequired"`
}

func TestJsonIterGet(t *testing.T) {
	bt, err := os.ReadFile(pathTestData)
	assert.NoError(t, err)
	assert.NotNil(t, bt)

	var data Adf
	err = Unmarshal(bt, &data)
	assert.NoError(t, err)

	assert.Equal(t, float64(2), data.Version)
	assert.Equal(t, 2, len(data.Sections))
	assert.Equal(t, 17, len(data.Sections[0].Fields))

	mergedFields := make([]Field, 0)
	mergedFields = append(mergedFields, data.Sections[0].Fields...)

	if len(data.Sections) > 1 {
		for i := 1; i < len(data.Sections); i++ {
			mergedFields = append(mergedFields, data.Sections[i].Fields...)
		}
	}

	assert.Equal(t, 18, len(mergedFields))
}
