package fsserver

import (
	"net/http"

	"github.com/hoisie/mustache"
)

func ServeError(w http.ResponseWriter, r *http.Request) {
	template := map[string]interface{}{"path": r.URL.Path}
	data, err := Asset("assets/404.mustache")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	content := mustache.Render(string(data), template)
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(content))
}
