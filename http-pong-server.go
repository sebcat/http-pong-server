package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var listenStr = flag.String("listen", ":8080", "listen directive")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong\n"))
	})

	log.Println("Starting server")
	log.Fatal(http.ListenAndServe(*listenStr, nil))

}
