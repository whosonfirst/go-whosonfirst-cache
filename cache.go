package cache

import (
	"io"
	_ "log"
)

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }

type Cache interface {
	Get(string) (io.ReadCloser, error)
	Set(string, io.ReadCloser) (io.ReadCloser, error)
	Unset(string) error
	Hits() int64
	Misses() int64
	Evictions() int64
	Size() int64
}
