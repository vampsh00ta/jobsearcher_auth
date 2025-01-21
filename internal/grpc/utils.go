package grpc

func nilToString(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}
