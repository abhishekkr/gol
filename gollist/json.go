package gollist

import "encoding/json"

// init: register JSONmap to DataMap
func init() {
	RegisterListEngine("json", new(JSONmap))
}

type JSONmap struct{}

/* convert list elements to single line JSON */
func List_to_json(list []string) string {
	jsonValues_bytes, _ := json.Marshal(list)
	return string(jsonValues_bytes)
}

/* JSONmap proxy for List_to_json */
func (jsonmap JSONmap) FromList(list []string) string {
	return List_to_json(list)
}

/* entertains single depth JSON as element of list */
func Json_to_list(jsonValues string) []string {
	var list []string

	jsonValues_bytes := []byte(jsonValues)
	json.Unmarshal(jsonValues_bytes, &list)

	return list
}

/* JSONmap proxy for Json_to_list */
func (jsonmap JSONmap) ToList(jsonValues string) []string {
	return Json_to_list(jsonValues)
}
