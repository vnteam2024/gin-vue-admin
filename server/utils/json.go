package utils

import (
	"encoding/json"
	"strings"
)

func GetJSONKeys(jsonStr string) (keys []string, err error) {
	// Use json.Decoder to record the order of keys during parsing
	dec := json.NewDecoder(strings.NewReader(jsonStr))
	t, err := dec.Token()
	if err != nil {
		return nil, err
	}
	// Make sure the data is an object
	if t != json.Delim('{') {
		return nil, err
	}
	for dec.More() {
		t, err = dec.Token()
		if err != nil {
			return nil, err
		}
		keys = append(keys, t.(string))

		// parse value
		var value interface{}
		err = dec.Decode(&value)
		if err != nil {
			return nil, err
		}
	}
	return keys, nil
}
