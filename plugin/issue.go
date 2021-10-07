package plugin

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/google/go-github/v35/github"
)

// issueClient ties the drone env data and github client together.
type issueClient struct {
	Client              *github.Client
	Context             context.Context
	Owner               string
	Repo                string
	Title               string
	Body                string
	BodyTextAttachments []string
	Assignees           []string
	Labels              []string
}

// createIssue creates an issue if one doesn't exist.
func (ic *issueClient) createIssue() (*github.Issue, error) {
	// check if issue exists before creating one
	issue, err := ic.getIssue()
	if err != nil {
		return nil, fmt.Errorf("failed to search for issue: %w", err)
	}

	if issue != nil {
		return issue, nil
	}

	// issue was not found in search, create an issue
	newIssue, err := ic.newIssue()
	if err != nil {
		return nil, fmt.Errorf("failed to create new issue: %w", err)
	}

	return newIssue, nil
}

// getIssue retrieves an issue if it exists.
func (ic *issueClient) getIssue() (*github.Issue, error) {
	searchOpts := &github.SearchOptions{
		ListOptions: github.ListOptions{
			PerPage: 1,
		},
	}
	searchQuery := fmt.Sprintf(
		"%s in:title repo:%s/%s type:issue state:open",
		ic.Title, ic.Owner, ic.Repo,
	)
	issues, _, err := ic.Client.Search.Issues(ic.Context, searchQuery, searchOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to search issues: %w", err)
	}

	if *issues.Total > 0 {
		issue := issues.Issues[0]

		fmt.Printf(
			"Found issue %d with title `%s`\n", issue.GetID(), ic.Title,
		)
		return issues.Issues[0], nil
	}

	fmt.Println("No issue found with the given title")
	return nil, nil
}

// newIssue creates a new issue.
func (ic *issueClient) newIssue() (*github.Issue, error) {
	issueBody, issueBodyErr := ic.getBodyString()
	if issueBodyErr != nil {
		return nil, fmt.Errorf("failed to create issue: %w", issueBodyErr)
	}

	issueReq := &github.IssueRequest{
		Title:     &ic.Title,
		Body:      &issueBody,
		Labels:    &ic.Labels,
		Assignees: &ic.Assignees,
	}

	issue, _, err := ic.Client.Issues.Create(
		ic.Context, ic.Owner, ic.Repo, issueReq,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create issue: %w", err)
	}

	fmt.Printf(
		"Successfully created issue %d with title `%s`\n",
		issue.GetID(), ic.Title,
	)
	return issue, nil
}

// getBody returns the issues body. The function will append the values of the text
// attachments to the body
func (ic *issueClient) getBodyString() (string, error) {
	bodyString := ic.Body

	for _, curAttachment := range ic.BodyTextAttachments {
		curAttContents, curAttContentsErr := ioutil.ReadFile(curAttachment)
		if curAttContentsErr != nil {
			return bodyString, curAttContentsErr
		}
		bodyString = fmt.Sprintf("%s\n%s", bodyString, string(curAttContents))
	}

	return bodyString, nil
}
