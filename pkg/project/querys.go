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

const QueryProjectForOrgByID = `{
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

const QueryProjectV2FieldsByPage = `{
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
              number
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

const QueryIssueRelatedProjectsV2 = `{
  repository(owner: "%s", name: "%s") {
    issue(number: %d) {
      number
      projectsV2(first: %d,after:"%s") {
        nodes {
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

const QueryIssueRelatedTextProjectV2Items = `{
  repository(owner: "%s", name: "%s") {
    issue(number: %d) {
      projectItems(includeArchived: %t, first: %d, after: "%s") {
        nodes {
          fieldValueByName(name: "%s") {
            ... on ProjectV2ItemFieldTextValue {
              text
              id
              databaseId
              createdAt
              updatedAt
              creator {
                login
              }
              field{
                ... on ProjectV2FieldCommon {
                  name
                  dataType
                  project{
                    id
                    number
                  }
                }
              }
            }
          }
        }
        pageInfo{
          hasNextPage
          hasPreviousPage
          startCursor
          endCursor
        }
      }
    }
  }
}`

const QueryIssueRelatedSingleSelectProjectV2Items = `{
  repository(owner: "%s", name: "%s") {
    issue(number: %d) {
      projectItems(includeArchived: %t, first: %d,after: "%s") {
        nodes {
          fieldValueByName(name: "%s") {
            ... on ProjectV2ItemFieldSingleSelectValue {
              name
              nameHTML
              color
              description
              descriptionHTML
              optionId
              id
              databaseId
              createdAt
              updatedAt
              creator {
                login
              }
              field{
                ... on ProjectV2FieldCommon {
                  name
                  dataType
                  project{
                    id
                    number
                  }
                }
              }
            }
          }
          
        }
        pageInfo{
          hasNextPage
          hasPreviousPage
          startCursor
          endCursor
        }
      }
    }
  }
}`
