package boundedparallelism

import (
	"fmt"
	"sort"
	"testing"
)

func TestBoundedParallelism(t *testing.T) {
	m, _ := MD5All("./")

	var paths []string
	for path := range m {
		paths = append(paths, path)
	}

	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x  %s\n", m[path], path)
	}

	_, err := MD5All("nosuchdir")
	fmt.Println(err)
}
