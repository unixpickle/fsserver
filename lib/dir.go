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
		if !fileInfo.IsDir() {
			continue
		}
		fileTemplate := map[string]string{
			"name": fileInfo.Name() + "/",
			"size": "-",
			"url":  path.Join(r.URL.Path, fileInfo.Name()) + "/",
		}
		fileTemplates = append(fileTemplates, fileTemplate)
	}
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}
		fileTemplate := map[string]string{
			"name": fileInfo.Name(),
			"size": strconv.FormatInt(fileInfo.Size(), 10),
			"url":  path.Join(r.URL.Path, fileInfo.Name()),
		}
		fileTemplates = append(fileTemplates, fileTemplate)
	}

	bootstrapData, err := Asset("assets/bootstrap.min.css")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	template := map[string]interface{}{
		"path":         r.URL.Path,
		"parent":       PathParent(r.URL.Path),
		"contents":     fileTemplates,
		"bootstrapCSS": string(bootstrapData),
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
