package main

import (
	"log"
	"net/http"

	"github.com/sant470/devetron/apis"
	handlers "github.com/sant470/devetron/apis/v1"
	"github.com/sant470/devetron/config"
	"github.com/sant470/devetron/services"
)

func main() {
	lgr := log.Default()
	lgr.Println("info: starting the server")
	// appConf := config.GetAppConfig("config", "./")
	router := config.InitRouters()
	searchSvc := services.NewSearchService(lgr)
	searchHlr := handlers.NewSearchHandler(lgr, searchSvc)
	apis.InitSerachRoutes(router, searchHlr)
	http.ListenAndServe("localhost:8000", router)
}
