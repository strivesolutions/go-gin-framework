package api

type ApiError struct {
	StatusCode  int   
	Message string `json:"message"`
	Path    string `json:"path"`
	Code    int    `json:"code"`
	Detail  string `json:"detail"`
}

type Exception struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Detail string `json:"detail"`
}


