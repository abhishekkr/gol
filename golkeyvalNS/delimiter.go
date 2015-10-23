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
Delimited struct with required db details.
Delimited is just easy quick namespacing style around stores with just key-val support like leveldb, such.
Not to be opted for performance oriented use-cases.
*/
type Delimited struct {
	db golkeyval.DBEngine
}

/*
Configure populates Delimited required configs.
*/
func (delimited *Delimited) Configure(db golkeyval.DBEngine) {
	delimited.db = db
}

/*
register delimited namespace engine // non-performant
*/
func init() {
	RegisterNSDBEngine("delimited", new(Delimited))
}

/*
ReadNS reads all direct child values in a given NameSpace
For e.g.:
  given keys a, a:b, a:b:1, a:b:2, a:b:2:3
  reads for a:b:1, a:b:2 if queried for a:b
*/
func (delimited *Delimited) ReadNS(key string) golhashmap.HashMap {
	var hmap golhashmap.HashMap
	hmap = make(golhashmap.HashMap)
	key = "key::" + key
	val := delimited.db.GetVal(key)
	if val == "" {
		return hmap
	}
	children := strings.Split(val, ",")
	for _, child := range children {

		childKey := strings.Split(child, "key::")[1]
		childVal := delimited.db.GetVal("val::" + childKey)
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
func (delimited *Delimited) ReadNSRecursive(key string) golhashmap.HashMap {
	var hmap golhashmap.HashMap
	hmap = make(golhashmap.HashMap)

	keyname := "key::" + key
	valname := "val::" + key
	keynameVal := delimited.db.GetVal(keyname)
	valnameVal := delimited.db.GetVal(valname)
	if valnameVal != "" {
		hmap[key] = valnameVal
	}
	if keynameVal == "" {
		return hmap
	}
	children := strings.Split(keynameVal, ",")

	for _, childValAsKey := range children {
		childKey := strings.Split(childValAsKey, "key::")[1]
		inhmap := delimited.ReadNSRecursive(childKey)
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
func (delimited *Delimited) ifChildExists(childKey string, parentValue string) bool {
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
func (delimited *Delimited) appendKey(parent string, child string) bool {
	parentKeyName := fmt.Sprintf("key::%s", parent)
	childKeyName := fmt.Sprintf("key::%s:%s", parent, child)
	status := true

	val := delimited.db.GetVal(parentKeyName)
	if val == "" {
		if !delimited.db.PushKeyVal(parentKeyName, childKeyName) {
			status = false
		}
	} else if delimited.ifChildExists(childKeyName, val) {
		if !delimited.db.PushKeyVal(parentKeyName, val) {
			status = false
		}
	} else {
		val = fmt.Sprintf("%s,%s", val, childKeyName)
		if !delimited.db.PushKeyVal(parentKeyName, val) {
			status = false
		}
	}
	return status
}

/*
CreateNS updates entry with its trail of namespace
given a keyname.
*/
func (delimited *Delimited) CreateNS(key string) bool {
	splitIndex := strings.LastIndexAny(key, NamespaceSeparator)
	if splitIndex >= 0 {
		parentKey := key[0:splitIndex]
		childKey := key[splitIndex+1:]

		if delimited.appendKey(parentKey, childKey) {
			return delimited.CreateNS(parentKey)
		}
		return false
	}
	return true
}

/*
PushNS feeds in namespace entries given namespace key and val.
*/
func (delimited *Delimited) PushNS(key string, val string) bool {
	delimited.CreateNS(key)
	key = "val::" + key
	return delimited.db.PushKeyVal(key, val)
}

/*
unrootNSParent responsible for returning updated parentKeynameVal.
*/
func (delimited *Delimited) unrootNSParent(selfKeyname string, parentKeynameVal string) string {
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
func (delimited *Delimited) UnrootNS(key string) bool {
	statusParentUnroot, statusParentUpdate := true, true
	splitIndex := strings.LastIndexAny(key, NamespaceSeparator)
	if splitIndex < 0 {
		return true
	}
	parentKey := key[0:splitIndex]
	selfKeyname := fmt.Sprintf("key::%s", key)
	parentKeyname := fmt.Sprintf("key::%s", parentKey)
	parentKeynameVal := delimited.db.GetVal(parentKeyname)
	if parentKeynameVal == "" {
		return true
	}

	parentKeynameVal = delimited.unrootNSParent(selfKeyname, parentKeynameVal)

	if parentKeynameVal == "" {
		statusParentUnroot = delimited.UnrootNS(parentKey)
	}

	statusParentUpdate = delimited.db.PushKeyVal(parentKeyname, parentKeynameVal)

	return statusParentUnroot && statusParentUpdate
}

/*
DeleteNSKey directly deletes a child key-val and unroot it from parent.
*/
func (delimited *Delimited) DeleteNSKey(key string) bool {
	selfVal := "val::" + key
	if delimited.db.DelKey(selfVal) {
		keyname := "key::" + key
		if delimited.db.DelKey(keyname) {
			delimited.UnrootNS(key)
			return true
		}
	}
	delimited.UnrootNS(key)
	return false
}

/*
deleteNSChildren deletes children of any keyname.
*/
func (delimited *Delimited) deleteNSChildren(nsKey string) bool {
	delStatus := true
	key := strings.Split(nsKey, "key::")[1]
	valueKey := "val::" + key

	nsValue := delimited.db.GetVal(nsKey)
	if nsValue == "" {
		if delimited.db.DelKey(nsKey) {
			delStatus = delStatus && delimited.UnrootNS(key)
		} else {
			delStatus = false
		}
	}

	if !delimited.db.DelKey(valueKey) && delStatus {
		delStatus = false
	}
	return delStatus
}

/*
DeleteNS deletes a namespace with all direct children with no further children
and unroots it.
For nodes with children, it just delete baselevel value, doesn't unroot.
*/
func (delimited *Delimited) DeleteNS(key string) bool {
	delStatus := true
	nsKey := "key::" + key
	valueKey := "val::" + key

	childrenKeys := delimited.db.GetVal(nsKey)
	childrenKeysArray := strings.Split(childrenKeys, ",")
	for _, childKey := range childrenKeysArray {
		delStatus = delStatus && delimited.deleteNSChildren(childKey)
	}

	if delStatus {
		if !delimited.db.DelKey(valueKey) {
			delStatus = delStatus && false
		}
	}
	return delStatus
}

/*
deleteNSRecursiveChildren deletes recursive children of any keyname.
*/
func (delimited *Delimited) deleteNSRecursiveChildren(val string) bool {
	if val == "" {
		return true
	}
	status := true
	children := strings.Split(val, ",")
	for _, childKey := range children {
		childVal := "val::" + strings.Split(childKey, "key::")[1]
		status = status && delimited.deleteNSRecursiveChildren(delimited.db.GetVal(childKey))
		status = status && delimited.db.DelKey(childKey)
		status = status && delimited.db.DelKey(childVal)
	}
	return status
}

/*
DeleteNSRecursive to delete a namespace with all children below and unroot it.
*/
func (delimited *Delimited) DeleteNSRecursive(key string) bool {
	keyname := "key::" + key
	valname := "val::" + key
	keynameVal := delimited.db.GetVal(keyname)
	if delimited.db.DelKey(keyname) {
		if delimited.db.DelKey(valname) {

			if keynameVal == "" {
				delimited.UnrootNS(key)
				return true
			}
			delimited.UnrootNS(key)
			return delimited.deleteNSRecursiveChildren(keynameVal)
		}
	}
	return false
}
