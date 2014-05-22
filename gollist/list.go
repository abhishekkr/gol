package gollist

type List []string

type ListEngine interface {
	ToList(data string) []string
	FromList(list []string) string
}

var ListEngines = make(map[string]ListEngine)

func RegisterListEngine(name string, listEngine ListEngine) {
	ListEngines[name] = listEngine
}

func GetListEngine(name string) ListEngine {
	return ListEngines[name]
}
