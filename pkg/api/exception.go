package api

type ApiError struct {
	StatusCode  int   
	Message string `json:"message"`
	Path    string `json:"path"`
	Code    int    `json:"code"`
	Detail  string `json:"detail"`
}

