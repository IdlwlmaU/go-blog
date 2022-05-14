package router

import "time"

func IsODD(num int) bool {
	if num&1 == 0 {
		return true
	}
	return false
}

func GetNextName(s []string, index int) string {
	return s[index+1]
}

func Date(layout string) string {
	return time.Now().Format(layout)
}
