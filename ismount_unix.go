//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris

package fsutil

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"syscall"
)

// IsMount returns true if pathname name is a mount point.
func IsMount(name string) (bool, error) {
	fileInfo, err := os.Lstat(name)
	if err != nil {
		return false, fmt.Errorf("could not stat %q: %w", name, err)
	}
	if fileInfo.Mode()&fs.ModeSymlink != 0 {
		return false, nil
	}
	parent := path.Join(name, "..")
	parentFileInfo, err := os.Lstat(parent)
	if err != nil {
		return false, fmt.Errorf("could not stat %q: %w", parent, err)
	}
	st1, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		return false, ErrUnsupported
	}
	st2, ok := parentFileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		return false, ErrUnsupported
	}

	if st1.Dev != st2.Dev {
		return true, nil
	}
	if st1.Ino == st2.Ino {
		return true, nil
	}
	return false, nil
}
