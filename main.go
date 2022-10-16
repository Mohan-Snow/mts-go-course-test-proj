package main

import (
	"sort"
	"strconv"
)

func main() {}

func shrinkString(input string) (result string) {
	var arrayString []string
	for _, v := range input {
		arrayString = append(arrayString, string(v))
	}
	sort.Strings(arrayString)
	var counter int
	counterMap := make(map[string]struct{})
	for i, v := range arrayString {
		_, ok := counterMap[v]
		if !ok {
			if i == 0 {
				counterMap[v] = struct{}{}
				result += v
				counter++
				continue
			}
			result += strconv.Itoa(counter)
			counter = 0
			counterMap[v] = struct{}{}
			result += v
			counter++
		} else {
			counter++
		}
	}
	result += strconv.Itoa(counter)
	return
}
