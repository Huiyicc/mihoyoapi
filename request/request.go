package request

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Https struct {
	Cookies *AppCookies
}

// NewClient 创建一个请求客户端
func NewClient(cookies *AppCookies) *Https {
	return &Https{
		Cookies: cookies,
	}
}

func (t *Https) SetCookies(cookies *AppCookies) {
	t.Cookies = cookies
}

// Get 用于内部请求
func (t *Https) Get(req RequestStruct, dsType int, headerFunc func(r *http.Request)) ([]byte, error) {
	uri := req.Url
	if req.Query != "" {
		uri += "?" + req.Query
	}
	var (
		resp *http.Response
		err  error
	)
	cli := &http.Client{}
	if proxy {
		httpProxy(cli)
	}
	requ, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	requ.Header = t.Cookies.GetHeadersMap(req, dsType)
	if headerFunc != nil {
		headerFunc(requ)
	}
	//处理返回结果
	resp, _ = cli.Do(requ)
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return data, nil
}

// Post 用于内部请求
func (t *Https) Post(req RequestStruct, dsType int, headerFunc func(*http.Request)) ([]byte, error) {
	uri := req.Url
	if req.Query != "" {
		uri += "?" + req.Query
	}
	var (
		resp *http.Response
		err  error
		cli  http.Client
	)
	if proxy {
		httpProxy(&cli)
	}
	requ, err := http.NewRequest("POST", uri, strings.NewReader(req.Body.GetData()))
	if err != nil {
		return nil, err
	}
	requ.Header = t.Cookies.GetHeadersMap(req, dsType)
	if headerFunc != nil {
		headerFunc(requ)
	}
	//处理返回结果
	resp, _ = cli.Do(requ)
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return data, nil
}

var proxy = true

func httpProxy(client *http.Client) {
	proxyStr := "http://127.0.0.1:8080"
	urls, _ := url.Parse(proxyStr)
	p := http.ProxyURL(urls)
	client.Transport = &http.Transport{
		Proxy: p,
	}
}
