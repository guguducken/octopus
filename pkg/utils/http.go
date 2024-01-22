package utils

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/guguducken/octopus/pkg/common"
	"github.com/guguducken/octopus/pkg/config"
)

var (
	ErrHttpResponse   = errors.New("github api server response is not 2xx")
	ErrHttpRequest    = errors.New("request failed")
	ErrHttpNewRequest = errors.New("new http request failed")
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

func GetWithRetryWithRateCheck(cfg *config.Config, url URL) (reply *Reply, err error) {
	retryTimes := cfg.GetRetryTimes()
	for j := 0; j < retryTimes; j++ {
		reply, err = GetWithRetry(cfg, url)
		if err != nil {
			if errors.Is(err, common.ErrRateLimited) {
				SleepTo(reply.RateLimit.Reset)
				continue
			}
			return nil, err
		}
		break
	}
	return reply, err
}

func PostWithRetryWithRateCheck(cfg *config.Config, url URL, body string) (reply *Reply, err error) {
	retryTimes := cfg.GetRetryTimes()
	for j := 0; j < retryTimes; j++ {
		reply, err = PostWithRetry(cfg, url, body)
		if err != nil {
			if errors.Is(err, common.ErrRateLimited) {
				SleepTo(reply.RateLimit.Reset)
				continue
			}
			return nil, err
		}
		break
	}
	return reply, err
}

func PostWithRetry(cfg *config.Config, url URL, body string) (reply *Reply, err error) {
	retryTimes := cfg.GetRetryTimes()
	retryDuration := cfg.GetRetryDuration()
	for i := 0; i < retryTimes; i++ {
		reply, err = Post(cfg, url, body)
		if err == nil {
			break
		}
		if err != nil && err != ErrHttpRequest {
			break
		}
		time.Sleep(retryDuration)
	}
	return reply, err

}

func GetWithRetry(cfg *config.Config, url URL) (reply *Reply, err error) {
	retryTimes := cfg.GetRetryTimes()
	retryDuration := cfg.GetRetryDuration()

	for i := 0; i < retryTimes; i++ {
		reply, err = Get(cfg, url)
		if err == nil {
			break
		}
		if err != nil && !errors.Is(err, ErrHttpRequest) {
			break
		}
		time.Sleep(retryDuration)
	}
	return reply, err
}

func Get(cfg *config.Config, url URL) (reply *Reply, err error) {
	req, err := http.NewRequest("GET", url.toRawURL(), nil)
	if err != nil {
		return reply, errors.Join(ErrHttpRequest, err)
	}
	return resolveReply(do(cfg, req))
}

func Post(cfg *config.Config, url URL, body string) (reply *Reply, err error) {
	req, err := http.NewRequest("POST", url.toRawURL(), strings.NewReader(body))
	if err != nil {
		return reply, errors.Join(ErrHttpRequest, err)
	}
	return resolveReply(do(cfg, req))
}

func do(cfg *config.Config, req *http.Request) (reply *Reply, err error) {
	reply = &Reply{}

	// set GitHub related headers, example: bearer token, api version
	setGithubHeader(req, cfg)

	// do http request
	resp, err := newHttpClient(cfg).Do(req)
	if err != nil {
		return reply, errors.Join(ErrHttpRequest, err)
	}

	// read http status code from resp
	reply.StatusCode = resp.StatusCode

	// read headers from resp
	reply.Headers = resp.Header
	reply.setRateLimit()

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

	if len(u.Path) != 0 {
		pathStart, pathEnd := 0, len(u.Path)
		if u.Path[0] == '/' {
			pathStart++
		}
		if u.Path[len(u.Path)-1] == '/' {
			pathEnd--
		}

		url += "/" + u.Path[pathStart:pathEnd]
	}

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
	r.RateLimit = &common.RateLimitUnit{}
	r.RateResource = r.Headers.Get("x-ratelimit-resource")
	r.RateLimit.Remaining, _ = strconv.Atoi(r.Headers.Get("x-ratelimit-remaining"))
	r.RateLimit.Limit, _ = strconv.Atoi(r.Headers.Get("x-ratelimit-limit"))
	r.RateLimit.Remaining, _ = strconv.Atoi(r.Headers.Get("x-ratelimit-remaining"))
	r.RateLimit.Reset, _ = strconv.ParseInt(r.Headers.Get("x-ratelimit-reset"), 10, 0)
}

func resolveReply(reply *Reply, err error) (*Reply, error) {
	if err != nil {
		return reply, err
	}

	// return error if rate limited
	if tryCheckWhetherIsRateLimited(reply) {
		return reply, common.ErrRateLimited
	}

	// return error if reply.statuscode != 2xx
	if reply.StatusCode >= 300 {
		return reply, errors.Join(ErrHttpResponse, errors.New(fmt.Sprintf("status: %d, resp message: %s", reply.StatusCode, string(reply.Body))))
	}

	return reply, nil
}

func SleepTo(target int64) {
	time.Sleep(time.Duration(target-time.Now().Unix()) * time.Second)
}

func tryCheckWhetherIsRateLimited(reply *Reply) bool {
	type Message struct {
		Message          string `json:"message"`
		DocumentationURL string `json:"documentation_url"`
	}
	m := &Message{}
	if err := json.Unmarshal(reply.Body, m); err == nil {
		return strings.Contains(m.DocumentationURL, "rate-limit") && reply.StatusCode == 403 && reply.RateLimit.Remaining == 0
	}
	return false
}
