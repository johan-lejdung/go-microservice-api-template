package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/johan-lejdung/go-microservice-api-template/bootstrap"
	"github.com/johan-lejdung/go-microservice-api-template/middleware"
	"github.com/urfave/negroni"

	log "github.com/sirupsen/logrus"
)

func main() {
	app := bootstrap.Service()

	n := negroni.New()
	n.Use(middleware.Logger{
		ExludePaths: []string{"/health/", "/health", "/readiness/", "/readiness"},
	})
	n.UseHandler(app.Router)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Content-Length", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	log.Info("Accepting connections on :8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headersOk, originsOk, methodsOk)(n)))
}
