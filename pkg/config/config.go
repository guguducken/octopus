package config

func New(token string) *Config {
	return &Config{
		Token: token,
		ApiConfig: GitHubAPI{
			GitHubAPIVersion: DEFAULT_GITHUB_API_VERSION,
			GitHubRestAPI:    DEFAULT_GITHUB_REST_API,
			GitHubGraphQLAPI: DEFAULT_GITHU_GRAPHQL_API,
		},
		ClientConfig: ClientConfig{
			SkipTLSVerify: false,
			ProxyEnable:   false,
		},
	}
}

func (c *Config) SkipTlsCheck() {
	c.ClientConfig.SkipTLSVerify = true
}

func (c *Config) EnableProxy() {
	c.ClientConfig.ProxyEnable = true
}
