package restfulapi

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (server *Server) Serve(l net.Listener) error {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	http.Serve(l, router)
	return nil
}
