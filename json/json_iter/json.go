package json_iter

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func Get(data []byte, key string) (size int, res interface{}, err error) {
	val := json.Get(data, key)
	if val.LastError() != nil {
		return 0, nil, val.LastError()
	}

	return val.Size(), val.ToString(), nil
}
