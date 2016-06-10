package ng_test

import (
	"log"
	"testing"

	"github.com/ngurajeka/ng"
)

func TestIntExist(t *testing.T) {

	var (
		needle   = 1
		haystack = []int{5, 6, 7, 10, 1}
		pos      = 4
	)

	log.Printf("Finding %d on %v", needle, haystack)
	if index, ok := ng.IntExist(needle, haystack); !ok {
		t.Errorf("%d Should be on the haystack", needle)
	} else if index != pos {
		t.Errorf(
			"%d Should be on have the %d position of the haystack",
			needle, pos, index,
		)
	}

}

func TestStringExist(t *testing.T) {

	var (
		needle   = "ady"
		haystack = []string{"ady", "rahmat", "ma"}
		pos      = 0
	)

	log.Printf("Finding %s on %v", needle, haystack)
	if index, ok := ng.StringExist(needle, haystack); !ok {
		t.Errorf("%s Should be on the haystack", needle)
	} else if index != pos {
		t.Errorf(
			"%s Should be on have the %d position of the haystack",
			needle, pos, index,
		)
	}
}

func TestIsSlice(t *testing.T) {

	var (
		mySlice  = []int{}
		notSlice = "not a slice"
	)

	log.Printf("Checking %v is a slice or not", mySlice)
	if ok := ng.IsSlice(mySlice); !ok {
		t.Errorf("%v Should be a slice of int", mySlice)
	}

	log.Printf("Checking %v is a slice or not", notSlice)
	if ok := ng.IsSlice(notSlice); ok {
		t.Errorf("%v Should not be a slice, but a string")
	}

}

func TestConvertion(t *testing.T) {

	var (
		mySlice = []int{1, 2, 3, 4, 5, 6}
	)

	_ = ng.SliceIntToSliceString(mySlice)

	stringify := ng.SliceIntToString(mySlice, ",")
	log.Println(stringify)

	sliceByte := ng.StringToSliceByte(stringify)
	log.Println(sliceByte)
}
