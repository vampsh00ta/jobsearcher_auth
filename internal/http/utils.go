package http

type responseError struct {
	Error error `json:"err"`
}

func nilToString(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}
