package utils

import (
	"crypto/tls"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/guguducken/octopus/pkg/common"
	"github.com/guguducken/octopus/pkg/config"
)

var (
	ErrHttpResponse = errors.New("github api server response is not 2xx")
)

type Reply struct {
	Body         []byte
	StatusCode   int
	Headers      http.Header
	RateResource string
	RateLimit    *common.RateLimitUnit
}

type URL struct {
	RawURL   string
	Endpoint string
	Path     string
	Params   map[string]string
}

func GetWithRetry(cfg *config.Config, url URL) (reply *Reply, err error) {
	req, err := http.NewRequest("GET", url.toRawURL(), nil)
	if err != nil {
		return reply, err
	}

	retryTimes := cfg.GetRetryTimes()
	retryDuration := cfg.GetRetryDuration()
	for i := 0; i < retryTimes; i++ {
		reply, err = do(cfg, req)
		if err == nil {
			break
		}
		if err != nil && err != http.ErrHandlerTimeout {
			break
		}
		time.Sleep(retryDuration)
	}
	return reply, err
}

func Get(cfg *config.Config, url URL) (reply *Reply, err error) {
	req, err := http.NewRequest("GET", url.toRawURL(), nil)
	if err != nil {
		return reply, err
	}
	return resolveReply(do(cfg, req))
}

func Post(cfg *config.Config, url URL, body string) (reply *Reply, err error) {
	req, err := http.NewRequest("POST", url.toRawURL(), strings.NewReader(body))
	if err != nil {
		return reply, err
	}
	return resolveReply(do(cfg, req))
}

func do(cfg *config.Config, req *http.Request) (reply *Reply, err error) {
	reply = &Reply{
		RateLimit: &common.RateLimitUnit{},
	}

	// set GitHub related headers, example: bearer token, api version
	setGithubHeader(req, cfg)

	// do http request
	resp, err := newHttpClient(cfg).Do(req)
	if err != nil {
		return reply, err
	}

	// read http status code from resp
	reply.StatusCode = resp.StatusCode

	// read headers from resp
	reply.Headers = resp.Header

	// read reply from resp
	reply.Body, err = io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return reply, err
	}
	return reply, nil
}

func setGithubHeader(req *http.Request, cfg *config.Config) {
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+cfg.GetToken())
	req.Header.Set("X-GitHub-Api-Version", cfg.GetGithubAPIVersion())
}

func newHttpClient(cfg *config.Config) *http.Client {
	c := http.Client{}

	// transport settings
	t := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: cfg.GetTLSCheck(),
		},
	}
	if cfg.GetProxyFromEnv() {
		t.Proxy = http.ProxyFromEnvironment
	}

	c.Transport = t
	c.Timeout = cfg.GetTimeOut()

	return &c
}

func (u URL) toRawURL() string {
	if len(u.RawURL) != 0 {
		return u.RawURL
	}
	url := u.Endpoint

	pathStart, pathEnd := 0, len(u.Path)
	if u.Path[0] == '/' {
		pathStart++
	}
	if u.Path[len(u.Path)-1] == '/' {
		pathEnd--
	}

	url += "/" + u.Path[pathStart:pathEnd]

	params := ""
	for key, value := range u.Params {
		params += "&" + key + "=" + value
	}

	if len(params) != 0 {
		url += "?" + params[1:]
	}
	return url
}

func (r *Reply) setRateLimit() {
	r.RateResource = r.Headers.Get("x-ratelimit-resource")
	r.RateLimit.Remaining, _ = strconv.Atoi(r.Headers.Get("x-ratelimit-remaining"))
	r.RateLimit.Limit, _ = strconv.Atoi(r.Headers.Get("x-ratelimit-limit"))
	r.RateLimit.Remaining, _ = strconv.Atoi(r.Headers.Get("x-ratelimit-remaining"))
	r.RateLimit.Reset, _ = strconv.Atoi(r.Headers.Get("x-ratelimit-reset"))
}

func resolveReply(reply *Reply, err error) (*Reply, error) {
	if err != nil {
		return reply, err
	}

	reply.setRateLimit()
	// return error if reply.statuscode != 2xx
	if reply.StatusCode >= 300 {
		return reply, ErrHttpResponse
	}

	// return error if rate limited
	if reply.RateLimit.Remaining == 0 {
		return reply, common.ErrRateLimited
	}
	return reply, nil
}
