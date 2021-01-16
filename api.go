package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Metrics struct {
	TotalProcessed uint64            `json:"totalProcessed"`
	ByTzProcessed  map[string]uint64 `json:"byTZProcessed"`
}

func InitMetrics() Metrics {
	m := Metrics{}
	m.ByTzProcessed = map[string]uint64{}
	return m
}

func decodeRequest(r *http.Request) (string, error) {
	var tz struct {
		Timezone string `json:"timezone"`
	}
	if err := json.NewDecoder(r.Body).Decode(&tz); err != nil {
		return "", err
	}
	return tz.Timezone, nil
}

func GetMetrics(w http.ResponseWriter, r *http.Request, metrics *Metrics) {
	w.Header().Set("Content-Type", "application/json")
	//TODO threadsafe read for metrics
	json.NewEncoder(w).Encode(metrics)
}

func GetTime(w http.ResponseWriter, r *http.Request, metrics *Metrics) {
	//TODO threadsafe update for metrics.TotalProcessed
	metrics.TotalProcessed += 1

	locname, err := decodeRequest(r)
	if err != nil {
		//invalid body
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//TODO cache for locations
	loc, err := time.LoadLocation(locname)
	if err != nil {
		//invalid location parameter
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//TODO threadsafe update for metrics.ByTzProcessed
	if _, ok := metrics.ByTzProcessed[locname]; !ok {
		metrics.ByTzProcessed[locname] = 0
	}
	metrics.ByTzProcessed[locname] += 1
	fmt.Fprint(w, time.Now().In(loc))
}
