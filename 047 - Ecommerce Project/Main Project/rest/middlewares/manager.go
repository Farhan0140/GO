package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}


func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
}


func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	h := next

	for _, middleware := range middlewares {
		h = middleware(h)
	}

	return h
}


func (mngr *Manager) WrapMux(next http.Handler) http.Handler {
	h := next

	for _, gblMiddleware := range mngr.globalMiddlewares {
		h = gblMiddleware(h)
	}

	return h
}