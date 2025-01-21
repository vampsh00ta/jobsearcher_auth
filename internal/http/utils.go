package http

type responseError struct {
	Error error `json:"err"`
}

func checkNilString(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}
