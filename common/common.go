package common

import "container/list"

//ConvertListToArray - convert list to an array of objects
func ConvertListToArray(given *list.List) *[]interface{} {
	docs := make([]interface{}, given.Len())
	index := 0
	for e := given.Front(); e != nil; e = e.Next() {
		docs[index] = e.Value
		index++
	}
	return &docs
}
