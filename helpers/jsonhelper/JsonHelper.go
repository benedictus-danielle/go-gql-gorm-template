package jsonhelper

import "encoding/json"

func SerializeMap(m map[string]interface{}) string {
	bytes, _ := json.Marshal(m)
	return string(bytes)
}
