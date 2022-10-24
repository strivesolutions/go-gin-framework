package api

import "github.com/gin-gonic/gin"

type ApiRoute struct {
	MethodType MethodType
	Anonymous  bool
	Path       string
	Handler    gin.HandlerFunc
}
