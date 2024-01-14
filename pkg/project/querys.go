package project

import (
	"fmt"
	"strings"
)

func QueryToJson(query string) (jsonQuery string) {
	// remove \n
	query = strings.Replace(query, "\n", " ", -1)
	// remove \t
	query = strings.Replace(query, "\t", " ", -1)
	for strings.Contains(query, "  ") {
		query = strings.Replace(query, "  ", " ", -1)
	}
	for strings.Contains(query, "} }") {
		query = strings.Replace(query, "} }", "}}", -1)
	}
	query = strings.Replace(query, `"`, `\"`, -1)
	return fmt.Sprintf(`{"query":"%s"}`, query)
}

const GetProjectForOrgByID = `{
  organization(login: "%s") {
    login
    id
    projectV2(number: %d) {
      title
      id
      closed
      closedAt
      createdAt
      creator {
        login
      }
      databaseId
      number
      owner {
        id
      }
      public
      readme
      resourcePath
      shortDescription
      template
      updatedAt
      url
      viewerCanClose
      viewerCanReopen
      viewerCanUpdate
    }
  }
}`

const GetProjectV2FieldsByPage = `{
  organization(login: "%s") {
    projectV2(number: %d) {
      fields(first: %d, after: "%s") {
        nodes {
        ... on ProjectV2FieldCommon {
            createdAt
            dataType
            databaseId
            id
            name
            project {
              id
            }
            updatedAt
        }
        }
        pageInfo {
          hasNextPage
          hasPreviousPage
          startCursor
          endCursor
        }
      }
    }
  }
}`
