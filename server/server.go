package server

import "net/http"

type Country struct {
	Name     string
	Language string
}

var countries []*Country

//This creates the new server in the specific port
func New(addr string) *http.Server {

	initRoutes()
	return &http.Server{
		Addr: addr,
	}
}
