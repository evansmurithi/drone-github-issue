package plugin

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-github/v35/github"
	"github.com/urfave/cli/v2"
	"golang.org/x/oauth2"
)

// Settings for the plugin.
type Settings struct {
	APIKey              string
	Title               string
	Body                string
	BodyTextAttachments cli.StringSlice
	Assignees           cli.StringSlice
	Labels              cli.StringSlice
	BaseURL             string

	baseURL *url.URL
}

// Validate handles the settings validation of the plugin.
func (p *Plugin) Validate() error {
	var err error

	if p.settings.APIKey == "" {
		return fmt.Errorf("no api key provided")
	}

	if p.settings.Title == "" {
		return fmt.Errorf("no github issue title provided")
	}

	if !strings.HasSuffix(p.settings.BaseURL, "/") {
		p.settings.BaseURL = p.settings.BaseURL + "/"
	}
	p.settings.baseURL, err = url.Parse(p.settings.BaseURL)
	if err != nil {
		return fmt.Errorf("failed to parse base url: %w", err)
	}

	return nil
}

// Execute provides the implementation of the plugin.
func (p *Plugin) Execute() error {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: p.settings.APIKey})
	httpClient := oauth2.NewClient(
		context.WithValue(context.Background(), oauth2.HTTPClient, p.network.Client),
		tokenSource,
	)

	ghClient := github.NewClient(httpClient)

	ghClient.BaseURL = p.settings.baseURL

	ic := issueClient{
		Client:              ghClient,
		Context:             p.network.Context,
		Owner:               p.pipeline.Repo.Owner,
		Repo:                p.pipeline.Repo.Name,
		Title:               p.settings.Title,
		Body:                p.settings.Body,
		BodyTextAttachments: p.settings.BodyTextAttachments.Value(),
		Assignees:           p.settings.Assignees.Value(),
		Labels:              p.settings.Labels.Value(),
	}

	_, err := ic.createIssue()
	if err != nil {
		return fmt.Errorf("failed to create the issue: %w", err)
	}

	return nil
}
