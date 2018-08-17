package bootstrap

import (
	"os"

	"github.com/facebookgo/inject"
	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"

	"github.com/johan-lejdung/go-microservice-api-template/api"
	"github.com/johan-lejdung/go-microservice-api-template/db"
	"github.com/johan-lejdung/go-microservice-api-template/goservice"
	"github.com/johan-lejdung/go-microservice-api-template/liveness"
	jlog "github.com/johan-lejdung/log"
	"github.com/joho/godotenv"
)

// App contains the context of the application
type App struct {
	Router      *mux.Router          `inject:""`
	GoAPI       *api.GoAPI           `inject:""`
	LivenessAPI liveness.LivenessAPI `inject:""`
}

// Service will bootstrap and create a new App for the service
func Service() App {
	setupEnv()
	setupLogging()

	app := bootstrapApp()

	setupRouter(app.Router, &app)

	return app
}

func setupEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Debug("Error loading .env file")
	}
}

func setupLogging() {
	log.SetFormatter(&jlog.FluentdFormatter{})
	if os.Getenv("ENV") == "dev" {
		log.SetFormatter(&log.TextFormatter{})
	}

	if level := os.Getenv("LOG_LEVEL"); len(level) > 0 {
		ll, err := log.ParseLevel(level)
		if err == nil {
			log.SetLevel(ll)
		} else {
			log.Info("Failed to parse log level. Will use default: %s", err.Error())
		}
	}
}

func setupRouter(router *mux.Router, app *App) {
	app.GoAPI.InitAPIRoute()

	app.LivenessAPI.InitHealthRouter()
	app.LivenessAPI.InitReadinessRouter()
}

func bootstrapApp() App {
	g := inject.Graph{}
	g.Logger = log.StandardLogger()

	app := App{}
	database, err := db.CreateDatabase()
	if err != nil {
		log.Fatal(err)
	}

	err = g.Provide(
		&inject.Object{Value: &app},
		&inject.Object{Value: &api.GoAPI{}},
		&inject.Object{Value: database},
		&inject.Object{Value: &liveness.API{}},
		&inject.Object{Value: &goservice.Service{}},
		&inject.Object{Value: mux.NewRouter().StrictSlash(true)},
	)
	if err != nil {
		log.Fatal(err)
	}

	if err = g.Populate(); err != nil {
		log.Fatal(err)
	}

	return app
}
