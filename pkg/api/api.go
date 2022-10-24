package api

type Api interface {
	GetRoutes() []ApiRoute
}
