package ui

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/Post-and-Play/gears/services"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	handleRoutes(r)
	return r
}

func RunServer() {
	done := services.MakeDoneSignal()
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	server := &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: Router(),
	}

	go func() {
		log.Printf("Server started at %s:%s", host, port)

		if err := server.ListenAndServe(); err != nil {
			log.Panicf("Error trying to start server: %+v", err)
		}
	}()

	<-done
	log.Println("Stopping server...")
}

func handleRoutes(r *gin.Engine) {
	for _, route := range healthCheck {
		r.Handle(route.Method, route.Path, route.Handler)
	}

	apiGroup := r.Group("/api")

	routers := [][]Route{cad, login}

	for _, router := range routers {
		for _, route := range router {
			apiGroup.Handle(route.Method, route.Path, route.Handler)
		}
	}
}
