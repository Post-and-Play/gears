package ui

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/Post-and-Play/gears/services"
	"github.com/gin-gonic/gin"

	docs "github.com/Post-and-Play/gears/docs"
)

func Router() *gin.Engine {
	r := gin.Default()
	handleRoutes(r)
	return r
}

func RunServer(env string) {
	var adr string

	if env == "PROD" {
		adr = os.Getenv("PROD_HOST")
	} else {
		host := os.Getenv("GIN_HOST")
		port := os.Getenv("GIN_PORT")
		adr = net.JoinHostPort(host, port)
	}

	done := services.MakeDoneSignal()

	server := &http.Server{
		Addr:    adr,
		Handler: Router(),
	}

	go func() {
		log.Printf("Server started at %s", adr)

		if err := server.ListenAndServe(); err != nil {
			log.Panicf("Error trying to start server: %+v", err)
		}
	}()

	<-done
	log.Println("Stopping server...")
}

func handleRoutes(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"

	for _, route := range healthCheck {
		r.Handle(route.Method, route.Path, route.Handler)
	}

	for _, route := range swagg {
		r.Handle(route.Method, route.Path, route.Handler)
	}

	apiGroup := r.Group("/api")

	routers := [][]Route{cad, login, game, favorite, review, like}

	for _, router := range routers {
		for _, route := range router {
			apiGroup.Handle(route.Method, route.Path, route.Handler)
		}
	}
}
