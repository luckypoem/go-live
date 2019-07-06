package restfulapi

import (
	"net"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (server *Server) Serve(l net.Listener) error {
	router := httprouter.New()

	http.Serve(l, router)
	return nil
}
