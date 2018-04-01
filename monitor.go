package main

import (
	"flag"
	"log"
	"net/http"
)

var p = flag.String("p", ":8080", "port to listen")
var d = flag.String("d", "/usr/share/doc", "directory to serve")

func main() {
	flag.Parse()
	log.Printf("Serving directory %s on port %s\n", *d, *p)
	log.Fatal(http.ListenAndServe(*p, http.FileServer(http.Dir(*d))))
}
