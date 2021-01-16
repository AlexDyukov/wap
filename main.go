// wap is a tiny REST api service with /metrics and /time endpoints
package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	metrics := InitMetrics()

	router := httprouter.New()

	router.GET("/metrics", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		GetMetrics(w, r, &metrics)
	})
	router.POST("/time", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		GetTime(w, r, &metrics)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
