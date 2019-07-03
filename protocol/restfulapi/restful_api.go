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

	router.POST("/api/apps", server.CreateAppHandler)

	http.Serve(l, router)
	return nil
}

func (server *Server) CreateAppHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
