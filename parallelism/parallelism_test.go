package parallelism

import (
	"fmt"
	"sort"
	"testing"
)

func TestParallelism(t *testing.T) {
	// ok case:
	m, _ := MD5All("./")

	var paths []string
	for path := range m {
		paths = append(paths, path)
	}

	sort.Strings(paths)

	for _, path := range paths {
		fmt.Printf("%x  %s\n", m[path], path)
	}

	// error case:
	_, err := MD5All("nosuchdir")
	fmt.Println(err)
}
