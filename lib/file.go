package fsserver

import "net/http"

func ServeFile(w http.ResponseWriter, r *http.Request, name string, f http.File) {
	stats, err := f.Stat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.ServeContent(w, r, name, stats.ModTime(), f)
}
