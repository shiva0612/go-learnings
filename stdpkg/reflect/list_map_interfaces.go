package main

import (
	"fmt"
	"reflect"
)

func ListAdd(list interface{}, value interface{}) {
	listValue := reflect.ValueOf(list).Elem()
	newValue := reflect.ValueOf(value)

	if listValue.Kind() != reflect.Slice {
		panic("list is not a slice")
	}

	if listValue.Type().Elem() != reflect.TypeOf(value) {
		panic("value has a different type than the list elements")
	}

	newSlice := reflect.Append(listValue, newValue)
	listValue.Set(newSlice)
}

func listadd_() {
	stringList := []string{"apple", "banana", "cherry"}
	ListAdd(&stringList, "date")
	fmt.Println("in main", stringList)

	intList := []int{1, 2, 3}
	ListAdd(&intList, 4)
	fmt.Println("in main", intList)
}

func ListClear(list interface{}) {
	rv := reflect.ValueOf(list)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Slice {
		panic("not a pointer to a slice")
	}
	rv.Elem().Set(reflect.Zero(rv.Elem().Type()))
}

func listclear_() {
	stringList := []string{"apple", "banana", "cherry"}
	intList := []int{1, 2, 3}

	ListClear(&stringList)
	ListClear(&intList)

	fmt.Println(stringList) // []
	fmt.Println(intList)    // []
}

func ListGet(list interface{}, i int) interface{} {
	listValue := reflect.ValueOf(list)

	if listValue.Kind() != reflect.Slice {
		panic("list is not a slice")
	}

	if i < 0 || i >= listValue.Len() {
		panic("index out of range")
	}

	elementValue := listValue.Index(i)

	return elementValue.Interface()
}

func listget_() {
	list_int := []int{1, 2, 3}
	var a interface{}
	a = ListGet(list_int, 0)
	fmt.Println(a, reflect.TypeOf(a))

	list_string := []string{"a", "b"}
	var b interface{}
	b = ListGet(list_string, 0)
	fmt.Println(b, reflect.TypeOf(b))
}

func MapCreate(myKeyType interface{}, myValueType interface{}) interface{} {
	keyType := reflect.TypeOf(myKeyType)
	valueType := reflect.TypeOf(myValueType)

	// Create the map type with the specified key and value types
	mapType := reflect.MapOf(keyType, valueType)

	// Create a new map with the constructed map type
	newMap := reflect.MakeMap(mapType).Interface()

	return newMap
}

func MapClear(myLittleMap interface{}) {
	mapValue := reflect.ValueOf(myLittleMap)

	if mapValue.Kind() != reflect.Ptr {
		panic("not a pointer to a map")
	}

	// Create a new empty map with the same key and value types
	newMap := MapCreate(reflect.Zero(mapValue.Elem().Type().Key()).Interface(), reflect.Zero(mapValue.Elem().Type().Elem()).Interface())

	// Set the original map's value to the newly created empty map
	mapValue.Elem().Set(reflect.ValueOf(newMap))
}

func mapclear_() {
	myMap := make(map[int]string)
	myMap[1] = "one"
	myMap[2] = "two"

	MapClear(&myMap)

	fmt.Println(myMap) // map[]
}

func MapGet(myMap interface{}, key interface{}) interface{} {
	mapValue := reflect.ValueOf(myMap)
	keyValue := reflect.ValueOf(key)

	if mapValue.Kind() != reflect.Map {
		return nil
	}

	// Get the map key type
	mapKeyType := mapValue.Type().Key()

	// Check if the key's type matches the map's key type
	if keyValue.Type() != mapKeyType {
		return nil
	}

	result := mapValue.MapIndex(keyValue)

	if !result.IsValid() {
		return nil
	}

	return result.Interface()
}

func mapget_() {
	mis := make(map[int]string)
	mis[0] = "zero"
	mis[1] = "one"
	fmt.Println(MapGet(mis, 0))      // "zero"
	fmt.Println(MapGet(mis, "zero")) // nil (key type mismatch)

	msi := make(map[string]int)
	msi["zero"] = 0
	msi["one"] = 1
	fmt.Println(MapGet(msi, "zero")) // 0
	fmt.Println(MapGet(msi, 0))      // nil (key type mismatch)
}

func MapKeys(myMap interface{}) interface{} {
	mapValue := reflect.ValueOf(myMap)

	if mapValue.Kind() != reflect.Map {
		panic("source is not a map")
	}

	keyType := mapValue.Type().Key()

	sliceType := reflect.SliceOf(keyType)
	keys := reflect.MakeSlice(sliceType, 0, mapValue.Len())

	for _, keyValue := range mapValue.MapKeys() {
		keys = reflect.Append(keys, keyValue)
	}

	return keys.Interface()
}

func mapkeys_() {
	mis := make(map[int]string)
	mis[0] = "zero"
	mis[1] = "one"
	keys := MapKeys(mis).([]int)
	fmt.Println(keys) // [0 1]

	msi := make(map[string]int)
	msi["zero"] = 0
	msi["one"] = 1
	keys2 := MapKeys(msi).([]string)
	fmt.Println(keys2) // ["zero" "one"]
}

func MapRemove(myMap interface{}, myKey interface{}) interface{} {
	mapValue := reflect.ValueOf(myMap)
	keyValue := reflect.ValueOf(myKey)

	if mapValue.Kind() != reflect.Ptr || mapValue.Elem().Kind() != reflect.Map {
		panic("source is not a pointer to a map")
	}

	if keyValue.Type() != mapValue.Elem().Type().Key() {
		panic("key type does not match the map's key type")
	}

	// mapInterface := mapValue.Interface()
	value := mapValue.Elem().MapIndex(keyValue).Interface()

	// Delete the key from the map
	mapValue.Elem().SetMapIndex(keyValue, reflect.Value{})

	return value
}

func mapremove_() {
	mis := make(map[int]string)
	mis[0] = "zero"
	mis[1] = "one"
	removed := MapRemove(&mis, 0)
	fmt.Println(removed) // "zero"
	fmt.Println(mis)     // map[1:one]

	msi := make(map[string]int)
	msi["zero"] = 0
	msi["one"] = 1
	removed2 := MapRemove(&msi, "zero")
	fmt.Println(removed2) // 0
	fmt.Println(msi)      // map[one:1]
}

func MapSet(myMap interface{}, myKey interface{}, myKeyValue interface{}) {
	mapValue := reflect.ValueOf(myMap)
	keyType := reflect.TypeOf(myKey)
	valueType := reflect.TypeOf(myKeyValue)

	if mapValue.Kind() != reflect.Map {
		panic("source is not a map")
	}

	if keyType != mapValue.Type().Key() {
		panic("key type does not match the map's key type")
	}

	if valueType != mapValue.Type().Elem() {
		panic("value type does not match the map's value type")
	}

	mapValue.SetMapIndex(reflect.ValueOf(myKey), reflect.ValueOf(myKeyValue))
}

func mapset_() {
	mis := make(map[int]string)
	MapSet(mis, 0, "zero")
	fmt.Println(mis) // map[0:zero]

	msi := make(map[string]int)
	MapSet(msi, "zero", 0)
	fmt.Println(msi) // map[zero:0]
}

func MapSize(myMap interface{}) int32 {
	mapValue := reflect.ValueOf(myMap)

	if mapValue.Kind() != reflect.Map {
		panic("source is not a map")
	}

	return int32(mapValue.Len())
}

func mapsize_() {
	mis := make(map[int]string)
	mis[0] = "zero"
	mis[1] = "one"
	size := MapSize(mis)
	fmt.Println(size) // 2

	msi := make(map[string]int)
	msi["zero"] = 0
	msi["one"] = 1
	size2 := MapSize(msi)
	fmt.Println(size2) // 2
}

func MapValues(myMap interface{}) interface{} {
	mapValue := reflect.ValueOf(myMap)

	if mapValue.Kind() != reflect.Map {
		panic("source is not a map")
	}

	valueType := mapValue.Type().Elem()

	sliceType := reflect.SliceOf(valueType)
	values := reflect.MakeSlice(sliceType, 0, mapValue.Len())

	for _, key := range mapValue.MapKeys() {
		value := mapValue.MapIndex(key)
		values = reflect.Append(values, value)
	}

	return values.Interface()
}

func mapvalues_() {
	mis := make(map[int]string)
	mis[0] = "zero"
	mis[1] = "one"
	values := MapValues(mis).([]string)
	fmt.Println(values) // ["zero" "one"]

	msi := make(map[string]int)
	msi["zero"] = 0
	msi["one"] = 1
	values2 := MapValues(msi).([]int)
	fmt.Println(values2) // [0 1]
}
