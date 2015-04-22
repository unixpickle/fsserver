package fsserver

import (
	"net/http"
	"path"
	"strconv"

	"github.com/hoisie/mustache"
)

func ServeDir(w http.ResponseWriter, r *http.Request, f http.File) {
	fileInfos, err := f.Readdir(0)
	if err != nil {
		ServeError(w, r)
		return
	}

	fileTemplates := []map[string]string{}
	for _, fileInfo := range fileInfos {
		name := fileInfo.Name()
		size := strconv.FormatInt(fileInfo.Size(), 10)
		if fileInfo.IsDir() {
			name += "/"
			size = "-"
		}
		fileTemplate := map[string]string{
			"name": name,
			"size": size,
			"url":  path.Join(r.URL.Path, fileInfo.Name()),
		}
		fileTemplates = append(fileTemplates, fileTemplate)
	}

	template := map[string]interface{}{
		"path":     r.URL.Path,
		"parent":   path.Dir(r.URL.Path),
		"contents": fileTemplates,
	}

	data, err := Asset("assets/dir.mustache")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	content := mustache.Render(string(data), template)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(content))
}
