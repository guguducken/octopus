package project

import (
	"fmt"
	"github.com/guguducken/octopus/pkg/common"
	"github.com/guguducken/octopus/pkg/config"
	"github.com/guguducken/octopus/pkg/issue"
	"time"
)

type FieldValuer interface {
	CommonFieldValue | TextFieldValue | NumberFieldValue | LabelFieldValue | SingleSelectFieldValue

	GenQuery(cfg *config.Config, issue *issue.Issue, filed *Field, includeArchived bool, perPage int, cursor string, subCursor string) (string, error)
	FilterProject(projectV2 *ProjectV2) bool
	GetSubPageInfo() *PageInfo
	IsNil() bool
}

type FieldValues[T FieldValuer] struct {
	fieldValues []T
	pageInfo    *PageInfo
}

func (fvs *FieldValues[T]) GetFieldValues() []T {
	return fvs.fieldValues
}

func GenFieldValuers[T FieldValuer]() *FieldValues[T] {
	return &FieldValues[T]{
		fieldValues: make([]T, 0, 10),
		pageInfo:    nil,
	}
}

func (fvs *FieldValues[T]) Add(value T) {
	fvs.fieldValues = append(fvs.fieldValues, value)
}
func (fvs *FieldValues[T]) AddSlice(values []T) {
	fvs.fieldValues = append(fvs.fieldValues, values...)
}

func (fvs *FieldValues[T]) SetPageInfo(pageInfo *PageInfo) {
	fvs.pageInfo = &PageInfo{
		HasNextPage:     pageInfo.HasNextPage,
		HasPreviousPage: pageInfo.HasPreviousPage,
		StartCursor:     pageInfo.StartCursor,
		EndCursor:       pageInfo.EndCursor,
	}
}

func (fvs *FieldValues[T]) GetPageInfo() *PageInfo {
	return fvs.pageInfo
}

type CommonFieldValue struct {
	ID         string     `json:"id"`
	DatabaseID int        `json:"databaseId"`
	CreatedAt  *time.Time `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
	// only get login for creator
	Creator *common.User `json:"creator"`
	Field   *Field       `json:"field"`
	// PageInfo for temp
	pageInfo *PageInfo
}

func (c CommonFieldValue) FilterProject(projectV2 *ProjectV2) bool {
	return projectV2.ID == c.Field.Project.ID
}

func (c CommonFieldValue) GenQuery(cfg *config.Config, issue *issue.Issue, filed *Field, includeArchived bool, perPage int, cursor string, subCursor string) (string, error) {
	return "", nil
}

func (c CommonFieldValue) GetSubPageInfo() *PageInfo {
	return nil
}

func (c CommonFieldValue) IsNil() bool {
	return c.Field == nil
}

type TextFieldValue struct {
	Text string `json:"text"`
	CommonFieldValue
}

func (t TextFieldValue) GenQuery(cfg *config.Config, issue *issue.Issue,
	field *Field, includeArchived bool, perPage int, cursor string, subCursor string) (string, error) {
	return QueryToJson(fmt.Sprintf(QueryIssueRelatedTextProjectV2Items,
		issue.Repository.Owner.Login,
		issue.Repository.Name,
		issue.Number,
		includeArchived,
		perPage,
		cursor,
		field.Name,
	)), nil
}

type NumberFieldValue struct {
	Number float32 `json:"number"`
	CommonFieldValue
}

type LabelFieldValue struct {
	CommonFieldValue
}

type SingleSelectFieldValue struct {
	CommonFieldValue

	Name            string `json:"name"`
	NameHTML        string `json:"nameHTML"`
	Color           string `json:"color"`
	Description     string `json:"description"`
	DescriptionHTML string `json:"descriptionHTML"`
	OptionID        string `json:"optionId"`
}

func (ss SingleSelectFieldValue) GenQuery(cfg *config.Config, issue *issue.Issue, field *Field, includeArchived bool, perPage int, cursor string, subCursor string) (string, error) {
	return QueryToJson(fmt.Sprintf(QueryIssueRelatedSingleSelectProjectV2Items,
		issue.Repository.Owner.Login,
		issue.Repository.Name,
		issue.Number,
		includeArchived,
		perPage,
		cursor,
		field.Name,
	)), nil
}
