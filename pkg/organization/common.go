package organization

import (
	"time"

	"github.com/guguducken/octopus/pkg/common"
)

type Organization struct {
	Login                                                 string       `json:"login"`
	ID                                                    int          `json:"id"`
	NodeID                                                string       `json:"node_id"`
	URL                                                   string       `json:"url"`
	ReposURL                                              string       `json:"repos_url"`
	EventsURL                                             string       `json:"events_url"`
	HooksURL                                              string       `json:"hooks_url"`
	IssuesURL                                             string       `json:"issues_url"`
	MembersURL                                            string       `json:"members_url"`
	PublicMembersURL                                      string       `json:"public_members_url"`
	AvatarURL                                             string       `json:"avatar_url"`
	Description                                           string       `json:"description"`
	Name                                                  string       `json:"name"`
	Company                                               string       `json:"company"`
	Blog                                                  string       `json:"blog"`
	Location                                              string       `json:"location"`
	Email                                                 string       `json:"email"`
	TwitterUsername                                       string       `json:"twitter_username"`
	IsVerified                                            bool         `json:"is_verified"`
	HasOrganizationProjects                               bool         `json:"has_organization_projects"`
	HasRepositoryProjects                                 bool         `json:"has_repository_projects"`
	PublicRepos                                           int          `json:"public_repos"`
	PublicGists                                           int          `json:"public_gists"`
	Followers                                             int          `json:"followers"`
	Following                                             int          `json:"following"`
	HTMLURL                                               string       `json:"html_url"`
	CreatedAt                                             *time.Time   `json:"created_at"`
	Type                                                  string       `json:"type"`
	TotalPrivateRepos                                     int          `json:"total_private_repos"`
	OwnedPrivateRepos                                     int          `json:"owned_private_repos"`
	PrivateGists                                          int          `json:"private_gists"`
	DiskUsage                                             int          `json:"disk_usage"`
	Collaborators                                         int          `json:"collaborators"`
	BillingEmail                                          string       `json:"billing_email"`
	Plan                                                  *common.Plan `json:"plan"`
	DefaultRepositoryPermission                           string       `json:"default_repository_permission"`
	MembersCanCreateRepositories                          bool         `json:"members_can_create_repositories"`
	TwoFactorRequirementEnabled                           bool         `json:"two_factor_requirement_enabled"`
	MembersAllowedRepositoryCreationType                  string       `json:"members_allowed_repository_creation_type"`
	MembersCanCreatePublicRepositories                    bool         `json:"members_can_create_public_repositories"`
	MembersCanCreatePrivateRepositories                   bool         `json:"members_can_create_private_repositories"`
	MembersCanCreateInternalRepositories                  bool         `json:"members_can_create_internal_repositories"`
	MembersCanCreatePages                                 bool         `json:"members_can_create_pages"`
	MembersCanCreatePublicPages                           bool         `json:"members_can_create_public_pages"`
	MembersCanCreatePrivatePages                          bool         `json:"members_can_create_private_pages"`
	MembersCanForkPrivateRepositories                     bool         `json:"members_can_fork_private_repositories"`
	WebCommitSignoffRequired                              bool         `json:"web_commit_signoff_required"`
	UpdatedAt                                             *time.Time   `json:"updated_at"`
	DependencyGraphEnabledForNewRepositories              bool         `json:"dependency_graph_enabled_for_new_repositories"`
	DependabotAlertsEnabledForNewRepositories             bool         `json:"dependabot_alerts_enabled_for_new_repositories"`
	DependabotSecurityUpdatesEnabledForNewRepositories    bool         `json:"dependabot_security_updates_enabled_for_new_repositories"`
	AdvancedSecurityEnabledForNewRepositories             bool         `json:"advanced_security_enabled_for_new_repositories"`
	SecretScanningEnabledForNewRepositories               bool         `json:"secret_scanning_enabled_for_new_repositories"`
	SecretScanningPushProtectionEnabledForNewRepositories bool         `json:"secret_scanning_push_protection_enabled_for_new_repositories"`
	SecretScanningPushProtectionCustomLink                string       `json:"secret_scanning_push_protection_custom_link"`
	SecretScanningPushProtectionCustomLinkEnabled         bool         `json:"secret_scanning_push_protection_custom_link_enabled"`
}
