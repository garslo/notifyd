package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/garslo/notifyd/api"
	"github.com/garslo/notifyd/runners"
	"github.com/garslo/notifyd/stores"
)

func main() {
	var (
		granularity time.Duration
	)
	flag.DurationVar(&granularity, "g", 500*time.Millisecond, "granularity")
	flag.Parse()

	store := stores.NewMemory()
	runner := runners.New(store, granularity)
	go runner.RunForever()

	web := api.New(store)
	http.HandleFunc("/", web.List)
	http.HandleFunc("/print", web.AddPrint)
	http.ListenAndServe(":8080", nil)
}
