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

// MarshalJSON marshals the map to JSON.
func Encode(m *swiss.Map[string, interface{}], finish int) ([]byte, error) {
	mp := map[string]interface{}{}
	if m.Count() == 0 {
		return json_iter.Marshal(mp)
	}

	if finish <= 0 {
		finish = m.Count()
	}

	i := 1
	m.Iter(func(k string, v interface{}) bool {
		if i <= m.Count() {
			if i <= finish {
				mp[k] = v
				i++
				return false
			}
		}

		return true
	})

	return json_iter.Marshal(mp)
}
