package fsserver

import (
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestRootFileSystemRejectsEscapingSymlink(t *testing.T) {
	parent := t.TempDir()
	rootDir := filepath.Join(parent, "root")
	if err := os.Mkdir(rootDir, 0o755); err != nil {
		t.Fatal(err)
	}
	secret := filepath.Join(parent, "secret")
	if err := os.WriteFile(secret, []byte("secret"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(secret, filepath.Join(rootDir, "escape")); err != nil {
		t.Skipf("cannot create symlink: %v", err)
	}

	fsys, err := NewRootFileSystem(rootDir)
	if err != nil {
		t.Fatal(err)
	}
	defer fsys.Close()
	if file, err := fsys.Open("/escape"); err == nil {
		file.Close()
		t.Fatal("opened a symlink which escapes the root")
	}
}

func TestRootFileSystemAllowsInternalSymlink(t *testing.T) {
	rootDir := t.TempDir()
	if err := os.WriteFile(filepath.Join(rootDir, "target"), []byte("data"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink("target", filepath.Join(rootDir, "link")); err != nil {
		t.Skipf("cannot create symlink: %v", err)
	}

	fsys, err := NewRootFileSystem(rootDir)
	if err != nil {
		t.Fatal(err)
	}
	defer fsys.Close()
	file, err := fsys.Open("/link")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "data" {
		t.Fatalf("read %q; want data", data)
	}
}
