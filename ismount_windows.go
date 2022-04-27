//go:build windows
// +build windows

package fsutil

import "syscall"

// IsMount returns true if pathname name is a mount point.
func IsMount(name string) (bool, error) {
	ptr, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return false, err
	}
	attrs, err := syscall.GetFileAttributes(ptr)
	if err != nil {
		return false, err
	}

	return attrs&syscall.FILE_ATTRIBUTE_REPARSE_POINT != 0, nil
}
