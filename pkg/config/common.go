package config

import (
	"time"
)

const (
	DefaultGitHubAPIVersion = "2022-11-28"
	DefaultGitHubRestAPI    = "https://api.github.com"
	DefaultGitHubGraphQLAPI = "https://api.github.com/graphql"
	DefaultRequestTimeOut   = 30 * time.Second
	DefaultRetryTimes       = 10
	DefaultRetryDuration    = 5 * time.Second
	DefaultPerPage          = 100
)

type Config struct {
	token         string
	timeOut       time.Duration
	retryTimes    int
	retryDuration time.Duration
	perPage       int
	apiConfig     gitHubAPI
	clientConfig  clientConfig
}

type gitHubAPI struct {
	gitHubAPIVersion string
	gitHubRestAPI    string
	gitHubGraphQLAPI string
}

type clientConfig struct {
	skipTLSVerify bool
	proxyFromEnv  bool
}
