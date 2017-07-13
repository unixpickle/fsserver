package fsserver

import "net/http"

func ServeNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	ServeTemplate(w, "404", map[string]interface{}{
		"parent": PathParent(r.URL.Path),
		"path":   r.URL.Path,
	})
}
