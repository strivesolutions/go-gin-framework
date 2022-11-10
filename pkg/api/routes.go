package api

import "github.com/gin-gonic/gin"

// type Route interface {
// 	GetMethod() MethodType
// 	GetIsAnonymous() bool
// 	GetSkipTrustFundCheck() bool
// 	GetPath() string
// 	GetHandler() gin.HandlerFunc
// }

type ApiRoute struct {
	MethodType         MethodType
	Anonymous          bool
	SkipTrustFundCheck bool
	Path               string
	Handler            gin.HandlerFunc
}

func CreateSubscriptionRoute(path string, handler gin.HandlerFunc) ApiRoute {
	return ApiRoute{
		MethodType:         POST,
		Anonymous:          true,
		SkipTrustFundCheck: true,
		Path:               path,
		Handler:            handler,
	}
}
