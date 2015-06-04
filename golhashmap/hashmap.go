package golhashmap

type HashMap map[string]string

type HashMapEngine interface {
	ToHashMap(data string) HashMap
	FromHashMap(hashmap HashMap) string
}

var HashMapEngines = make(map[string]HashMapEngine)

func RegisterHashMapEngine(name string, hashmapEngine HashMapEngine) {
	HashMapEngines[name] = hashmapEngine
}

func GetHashMapEngine(name string) HashMapEngine {
	return HashMapEngines[name]
}

func (hmap HashMap) Keys() []string {
	var keys = make([]string, len(hmap))
	idx := 0
	for _k, _ := range hmap {
		keys[idx] = _k
		idx++
	}
	return keys
}

func (hmap HashMap) Values() []string {
	var values = make([]string, len(hmap))
	idx := 0
	for _, _v := range hmap {
		values[idx] = _v
		idx++
	}
	return values
}

func (hmap HashMap) Items() [][]string {
	var items = make([][]string, len(hmap))
	idx := 0
	for _k, _v := range hmap {
		items[idx] = make([]string, 2)
		items[idx][0], items[idx][1] = _k, _v
		idx++
	}
	return items
}

func (hmap HashMap) Count() int {
	return len(hmap)
}
