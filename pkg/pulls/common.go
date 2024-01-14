package pulls

import (
	"time"

	"github.com/guguducken/octopus/pkg/common"
	"github.com/guguducken/octopus/pkg/organization"
	"github.com/guguducken/octopus/pkg/repository"
)

type PullRequest struct {
	URL                 string              `json:"url"`
	ID                  int                 `json:"id"`
	NodeID              string              `json:"node_id"`
	HTMLURL             string              `json:"html_url"`
	DiffURL             string              `json:"diff_url"`
	PatchURL            string              `json:"patch_url"`
	IssueURL            string              `json:"issue_url"`
	CommitsURL          string              `json:"commits_url"`
	ReviewCommentsURL   string              `json:"review_comments_url"`
	ReviewCommentURL    string              `json:"review_comment_url"`
	CommentsURL         string              `json:"comments_url"`
	StatusesURL         string              `json:"statuses_url"`
	Number              int                 `json:"number"`
	State               string              `json:"state"`
	Locked              bool                `json:"locked"`
	Title               string              `json:"title"`
	User                *common.User        `json:"user"`
	Body                string              `json:"body"`
	Labels              []common.Label      `json:"labels"`
	Milestone           *common.Milestone   `json:"milestone"`
	ActiveLockReason    string              `json:"active_lock_reason"`
	CreatedAt           *time.Time          `json:"created_at"`
	UpdatedAt           *time.Time          `json:"updated_at"`
	ClosedAt            *time.Time          `json:"closed_at"`
	MergedAt            *time.Time          `json:"merged_at"`
	MergeCommitSha      string              `json:"merge_commit_sha"`
	Assignee            *common.User        `json:"assignee"`
	Assignees           []common.User       `json:"assignees"`
	RequestedReviewers  []common.User       `json:"requested_reviewers"`
	RequestedTeams      []organization.Team `json:"requested_teams"`
	Head                *PullRef            `json:"head"`
	Base                *PullRef            `json:"base"`
	Links               *Links              `json:"_links"`
	AuthorAssociation   string              `json:"author_association"`
	AutoMerge           bool                `json:"auto_merge"`
	Draft               bool                `json:"draft"`
	Merged              bool                `json:"merged"`
	Mergeable           bool                `json:"mergeable"`
	Rebaseable          bool                `json:"rebaseable"`
	MergeableState      string              `json:"mergeable_state"`
	MergedBy            *common.User        `json:"merged_by"`
	Comments            int                 `json:"comments"`
	ReviewComments      int                 `json:"review_comments"`
	MaintainerCanModify bool                `json:"maintainer_can_modify"`
	Commits             int                 `json:"commits"`
	Additions           int                 `json:"additions"`
	Deletions           int                 `json:"deletions"`
	ChangedFiles        int                 `json:"changed_files"`
}

type PullRef struct {
	Label string                 `json:"label"`
	Ref   string                 `json:"ref"`
	Sha   string                 `json:"sha"`
	User  *common.User           `json:"user"`
	Repo  *repository.Repository `json:"repo"`
}
type Links struct {
	Self           *Link `json:"self"`
	HTML           *Link `json:"html"`
	Issue          *Link `json:"issue"`
	Comments       *Link `json:"comments"`
	ReviewComments *Link `json:"review_comments"`
	ReviewComment  *Link `json:"review_comment"`
	Commits        *Link `json:"commits"`
	Statuses       *Link `json:"statuses"`
}

type Link struct {
	Href string `json:"href"`
}
