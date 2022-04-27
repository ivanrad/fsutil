//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris

package fsutil_test

import (
	"testing"

	. "github.com/ivanrad/fsutil"
)

func TestIsMount(t *testing.T) {
	dir := "/"
	got, err := IsMount(dir)
	if err != nil {
		t.Fatalf("IsMount(%q) returned an error: %v", dir, err)
	}
	if got != true {
		t.Errorf("IsMount(%q) = %v; want true", dir, got)
	}
}
