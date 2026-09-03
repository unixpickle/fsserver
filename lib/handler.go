package fsserver

import (
	"crypto/sha256"
	"crypto/subtle"
	"log"
	"net/http"
	"path"
	"strings"
)

type Handler struct {
	FileSystem http.FileSystem
	IndexName  string
	Silent     bool
	BasicAuth  string
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !h.authenticate(r) {
		if !h.Silent {
			log.Print("Unauthenticated request from: " + r.RemoteAddr)
		}
		w.Header().Set("WWW-Authenticate", "Basic realm=\"fsserver\"")
		http.Error(w, "Login incorrect.", http.StatusUnauthorized)
	} else if h.serveOrFail(w, r) != nil {
		ServeNotFound(w, r)
	}
}

func (h *Handler) authenticate(r *http.Request) bool {
	if h.BasicAuth == "" {
		return true
	}
	username, password, ok := r.BasicAuth()
	if !ok {
		return false
	}
	expectedUsername, expectedPassword, valid := strings.Cut(h.BasicAuth, ":")
	if !valid {
		return false
	}

	usernameHash := sha256.Sum256([]byte(username))
	expectedUsernameHash := sha256.Sum256([]byte(expectedUsername))
	passwordHash := sha256.Sum256([]byte(password))
	expectedPasswordHash := sha256.Sum256([]byte(expectedPassword))

	usernameMatch := subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:])
	passwordMatch := subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:])
	return usernameMatch&passwordMatch == 1
}

func (h *Handler) serveDir(w http.ResponseWriter, r *http.Request,
	file http.File) {
	indexFile, err := h.FileSystem.Open(path.Join(r.URL.Path, h.IndexName))
	if err != nil {
		if r.URL.Query().Get("download") != "" {
			ServeDirTar(w, r, h.FileSystem, r.URL.Path)
		} else {
			ServeDir(w, r, file)
		}
	} else {
		defer indexFile.Close()
		ServeFile(w, r, h.IndexName, indexFile)
	}
}

func (h *Handler) serveOrFail(w http.ResponseWriter, r *http.Request) error {
	if !h.Silent {
		log.Print("Got URL request: ", r.URL)
	}

	file, err := h.FileSystem.Open(r.URL.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	if info.IsDir() {
		if !strings.HasSuffix(r.URL.Path, "/") {
			http.Redirect(w, r, r.URL.Path+"/", http.StatusTemporaryRedirect)
			return nil
		}
		h.serveDir(w, r, file)
	} else {
		ServeFile(w, r, path.Base(r.URL.Path), file)
	}

	return nil
}
