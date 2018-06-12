package cache

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
)

type FSCache struct {
	Cache
	root   string
	misses int64
	hits   int64
	mu *sync.RWMutex
}

func NewFSCache(root string) (Cache, error) {

	abs_root, err := filepath.Abs(root)

	if err != nil {
		return nil, err
	}

	info, err := os.Stat(abs_root)

	if os.IsNotExist(err) {
		return nil, errors.New("Root doesn't exist")
	}

	if !info.IsDir() {
		return nil, errors.New("Root is not a directory")
	}

	mu := new(sync.RWMutex)

	c := FSCache{
		root:   abs_root,
		hits:   int64(0),
		misses: int64(0),
		mu:     mu,
	}

	return &c, nil
}

func (c *FSCache) Get(key string) (io.ReadCloser, error) {

	c.mu.RLock()
	defer c.mu.RUnlock()

	abs_path := filepath.Join(c.root, key)
	fh, err := os.Open(abs_path)

	if err != nil {
		atomic.AddInt64(&c.misses, 1)
		return nil, err
	}

	atomic.AddInt64(&c.hits, 1)
	return fh, nil
}

func (c *FSCache) Set(key string, fh io.ReadCloser) (io.ReadCloser, error) {

	c.mu.Lock()
	defer c.mu.Unlock()

	abs_path := filepath.Join(c.root, key)
	abs_root := filepath.Dir(abs_path)

	_, err := os.Stat(abs_root)

	if os.IsNotExist(err) {

		err = os.MkdirAll(abs_root, 0755)

		if err != nil {
			return nil, err
		}
	}

	out, err := os.OpenFile(abs_path, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		return nil, err
	}

	_, err = io.Copy(out, fh)

	out.Close()

	if err != nil {
		return nil, err
	}

	return c.Get(key)
}

func (c *FSCache) Unset(key string) error {

	c.mu.Lock()
	defer c.mu.Unlock()

	abs_path := filepath.Join(c.root, key)
	abs_root := filepath.Dir(abs_path)

	_, err := os.Stat(abs_root)

	if os.IsNotExist(err) {
		return nil
	}

	return os.Remove(abs_path)
}

// TO DO: walk c.root

func (c *FSCache) Size() int64 {
	return 0
}

func (c *FSCache) Hits() int64 {
	return c.hits
}

func (c *FSCache) Misses() int64 {
	return c.misses
}

func (c *FSCache) Evictions() int64 {
	return 0
}
