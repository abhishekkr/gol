package golkeyvalNS

import (
	"fmt"
	"strings"

	golhashmap "github.com/abhishekkr/gol/golhashmap"
	golkeyval "github.com/abhishekkr/gol/golkeyval"
)

/* NameSpaceSeparator could be modified if something other than colon-char ":"
is to be used as separator symbol for NameSpace. */
var (
	NamespaceSeparator = ":"
)

/*
ReadNS reads all direct child values in a given NameSpace
For e.g.:
  given keys a, a:b, a:b:1, a:b:2, a:b:2:3
  reads for a:b:1, a:b:2 if queried for a:b
*/
func ReadNS(key string, db golkeyval.DBEngine) golhashmap.HashMap {
	var hmap golhashmap.HashMap
	hmap = make(golhashmap.HashMap)
	key = "key::" + key
	val := db.GetVal(key)
	if val == "" {
		return hmap
	}
	children := strings.Split(val, ",")
	for _, child := range children {

		childKey := strings.Split(child, "key::")[1]
		childVal := db.GetVal("val::" + childKey)
		if childVal != "" {
			hmap[childKey] = childVal
		}
	}
	return hmap
}

/*
ReadNSRecursive reads all values belonging to tree below given NameSpace
For e.g.:
  given keys a, a:b, a:b:1, a:b:2, a:b:2:3
  reads for a:b:1, a:b:2, a:b:2:3 if queried for a:b
*/
func ReadNSRecursive(key string, db golkeyval.DBEngine) golhashmap.HashMap {
	var hmap golhashmap.HashMap
	hmap = make(golhashmap.HashMap)

	keyname := "key::" + key
	valname := "val::" + key
	keynameVal := db.GetVal(keyname)
	valnameVal := db.GetVal(valname)
	if valnameVal != "" {
		hmap[key] = valnameVal
	}
	if keynameVal == "" {
		return hmap
	}
	children := strings.Split(keynameVal, ",")

	for _, childValAsKey := range children {
		childKey := strings.Split(childValAsKey, "key::")[1]
		inhmap := ReadNSRecursive(childKey, db)
		for inHmapKey, inHmapVal := range inhmap {
			hmap[inHmapKey] = inHmapVal
		}
	}

	return hmap
}

/*
Given all full child keynames of a given NameSpace reside as string separated
by a comma(","). This method checks for a given keyname being a child keyname
for provided for group string of all child keynames.
Return:
  true if given keyname is present as child in group-val of child keynames
  false if not
*/
func ifChildExists(childKey string, parentValue string) bool {
	children := strings.Split(parentValue, ",")
	for _, child := range children {
		if child == childKey {
			return true
		}
	}
	return false
}

/*
appendKey updates the group-val for child keynames
of a parent keyname as required
given a parent keyname and child keyname.
*/
func appendKey(parent string, child string, db golkeyval.DBEngine) bool {
	parentKeyName := fmt.Sprintf("key::%s", parent)
	childKeyName := fmt.Sprintf("key::%s:%s", parent, child)
	status := true

	val := db.GetVal(parentKeyName)
	if val == "" {
		if !db.PushKeyVal(parentKeyName, childKeyName) {
			status = false
		}
	} else if ifChildExists(childKeyName, val) {
		if !db.PushKeyVal(parentKeyName, val) {
			status = false
		}
	} else {
		val = fmt.Sprintf("%s,%s", val, childKeyName)
		if !db.PushKeyVal(parentKeyName, val) {
			status = false
		}
	}
	return status
}

/*
CreateNS updates entry with its trail of namespace
given a keyname.
*/
func CreateNS(key string, db golkeyval.DBEngine) bool {
	splitIndex := strings.LastIndexAny(key, NamespaceSeparator)
	if splitIndex >= 0 {
		parentKey := key[0:splitIndex]
		childKey := key[splitIndex+1:]

		if appendKey(parentKey, childKey, db) {
			return CreateNS(parentKey, db)
		}
		return false
	}
	return true
}

/*
PushNS feeds in namespace entries given namespace key and val.
*/
func PushNS(key string, val string, db golkeyval.DBEngine) bool {
	CreateNS(key, db)
	key = "val::" + key
	return db.PushKeyVal(key, val)
}

/*
unrootNSParent responsible for returning updated parentKeynameVal.
*/
func unrootNSParent(selfKeyname string, parentKeynameVal string) string {
	parentKeynameValElem := strings.Split(parentKeynameVal, ",")

	_tmpArray := make([]string, len(parentKeynameValElem))
	_tmpArrayIdx := 0
	for _, elem := range parentKeynameValElem {
		if elem == selfKeyname || elem == "" {
			continue
		}
		_tmpArray[_tmpArrayIdx] = elem
		_tmpArrayIdx++
	}

	if _tmpArrayIdx > 1 {
		return strings.Join(_tmpArray[0:_tmpArrayIdx], ",")
	}
	return _tmpArray[0]
}

/*
UnrootNS update key's presence from it's parent's  group-val of child key names.
*/
func UnrootNS(key string, db golkeyval.DBEngine) bool {
	statusParentUnroot, statusParentUpdate := true, true
	splitIndex := strings.LastIndexAny(key, NamespaceSeparator)
	if splitIndex < 0 {
		return true
	}
	parentKey := key[0:splitIndex]
	selfKeyname := fmt.Sprintf("key::%s", key)
	parentKeyname := fmt.Sprintf("key::%s", parentKey)
	parentKeynameVal := db.GetVal(parentKeyname)
	if parentKeynameVal == "" {
		return true
	}

	parentKeynameVal = unrootNSParent(selfKeyname, parentKeynameVal)

	if parentKeynameVal == "" {
		statusParentUnroot = UnrootNS(parentKey, db)
	}

	statusParentUpdate = db.PushKeyVal(parentKeyname, parentKeynameVal)

	return statusParentUnroot && statusParentUpdate
}

/*
RootNS update key's presence in it's parent's group-val of child key names.
*/
func RootNS(key string, childKey string, db golkeyval.DBEngine) bool {
	statusParentUnroot, statusParentUpdate := true, true
	return statusParentUnroot && statusParentUpdate
}

/*
DeleteNSKey directly deletes a child key-val and unroot it from parent.
*/
func DeleteNSKey(key string, db golkeyval.DBEngine) bool {
	selfVal := "val::" + key
	if db.DelKey(selfVal) {
		keyname := "key::" + key
		if db.DelKey(keyname) {
			UnrootNS(key, db)
			return true
		}
	}
	UnrootNS(key, db)
	return false
}

/*
deleteNSChildren deletes children of any keyname.
*/
func deleteNSChildren(nsKey string, db golkeyval.DBEngine) bool {
	delStatus := true
	key := strings.Split(nsKey, "key::")[1]
	valueKey := "val::" + key

	nsValue := db.GetVal(nsKey)
	if nsValue == "" {
		if db.DelKey(nsKey) {
			delStatus = delStatus && UnrootNS(key, db)
		} else {
			delStatus = false
		}
	}

	if !db.DelKey(valueKey) && delStatus {
		delStatus = false
	}
	return delStatus
}

/*
DeleteNS deletes a namespace with all direct children with no further children
and unroots it.
For nodes with children, it just delete baselevel value, doesn't unroot.
*/
func DeleteNS(key string, db golkeyval.DBEngine) bool {
	delStatus := true
	nsKey := "key::" + key
	valueKey := "val::" + key

	childrenKeys := db.GetVal(nsKey)
	childrenKeysArray := strings.Split(childrenKeys, ",")
	for _, childKey := range childrenKeysArray {
		delStatus = delStatus && deleteNSChildren(childKey, db)
	}

	if delStatus {
		if !db.DelKey(valueKey) {
			delStatus = delStatus && false
		}
	}
	return delStatus
}

/*
deleteNSRecursiveChildren deletes recursive children of any keyname.
*/
func deleteNSRecursiveChildren(val string, db golkeyval.DBEngine) bool {
	if val == "" {
		return true
	}
	status := true
	children := strings.Split(val, ",")
	for _, childKey := range children {
		childVal := "val::" + strings.Split(childKey, "key::")[1]
		status = status && deleteNSRecursiveChildren(db.GetVal(childKey), db)
		status = status && db.DelKey(childKey)
		status = status && db.DelKey(childVal)
	}
	return status
}

/*
DeleteNSRecursive to delete a namespace with all children below and unroot it.
*/
func DeleteNSRecursive(key string, db golkeyval.DBEngine) bool {
	keyname := "key::" + key
	valname := "val::" + key
	keynameVal := db.GetVal(keyname)
	if db.DelKey(keyname) {
		if db.DelKey(valname) {

			if keynameVal == "" {
				UnrootNS(key, db)
				return true
			}
			UnrootNS(key, db)
			return deleteNSRecursiveChildren(keynameVal, db)
		}
	}
	return false
}
