package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Server() {
	config, _ := Config()

	r := mux.NewRouter()

	r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "http-runner is running here.")
	})

	r.HandleFunc("/api/run/{id}", func (w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if !ScriptExists(id) {
			w.WriteHeader(http.StatusNotFound)
			_, _ = fmt.Fprintf(w, "Script %v not exists\n", vars["id"])
			return
		}

		go RunScript(id)

		_, _ = fmt.Fprintf(w, "Script %v started.\n", vars["id"])
		w.WriteHeader(http.StatusOK)
		//runCmd()
	})

	http.Handle("/", r)

	fmt.Println("HTTP Server running on http://" + config.Host + ":" + config.Port + "/")
	err := http.ListenAndServe(config.Host + ":"+ config.Port, nil)
	if err != nil {
		fmt.Printf("Error during starting http server : %s", err)
	}

}

