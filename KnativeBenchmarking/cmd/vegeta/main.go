package main

import (
	"log"
	"net/http"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	//var metrics vegeta.Metrics
	for i := 0; i < 1000; i++ {
		log.Printf("Starting attacker #%d\n", i)
		go func() {
			rate := vegeta.Rate{Freq: 100, Per: time.Second}
			duration := 2 * time.Hour
			targeter := vegeta.NewStaticTargeter(vegeta.Target{
				Method: "GET",
				URL:    "http://35.223.15.23/",
				Header: http.Header{
					"Host": []string{"helloworld-go.default.example.com"},
				},
			})
			attacker := vegeta.NewAttacker()

			for _ = range attacker.Attack(targeter, rate, duration, "Knative") {
				//metrics.Add(res)
			}
		}()
		if i%10 == 9 {
			time.Sleep(2 * time.Minute)
		} else {
			time.Sleep(10 * time.Second)
		}
	}

	//metrics.Close()

	//fmt.Printf("Min: %s\n", metrics.Latencies.Min)
	//fmt.Printf("P50: %s\n", metrics.Latencies.P50)
	//fmt.Printf("P90: %s\n", metrics.Latencies.P90)
	//fmt.Printf("P95: %s\n", metrics.Latencies.P95)
	//fmt.Printf("P99: %s\n", metrics.Latencies.P99)
	//fmt.Printf("Max: %s\n", metrics.Latencies.Max)
	//fmt.Printf("Throughput: %f\n", metrics.Throughput)
	//fmt.Printf("Errors: %v\n", metrics.Errors)
}
