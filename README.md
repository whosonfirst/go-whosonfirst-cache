# go-whosonfirst-cache

Go package for reading and writing Who's On First documents from a variety of sources.

## Important

This package has been officially deprecated. Please use the [go-cache](https://github.com/whosonfirst/go-cache) package instead.

## Install

You will need to have both `Go` (specifically version [1.12](https://golang.org/dl/) or higher) and the `make` programs installed on your computer. Assuming you do just type:

```
make tools
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Interfaces

```
type Cache interface {
     	Name() string
	Get(string) (io.ReadCloser, error)
	Set(string, io.ReadCloser) (io.ReadCloser, error)
	Unset(string) error
	Hits() int64
	Misses() int64
	Evictions() int64
	Size() int64
}
```

## See also

