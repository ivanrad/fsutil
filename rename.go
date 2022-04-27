package fsutil

import (
	"fmt"
	"os"
	"path"
)

func fdirsync(dirname string) error {
	fd, err := os.Open(dirname)
	if err != nil {
		return fmt.Errorf("could not open %q: %w", dirname, err)
	}

	if err = fd.Sync(); err != nil {
		_ = fd.Close()
		return fmt.Errorf("sync failed: %w", err)
	}

	if err = fd.Close(); err != nil {
		return fmt.Errorf("close failed: %w", err)
	}
	return nil
}

// Rename renames/moves oldpath to newpath with dirsync-ish semantics.
func Rename(oldpath, newpath string) error {
	if err := os.Rename(oldpath, newpath); err != nil {
		return fmt.Errorf("could not rename %q to %q: %w", oldpath, newpath, err)
	}

	newDirname := path.Dir(newpath)
	if err := fdirsync(newDirname); err != nil {
		return err
	}
	oldDirname := path.Dir(oldpath)
	if oldDirname != newDirname {
		return fdirsync(oldDirname)
	}

	return nil
}
