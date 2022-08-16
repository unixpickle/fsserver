package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	fsserver "github.com/unixpickle/fsserver/lib"
)

func main() {
	var handler fsserver.Handler

	port := flag.Int("port", 80, "server port number")
	path := flag.String("path", ".", "the directory to serve")
	flag.StringVar(&handler.IndexName, "index", "index.html", "the index filename")
	flag.StringVar(&handler.BasicAuth, "auth", "", "username:password")
	flag.BoolVar(&handler.Silent, "silent", false, "disable logging")
	flag.Parse()

	handler.FileSystem = http.Dir(*path)
	addr := ":" + strconv.Itoa(*port)

	err := http.ListenAndServe(addr, &handler)
	if err != nil && !handler.Silent {
		log.Fatal("ListenAndServe: ", err)
	}
}
