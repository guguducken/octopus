package config

import "time"

const (
	DEFAULT_GITHUB_API_VERSION = "2022-11-28"
	DEFAULT_GITHUB_REST_API    = "https://api.github.com"
	DEFAULT_GITHU_GRAPHQL_API  = "https://api.github.com/graphql"
)

type Config struct {
	Token        string
	TimeOut      time.Duration
	ApiConfig    GitHubAPI
	ClientConfig ClientConfig

	User User
}

type User struct {
	Login string
	Email string
}

type GitHubAPI struct {
	GitHubAPIVersion string
	GitHubRestAPI    string
	GitHubGraphQLAPI string
}

type ClientConfig struct {
	SkipTLSVerify bool
	ProxyEnable   bool
}
