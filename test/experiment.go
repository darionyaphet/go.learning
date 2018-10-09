package test

import "strings"

func Exist(parameters *string, key string, target string) bool {
	pairs := strings.Split(*parameters, ";")
	for _, pair := range pairs {
		array := strings.Split(pair, "=")
		if array[0] == key && array[1] == target {
			return true
		}
	}
	return false
}
