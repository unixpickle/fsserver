package fsserver

import (
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/unixpickle/seektar"
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

func ServeDirTar(w http.ResponseWriter, r *http.Request, vfs http.FileSystem, p string) {
	filename := path.Base(p)
	if filename == "/" {
		filename = "root"
	}
	filename = filename + ".tar"
	tar, err := seektar.TarHTTP(vfs, p, path.Base(p))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rs, err := tar.Open()
	defer rs.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/x-tar")
	w.Header().Set("Content-Disposition", "attachment; filename*=UTF-8''"+url.PathEscape(filename))
	http.ServeContent(w, r, filename, time.Now(), rs)
}
