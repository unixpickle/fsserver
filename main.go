package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/unixpickle/fsserver/lib"
)

func main() {
	port := flag.Int("port", 80, "server port number")
	indexName := flag.String("index", "index.html", "the index filename")
	path := flag.String("path", ".", "the directory to serve")
	silent := flag.Bool("silent", false, "disable logging")
	flag.Parse()

	addr := ":" + strconv.Itoa(*port)
	handler := &fsserver.Handler{http.Dir(*path), *indexName, *silent}
	err := http.ListenAndServe(addr, handler)
	if err != nil && !*silent {
		log.Fatal("ListenAndServe: ", err)
	}
}
