package config

import (
	"strconv"
	"time"
)

func New(token string) *Config {
	return &Config{
		token: token,
		apiConfig: gitHubAPI{
			gitHubAPIVersion: DefaultGitHubAPIVersion,
			gitHubRestAPI:    DefaultGitHubRestAPI,
			gitHubGraphQLAPI: DefaultGitHubGraphQLAPI,
		},
		clientConfig: clientConfig{
			skipTLSVerify: false,
			proxyFromEnv:  false,
		},
		timeOut:       DefaultRequestTimeOut,
		retryTimes:    DefaultRetryTimes,
		retryDuration: DefaultRetryDuration,
		perPage:       strconv.Itoa(DefaultPerPage),
	}
}

func (c *Config) SkipTlsCheck() *Config {
	c.clientConfig.skipTLSVerify = true
	return c
}

func (c *Config) EnableProxyFromEnv() *Config {
	c.clientConfig.proxyFromEnv = true
	return c
}

func (c *Config) SetTimeOut(timoeOut time.Duration) *Config {
	c.timeOut = timoeOut
	return c
}

func (c *Config) SetRetryTimes(retryTimes int) *Config {
	c.retryTimes = retryTimes
	return c
}

func (c *Config) SetRetryDuration(retryDuration time.Duration) *Config {
	c.retryDuration = retryDuration
	return c
}

func (c *Config) SetPerPage(perPage int) *Config {
	if perPage <= 0 {
		perPage = DefaultPerPage
	}
	c.perPage = strconv.Itoa(perPage)
	return c
}

func (c *Config) GetTLSCheck() bool {
	return c.clientConfig.skipTLSVerify
}

func (c *Config) GetProxyFromEnv() bool {
	return c.clientConfig.proxyFromEnv
}

func (c *Config) GetTimeOut() time.Duration {
	return c.timeOut
}

func (c *Config) GetRetryTimes() int {
	return c.retryTimes
}

func (c *Config) GetToken() string {
	return c.token
}

func (c *Config) GetGithubRestAPI() string {
	return c.apiConfig.gitHubRestAPI
}

func (c *Config) GetGithubAPIVersion() string {
	return c.apiConfig.gitHubAPIVersion
}

func (c *Config) GetRetryDuration() time.Duration {
	return c.retryDuration
}

func (c *Config) GetPerPage() string {
	return c.perPage
}
