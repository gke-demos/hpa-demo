/*
Copyright 2025 Google LLC.
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// simple request counter metric
var counter= prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "hpa_demo_requests",
		Help: "No of requests handled by compute handler",
	},
)



// check if this is a COR preflight request
func isPreflight(r *http.Request) bool {
	return r.Method == "OPTIONS" &&
	  r.Header.Get("Origin") != "" &&
	  r.Header.Get("Access-Control-Request-Method") != ""
}


func compute(w http.ResponseWriter, req *http.Request){
	counter.Inc()

	// lazy CORS handling
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// lazy way to return CORS preflight headers
	if isPreflight(req){
		 w.Header().Set("Access-Control-Allow-Methods","GET, OPTIONS")
		 if reqHeaders, found := req.Header["Access-Control-Request-Headers"]; found {
			w.Header().Set("Access-Control-Allow-Headers",strings.Join(reqHeaders,", "))
		 }
	} else {
		/*
		Similar to code used for https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough
		*/
		x := 0.
		rnd := rand.Intn(1000)
		for i := 0; i <= 10_000*rnd; i++ {
			x += math.Sqrt(float64(i))
		}
	}
	
	fmt.Fprint(w, "OK!")

}

func main() {
	log.SetOutput(os.Stdout)

	prometheus.MustRegister(counter)

	http.HandleFunc("/", compute)
	http.Handle("/metrics", promhttp.Handler())
	log.Print("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	log.Print("Server running on port 8080")
}
