package main

import (
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Server() {

	r := mux.NewRouter()

	config, _ := Config()

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

		logname := CreateLog(id)
		body, _ := ioutil.ReadAll(r.Body)
		go RunScript(id, logname, "", string(body))

		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, "{\"started\": true, \"logs\": \"%s\", \"logs_entrypoint\": \"/api/logs/%s/%s\"}", logname, id, logname)
		w.WriteHeader(http.StatusOK)
	})

	r.HandleFunc("/api/logs/{id}/{file}", func (w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		file := vars["file"]

		if !ScriptExists(id) {
			w.WriteHeader(http.StatusNotFound)
			_, _ = fmt.Fprintf(w, "Script %v not exists\n", vars["id"])
			return
		}

		w.WriteHeader(http.StatusOK)

		content := ContentLogFile(id, file)
		_, _ = w.Write(content)

	})

	r.HandleFunc("/api/logs/{id}", func (w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if !ScriptExists(id) {
			w.WriteHeader(http.StatusNotFound)
			_, _ = fmt.Fprintf(w, "Script %v not exists\n", vars["id"])
			return
		}

		w.WriteHeader(http.StatusOK)

		list := ListLogFiles(id)
		json, _ := json.Marshal(list)

		_, _ = w.Write(json)

	})

	if config.Security.Auth_type == "BASIC_AUTH" {
		r.Use(basicAuthMiddleware(config.Security.Basic_auth.Login, config.Security.Basic_auth.Password, "http-runner auth"))
	}

	if len(config.Security.Ip_authorised) > 0 {
		r.Use(testIpMiddleware(config.Security.Ip_authorised))
	}

	http.Handle("/", r)

	fmt.Println("HTTP Server running on http://" + config.Host + ":" + config.Port + "/")
	err := http.ListenAndServe(config.Host + ":"+ config.Port, nil)
	if err != nil {
		fmt.Printf("Error during starting http server : %s", err)
	}

}

func basicAuthMiddleware(username, password, realm string) mux.MiddlewareFunc {

	return func (next http.Handler) http.Handler {

		return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){

			user, pass, ok := r.BasicAuth()

			if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1{
				w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
				w.WriteHeader(401)
				w.Write([]byte("Unauthorised.\n"))
				return
			}

			log.Println(r.RequestURI)
			next.ServeHTTP(w, r)
		})
	}
}

func testIpMiddleware(ipAuthorised []string) mux.MiddlewareFunc {

	return func (next http.Handler) http.Handler {

		return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){

			clientIp := strings.Split(r.RemoteAddr, ":")[0]

			if !ArrayContains(ipAuthorised, clientIp) {
				log.Println("Unautorised IP connexion : " + clientIp)
				w.WriteHeader(401)
				w.Write([]byte("IP Unauthorised.\n"))
				return
			}

			log.Println(r.RequestURI)
			next.ServeHTTP(w, r)
		})
	}
}
