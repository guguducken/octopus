package project

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/guguducken/octopus/pkg/common"
)

var (
	ErrGraphQLResult = errors.New("the response from graphql server contain error")
)

type ProjectReply[T FieldValue] struct {
	Data   *Data[T]      `json:"data"`
	Errors GraphQLErrors `json:"errors"`
}

type Data[T FieldValue] struct {
	Organization *Organization  `json:"organization"`
	Repository   *Repository[T] `json:"repository"`
}

type Organization struct {
	Login     string     `json:"login"`
	ID        string     `json:"id"`
	ProjectV2 *ProjectV2 `json:"projectV2"`
}

type Repository[T FieldValue] struct {
	Issue *Issue[T] `json:"issue"`
}
type Issue[T FieldValue] struct {
	ProjectsV2   *IssueRelatedProjectReply `json:"projectsV2"`
	ProjectItems *Items[T]                 `json:"projectItems"`
}
type IssueRelatedProjectReply struct {
	Nodes    []ProjectV2 `json:"nodes"`
	PageInfo *PageInfo   `json:"pageInfo"`
}

type Items[T FieldValue] struct {
	Nodes []FieldValueByName[T] `json:"nodes"`
}

type FieldValueByName[T FieldValue] struct {
	FieldValueByName T `json:"fieldValueByName"`
}
type ProjectV2 struct {
	Title     string     `json:"title"`
	ID        string     `json:"id"`
	Closed    bool       `json:"closed"`
	ClosedAt  *time.Time `json:"closedAt"`
	CreatedAt *time.Time `json:"createdAt"`
	// we only get login for creator
	Creator          *common.User `json:"creator"`
	DatabaseID       int          `json:"databaseId"`
	Number           int          `json:"number"`
	Owner            Owner        `json:"owner"`
	Public           bool         `json:"public"`
	Readme           string       `json:"readme"`
	ResourcePath     string       `json:"resourcePath"`
	ShortDescription string       `json:"shortDescription"`
	Template         bool         `json:"template"`
	UpdatedAt        *time.Time   `json:"updatedAt"`
	URL              string       `json:"url"`
	ViewerCanClose   bool         `json:"viewerCanClose"`
	ViewerCanReopen  bool         `json:"viewerCanReopen"`
	ViewerCanUpdate  bool         `json:"viewerCanUpdate"`
	Fields           *FieldsReply `json:"fields"`
	// temp organization struct for record project's org information
	Organization struct {
		Login string
		ID    string
	}
}
type Owner struct {
	ID string `json:"id"`
}

type GraphQLErrors []GraphQLError

func (g GraphQLErrors) ToJson() string {
	result, _ := json.Marshal(&g)
	return string(result)
}

type GraphQLError struct {
	Type      string           `json:"type"`
	Path      []string         `json:"path"`
	Locations []ErrorLocations `json:"locations"`
	Message   string           `json:"message"`
}
type ErrorLocations struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

type FieldsReply struct {
	Nodes    []Field   `json:"nodes"`
	PageInfo *PageInfo `json:"pageInfo"`
}
type Field struct {
	CreatedAt  *time.Time `json:"createdAt"`
	DataType   string     `json:"dataType"`
	DatabaseID int        `json:"databaseId"`
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	// will only get id and number
	Project   *ProjectV2 `json:"project"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type PageInfo struct {
	HasNextPage     bool   `json:"hasNextPage"`
	HasPreviousPage bool   `json:"hasPreviousPage"`
	StartCursor     string `json:"startCursor"`
	EndCursor       string `json:"endCursor"`
}
