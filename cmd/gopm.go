package main

import (
	"flag"
	"gopm"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "listen to address")
var domain = flag.String("domain", "gopm.in", "The serving domain")

func main() {
	flag.Parse()
	gopm.SetDomain(*domain)
	http.HandleFunc("/", gopm.MainHandler)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
