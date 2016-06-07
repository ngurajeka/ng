package ng

import (
	"reflect"
	"strconv"
	"strings"
)

//IntExist check if needle is exist on the haystack.
//it will return the index as the first return value
//and the boolean expression on the second return value
func IntExist(needle int, haystack []int) (int, bool) {
	var (
		exist bool
		index int
	)

	for i, n := range haystack {
		if n == needle {
			exist = true
			index = i
			break
		}
	}

	return index, exist
}

//StringExist check if needle is exist on the haystack
//it will return the index as the first return value
//and the boolean expression on the second return value
func StringExist(needle string, haystack []string) (int, bool) {
	var (
		exist bool
		index int
	)

	for i, n := range haystack {
		if n == needle {
			exist = true
			index = i
			break
		}
	}

	return index, exist
}

//IsSlice check if v->interface is a slice or not
//returning the boolean expression
func IsSlice(v interface{}) bool {
	if reflect.ValueOf(v).Kind() == reflect.Slice {
		return true
	}
	return false
}

//InterfaceToSliceString convert an interface into a
//slice of string
func InterfaceToSliceString(v interface{}) []string {
	var stringSlice []string
	if IsSlice(v) {
		switch reflect.TypeOf(stringSlice).Elem().Kind() {
		case reflect.String:
			stringSlice = v.([]string)
		case reflect.Int:
			stringSlice = SliceIntToSliceString(v.([]int))
			//default:
			//TODO we need to handle this one properly
		}
	} else {
		if value, ok := v.(string); ok {
			if value != "" {
				stringSlice = append(stringSlice, value)
			}
		}
	}
	return stringSlice
}

//SliceIntToSliceString convert a slice of int into
//a slice of string
func SliceIntToSliceString(slice []int) []string {
	var stringSlice []string
	for _, n := range slice {
		stringSlice = append(stringSlice, strconv.Itoa(n))
	}
	return stringSlice
}

//SliceIntToString convert a slice of int into
//a string, using delimiter as parameter.
//"," is the default delimiter
func SliceIntToString(slice []int, delimiter string) string {
	if delimiter == "" {
		delimiter = ","
	}
	stringSlice := SliceIntToSliceString(slice)
	return strings.Join(stringSlice, delimiter)
}

//StringToSliceByte convert string into
//a slice of byte. Using copy
func StringToSliceByte(something string) []byte {
	sliceByte := make([]byte, len(something))
	copy(sliceByte, something)
	return sliceByte
}
