package pokecache

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

type TestCache struct {
	dir string
}

func NewTestCache() *TestCache {
	cacheDir := getCacheDir()

	_, err := os.Stat(cacheDir)
	if errors.Is(err, os.ErrNotExist) {
		os.Mkdir(cacheDir, 0777)
	}
	return &TestCache{
		dir: cacheDir,
	}
}

func (c *TestCache) Add(key string, val []byte) {
	fileName := fileNameHash(key)
	os.WriteFile(c.dir+"/"+fileName, val, 0666)
}

func (c *TestCache) Get(url string) ([]byte, error) {
	defer logResponseTime(time.Now())
	rawResponse, ok := c.readFromCache(url)
	if !ok {
		res, err := rawFromWeb(url)
		if err != nil {
			return res, err
		}
		rawResponse = res
	}

	c.Add(url, rawResponse)
	res, err := responseFromRaw(rawResponse)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 399 {
		return []byte{}, errors.New(res.Status)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return body, err
	}
	return body, nil
}

func (c *TestCache) readFromCache(url string) ([](byte), bool) {
	var bytes []byte
	fileName := fileNameHash(url)
	bytes, err := os.ReadFile(fileName)
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
	dir := path.Join(path.Dir(filename), ".") + "/testcache"
	return dir
}

func (c *TestCache) Dump() string {
	return "This is a cache for testing. Cache is saved in the internal/pokecache/testcache directory"
}
