package gollist

import "encoding/json"

// init: register JSONmap to DataMap
func init() {
	RegisterListEngine("json", new(JSONmap))
}

type JSONmap struct{}

/* convert list elements to single line JSON */
func ListToJSON(list []string) string {
	jsonValues_bytes, _ := json.Marshal(list)
	return string(jsonValues_bytes)
}

/* JSONmap proxy for ListToJSON */
func (jsonmap JSONmap) FromList(list []string) string {
	return ListToJSON(list)
}

/* entertains single depth JSON as element of list */
func JSONToList(jsonValues string) []string {
	var list []string

	jsonValues_bytes := []byte(jsonValues)
	json.Unmarshal(jsonValues_bytes, &list)

	return list
}

/* JSONmap proxy for JSONToList */
func (jsonmap JSONmap) ToList(jsonValues string) []string {
	return JSONToList(jsonValues)
}
