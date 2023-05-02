package muxxer

import "net/http"

type Router struct {
	mux *http.ServeMux
}

func New() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

func (router *Router) GET(path string, handler http.HandlerFunc) {
	router.mux.HandleFunc(path, register(http.MethodGet, handler))
}

func (router Router) POST(path string, handler http.HandlerFunc) {
	router.mux.HandleFunc(path, register(http.MethodPost, handler))
}

func (router Router) DELETE(path string, handler http.HandlerFunc) {
	router.mux.HandleFunc(path, register(http.MethodDelete, handler))
}

func (router Router) PUT(path string, handler http.HandlerFunc) {
	router.mux.HandleFunc(path, register(http.MethodPut, handler))
}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.mux.ServeHTTP(w, r)
}

func register(method string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			handler(w, r)
		}
	}
}
