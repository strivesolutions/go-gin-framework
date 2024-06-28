package api

type ApiError struct {
	Message string      `json:"message"`
	Path    string      `json:"path"`
	Code    int         `json:"code"`
	Detail  interface{} `json:"detail"`
}
