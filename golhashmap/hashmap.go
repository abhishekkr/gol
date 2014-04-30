package golhashmap

type HashMap map[string]string

type DataMap interface {
	GetHashMap() HashMap
	SetHashMap(hashmap HashMap)

	GetDataMap() []byte
	SetDataMap(data_bytes []byte)

	EncodeToHashMap()
	DecodeToDataMap()
}

var DataMapList = make(map[string]DataMap)

func RegisterDataMap(name string, datamap DataMap) {
	DataMapList[name] = datamap
}

func GetDataMap(name string) DataMap {
	return DataMapList[name]
}
