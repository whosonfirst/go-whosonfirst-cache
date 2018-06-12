package main

import (
	"flag"
	"github.com/whosonfirst/go-whosonfirst-cache"
	"log"
	"os"
)

func main() {

	null_cache := flag.Bool("null-cache", false, "...")
	go_cache := flag.Bool("go-cache", false, "...")
	fs_cache := flag.Bool("fs-cache", false, "...")
	fs_root := flag.String("fs-root", "", "...")

	flag.Parse()

	caches := make([]cache.Cache, 0)

	if *null_cache {

		c, err := cache.NewNullCache()

		if err != nil {
			log.Fatal(err)
		}

		caches = append(caches, c)
	}

	if *go_cache {

		opts, err := cache.DefaultGoCacheOptions()

		if err != nil {
			log.Fatal(err)
		}

		c, err := cache.NewGoCache(opts)

		if err != nil {
			log.Fatal(err)
		}

		caches = append(caches, c)
	}

	if *fs_cache {

		if *fs_root == "" {

			cwd, err := os.Getwd()

			if err != nil {
				log.Fatal(err)
			}

			*fs_root = cwd
		}

		c, err := cache.NewFSCache(*fs_root)

		if err != nil {
			log.Fatal(err)
		}

		caches = append(caches, c)
	}

	c, err := cache.NewMultiCache(caches)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(c)
}
