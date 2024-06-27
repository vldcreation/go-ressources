package swis_map

import (
	"github.com/dolthub/swiss"
	"github.com/vldcration/go-ressources/json/json_iter"
)

func NewMap(data []byte) (*swiss.Map[string, interface{}], error) {
	if len(data) == 0 {
		return swiss.NewMap[string, interface{}](0), nil
	}

	mp := map[string]interface{}{}
	err := json_iter.Unmarshal(data, &mp)
	if err != nil {
		return nil, err
	}

	m := swiss.NewMap[string, interface{}](0)
	for k, v := range mp {
		m.Put(k, v)
	}

	return m, nil
}
