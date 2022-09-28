package grouprouter

type RouterGroup struct {
	BaseRouter
	AdminRouters
}

var RouterGroupApp = new(RouterGroup)
