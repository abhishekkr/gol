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
