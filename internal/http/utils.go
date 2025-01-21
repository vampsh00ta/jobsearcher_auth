package http

type responseError struct {
	Error error `json:"err"`
}
