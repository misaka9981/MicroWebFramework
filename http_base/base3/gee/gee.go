package gee

import (
	"fmt"
	"net/http"
)

// HandlerFunc defines the request handler
type  HandlerFunc func(http.ResponseWriter, *http.Request)

// implement the interface of ServeHTTP
type Engine struct {
	router map[string]HandlerFunc
}

// Constructor of Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc)  {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// Get defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}
// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// RUN defines the method to start a http server
func (engine *Engine) RUn(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	}else{
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL.Path)
	}
}

