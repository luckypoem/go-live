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

	// APP Restful API
	router.POST("/app/:appname/create", CreateAppHandler)
	router.GET("/app/", ListAppsHandler)
	router.GET("/app/:appid/get", GetAppByIdHandler)
	router.DELETE("/app/:appid/del", DeleteAppByIdHandler)

	// Live Restful API
	router.POST("/live/:appname/create", CreateLiveHandler)
	router.GET("/live/", ListLivesHandler)
	router.GET("/live/:appid/create", GetLiveByIdHandler)
	router.PUT("/live/:appid/token", UpdateLiveTokenByIdHandler)
	router.DELETE("/live/:appid", DeleteLiveByIdHandler)

	http.Serve(l, router)
	return nil
}
