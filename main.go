package main

import (
	"flag"
	"github.com/luckypoem/go-live/protocol/hls"
	"github.com/luckypoem/go-live/protocol/httpflv"
	"github.com/luckypoem/go-live/protocol/httpopera"
	"github.com/luckypoem/go-live/protocol/restfulapi"
	"github.com/luckypoem/go-live/protocol/rtmp"
	"log"
	"net"
	"time"
)

var (
	rtmpAddr    = flag.String("rtmp-addr", ":1935", "RTMP server listen address")
	httpFlvAddr = flag.String("httpflv-addr", ":7001", "HTTP-FLV server listen address")
	hlsAddr     = flag.String("hls-addr", ":7002", "HLS server listen address")
	operaAddr   = flag.String("manage-addr", ":8090", "HTTP manage interface server listen address")
	apiAddr     = flag.String("api-addr", ":8040", "HTTP Restful API listen address")
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
	flag.Parse()
}

func startHls() *hls.Server {
	hlsListen, err := net.Listen("tcp", *hlsAddr)
	if err != nil {
		log.Fatal(err)
	}

	hlsServer := hls.NewServer()
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println("HLS server panic: ", r)
			}
		}()
		log.Println("HLS listen On", *hlsAddr)
		hlsServer.Serve(hlsListen)
	}()
	return hlsServer
}

func startRtmp(stream *rtmp.RtmpStream, hlsServer *hls.Server) {
	rtmpListen, err := net.Listen("tcp", *rtmpAddr)
	if err != nil {
		log.Fatal(err)
	}

	var rtmpServer *rtmp.Server

	rtmpServer = rtmp.NewRtmpServer(stream, hlsServer)

	defer func() {
		if r := recover(); r != nil {
			log.Println("RTMP server panic: ", r)
		}
	}()
	log.Println("RTMP Listen On", *rtmpAddr)
	rtmpServer.Serve(rtmpListen)
}

func startAPI() {
	apiListen, err := net.Listen("tcp", *apiAddr)
	if err != nil {
		log.Fatal(err)
	}

	var apiServer *restfulapi.Server

	apiServer = restfulapi.NewServer()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println("API server panic: ", r)
			}
		}()

		log.Println("API Listen On", *apiAddr)
		apiServer.Serve(apiListen)
	}()
}

func startHTTPFlv(stream *rtmp.RtmpStream) {
	flvListen, err := net.Listen("tcp", *httpFlvAddr)
	if err != nil {
		log.Fatal(err)
	}

	hdlServer := httpflv.NewServer(stream)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println("HTTP-FLV server panic: ", r)
			}
		}()
		log.Println("HTTP-FLV listen On", *httpFlvAddr)
		hdlServer.Serve(flvListen)
	}()
}

func startHTTPOpera(stream *rtmp.RtmpStream) {
	if *operaAddr != "" {
		opListen, err := net.Listen("tcp", *operaAddr)
		if err != nil {
			log.Fatal(err)
		}
		opServer := httpopera.NewServer(stream, *rtmpAddr)
		go func() {
			defer func() {
				if r := recover(); r != nil {
					log.Println("HTTP-Operation server panic: ", r)
				}
			}()
			log.Println("HTTP-Operation listen On", *operaAddr)
			opServer.Serve(opListen)
		}()
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("go-live panic: ", r)
			time.Sleep(1 * time.Second)
		}
	}()

	stream := rtmp.NewRtmpStream()
	hlsServer := startHls()
	startHTTPFlv(stream)
	startHTTPOpera(stream)
	startAPI()
	startRtmp(stream, hlsServer)
}
