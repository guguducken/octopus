package utils

import (
	"crypto/tls"
	"io"
	"net/http"
	"octopus/pkg/config"
	"strings"
)

type Reply struct {
	Body       []byte
	StatusCode int
	Headers    http.Header
}

type URL struct {
	Endpoint string
	Path     string
	Params   map[string]string
}

func Get(cfg *config.Config, url URL) (reply Reply, err error) {
	req, err := http.NewRequest("GET", url.toRawURL(), nil)
	if err != nil {
		return reply, err
	}
	return do(cfg, req)
}

func Post(cfg *config.Config, url URL, body string) (reply Reply, err error) {
	req, err := http.NewRequest("POST", url.toRawURL(), strings.NewReader(body))
	if err != nil {
		return reply, err
	}
	return do(cfg, req)
}

func do(cfg *config.Config, req *http.Request) (reply Reply, err error) {

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
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Authorization", "Bearer "+cfg.Token)
	req.Header.Set("X-GitHub-Api-Version", cfg.ApiConfig.GitHubAPIVersion)
}

func newHttpClient(cfg *config.Config) *http.Client {
	c := http.Client{}

	// transport settings
	t := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: cfg.ClientConfig.SkipTLSVerify,
		},
	}
	if cfg.ClientConfig.ProxyEnable {
		t.Proxy = http.ProxyFromEnvironment
	}

	c.Transport = t
	c.Timeout = cfg.TimeOut

	return &c
}

func (u URL) toRawURL() string {
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
