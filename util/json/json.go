package json

import "encoding/json"

func Encode(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}

func Decode(bytes []byte, v interface{}) error {
	return json.Unmarshal(bytes, v)
}