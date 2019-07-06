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

	router.POST("/app/:appname", CreateAppHandler)
	router.GET("/app", ListAppsHandler)
	router.GET("/app/:appid", GetAppByIdHandler)
	router.PUT("/app/:appid", UpdateAppByIdHandler)
	router.DELETE("/app/:appid", DeleteAppByIdHandler)

	http.Serve(l, router)
	return nil
}
