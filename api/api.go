package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/garslo/notifyd/jobs"
	"github.com/garslo/notifyd/stores"
)

type NotifydApi struct {
	store stores.Store
}

func New(store stores.Store) *NotifydApi {
	return &NotifydApi{
		store: store,
	}
}

type PrintJob struct {
	Message string `json:"message"`
	Seconds int    `json:"seconds"`
}

func (me *NotifydApi) AddPrint(w http.ResponseWriter, req *http.Request) {
	log.Printf("AddPrint from %s", req.RemoteAddr)
	printJob := &PrintJob{}
	err := json.NewDecoder(req.Body).Decode(printJob)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	dur := time.Duration(printJob.Seconds) * time.Second
	me.store.Add(jobs.NewPrint(printJob.Message, dur))
	w.Write([]byte(`{"ok":true}`))
}

func (me *NotifydApi) List(w http.ResponseWriter, req *http.Request) {
	log.Printf("List from %s", req.RemoteAddr)
	js, err := me.store.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(js)
}
