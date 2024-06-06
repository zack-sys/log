package util

import "time"

func GetString(a any) string {
	e, ok := a.(string)
	if ok {
		return e
	}
	return ""
}
func GetIndex(index string) string {
	return index + "-" + time.Now().Format("2006-01-02")
}
