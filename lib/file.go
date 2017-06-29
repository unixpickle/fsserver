package fsserver

import (
	"net/http"
	"path"
)

func ServeFile(w http.ResponseWriter, r *http.Request, f http.File) {
	stats, err := f.Stat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.ServeContent(w, r, path.Base(r.URL.Path), stats.ModTime(), f)
}
