package fsserver

import (
	"log"
	"net/http"
)

type Handler struct {
	FileSystem http.FileSystem
	IndexName  string
	Silent     bool
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.serveOrFail(w, r) != nil {
		ServeError(w, r)
	}
}

func (h *Handler) serveOrFail(w http.ResponseWriter, r *http.Request) error {
	if !h.Silent {
		log.Print("Got URL request: ", r.URL)
	}

	file, err := h.FileSystem.Open(r.URL.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	if info.IsDir() {
		ServeDir(w, r, file)
	} else {
		ServeFile(w, r, file)
	}

	return nil
}
