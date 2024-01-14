package common

import (
	"errors"
	"time"
)

var (
	ErrRateLimited = errors.New("")
)

type User struct {
	Login                   string     `json:"login"`
	ID                      int        `json:"id"`
	NodeID                  string     `json:"node_id"`
	AvatarURL               string     `json:"avatar_url"`
	GravatarID              string     `json:"gravatar_id"`
	URL                     string     `json:"url"`
	HTMLURL                 string     `json:"html_url"`
	FollowersURL            string     `json:"followers_url"`
	FollowingURL            string     `json:"following_url"`
	GistsURL                string     `json:"gists_url"`
	StarredURL              string     `json:"starred_url"`
	SubscriptionsURL        string     `json:"subscriptions_url"`
	OrganizationsURL        string     `json:"organizations_url"`
	ReposURL                string     `json:"repos_url"`
	EventsURL               string     `json:"events_url"`
	ReceivedEventsURL       string     `json:"received_events_url"`
	Type                    string     `json:"type"`
	SiteAdmin               bool       `json:"site_admin"`
	Name                    string     `json:"name"`
	Company                 string     `json:"company"`
	Blog                    string     `json:"blog"`
	Location                string     `json:"location"`
	Email                   string     `json:"email"`
	Hireable                bool       `json:"hireable"`
	Bio                     string     `json:"bio"`
	TwitterUsername         string     `json:"twitter_username"`
	PublicRepos             int        `json:"public_repos"`
	PublicGists             int        `json:"public_gists"`
	Followers               int        `json:"followers"`
	Following               int        `json:"following"`
	CreatedAt               *time.Time `json:"created_at"`
	UpdatedAt               *time.Time `json:"updated_at"`
	PrivateGists            int        `json:"private_gists"`
	TotalPrivateRepos       int        `json:"total_private_repos"`
	OwnedPrivateRepos       int        `json:"owned_private_repos"`
	DiskUsage               int        `json:"disk_usage"`
	Collaborators           int        `json:"collaborators"`
	TwoFactorAuthentication bool       `json:"two_factor_authentication"`
	Plan                    *Plan      `json:"plan"`
}

type Plan struct {
	Name         string `json:"name"`
	Space        int    `json:"space"`
	PrivateRepos int    `json:"private_repos"`
	// User only
	Collaborators int `json:"collaborators"`
	// Organization only
	FilledSeats int `json:"filled_seats"`
	Seats       int `json:"seats"`
}

type Label struct {
	ID          int    `json:"id"`
	NodeID      string `json:"node_id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Default     bool   `json:"default"`
}
type Milestone struct {
	URL          string     `json:"url"`
	HTMLURL      string     `json:"html_url"`
	LabelsURL    string     `json:"labels_url"`
	ID           int        `json:"id"`
	NodeID       string     `json:"node_id"`
	Number       int        `json:"number"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Creator      *User      `json:"creator"`
	OpenIssues   int        `json:"open_issues"`
	ClosedIssues int        `json:"closed_issues"`
	State        string     `json:"state"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	DueOn        *time.Time `json:"due_on"`
	ClosedAt     *time.Time `json:"closed_at"`
}

type Reactions struct {
	URL        string `json:"url"`
	TotalCount int    `json:"total_count"`
	Add        int    `json:"+1"`
	Sub        int    `json:"-1"`
	Laugh      int    `json:"laugh"`
	Hooray     int    `json:"hooray"`
	Confused   int    `json:"confused"`
	Heart      int    `json:"heart"`
	Rocket     int    `json:"rocket"`
	Eyes       int    `json:"eyes"`
}

type RateLimit struct {
	Resources Resources `json:"resources"`
	// deprecated: please use resource
	Rate *RateLimitUnit `json:"rate"`
}
type RateLimitUnit struct {
	Limit     int   `json:"limit"`
	Used      int   `json:"used"`
	Remaining int   `json:"remaining"`
	Reset     int64 `json:"reset"`
}

type Resources struct {
	CoreRateLimit                      *RateLimitUnit `json:"core"`
	SearchRateLimit                    *RateLimitUnit `json:"search"`
	GraphqlRateLimit                   *RateLimitUnit `json:"graphql"`
	IntegrationManifestRateLimit       *RateLimitUnit `json:"integration_manifest"`
	SourceImportRateLimit              *RateLimitUnit `json:"source_import"`
	CodeScanningUploadRateLimit        *RateLimitUnit `json:"code_scanning_upload"`
	ActionsRunnerRegistrationRateLimit *RateLimitUnit `json:"actions_runner_registration"`
	ScimRateLimit                      *RateLimitUnit `json:"scim"`
	DependencySnapshotsRateLimit       *RateLimitUnit `json:"dependency_snapshots"`
	CodeSearchRateLimit                *RateLimitUnit `json:"code_search"`
}
