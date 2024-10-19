package testcache

/*
testcache is used to cache the results of hitting the pokemon api
the cache values are added to the ./cache directory which is checked
in to source controll so that most tests can run without
hitting the external api.
*/

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	dir string
}

func NewCache() *Cache {
	cacheDir := getCacheDir()

	_, err := os.Stat(cacheDir)
	if errors.Is(err, os.ErrNotExist) {
		os.Mkdir(cacheDir, 0777)
	}
	return &Cache{
		dir: cacheDir,
	}
}

func (c *Cache) Add(key string, val []byte) {
	fileName := fileNameHash(key)
	fmt.Println(c.dir + "/" + fileName)
	os.WriteFile(c.dir+"/"+fileName, val, 0666)
}

func (c *Cache) Get(key string) ([]byte, bool) {
	var bytes []byte
	fileName := fileNameHash(key)
	bytes, err := os.ReadFile(c.dir + "/" + fileName)
	if err != nil {
		return bytes, false
	}
	return bytes, true
}

func fileNameHash(key string) string {
	h := md5.New()
	_, err := io.WriteString(h, key)
	if err != nil {
		log.Fatal("couldn't generate file name")
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func getCacheDir() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), ".") + "/cache"
	return dir
}

func (c *Cache) Dump() string {
	return "this is a cache for testing"
}
