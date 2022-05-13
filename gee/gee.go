package gee

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	*RouterGroup
	router *router
	group  []*RouterGroup
}

func New() *Engine {

	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.group = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (engine *Engine) GET(path string, handler HandlerFunc) {
	engine.router.addRoute("GET", path, handler)
}

func (engine *Engine) POST(path string, handler HandlerFunc) {
	engine.router.addRoute("POST", path, handler)
}

func (engine *Engine) Run(addr string) error {

	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	engine.router.handle(c)
}
