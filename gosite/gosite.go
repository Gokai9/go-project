package gosite

import (
	"fmt"
	"log"
	"net/http"
)

type HandlerRoute func(w http.ResponseWriter, req *http.Request)

// menambah route dan method sebagai string dan fungsi yang akan dieksekusi ke handler rute
type Router struct {
	router map[string]HandlerRoute
}

func NewRouter() *Router {
	return &Router{
		router: make(map[string]HandlerRoute),
	}
}

func (r *Router) addRoute(method string, route string, handler HandlerRoute) {
	key := method + "-" + route
	r.router[key] = handler
}

// cek apakah route berawalan / atau tidak
func (r *Router) Get(route string, handler HandlerRoute) {
	r.addRoute("GET", route, handler)
}

func (r *Router) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, r))
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := r.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "Error 404 -- rute %s tidak dapat ditemukan\n", req.URL.Path)
	}
}
