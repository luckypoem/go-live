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
	router.POST("/app/:appname", CreateAppHandler)
	router.GET("/app/", ListAppsHandler)
	router.GET("/app/:appid", GetAppByIdHandler)
	router.PUT("/app/:appid", UpdateAppByIdHandler)
	router.DELETE("/app/:appid", DeleteAppByIdHandler)

	// Live Restful API
	router.POST("/live/:appname", CreateLiveHandler)
	router.GET("/live/", ListLivesHandler)
	router.GET("/live/:appid", GetLiveByIdHandler)
	router.PUT("/live/:appid", UpdateLiveByIdHandler)
	router.PUT("/live/:appid/token", UpdateLiveTokenByIdHandler)
	router.DELETE("/live/:appid", DeleteLiveByIdHandler)

	http.Serve(l, router)
	return nil
}
