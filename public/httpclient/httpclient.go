package httpclient

import (
	"bytes"
	"crypto/tls"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

var (
	gCurCookieJar  *cookiejar.Jar
	ErrPoolContent                   = errors.New("pool content type error")
	readTimeout                      = 30 * time.Second
	tr             http.RoundTripper = &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}
)

type Config struct {
	Timeout  time.Duration
	PoolSize int
}

func init() {
	gCurCookieJar, _ = cookiejar.New(nil)
}

func getTransport() http.RoundTripper {
	return tr
}

func Init(conf Config) error {
	readTimeout = conf.Timeout * time.Second
	if conf.PoolSize <= 0 {
		return nil
	}
	tr = &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		DisableCompression:  true,
		MaxIdleConnsPerHost: conf.PoolSize,
	}
	return nil
}

func GetCookies(url *url.URL) []*http.Cookie {
	return gCurCookieJar.Cookies(url)
}

func PostDataCookie(url string, params url.Values, headers map[string]string) ([]byte, *http.Response, error) {
	cli := &http.Client{
		Jar:       gCurCookieJar,
		Transport: getTransport(),
	}
	var ioParams io.Reader
	if params != nil {
		ioParams = bytes.NewReader([]byte(params.Encode()))
	}
	req, err := http.NewRequest(http.MethodPost, url, ioParams)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return buff, resp, nil
}

func GetDataCookie(url string, params url.Values, headers map[string]string) ([]byte, *http.Response, error) {
	cli := &http.Client{
		Jar:       gCurCookieJar,
		Transport: getTransport(),
	}
	buf := bytes.NewBufferString(url)
	if params != nil {
		buf.WriteString("?")
		buf.WriteString(params.Encode())
	}
	req, err := http.NewRequest(http.MethodGet, buf.String(), nil)
	if err != nil {
		return nil, nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return buff, resp, nil
}

func GetData(url string, params url.Values) ([]byte, error) {
	cli := &http.Client{
		Timeout:   readTimeout,
		Transport: getTransport(),
	}
	buf := bytes.NewBufferString(url)
	if params != nil {
		buf.WriteString("?")
		buf.WriteString(params.Encode())
	}
	req, err := http.NewRequest(http.MethodGet, buf.String(), nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.75 Safari/537.36")
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return buff, nil
}

func Get(url string, params url.Values) (*http.Response, error) {
	cli := &http.Client{
		Timeout:   readTimeout,
		Transport: getTransport(),
	}
	buf := bytes.NewBufferString(url)
	if params != nil {
		buf.WriteString("?")
		buf.WriteString(params.Encode())
	}
	req, err := http.NewRequest(http.MethodGet, buf.String(), nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.75 Safari/537.36")
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
func GetDataWithHeader(url string, params url.Values, headers map[string]string) ([]byte, *http.Response, error) {
	cli := &http.Client{
		Transport: getTransport(),
	}
	buf := bytes.NewBufferString(url)
	if params != nil {
		buf.WriteString("?")
		buf.WriteString(params.Encode())
	}
	req, err := http.NewRequest(http.MethodGet, buf.String(), nil)
	if err != nil {
		return nil, nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return buff, resp, nil
}

func PostData(url string, params url.Values) ([]byte, error) {
	cli := &http.Client{
		Timeout:   readTimeout,
		Transport: getTransport(),
	}
	var ioParams io.Reader
	if params != nil {
		ioParams = bytes.NewReader([]byte(params.Encode()))
	}
	req, err := http.NewRequest(http.MethodPost, url, ioParams)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.75 Safari/537.36")
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return buff, nil
}

func PostDataWithHeader(url string, params url.Values, headers map[string]string) ([]byte, *http.Response, error) {
	cli := &http.Client{
		Transport: getTransport(),
	}
	var ioParams io.Reader
	if params != nil {
		ioParams = bytes.NewReader([]byte(params.Encode()))
	}
	req, err := http.NewRequest(http.MethodPost, url, ioParams)
	if err != nil {
		return nil, nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return buff, resp, nil
}

func PostBinary(url string, params url.Values) ([]byte, error) {
	cli := &http.Client{
		Timeout:   readTimeout,
		Transport: getTransport(),
	}
	var ioParams io.Reader
	if params != nil {
		ioParams = bytes.NewReader([]byte(params.Encode()))
	}
	req, err := http.NewRequest(http.MethodPost, url, ioParams)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.75 Safari/537.36")
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return buff, nil
}

func PostBody(url string, body io.Reader) ([]byte, error) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"

	buff, _, err := PostBodyWithHeader(url, body, headers)
	if err != nil {
		return nil, err
	}
	return buff, nil
}

func PostBodyWithHeader(url string, body io.Reader, headers map[string]string) ([]byte, *http.Response, error) {
	cli := &http.Client{
		Timeout:   readTimeout,
		Transport: getTransport(),
	}

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return buff, resp, nil
}

func SetRoundTripper(rt http.RoundTripper) {
	tr = rt
}
