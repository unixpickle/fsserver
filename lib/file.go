package fsserver

import (
	"io"
	"mime"
	"net/http"
	"path"
)

func ServeFile(w http.ResponseWriter, r *http.Request, f http.File) {
	// TODO: set the Content-Length
	ext := path.Ext(r.URL.Path)
	mimeType := mime.TypeByExtension(ext)
	w.Header().Set("Content-Type", mimeType)
	io.Copy(w, f)
}
