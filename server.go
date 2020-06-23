package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	port := os.Getenv("PORT")

	http.HandleFunc("/", logRequest)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func logRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s\n", time.Now().Format("2006-01-02 15:04:05 -0700"))
	log.Println("***************************")
	log.Printf("method:           %s", r.Method)
	log.Printf("host:             %s", r.Host)
	log.Printf("remote address:   %s", r.RemoteAddr)
	log.Printf("url:              %s", r.URL.String())
	log.Print("\nheader:\n")
	for k, vv := range r.Header {
		fmt.Printf("  %s", k)
		for _, v := range vv {
			if len(k) < 30 {
				s := strings.Repeat(" ", 30-len(k))
				fmt.Printf("  %s: ", s)
			} else {
				fmt.Print(": ")
			}
			fmt.Println(v)
		}
	}

	if r.Method == http.MethodOptions {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE ,OPTIONS")
		w.WriteHeader(http.StatusOK)
		return
	}


	buf := new(bytes.Buffer)
	i, err := buf.ReadFrom(r.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if i > 0 {
		fmt.Print("\nbody:\n")
		fmt.Println(buf.String())
	}
	w.WriteHeader(http.StatusOK)
}
