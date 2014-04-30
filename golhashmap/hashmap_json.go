package golhashmap

import (
	"encoding/json"
)

// init: register JSONMap to DataMap
func init() {
	RegisterDataMap("json", new(JSONMap))
}

// JSON's DataMap struct
type JSONMap struct {
	HMap      HashMap
	JSONBytes []byte
}

// getter JSONMap's HashMap
func (jsonmap *JSONMap) GetHashMap() HashMap {
	return jsonmap.HMap
}

// getter JSONMap's JSONBytes
func (jsonmap *JSONMap) GetDataMap() []byte {
	return jsonmap.JSONBytes
}

// setter JSONMap's HashMap
func (jsonmap *JSONMap) SetHashMap(hashmap HashMap) {
	jsonmap.HMap = make(HashMap)
	jsonmap.HMap = hashmap
}

// setter JSONMap's JSONBytes
func (jsonmap *JSONMap) SetDataMap(json_bytes []byte) {
	jsonmap.JSONBytes = json_bytes
}

// JSON DataMap's Encode json to hashmap
func (jsonmap *JSONMap) EncodeToHashMap() {
	var hashmap HashMap
	hashmap = make(HashMap)

	json_bytes := []byte(jsonmap.JSONBytes)
	err := json.Unmarshal(json_bytes, &hashmap)

	if err == nil {
		jsonmap.HMap = hashmap
	} else {
		jsonmap.HMap = make(HashMap)
	}
}

// JSON DataMap's Decode hashmap to json
func (jsonmap *JSONMap) DecodeToDataMap() {
	json_bytes, err := json.Marshal(jsonmap.HMap)

	if err == nil {
		jsonmap.JSONBytes = json_bytes
	}
}
