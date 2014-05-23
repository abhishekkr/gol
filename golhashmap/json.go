package golhashmap

import "encoding/json"

// init: register JSONmap to DataMap
func init() {
	RegisterHashMapEngine("json", new(JSONmap))
}

type JSONmap struct{}

func HashMapToJSON(hmap HashMap) string {
	jsonalues_bytes, _ := json.Marshal(hmap)
	return string(jsonalues_bytes)
}

func (jsonmap JSONmap) FromHashMap(hmap HashMap) string {
	return HashMapToJSON(hmap)
}

func JSONToHashMap(jsonalues string) HashMap {
	var hmap HashMap
	hmap = make(HashMap)

	jsonalues_bytes := []byte(jsonalues)
	json.Unmarshal(jsonalues_bytes, &hmap)

	return hmap
}

func (jsonmap JSONmap) ToHashMap(jsonalues string) HashMap {
	return JSONToHashMap(jsonalues)
}
