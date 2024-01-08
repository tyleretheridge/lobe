package utils

import "errors"

// IntSliceToSet converts a slice of integers to a map of integers to empty structs that is
// functionally a mathematical set of integers
func IntSliceToSet(intList []int) (map[int]struct{}, error) {
	if intList == nil {
		return nil, errors.New("unable to convert nil to a set")
	}
	set := map[int]struct{}{}
	for i := range intList {
		set[intList[i]] = struct{}{}
	}
	return set, nil
}
