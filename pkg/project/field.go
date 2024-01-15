package project

import (
	"fmt"
	"github.com/guguducken/octopus/pkg/common"
	"github.com/guguducken/octopus/pkg/config"
	"github.com/guguducken/octopus/pkg/issue"
	"time"
)

type FieldValue interface {
	CommonFieldValue | TextFieldValue | NumberFiledValue | LabelFieldValue
	GenQuery(cfg *config.Config, issue *issue.Issue, filed *Field, includeArchived bool, perPage int, cursor string, subCursor string) (string, error)
	FilterProject(projectV2 *ProjectV2) bool
}

type FieldValues[T FieldValue] []T

type CommonFieldValue struct {
	ID         string       `json:"id"`
	DatabaseID int          `json:"databaseId"`
	CreatedAt  *time.Time   `json:"createdAt"`
	UpdatedAt  *time.Time   `json:"updatedAt"`
	Creator    *common.User `json:"creator"`
	Field      *Field       `json:"field"`
	ProjectID  *ProjectID   `json:"projectID"`
}

type ProjectID struct {
	Project *ProjectV2 `json:"project"`
}

func (c CommonFieldValue) FilterProject(projectV2 *ProjectV2) bool {
	return projectV2.ID == c.ProjectID.Project.ID
}

func (c CommonFieldValue) GenQuery(cfg *config.Config, issue *issue.Issue, filed *Field, includeArchived bool, perPage int, cursor string, subCursor string) (string, error) {
	return "", nil
}

type TextFieldValue struct {
	Text string `json:"text"`
	CommonFieldValue
}

func (t TextFieldValue) GenQuery(cfg *config.Config, issue *issue.Issue, field *Field, includeArchived bool, perPage int, cursor string, subCursor string) (string, error) {
	return fmt.Sprintf(QueryIssueRelatedTextProjectV2Items,
		issue.Repository.Owner.Login,
		issue.Repository.Name,
		issue.Number,
		includeArchived,
		perPage,
		cursor,
		field.Name,
	), nil
}

type NumberFiledValue struct {
	Number float32 `json:"number"`
	CommonFieldValue
}

type LabelFieldValue struct {
	CommonFieldValue
}
