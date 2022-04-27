package fsutil_test

import (
	"os"
	"testing"

	. "github.com/ivanrad/fsutil"
)

func TestRename(t *testing.T) {
	f, err := os.CreateTemp(".", "test_rename_1")
	if err != nil {
		t.Fatal(err)
	}
	tempFilename := f.Name()
	f.Close()
	newFilename := "testdata/" + tempFilename + "_renamed"

	if err := Rename(tempFilename, "testdata/"+tempFilename+"_renamed"); err != nil {
		t.Errorf("Rename(%q, %q) returned an error %v; want nil", tempFilename,
			newFilename, err)
	}

	// cleanup
	if err := os.Remove(newFilename); err != nil {
		t.Fatal(err)
	}
}
