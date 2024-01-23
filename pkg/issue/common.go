package issue

import (
	"errors"
	"os/user"
	"time"

	"github.com/guguducken/octopus/pkg/pulls"
	"github.com/guguducken/octopus/pkg/repository"

	"github.com/guguducken/octopus/pkg/common"
)

const (
	EventMentioned       = "mentioned"
	EventSubscribed      = "subscribed"
	EventLabeled         = "labeled"
	EventUnLabeled       = "unlabeled"
	EventAssigned        = "assigned"
	EventUnAssigned      = "unassigned"
	EventMilestoned      = "milestoned"
	EventDeMilestoned    = "demilestoned"
	EventRenamed         = "renamed"
	EventCommented       = "commented"
	EventCrossReferenced = "cross-referenced"
	EventClosed          = "closed"
)

var (
	ErrNotIssue = errors.New("this is not an issue, maybe is a pull request")
)

type Issues []Issue

type Issue struct {
	URL                   string `json:"url"`
	RepositoryURL         string `json:"repository_url"`
	Repository            *repository.Repository
	LabelsURL             string             `json:"labels_url"`
	CommentsURL           string             `json:"comments_url"`
	EventsURL             string             `json:"events_url"`
	HTMLURL               string             `json:"html_url"`
	ID                    int                `json:"id"`
	NodeID                string             `json:"node_id"`
	Number                int                `json:"number"`
	Title                 string             `json:"title"`
	User                  *common.User       `json:"user"`
	Labels                []common.Label     `json:"labels"`
	State                 string             `json:"state"`
	Locked                bool               `json:"locked"`
	Assignee              *common.User       `json:"assignee"`
	Assignees             []common.User      `json:"assignees"`
	Milestone             *common.Milestone  `json:"milestone"`
	Comments              int                `json:"comments"`
	CreatedAt             *time.Time         `json:"created_at"`
	UpdatedAt             *time.Time         `json:"updated_at"`
	ClosedAt              *time.Time         `json:"closed_at"`
	AuthorAssociation     string             `json:"author_association"`
	ActiveLockReason      string             `json:"active_lock_reason"`
	Body                  string             `json:"body"`
	ClosedBy              *common.User       `json:"closed_by"`
	Reactions             *common.Reactions  `json:"reactions"`
	TimelineURL           string             `json:"timeline_url"`
	PerformedViaGithubApp bool               `json:"performed_via_github_app"`
	StateReason           string             `json:"state_reason"`
	PullRequest           *pulls.PullRequest `json:"pull_request"`
}

type Events []Event

type Event struct {
	// event: mentioned,subscribed
	ID                    int64        `json:"id"`
	NodeID                string       `json:"node_id"`
	URL                   string       `json:"url"`
	Actor                 *common.User `json:"actor"`
	Event                 string       `json:"event"`
	CommitID              string       `json:"commit_id"`
	CommitURL             string       `json:"commit_url"`
	CreatedAt             *time.Time   `json:"created_at"`
	PerformedViaGithubApp bool         `json:"performed_via_github_app"`
	// event: labeled,unlabeled
	LabeledEvent
	// event: assigned,unassigned
	AssignedEvent
	// event: milestoned,demilestoned
	MilestonedEvent
	// event: renamed
	RenameEvent
	// event: commented
	CommentedEvent
	// event: cross-referenced
	CrossReferencedEvent
	// event: closed
	ClosedEvent
}

type LabeledEvent struct {
	Label *common.Label `json:"label"`
}

type AssignedEvent struct {
	Assignee *user.User `json:"assignee"`
}

type MilestonedEvent struct {
	Milestone *common.Milestone `json:"milestone"`
}

type RenameEvent struct {
	Rename *Rename `json:"rename"`
}
type Rename struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type CommentedEvent struct {
	AuthorAssociation string            `json:"author_association"`
	UpdatedAt         *time.Time        `json:"updated_at"`
	Body              string            `json:"body"`
	Reactions         *common.Reactions `json:"reactions"`
	User              *common.User      `json:"user"`
	HTMLURL           string            `json:"html_url"`
	IssueURL          string            `json:"issue_url"`
}

type CrossReferencedEvent struct {
	Source *CrossReferencedSource `json:"source"`
}

type CrossReferencedSource struct {
	Type  string `json:"type"`
	Issue *Issue `json:"issue"`
}

type ClosedEvent struct {
	StateReason string `json:"state_reason"`
}
