package fsserver

import (
	"net/http"
	"os"
	"path"
	"strings"
)

// RootFileSystem is an HTTP file system confined to a directory tree. It
// follows symlinks within the tree but rejects symlinks which escape it.
type RootFileSystem struct {
	root *os.Root
}

// NewRootFileSystem opens name as the root of a confined HTTP file system.
func NewRootFileSystem(name string) (*RootFileSystem, error) {
	root, err := os.OpenRoot(name)
	if err != nil {
		return nil, err
	}
	return &RootFileSystem{root: root}, nil
}

// Open implements http.FileSystem.
func (r *RootFileSystem) Open(name string) (http.File, error) {
	// HTTP file-system paths are slash-separated and conventionally begin with
	// a slash, while os.Root accepts paths relative to the root.
	name = strings.TrimPrefix(path.Clean("/"+name), "/")
	if name == "" {
		name = "."
	}
	return r.root.Open(name)
}

// Close releases the operating-system resources held by the root.
func (r *RootFileSystem) Close() error {
	return r.root.Close()
}
