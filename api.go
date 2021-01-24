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
	total          chan int
	bytz           chan string
}

func (m *Metrics) Init() {
	if m.ByTzProcessed != nil {
		return
	}
	m.ByTzProcessed = map[string]uint64{}
	m.total = make(chan int)
	m.bytz = make(chan string)
	go func(c <-chan int) {
		for {
			_, ok := <-c
			if !ok {
				return
			}
			m.TotalProcessed += 1
		}
	}(m.total)
	go func(c <-chan string) {
		for {
			locname, ok := <-c
			if !ok {
				return
			}
			if _, ok := m.ByTzProcessed[locname]; !ok {
				m.ByTzProcessed[locname] = 0
			}
			m.ByTzProcessed[locname] += 1
		}
	}(m.bytz)
}

func (m *Metrics) TotalInc() {
	m.total <- 1
}

func (m *Metrics) ByTzInc(locname string) {
	m.bytz <- locname
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
	if err := json.NewEncoder(w).Encode(metrics); err != nil {
		fmt.Println(err)
	}
}

func GetTime(w http.ResponseWriter, r *http.Request, metrics *Metrics) {
	metrics.TotalInc()

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

	metrics.ByTzInc(locname)
	fmt.Fprint(w, time.Now().In(loc))
}
