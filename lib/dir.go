package fsserver

import (
	"net/http"
	"path"

	"github.com/dustin/go-humanize"
)

func ServeDir(w http.ResponseWriter, r *http.Request, f http.File) {
	fileInfos, err := f.Readdir(0)
	if err != nil {
		ServeNotFound(w, r)
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
			"size": humanize.Bytes(uint64(fileInfo.Size())),
			"url":  path.Join(r.URL.Path, fileInfo.Name()),
		}
		fileTemplates = append(fileTemplates, fileTemplate)
	}

	ServeTemplate(w, "dir", map[string]interface{}{
		"path":     r.URL.Path,
		"parent":   PathParent(r.URL.Path),
		"contents": fileTemplates,
	})
}
