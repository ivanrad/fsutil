package fsutil_test

import (
	"log"

	. "github.com/ivanrad/fsutil"
)

func ExampleWriteFile() {
	filename := "testdata/fsutil_writefile_0644.txt"
	if err := WriteFile(filename, []byte("hello world!\n"), 0o644); err != nil {
		log.Fatal(err)
	}
	// Output:
}
