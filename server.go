package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/boltdb/bolt"
)

type uriDB struct {
	db *bolt.DB
}

func (u *uriDB) queryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(os.Stdout, r.URL.Path)
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "OK")
		return
	}
	uri := strings.Split(r.URL.Path, "/")[1]
	fullURL, err := getLongURL(u.db, uri)
	if err != nil {
		http.NotFound(w, r)
	}
	if fullURL == "" {
		http.NotFound(w, r)

	} else {
		http.Redirect(w, r, fullURL, 302)

	}

}

func httpServer() {

	ud := uriDB{db: getDB()}
	mux := http.NewServeMux()

	mux.HandleFunc("/", ud.queryHandler)

	server := http.Server{Addr: ":80"}
	server.Handler = logRequest(mux)
	err := server.ListenAndServe()

	if err != nil {
		panic("unable to start server" + err.Error())

	}
	fmt.Println("Server listening on port 80.")
}

func logRequest(handler http.Handler) http.Handler {
	log := getLogger()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("ipaddress=%s httpMethod=%s URL=%s Agent=%s", r.RemoteAddr, r.Method, r.URL, r.UserAgent())
		handler.ServeHTTP(w, r)

	})
}
