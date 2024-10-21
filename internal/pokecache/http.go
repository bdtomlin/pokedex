package pokecache

import (
	"bufio"
	"bytes"
	"net/http"
	"net/http/httputil"
)

func rawFromWeb(url string) ([](byte), error) {
	var rawResponse []byte
	res, err := http.Get(url)
	if err != nil {
		return rawResponse, err
	}
	defer res.Body.Close()
	rawResponse, err = httputil.DumpResponse(res, true)
	if err != nil {
		return rawResponse, err
	}
	return rawResponse, nil
}

func responseFromRaw(rawResponse []byte) (*http.Response, error) {
	r := bufio.NewReader(bytes.NewReader(rawResponse))
	res, err := http.ReadResponse(r, &http.Request{})
	if err != nil {
		return &http.Response{}, err
	}
	return res, nil
}

// func logResponseTime(start time.Time) {
// end := time.Now()
// fmt.Printf("\nResponse time: %v\n", end.Sub(start))
// }
