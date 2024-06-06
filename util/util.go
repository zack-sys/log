package util

func GetString(a any) string {
	e, ok := a.(string)
	if ok {
		return e
	}
	return ""
}
