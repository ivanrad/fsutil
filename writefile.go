package fsutil

import (
	"errors"
	"fmt"
	"os"
	"path"
)

var ErrInvalidFilename = errors.New("invalid filename")

// WriteFile writes out data to a temporary file, sets permission, and renames
// it to the target filename in an atomic fashion.
func WriteFile(filename string, data []byte, perm os.FileMode) (err error) {
	dir := path.Dir(filename)
	file := path.Base(filename)
	if file == "." {
		return ErrInvalidFilename
	}

	fd, err := os.Open(dir)
	if err != nil {
		return fmt.Errorf("cannot open %q: %w", dir, err)
	}

	f, err := os.CreateTemp(dir, file)
	if err != nil {
		return fmt.Errorf("failed to create temp file %q: %w", filename, err)
	}
	tempFilename := f.Name()

	defer func() {
		if err != nil {
			_ = f.Close()
			_ = fd.Close()
			_ = os.Remove(tempFilename)
			_ = os.Remove(filename)
		}
	}()

	if _, err = f.Write(data); err != nil {
		return fmt.Errorf("write failed: %w", err)
	}
	if err = f.Chmod(perm); err != nil {
		return fmt.Errorf("chmod failed: %w", err)
	}
	if err = f.Sync(); err != nil {
		return fmt.Errorf("sync failed: %w", err)
	}
	if err = f.Close(); err != nil {
		return fmt.Errorf("close failed: %w", err)
	}
	if err = os.Rename(tempFilename, filename); err != nil {
		return fmt.Errorf("could not rename %q to %q: %w", tempFilename, filename, err)
	}
	// If rename succeeded, keep target file even if dir sync()/close() fails
	if err := fd.Sync(); err != nil {
		return fmt.Errorf("sync failed: %w", err)
	}
	if err := fd.Close(); err != nil {
		return fmt.Errorf("close failed: %w", err)
	}

	return nil
}
