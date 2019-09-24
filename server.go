package main

import (
	"crypto/subtle"
	"fmt"
	"github.com/gorilla/mux"
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

		go RunScript(id)

		_, _ = fmt.Fprintf(w, "Script %v started.\n", vars["id"])
		w.WriteHeader(http.StatusOK)
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
