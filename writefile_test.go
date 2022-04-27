package fsutil_test

import (
	"errors"
	"os"
	"testing"

	. "github.com/ivanrad/fsutil"
)

func TestWriteFile(t *testing.T) {
	filename := "fsutil_writefile_0644.txt"
	defer os.Remove(filename)
	if err := WriteFile(filename, []byte("hello world!\n"), 0o644); err != nil {
		t.Fatalf("WriteFile(%q) returned an error: %v; want nil", filename, err)
	}
}

func TestWriteFile_TmpDir(t *testing.T) {
	filename := "/tmp/fsutil_writefile_0644.txt"
	defer os.Remove(filename)
	if err := WriteFile(filename, []byte("hello world!\n"), 0o644); err != nil {
		t.Fatalf("WriteFile(%q) returned an error: %v; want nil", filename, err)
	}
}

func TestWriteFile_EmptyFilename(t *testing.T) {
	filename := ""
	err := WriteFile(filename, []byte("hello world!\n"), 0o644)
	if !errors.Is(err, ErrInvalidFilename) {
		t.Fatalf("WriteFileName(%q) returned error %v; expected %v\n", filename,
			err, ErrInvalidFilename)
	}
}
