package fsserver

import (
	"io"
	"mime"
	"net/http"
	"path"
	"strconv"
)

func ServeFile(w http.ResponseWriter, r *http.Request, f http.File) {
	stats, err := f.Stat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	ext := path.Ext(r.URL.Path)
	mimeType := mime.TypeByExtension(ext)
	if mimeType != "" {
		w.Header().Set("Content-Type", mimeType)
	}
	w.Header().Set("Content-Length", strconv.FormatInt(stats.Size(), 10))
	io.Copy(w, f)
}
