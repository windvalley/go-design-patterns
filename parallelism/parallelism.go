package parallelism

import (
	"crypto/md5"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

func sumFiles(done <-chan struct{}, root string) (<-chan result, <-chan error) {
	c := make(chan result)
	errc := make(chan error, 1)

	go func() {
		var wg sync.WaitGroup
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}

			wg.Add(1)
			go func() {
				data, err := ioutil.ReadFile(path)
				select {
				case c <- result{path, md5.Sum(data), err}:
				case <-done:
				}
				wg.Done()
			}()

			// Abort the walk if done is closed.
			select {
			case <-done:
				return errors.New("walk canceled")
			default:
				return nil
			}
		})

		// Walk has returned, so all calls to wg.Add are done. Start a
		// goroutine to close c once all the sends are done.
		go func() {
			wg.Wait()
			close(c)
		}()

		// No select needed here, since errc is buffered.
		errc <- err
	}()

	return c, errc
}

// MD5All ...
func MD5All(root string) (map[string][md5.Size]byte, error) {
	done := make(chan struct{})
	defer close(done)

	c, errc := sumFiles(done, root)

	m := make(map[string][md5.Size]byte)
	for r := range c {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}

	if err := <-errc; err != nil {
		return nil, err
	}

	return m, nil
}
