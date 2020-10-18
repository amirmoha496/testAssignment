package main

import (
	"context"
	"net/http"
)

var hm HandlerManager = HandlerManager{}

func main() {
	mux := http.NewServeMux()
	srv := http.Server{Addr: port, Handler: mux}

	mux.HandleFunc("/hash", hm.EncodePasswordHandler)
	mux.HandleFunc("/hash/1", hm.GetHashHandler)
	mux.HandleFunc("/stats", hm.StatsHandler)
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		getLogger().Info("Shutting Down Server")
		srv.Shutdown(context.Background())
	})

	err := srv.ListenAndServe()

	if err != nil {
		getLogger().Error("Failed to start Server:" + err.Error())
	} else {
		getLogger().Info("Server started at port 8080")
	}

}

const port string = ":8080"
