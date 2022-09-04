package request

import (
	"io"
	"net/http"
)

func GetUrl(url, cookies string) ([]byte, error) {
	cli := http.Client{}
	header, _ := http.NewRequest("GET", url, nil)
	header.Header.Set("cookie", cookies)
	req, err := cli.Do(header)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(req.Body)
}
